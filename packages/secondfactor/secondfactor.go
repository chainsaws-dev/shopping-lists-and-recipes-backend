package secondfactor

import (
	"errors"
	"net/http"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
)

// SecondFactor - обработчик для работы с настройками двухфакторной авторизации, принимает http запросы GET, POST и DELETE
//
// Аутентификация
//
//  Куки
//  Session - шифрованная сессия
//	Email - шифрованный электронный адрес пользователя
//
//  или
//
//	Заголовки:
//  Auth - Токен доступа
//
//	и
//
//	ApiKey - Постоянный ключ доступа к API *
//
// GET
//
// Ничего не требуется
//
// POST
//
// 	тело запроса должно быть заполнено JSON объектом
// 	идентичным по структуре TOTPSecret
//
// DELETE
//
// 	ожидается заголовок UserID с UUID пользователя пропущенным через encodeURIComponent и btoa (закодированным base64)
func SecondFactor(w http.ResponseWriter, req *http.Request) {
	found, err := signinupout.CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		// Проверка токена и получение роли
		issued, role := signinupout.TwoWayAuthentication(w, req)

		if issued {

			switch {
			case req.Method == http.MethodGet:
				if setup.ServerSettings.CheckRoleForRead(role, "SecondFactor") {
				} else {
					shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
				}

			case req.Method == http.MethodPost:

				if setup.ServerSettings.CheckRoleForChange(role, "SecondFactor") {

				} else {
					shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
				}

			case req.Method == http.MethodDelete:

				if setup.ServerSettings.CheckRoleForDelete(role, "SecondFactor") {

				} else {
					shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
				}

			default:
				shared.HandleOtherError(w, "Method is not allowed", signinupout.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
			}

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusUnauthorized)
		}
	}
}

// GetQRCode - обработчик возвращающий png QR кода для привязки приложения временного токена
//
// Аутентификация
//
//  Куки
//  Session - шифрованная сессия
//	Email - шифрованный электронный адрес пользователя
//
//  или
//
//	Заголовки:
//  Auth - Токен доступа
//
//	и
//
//	ApiKey - Постоянный ключ доступа к API *
//
// GET
//
// Ничего не требуется
//
func GetQRCode(w http.ResponseWriter, req *http.Request) {
	found, err := signinupout.CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		// Проверка токена и получение роли
		issued, role := signinupout.TwoWayAuthentication(w, req)

		if issued {

			switch {
			case req.Method == http.MethodGet:
				if setup.ServerSettings.CheckRoleForRead(role, "GetQRCode") {

					// Получаем данные текущего пользователя

					Email := signinupout.GetCurrentUserEmail(w, req)

					err := setup.ServerSettings.SQL.Connect(role)

					if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
						return
					}
					defer setup.ServerSettings.SQL.Disconnect()

					FoundUser, err := databases.PostgreSQLGetUserByEmail(Email)

					if err != nil {
						if errors.Is(databases.ErrNoUserWithEmail, err) {
							shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
							return
						}
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}

					var usf UserSecondFactor

					usf.User = FoundUser
					usf.URL = shared.CurrentPrefix + req.Host

					b, err := usf.GetQR(200, 200)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					shared.WriteBufferToPNG(w, b)

				} else {
					shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
				}

			default:
				shared.HandleOtherError(w, "Method is not allowed", signinupout.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
			}

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusUnauthorized)
		}
	}
}

// CheckSecondFactor - проверяет второй фактор для авторизации
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
//  TODO
func CheckSecondFactor(w http.ResponseWriter, req *http.Request) {
	found, err := signinupout.CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		// Проверка токена и получение роли
		issued, role := signinupout.TwoWayAuthentication(w, req)

		if issued {

			switch {
			case req.Method == http.MethodPost:
				if setup.ServerSettings.CheckRoleForRead(role, "CheckSecondFactor") {

					// TODO

				} else {
					shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
				}

			default:
				shared.HandleOtherError(w, "Method is not allowed", signinupout.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
			}

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", signinupout.ErrWrongKeyInParams, http.StatusBadRequest)
	}
}
