package signinupout

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"shopping-lists-and-recipes/packages/admin"
	"shopping-lists-and-recipes/packages/authentication"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"strconv"
	"strings"
	"time"

	"encoding/base64"
)

// Список типовых ошибок
var (
	ErrNotAllowedMethod = errors.New("Запрошен недопустимый метод при авторизации")
	ErrNoKeyInParams    = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams = errors.New("API ключ не зарегистрирован")
	ErrPasswordTooShort = errors.New("Выбран слишком короткий пароль")
	ErrNotAuthorized    = errors.New("Неверный логин или пароль")
	ErrForbidden        = errors.New("Доступ запрещён")
	ErrBadEmail         = errors.New("Указана некорректная электронная почта")
	ErrHeadersNotFilled = errors.New("Не заполнены обязательные параметры запроса")
)

// TokenList - список активных токенов
var TokenList []authentication.ActiveToken

// SignIn - обработчик для авторизации пользователя
func SignIn(w http.ResponseWriter, req *http.Request) {

	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		shared.HandleOtherError(w, ErrNoKeyInParams.Error(), ErrNoKeyInParams, http.StatusBadRequest)
		return
	}

	key := keys[0]

	_, found := shared.FindInStringSlice(setup.APIkeys, key)

	if found {
		switch {
		case req.Method == http.MethodPost:

			w.Header().Set("Content-Type", "application/json")

			// Читаем тело запроса в структуру
			var AuthRequest authentication.AuthRequestData

			err := json.NewDecoder(req.Body).Decode(&AuthRequest)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Разбираем зашифрованные base64 логин
			resbytelog, err := base64.StdEncoding.DecodeString(AuthRequest.Email)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Проверяем против регулярного выражения, что это почта
			AuthRequest.Email = string(resbytelog)

			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !re.MatchString(AuthRequest.Email) {
				shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
				return
			}

			// Разбираем и декодируем зашифрованный base64 пароль
			resbytepas, err := base64.StdEncoding.DecodeString(AuthRequest.Password)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			AuthRequest.Password = string(resbytepas)
			AuthRequest.Password, err = url.QueryUnescape(AuthRequest.Password)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Вызываем проверку пароля и выдачу токена
			secretauth(w, req, AuthRequest)

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}

// SignUp - обработчик для регистрации пользователя
func SignUp(w http.ResponseWriter, req *http.Request) {

	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		shared.HandleOtherError(w, ErrNoKeyInParams.Error(), ErrNoKeyInParams, http.StatusBadRequest)
		return
	}

	key := keys[0]

	_, found := shared.FindInStringSlice(setup.APIkeys, key)

	if found {

		switch {
		case req.Method == http.MethodPost:
			w.Header().Set("Content-Type", "application/json")

			// Читаем тело запроса в структуру
			var SignUpRequest authentication.AuthSignUpRequestData

			err := json.NewDecoder(req.Body).Decode(&SignUpRequest)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Разбираем зашифрованные base64 логин
			resbytelog, err := base64.StdEncoding.DecodeString(SignUpRequest.Email)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Проверяем против регулярного выражения, что это почта
			SignUpRequest.Email = string(resbytelog)

			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !re.MatchString(SignUpRequest.Email) {
				shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
				return
			}

			// Разбираем и декодируем зашифрованный base64 пароль
			resbytepas, err := base64.StdEncoding.DecodeString(SignUpRequest.Password)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			SignUpRequest.Password = string(resbytepas)
			SignUpRequest.Password, err = url.QueryUnescape(SignUpRequest.Password)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Разбираем и декодируем зашифрованный base64 имя пользователя
			resbytenam, err := base64.StdEncoding.DecodeString(SignUpRequest.Name)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			SignUpRequest.Name = string(resbytenam)
			SignUpRequest.Name, err = url.QueryUnescape(SignUpRequest.Name)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			err = admin.CreateUser(&setup.ServerSettings.SQL, SignUpRequest.Name, SignUpRequest.Email, SignUpRequest.Password)

			if err != nil {
				if err.Error() == "Указанный адрес электронной почты уже занят" {
					shared.HandleOtherError(w, "Указанный адрес электронной почты уже занят", err, http.StatusInternalServerError)
					return
				}
			}

			if shared.HandleInternalServerError(w, err) {
				return
			}

			// Авторизация пользователя
			secretauth(w, req, ConvertToSignInRequest(SignUpRequest))

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}

	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}

}

// HandleUsers - обработчик для работы с пользователями
func HandleUsers(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		shared.HandleOtherError(w, ErrNoKeyInParams.Error(), ErrNoKeyInParams, http.StatusBadRequest)
		return
	}

	key := keys[0]

	_, found := shared.FindInStringSlice(setup.APIkeys, key)

	if found {
		// Проверка токена и получение роли
		issued, role := CheckTokenIssued(*req)

		if issued {

			if role == "admin_role_CRUD" {
				switch {
				case req.Method == http.MethodGet:

					w.Header().Set("Content-Type", "application/json")

					PageStr := req.Header.Get("Page")
					LimitStr := req.Header.Get("Limit")

					var usersresp databases.UsersResponse
					var err error

					// Авторизация под ролью пользователя
					err = setup.ServerSettings.SQL.Connect(role)

					if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
						return
					}
					defer setup.ServerSettings.SQL.Disconnect()

					if PageStr != "" && LimitStr != "" {

						Page, err := strconv.Atoi(PageStr)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						Limit, err := strconv.Atoi(LimitStr)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						usersresp, err = databases.PostgreSQLUsersSelect(Page, Limit)

					} else {
						shared.HandleOtherError(w, ErrHeadersNotFilled.Error(), ErrHeadersNotFilled, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}

					js, err := json.Marshal(usersresp)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					_, err = w.Write(js)

					if shared.HandleInternalServerError(w, err) {
						return
					}

				case req.Method == http.MethodPost:

				default:
					shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
				}
			} else {
				shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
			}
		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusUnauthorized)
		}
	}

}

// secretauth - внутренняя функция для проверки пароля и авторизации
func secretauth(w http.ResponseWriter, req *http.Request, AuthRequest authentication.AuthRequestData) {
	// Авторизация под ограниченной ролью
	err := setup.ServerSettings.SQL.Connect("guest_role_read_only")

	if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
		return
	}
	defer setup.ServerSettings.SQL.Disconnect()

	if len(AuthRequest.Password) < 6 {
		shared.HandleOtherError(w, "Пароль должен быть длиной свыше шести символов", ErrNotAllowedMethod, http.StatusBadRequest)
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

		CleanOldTokens()

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
		te := tb.Add(3600 * time.Second)

		TokenList = append(TokenList, authentication.ActiveToken{
			Email:     AuthRequest.Email,
			Token:     AuthResponse.Token,
			IssDate:   tb,
			ExpDate:   te,
			Role:      strrole,
			UserAgent: UserAgent,
			IP:        ClientIP,
		})

		if !AuthRequest.ReturnSecureToken {
			AuthResponse.Token = ""
		}

		js, err := json.Marshal(AuthResponse)

		if shared.HandleInternalServerError(w, err) {
			return
		}

		_, err = w.Write(js)

		if shared.HandleInternalServerError(w, err) {
			return
		}
	} else {
		shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
	}
}

// CleanOldTokens - удаляет старые токены из списка
// а также те, которые выданы заданному пользователю
// если указано "_" вместо почты, удаляются только старые
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

// SearchIssuedTokens - ищет уже выданные токены
func SearchIssuedTokens(Email string) (authentication.ActiveToken, bool) {
	if len(Email) != 0 {
		for _, t := range TokenList {

			ct := time.Now()

			if ct.Before(t.ExpDate) && t.Email == Email {
				return t, true
			}
		}
	}

	return authentication.ActiveToken{}, false
}

// SliceDelete - удаляет элемент из списка токенов
func SliceDelete(idx int) {
	l := len(TokenList)

	TokenList[idx] = TokenList[l-1]
	TokenList[l-1] = authentication.ActiveToken{}
	TokenList = TokenList[:l-1]
}

// CheckTokenIssued - проверяет что токен есть в списке и не протух
func CheckTokenIssued(req http.Request) (bool, string) {

	CleanOldTokens()

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
