package secondfactor

import (
	"encoding/json"
	"errors"
	"net/http"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
)

// Список типовых ошибок
var (
	ErrAlreadySetSecondFactor = errors.New("Двухфакторная авторизация уже настроена")
	ErrSecondFactorInactive   = errors.New("Двухфакторная авторизация неактивна")
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
//	ожидается заголовок Passcode с ключом с токена
// 	тело запроса должно быть заполнено JSON объектом
// 	идентичным по структуре User
//
//
// DELETE
//
// 	Ничего не требуется
func SecondFactor(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:
		if setup.ServerSettings.CheckRoleForRead(role, "SecondFactor") {

			var result databases.TOTPResponse

			// Получаем данные текущего пользователя

			Email := signinupout.GetCurrentUserEmail(w, req)

			// Подключение к базе данных
			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
				return
			}
			defer dbc.Close()

			FoundUser, err := databases.PostgreSQLGetUserByEmail(Email, dbc)

			if err != nil {
				if errors.Is(databases.ErrNoUserWithEmail, err) {
					shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
					return
				}
			}

			if shared.HandleInternalServerError(w, err) {
				return
			}

			Totp, err := databases.PostgreSQLGetSecretByUserID(FoundUser.GUID, dbc)

			if err != nil {
				if errors.Is(databases.ErrUserTOTPNotFound, err) {
					result.Confirmed = false
					result.UserID = FoundUser.GUID
				} else {
					if shared.HandleInternalServerError(w, err) {
						return
					}
				}
			} else {
				result.Confirmed = Totp.Confirmed
				result.UserID = Totp.UserID
			}

			shared.WriteObjectToJSON(w, result)

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
			return
		}

	case req.Method == http.MethodPost:

		if setup.ServerSettings.CheckRoleForChange(role, "SecondFactor") {

			PassStr := req.Header.Get("Passcode")

			if len(PassStr) > 0 {

				var CurUser databases.User

				err := json.NewDecoder(req.Body).Decode(&CurUser)

				if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
					return
				}

				// Подключение к базе данных
				dbc := setup.ServerSettings.SQL.Connect(w, role)
				if dbc == nil {
					shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
					return
				}
				defer dbc.Close()

				err = EnableTOTP(PassStr, CurUser, dbc)

				if err != nil {
					if errors.Is(ErrSecretNotSaved, err) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if errors.Is(databases.ErrUserTOTPNotFound, err) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}
				}

				shared.HandleSuccessMessage(w, "Второй фактор успешно настроен")

			} else {
				shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest)
				return
			}

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
			return
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "SecondFactor") {

			// Получаем данные текущего пользователя
			Email := signinupout.GetCurrentUserEmail(w, req)

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
				return
			}
			defer dbc.Close()

			FoundUser, err := databases.PostgreSQLGetUserByEmail(Email, dbc)

			if err != nil {
				if errors.Is(databases.ErrNoUserWithEmail, err) {
					shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
					return
				}
			}

			if shared.HandleInternalServerError(w, err) {
				return
			}

			err = databases.PostgreSQLDeleteSecondFactorSecret(FoundUser.GUID, dbc)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.HandleSuccessMessage(w, "Второй фактор успешно удалён")

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
			return
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", signinupout.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		return
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

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	switch {
	case req.Method == http.MethodGet:
		if setup.ServerSettings.CheckRoleForRead(role, "GetQRCode") {

			// Получаем данные текущего пользователя

			Email := signinupout.GetCurrentUserEmail(w, req)

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
				return
			}
			defer dbc.Close()

			FoundUser, err := databases.PostgreSQLGetUserByEmail(Email, dbc)

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

			b, err := usf.GetQR(200, 200, dbc)

			if err != nil {
				if errors.Is(databases.ErrTOTPConfirmed, err) {
					shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
					return
				}

				if shared.HandleInternalServerError(w, err) {
					return
				}
			}

			shared.WriteBufferToPNG(w, b)

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", signinupout.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}

// CheckSecondFactor - проверяет второй фактор для авторизации
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
//	ожидается заголовок Passcode с ключом с токена
func CheckSecondFactor(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthNoSecondFactor(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodPost:
		if setup.ServerSettings.CheckRoleForRead(role, "CheckSecondFactor") {

			PassStr := req.Header.Get("Passcode")

			if len(PassStr) > 0 {

				// Получаем данные текущего пользователя
				Email := signinupout.GetCurrentUserEmail(w, req)

				dbc := setup.ServerSettings.SQL.Connect(w, role)
				if dbc == nil {
					shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
					return
				}
				defer dbc.Close()

				FoundUser, err := databases.PostgreSQLGetUserByEmail(Email, dbc)

				if err != nil {
					if errors.Is(databases.ErrNoUserWithEmail, err) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}
				}

				if shared.HandleInternalServerError(w, err) {
					return
				}

				if !FoundUser.SecondFactor {
					shared.HandleOtherError(w, ErrSecondFactorInactive.Error(), ErrSecondFactorInactive, http.StatusBadRequest)
					return
				}

				Correct, err := Validate(PassStr, FoundUser, dbc)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				if Correct {
					at, err := signinupout.GetCurrentSession(w, req)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					at.SecondFactor.CheckResult = Correct

					signinupout.SetTokenStrict(at)

					shared.HandleSuccessMessage(w, "Двухфакторная авторизация успешно пройдена")
				} else {
					shared.HandleOtherError(w, "Указан неверный ключ", shared.ErrNotAuthorized, http.StatusUnauthorized)
				}

			} else {
				shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest)
				return
			}

		} else {
			shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", signinupout.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
