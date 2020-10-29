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
	found, err := signinupout.CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		// Проверка токена и получение роли
		issued, role := signinupout.TwoWayAuthentication(w, req)

		// Проверка прохождения двухфакторной авторизации
		sf := signinupout.SecondFactorAuthenticationCheck(w, req)

		if issued {
			if sf {
				switch {
				case req.Method == http.MethodGet:
					if setup.ServerSettings.CheckRoleForRead(role, "SecondFactor") {

						// Получаем данные текущего пользователя

						Email := signinupout.GetCurrentUserEmail(w, req)

						// Подключение к базе данных
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

						Totp, err := databases.PostgreSQLGetSecretByUserID(FoundUser.GUID)

						if err != nil {
							if errors.Is(databases.ErrUserTOTPNotFound, err) {
								shared.HandleOtherError(w, err.Error(), err, http.StatusNotFound)
							}

							if shared.HandleInternalServerError(w, err) {
								return
							}
						}

						var result databases.TOTPResponse

						result.Confirmed = Totp.Confirmed
						result.UserID = Totp.UserID

						shared.WriteObjectToJSON(w, result)

					} else {
						shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
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
							err = setup.ServerSettings.SQL.Connect(role)

							if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
								return
							}
							defer setup.ServerSettings.SQL.Disconnect()

							err = EnableTOTP(PassStr, CurUser)

							if err != nil {
								if errors.Is(ErrSecretNotSaved, err) {
									shared.HandleOtherError(w, err.Error(), err, http.StatusUnauthorized)
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
					}

				case req.Method == http.MethodDelete:

					if setup.ServerSettings.CheckRoleForDelete(role, "SecondFactor") {

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

						err = databases.PostgreSQLDeleteSecondFactorSecret(FoundUser.GUID)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						shared.HandleSuccessMessage(w, "Второй фактор успешно удалён")

					} else {
						shared.HandleOtherError(w, signinupout.ErrForbidden.Error(), signinupout.ErrForbidden, http.StatusForbidden)
					}

				default:
					shared.HandleOtherError(w, "Method is not allowed", signinupout.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
				}
			} else {
				shared.HandleOtherError(w, shared.ErrNotAuthorizedTwoFactor.Error(), shared.ErrNotAuthorizedTwoFactor, http.StatusUnauthorized)
			}
		} else {
			shared.HandleOtherError(w, shared.ErrNotAuthorized.Error(), shared.ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", shared.ErrWrongKeyInParams, http.StatusBadRequest)
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

		// Проверка прохождения двухфакторной авторизации
		sf := signinupout.SecondFactorAuthenticationCheck(w, req)

		if issued {
			if sf {
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
			} else {
				shared.HandleOtherError(w, shared.ErrNotAuthorizedTwoFactor.Error(), shared.ErrNotAuthorizedTwoFactor, http.StatusUnauthorized)
			}
		} else {
			shared.HandleOtherError(w, shared.ErrNotAuthorized.Error(), shared.ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", shared.ErrWrongKeyInParams, http.StatusBadRequest)
	}
}

// CheckSecondFactor - проверяет второй фактор для авторизации
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
//	ожидается заголовок Passcode с ключом с токена
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

					PassStr := req.Header.Get("Passcode")

					if len(PassStr) > 0 {

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

						Correct, err := Validate(PassStr, FoundUser)

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

		} else {
			shared.HandleOtherError(w, shared.ErrNotAuthorized.Error(), shared.ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", shared.ErrWrongKeyInParams, http.StatusBadRequest)
	}
}
