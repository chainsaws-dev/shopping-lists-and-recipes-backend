package signinupout

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"myprojects/Shopping-lists-and-recipes/packages/authentication"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"net/http"
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

			// Разбираем зашифрованные base64 логин и пароль
			resbytelog, err := base64.StdEncoding.DecodeString(AuthRequest.Email)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			AuthRequest.Email = string(resbytelog)

			resbytepas, err := base64.StdEncoding.DecodeString(AuthRequest.Password)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			AuthRequest.Password = string(resbytepas)

			// Авторизация под ограниченной ролью
			err = setup.ServerSettings.SQL.Connect("guest_role_read_only")

			if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
				return
			}
			defer setup.ServerSettings.SQL.Disconnect()

			if len(AuthRequest.Password) < 6 {
				shared.HandleOtherError(w, "Пароль должен быть длиной свыше шести символов", ErrNotAllowedMethod, http.StatusBadRequest)
				return
			}

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

				tokenb, err := authentication.GenerateRandomBytes(32)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				var AuthResponse authentication.AuthResponseData
				AuthResponse.Token = hex.EncodeToString(tokenb)
				AuthResponse.Email = base64.StdEncoding.EncodeToString([]byte(AuthRequest.Email))
				AuthResponse.ExpiresIn = 3600
				AuthResponse.Registered = true
				AuthResponse.Role = base64.StdEncoding.EncodeToString([]byte(strrole))

				tb := time.Now()
				te := tb.Add(3600 * time.Second)

				TokenList = append(TokenList, authentication.ActiveToken{
					Token:   AuthResponse.Token,
					IssDate: tb,
					ExpDate: te,
					Role:    strrole,
				})

				CleanOldTokens()

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

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
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

// CheckTokenIssued - проверяет что токен есть в списке и не протух
func CheckTokenIssued(Token string) (bool, string) {
	CleanOldTokens()

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

// SignUp - обработчик для регистрации пользователя
func SignUp(w http.ResponseWriter, req *http.Request) {
	// TODO
	shared.HandleOtherError(w, "Method is not implemented", ErrNotAllowedMethod, http.StatusNotImplemented)
}
