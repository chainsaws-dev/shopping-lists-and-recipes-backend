package signinupout

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"shopping-lists-and-recipes/packages/authentication"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/securecookies"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
)

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

		log.Println(securecookies.ErrPairNotFound)
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

// CompareSessions - выполняет сравнение сессий
func CompareSessions(at authentication.ActiveToken, SessionToCompare []byte) bool {
	res := bytes.Compare(at.Session, SessionToCompare)

	if res == 0 {
		return true
	}

	return false
}

// CleanOldTokens - удаляет старые токены из списка
func CleanOldTokens() {
	todel := []int{}

	for i, t := range TokenList {
		ct := time.Now()

		if ct.After(t.ExpDate) {
			todel = append(todel, i)
		}

	}

	for _, idx := range todel {
		SliceDelete(idx)
	}
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
	// Порт нас не интересует
	if idx := strings.IndexByte(IPAddress, ':'); idx >= 0 {
		IPAddress = IPAddress[:idx]
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
