package authenticationhandler

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"net/http"
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
var TokenList []AuthResponseData

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

			var AuthRequest AuthRequestData

			err := json.NewDecoder(req.Body).Decode(&AuthRequest)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// TODO
			// Роль для поиска должна назначаться аутентификацией
			err = setup.ServerSettings.SQL.Connect("admin_role_CRUD")

			if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
				return
			}
			defer setup.ServerSettings.SQL.Disconnect()

			if len(AuthRequest.Password) < 6 {
				shared.HandleOtherError(w, "Пароль должен быть длиной свыше шести символов", ErrNotAllowedMethod, http.StatusBadRequest)
				return
			}

			// Получаем хеш из базы данных
			strhash, err := databases.PostgreSQLGetTokenForUser(AuthRequest.Email)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			// Проверяем пароль против хеша
			match, err := Argon2ComparePasswordAndHash(AuthRequest.Password, strhash)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			if match {

				tokenb, err := GenerateRandomBytes(32)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				var AuthResponse AuthResponseData
				AuthResponse.Token = hex.EncodeToString(tokenb)
				AuthResponse.Email = AuthRequest.Email
				AuthResponse.ExpiresIn = 3600

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

// SignUp - обработчик для регистрации пользователя
func SignUp(w http.ResponseWriter, req *http.Request) {

}
