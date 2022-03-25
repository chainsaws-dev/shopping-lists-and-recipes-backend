package signinupout

import (
	"bytes"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"time"

	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/internal/setup"
	"shopping-lists-and-recipes/packages/authentication"
	"shopping-lists-and-recipes/packages/securecookies"
	"shopping-lists-and-recipes/packages/shared"

	"github.com/gorilla/securecookie"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Список типовых ошибок
var (
	ErrPasswordTooShort         = errors.New("password is too short")
	ErrPasswordLength           = errors.New("password must be more than six characters")
	ErrPasswordNewUserNotFilled = errors.New("password of the new user must be set")
	ErrNotAuthorized            = errors.New("wrong login or password")
	ErrBadEmail                 = errors.New("invalid email specified")
	ErrBadPhone                 = errors.New("invalid phone number specified")
	ErrBadRole                  = errors.New("invalid role specified")
	ErrSessionNotFoundByEmail   = errors.New("session is not found for specified email")
	ErrSessionNotFoundByToken   = errors.New("session is not found for specified token")
	ErrUserDisabled             = errors.New("you are denied access to the resource")
	ErrPasswdHashCalc           = errors.New("error calculating password hash")
	ErrEmailRegistered          = errors.New("specified email address is already taken")
	ErrIncorrectUserID          = errors.New("invalid user id specified")
	ErrUnableToDeleteAbsent     = errors.New("user not found, unable to delete")
)

var (
	MesEmailSent       = "email was sent"
	MesEmailConfirmed  = "email successfully confirmed"
	MesPasswdChanged   = "password changed"
	MesUserDeleted     = "user deleted"
	MesSessionsDeleted = "sessions deleted"
	MesSessionDeleted  = "session deleted"
)

// AuthBasic - базовая аутентификация проверка API ключа
func AuthBasic(w http.ResponseWriter, req *http.Request) bool {

	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest) {
			return false
		}
	}

	if found {
		return true
	}
	shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, shared.ErrWrongKeyInParams)
	return false
}

// AuthNoSecondFactor - Полная аутентификация пользователя используется только для проверки 2 фактора
func AuthNoSecondFactor(w http.ResponseWriter, req *http.Request) (string, bool) {
	if !AuthBasic(w, req) {
		return "", false
	}

	// Проверка токена и получение роли
	issued, role := TwoWayAuthentication(w, req)

	if issued {
		return role, true
	}

	shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAuthorized.Error(), shared.ErrNotAuthorized, http.StatusUnauthorized)
	return "", false

}

// AuthGeneral - Полная аутентификация пользователя для админки
func AuthGeneral(w http.ResponseWriter, req *http.Request) (string, bool) {
	if !AuthBasic(w, req) {
		return "", false
	}

	// Проверка токена и получение роли
	issued, role := TwoWayAuthentication(w, req)

	// Проверка прохождения двухфакторной авторизации
	sf := SecondFactorAuthenticationCheck(w, req)

	if issued {
		if sf {
			return role, true
		}

		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAuthorizedTwoFactor.Error(), shared.ErrNotAuthorizedTwoFactor, http.StatusUnauthorized)
		return "", false
	}

	shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAuthorized.Error(), shared.ErrNotAuthorized, http.StatusUnauthorized)
	return "", false

}

// secretauth - внутренняя функция для проверки пароля и авторизации
// (если ReturnToken=false - то куки)
func secretauth(w http.ResponseWriter, req *http.Request, AuthRequest authentication.AuthRequestData) {

	if len(AuthRequest.Password) < 6 {
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrPasswordLength.Error(), ErrPasswordLength, http.StatusBadRequest)
		return
	}

	UserAgent := req.Header.Get("User-Agent")
	ClientIP := GetIP(req)

	// Получаем хеш из базы данных
	strhash, strrole, err := databases.PostgreSQLGetTokenForUser(AuthRequest.Email, setup.ServerSettings.SQL.ConnPool)

	if err != nil {
		if shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusTeapot) {
			return
		}
	}

	// Проверяем пароль против хеша
	match, err := authentication.Argon2ComparePasswordAndHash(AuthRequest.Password, strhash)

	if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
		return
	}

	if match {

		// Удаляем просроченые токены
		err = databases.PostgreSQLDeleteExpiredSessions(setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			if !errors.Is(err, databases.ErrSessionsNotFoundExpired) {
				log.Println(err)
			}
		}

		if CountTokensByEmail(AuthRequest.Email) > 1 {
			// Удаляем остальные токены если число сессий превышает 2
			DeleteSessionByEmail(AuthRequest.Email)
		}

		// Получаем текущего пользователя по электронной почте
		FoundUser, err := databases.PostgreSQLGetUserByEmail(AuthRequest.Email, setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			if errors.Is(databases.ErrNoUserWithEmail, err) {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
			}
		}

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		if FoundUser.Disabled {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrUserDisabled.Error(), ErrUserDisabled, http.StatusForbidden)
			return
		}

		// Генерим случайные 32 байта
		tokenb, err := authentication.GenerateRandomBytes(32)

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		// Формируем ответ
		AuthResponse := authentication.AuthResponseData{
			Token:      hex.EncodeToString(tokenb),
			Email:      AuthRequest.Email,
			ExpiresIn:  3600,
			Registered: true,
			Role:       strrole,
			SecondFactor: authentication.TOTP{
				Enabled:     FoundUser.SecondFactor,
				CheckResult: false,
			},
			Locale: FoundUser.Lang,
		}

		// Формируем и запоминаем сессию
		tb := time.Now()
		te := tb.Add(time.Hour)

		NewActiveToken := authentication.ActiveToken{
			Email:     AuthRequest.Email,
			Token:     AuthResponse.Token,
			Session:   securecookie.GenerateRandomKey(64),
			IssDate:   tb,
			ExpDate:   te,
			Role:      strrole,
			UserAgent: UserAgent,
			IP:        ClientIP,
			SecondFactor: authentication.TOTP{
				Enabled:     FoundUser.SecondFactor,
				CheckResult: false,
			},
		}

		err = databases.PostgreSQLSessionsInsert(NewActiveToken, setup.ServerSettings.SQL.ConnPool)

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		// Если не возвращаем токен, то пишем куки
		if !AuthRequest.ReturnSecureToken {
			AuthResponse.Token = ""

			err = securecookies.SetCookies(te, NewActiveToken, w)

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}
		}

		shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, AuthResponse)

	} else {
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
	}
}

// CheckAPIKey - проверяет API ключ
func CheckAPIKey(w http.ResponseWriter, req *http.Request) (bool, error) {

	APIKey := req.Header.Get("ApiKey")

	if len(APIKey) < 1 {
		return false, shared.ErrNoKeyInParams
	}

	_, found := shared.FindInStringSlice(setup.APIkeys, APIKey)

	if !found {
		return false, shared.ErrWrongKeyInParams
	}

	return true, nil

}

// TwoWayAuthentication - выполняет аутентификацию как с помощью заголовка Auth, так и с помощью куки
func TwoWayAuthentication(w http.ResponseWriter, req *http.Request) (issued bool, role string) {
	// Освобождаем память от истекших токенов
	err := databases.PostgreSQLDeleteExpiredSessions(setup.ServerSettings.SQL.ConnPool)

	if err != nil {
		if !errors.Is(err, databases.ErrSessionsNotFoundExpired) {
			log.Println(err)
		}
	}
	// Проверка кук и получение роли
	cookiefound, role := CheckCookiesIssued(w, req)

	if cookiefound {
		return cookiefound, role
	}

	// Проверка токена и получение роли (если нет кук)
	return CheckTokensIssued(req)

}

// SecondFactorAuthenticationCheck - проверяет значение двухфакторной авторизации
func SecondFactorAuthenticationCheck(w http.ResponseWriter, req *http.Request) bool {

	result, err := GetCurrentSession(w, req)

	if err != nil {
		log.Println(err)
		return false
	}

	if result.SecondFactor.Enabled && result.SecondFactor.CheckResult {
		return true
	}

	if !result.SecondFactor.Enabled {
		return true
	}

	return false
}

// CheckCookiesIssued - проверяет выпущенные куки
func CheckCookiesIssued(w http.ResponseWriter, req *http.Request) (issued bool, role string) {
	mycookies, err := securecookies.GetCookies(w, req)

	if err != nil {

		if !errors.Is(http.ErrNoCookie, err) {

			if errors.Is(securecookie.ErrMacInvalid, err) {
				log.Println("Невозможно расширфровать куки (HMAC устарел)")
				return false, ""
			}

			log.Println(err)
			return false, ""

		}
	}

	if len(mycookies.Email) > 0 && len(mycookies.Session) > 0 {

		// Ищем живые токены по сессии
		at, found := SearchIssuedSessions(mycookies.Email, mycookies.Session)

		if found {
			return found, at.Role
		}

		log.Println(securecookies.ErrAuthCookiesNotFound)
		return false, ""
	}

	return false, ""
}

// CheckTokensIssued - ищет активный токен по электронной почте
func CheckTokensIssued(req *http.Request) (issued bool, role string) {

	Token := req.Header.Get("Auth")

	if len(Token) > 0 {

		at, err := databases.PostgreSQLGetActiveTokenByToken(Token, setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			log.Println(err)
			return false, ""
		} else {
			return true, at.Role
		}

	}

	return false, ""
}

// SearchIssuedSessions - ищет активный токен по электронной почте и с совпадающей сессией
func SearchIssuedSessions(Email string, Session []byte) (authentication.ActiveToken, bool) {
	if len(Email) != 0 {

		t, err := databases.PostgreSQLGetActiveTokenBySession(Email, Session, setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			log.Println(err)
			return authentication.ActiveToken{}, false
		} else {
			return t, true
		}
	}

	return authentication.ActiveToken{}, false
}

// DeleteSessionByEmail - ищет сессию по электронной почте и удаляет её
func DeleteSessionByEmail(Email string) error {

	return databases.PostgreSQLDeleteSessionsByEmail(Email, setup.ServerSettings.SQL.ConnPool)

}

// DeleteSessionByToken - ищет сессию по токену и удаляет её
func DeleteSessionByToken(Token string) error {

	if len(Token) > 0 {
		return databases.PostgreSQLDeleteSessionsByToken(Token, setup.ServerSettings.SQL.ConnPool)
	} else {
		return ErrSessionNotFoundByToken
	}

}

// CompareSessions - выполняет сравнение сессий
func CompareSessions(at authentication.ActiveToken, SessionToCompare []byte) bool {
	res := bytes.Compare(at.Session, SessionToCompare)

	if res == 0 {
		return true
	}

	return false
}

// CountTokensByEmail - cчитает количество токенов с одним Email в списке сессий
func CountTokensByEmail(Email string) int {

	if len(Email) > 0 {

		count, err := databases.PostgreSQLCountTokensByEmail(Email, setup.ServerSettings.SQL.ConnPool)

		if err == nil {
			return count
		} else {
			log.Println(err)
			return 0
		}

	}

	return 0

}

// GetIP - получает IP адрес клиента
func GetIP(r *http.Request) string {

	IPAddress := r.Header.Get("X-Real-Ip")

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	return IPAddress
}

// ConvertToSignInRequest - преобразует тип запроса регистрации в тип запроса авторизации
func ConvertToSignInRequest(SignUpRequest authentication.AuthSignUpRequestData) authentication.AuthRequestData {
	return authentication.AuthRequestData{
		Email:             SignUpRequest.Email,
		Password:          SignUpRequest.Password,
		ReturnSecureToken: true,
	}
}

// RegularConfirmTokensCleanup - в фоновом режиме удаляет устаревшие токены
func RegularConfirmTokensCleanup(dbc *pgxpool.Pool) {
	for {
		log.Println("Очистка истекших токенов...")

		// Освобождаем память от истекших токенов
		err := databases.PostgreSQLDeleteExpiredSessions(dbc)

		if err != nil {
			if !errors.Is(err, databases.ErrSessionsNotFoundExpired) {
				log.Println(err)
			}
		}

		databases.PostgreSQLCleanAccessTokens(dbc)

		log.Println("Таблица токенов очищена!")

		// Ждем пять минут
		time.Sleep(time.Minute * 5)
	}
}

// GetCurrentUserEmail - получает текущую почту пользователя
func GetCurrentUserEmail(w http.ResponseWriter, req *http.Request) (Email string) {

	// Освобождаем память от истекших токенов
	err := databases.PostgreSQLDeleteExpiredSessions(setup.ServerSettings.SQL.ConnPool)

	if err != nil {
		log.Println(err)
	}

	result := GetEmailBasedOnCookies(w, req)

	if len(result) > 0 {
		return result
	}

	return GetEmailBasedOnToken(req)

}

// GetEmailBasedOnCookies - получает электронную почту из списка сессий по куки
func GetEmailBasedOnCookies(w http.ResponseWriter, req *http.Request) (Email string) {

	mycookies, err := securecookies.GetCookies(w, req)

	if err != nil {

		if !errors.Is(http.ErrNoCookie, err) {
			log.Println(err)
			return ""
		}
	}

	if len(mycookies.Email) > 0 && len(mycookies.Session) > 0 {

		// Ищем живые токены по сессии
		at, found := SearchIssuedSessions(mycookies.Email, mycookies.Session)

		if found {
			return at.Email
		}

		log.Println(securecookies.ErrAuthCookiesNotFound)
		return ""
	}

	return ""
}

// GetEmailBasedOnToken - получает почту пользователя из списка сессий
func GetEmailBasedOnToken(req *http.Request) (Email string) {

	Token := req.Header.Get("Auth")

	if len(Token) > 0 {

		at, err := databases.PostgreSQLGetActiveTokenByToken(Token, setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			log.Println(err)
			return ""
		} else {
			return at.Email
		}

	}

	return ""
}

// GetCurrentSession - получаем текущую сессию пользователя
func GetCurrentSession(w http.ResponseWriter, req *http.Request) (authentication.ActiveToken, error) {
	// Освобождаем память от истекших токенов
	err := databases.PostgreSQLDeleteExpiredSessions(setup.ServerSettings.SQL.ConnPool)

	if err != nil {
		if !errors.Is(err, databases.ErrSessionsNotFoundExpired) {
			log.Println(err)
		}
	}

	result, err := GetTokenBasedOnCookies(w, req)

	if err == nil {
		return result, err
	}

	return GetTokenBasedOnToken(req)

}

// GetTokenBasedOnCookies - получает текущую сессию на основе куки
func GetTokenBasedOnCookies(w http.ResponseWriter, req *http.Request) (authentication.ActiveToken, error) {

	result := authentication.ActiveToken{}

	mycookies, err := securecookies.GetCookies(w, req)

	if err != nil {

		if !errors.Is(http.ErrNoCookie, err) {
			log.Println(err)
			return result, err
		}
	}

	if len(mycookies.Email) > 0 && len(mycookies.Session) > 0 {

		// Ищем живые токены по сессии
		result, found := SearchIssuedSessions(mycookies.Email, mycookies.Session)

		if found {
			return result, nil
		}

		return result, securecookies.ErrAuthCookiesNotFound
	}

	return result, securecookies.ErrAuthCookiesNotFound
}

// GetTokenBasedOnToken - получает текущую сессию по токену
func GetTokenBasedOnToken(req *http.Request) (authentication.ActiveToken, error) {

	result := authentication.ActiveToken{}

	Token := req.Header.Get("Auth")

	if len(Token) > 0 {

		return databases.PostgreSQLGetActiveTokenByToken(Token, setup.ServerSettings.SQL.ConnPool)

	}

	return result, ErrSessionNotFoundByToken
}
