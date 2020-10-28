package signinupout

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"math"
	"net/http"
	"time"

	"shopping-lists-and-recipes/packages/authentication"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/securecookies"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"

	"github.com/gorilla/securecookie"
)

// secretauth - внутренняя функция для проверки пароля и авторизации
// (если ReturnToken=false - то куки)
func secretauth(w http.ResponseWriter, req *http.Request, AuthRequest authentication.AuthRequestData) {
	// Авторизация под ограниченной ролью
	err := setup.ServerSettings.SQL.Connect("guest_role_read_only")

	if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
		return
	}
	defer setup.ServerSettings.SQL.Disconnect()

	if len(AuthRequest.Password) < 6 {
		shared.HandleOtherError(w, "Пароль должен быть более шести символов", ErrPasswordTooShort, http.StatusBadRequest)
		return
	}

	UserAgent := req.Header.Get("User-Agent")
	ClientIP := GetIP(req)

	// Получаем хеш из базы данных
	strhash, strrole, err := databases.PostgreSQLGetTokenForUser(AuthRequest.Email)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusTeapot) {
			return
		}
	}

	// Проверяем пароль против хеша
	match, err := authentication.Argon2ComparePasswordAndHash(AuthRequest.Password, strhash)

	if shared.HandleInternalServerError(w, err) {
		return
	}

	if match {

		// Удаляем просроченые токены
		CleanOldTokens()

		if CountTokensByEmail(AuthRequest.Email) > 1 {
			// Удаляем остальные токены если число сессий превышает 2
			DeleteSessionByEmail(AuthRequest.Email)
		}

		var AuthResponse authentication.AuthResponseData

		tokenb, err := authentication.GenerateRandomBytes(32)

		if shared.HandleInternalServerError(w, err) {
			return
		}

		AuthResponse = authentication.AuthResponseData{
			Token:      hex.EncodeToString(tokenb),
			Email:      base64.StdEncoding.EncodeToString([]byte(AuthRequest.Email)),
			ExpiresIn:  3600,
			Registered: true,
			Role:       base64.StdEncoding.EncodeToString([]byte(strrole)),
		}

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
		}

		TokenList = append(TokenList, NewActiveToken)

		// Если не возвращаем токен, то пишем куки
		if !AuthRequest.ReturnSecureToken {
			AuthResponse.Token = ""

			err = securecookies.SetCookies(te, NewActiveToken, w)

			if shared.HandleInternalServerError(w, err) {
				return
			}
		}

		FoundUser, err := databases.PostgreSQLGetUserByEmail(AuthRequest.Email)

		if err != nil {
			if errors.Is(databases.ErrNoUserWithEmail, err) {
				shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
			}
		}

		if shared.HandleInternalServerError(w, err) {
			return
		}

		AuthResponse.SecondFactor = FoundUser.SecondFactor

		shared.WriteObjectToJSON(w, AuthResponse)

	} else {
		shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
	}
}

// CheckAPIKey - проверяет API ключ
func CheckAPIKey(w http.ResponseWriter, req *http.Request) (bool, error) {

	APIKey := req.Header.Get("ApiKey")

	if len(APIKey) < 1 {
		return false, ErrNoKeyInParams
	}

	_, found := shared.FindInStringSlice(setup.APIkeys, APIKey)

	if !found {
		return false, ErrWrongKeyInParams
	}

	return true, nil

}

// TwoWayAuthentication - выполняет аутентификацию как с помощью заголовка Auth, так и с помощью куки
func TwoWayAuthentication(w http.ResponseWriter, req *http.Request) (issued bool, role string) {
	// Освобождаем память от истекших токенов
	CleanOldTokens()

	// Проверка кук и получение роли
	cookiefound, role := CheckCookiesIssued(w, req)

	if cookiefound {
		return cookiefound, role
	}

	// Проверка токена и получение роли (если нет кук)
	return CheckTokensIssued(req)

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
		for _, t := range TokenList {

			ct := time.Now()

			if ct.Before(t.ExpDate) && t.Token == Token {
				return true, t.Role
			}
		}
	}

	return false, ""
}

// SearchIssuedSessions - ищет активный токен по электронной почте и с совпадающей сессией
func SearchIssuedSessions(Email string, Session []byte) (authentication.ActiveToken, bool) {
	if len(Email) != 0 {
		for _, t := range TokenList {

			ct := time.Now()

			if ct.Before(t.ExpDate) && t.Email == Email && CompareSessions(t, Session) {
				return t, true
			}
		}
	}

	return authentication.ActiveToken{}, false
}

// GetSessionsList - получает список сессий в постраничной разбивке
func GetSessionsList(page int, limit int) (SessionsResponse, error) {

	var result SessionsResponse

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	result.Total = len(TokenList)
	result.Limit = limit
	result.Offset = offset

	if databases.PostgreSQLCheckLimitOffset(limit, offset) &&
		result.Total > result.Offset {

		if offset+limit >= result.Total {
			result.Sessions = TokenList[offset:]
		} else {
			result.Sessions = TokenList[offset : offset+limit]
		}

	} else {
		return result, ErrLimitOffsetInvalid
	}

	return result, nil

}

// DeleteSessionByEmail - ищет сессию по электронной почте и удаляет её
func DeleteSessionByEmail(Email string) error {

	var idx int
	var found bool

	for idx >= 0 {

		idx = FindSessionIdxByEmail(Email)
		if idx >= 0 {
			found = true
			SliceDelete(idx)
		}
	}

	if !found {
		return ErrSessionNotFoundByEmail
	}

	return nil
}

// DeleteSessionByToken - ищет сессию по токену и удаляет её
func DeleteSessionByToken(Token string) error {
	var idx int
	var found bool

	for idx >= 0 {

		idx = FindSessionIdxByToken(Token)
		if idx >= 0 {
			found = true
			SliceDelete(idx)
		}
	}

	if !found {
		return ErrSessionNotFoundByToken
	}

	return nil
}

// FindSessionIdxByEmail - ищет сессию по электронному адресу и возвращает индекс
func FindSessionIdxByEmail(Email string) int {

	for idx, session := range TokenList {
		if session.Email == Email {
			return idx
		}
	}

	return -1
}

// FindSessionIdxByToken - ищет сессию по токену и возвращает индекс
func FindSessionIdxByToken(Token string) int {

	for idx, session := range TokenList {
		if session.Token == Token {
			return idx
		}
	}

	return -1
}

// FindSessionIdxExpired - ищет первую попавшуюся истёкшую сессию и возвращает её индекс
func FindSessionIdxExpired() int {

	ct := time.Now()

	for idx, session := range TokenList {

		if ct.After(session.ExpDate) {
			return idx
		}
	}

	return -1
}

// CompareSessions - выполняет сравнение сессий
func CompareSessions(at authentication.ActiveToken, SessionToCompare []byte) bool {
	res := bytes.Compare(at.Session, SessionToCompare)

	if res == 0 {
		return true
	}

	return false
}

// CleanOldTokens - удаляет старые токены из списка
func CleanOldTokens() error {

	var idx int

	for idx >= 0 {

		idx = FindSessionIdxExpired()

		if idx >= 0 {
			SliceDelete(idx)
		}
	}

	return nil

}

// CountTokensByEmail - cчитает количество токенов с одним Email в списке сессий
func CountTokensByEmail(Email string) int {

	var result int

	for _, t := range TokenList {

		if t.Email == Email {
			result++
		}

	}

	return result
}

// SliceDelete - удаляет элемент из списка токенов
func SliceDelete(idx int) {
	l := len(TokenList)

	TokenList[idx] = TokenList[l-1]
	TokenList[l-1] = authentication.ActiveToken{}
	TokenList = TokenList[:l-1]
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
func RegularConfirmTokensCleanup() {
	for {
		log.Println("Очистка истекших токенов...")

		// Освобождаем память от истекших токенов
		CleanOldTokens()

		err := setup.ServerSettings.SQL.Connect("admin_role_CRUD")

		if err != nil {
			log.Fatalln(err)
		}

		defer setup.ServerSettings.SQL.Disconnect()

		databases.PostgreSQLCleanAccessTokens()

		log.Println("Таблица токенов очищена!")

		// Ждем пять минут
		time.Sleep(time.Minute * 5)
	}
}

// GetCurrentUserEmail - получает текущую почту пользователя
func GetCurrentUserEmail(w http.ResponseWriter, req *http.Request) (Email string) {

	// Освобождаем память от истекших токенов
	CleanOldTokens()

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
		for _, t := range TokenList {

			ct := time.Now()

			if ct.Before(t.ExpDate) && t.Token == Token {
				return t.Email
			}
		}
	}

	return ""
}
