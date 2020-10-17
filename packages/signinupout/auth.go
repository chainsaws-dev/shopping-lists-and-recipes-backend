// Package signinupout - отвечает за авторизацию, регистрацию, админку и функции для восстановления пароля и подтверждение почты
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
	"shopping-lists-and-recipes/packages/messages"
	"shopping-lists-and-recipes/packages/securecookies"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"strconv"
	"time"

	"github.com/gorilla/securecookie"
	uuid "github.com/satori/go.uuid"

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
	ErrBadPhone         = errors.New("Указан некорректный телефонный номер")
	ErrBadRole          = errors.New("Указана некорректная роль")
	ErrHeadersNotFilled = errors.New("Не заполнены обязательные параметры запроса")
)

// TokenList - список активных токенов
var TokenList []authentication.ActiveToken

// SignIn - обработчик для авторизации пользователя POST запросом
//
// POST
//
// 	ожидается параметр key с API ключом
// 	в теле запроса JSON объект AuthRequestData
//	Email и пароль должны быть пропущены через через encodeURIComponent и btoa
func SignIn(w http.ResponseWriter, req *http.Request) {

	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		switch {
		case req.Method == http.MethodPost:

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

// SignUp - обработчик для регистрации пользователя POST запросом
//
// POST
//
// 	ожидается параметр key с API ключом
// 	в теле запроса JSON объект AuthSignUpRequestData
//	Email, имя пользователя и пароль должны быть
//	пропущены через через encodeURIComponent и btoa
func SignUp(w http.ResponseWriter, req *http.Request) {

	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {

		switch {
		case req.Method == http.MethodPost:

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

			// Создаём пользователя
			err = admin.CreateUser(&setup.ServerSettings.SQL, SignUpRequest.Name, SignUpRequest.Email, SignUpRequest.Password, setup.ServerSettings.SMTP.Use)

			if err != nil {

				if errors.Is(err, databases.ErrEmailIsOccupied) {
					shared.HandleOtherError(w, "Указанный адрес электронной почты уже занят", err, http.StatusInternalServerError)
					return
				}
			}

			if shared.HandleInternalServerError(w, err) {
				return
			}

			// Авторизация пользователя
			secretauth(w, req, ConvertToSignInRequest(SignUpRequest))

			// Отправляем письмо-подтверждение
			messages.SendEmailConfirmationLetter(&setup.ServerSettings.SQL, SignUpRequest.Email, shared.CurrentPrefix+req.Host)

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}

	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}

}

// ResendEmail - отправляет письмо подтверждение повторно в ответ на POST запрос
//
// POST
//
// 	ожидается параметр key с API ключом
// 	ожидается заголовок Email с электронной почтой
func ResendEmail(w http.ResponseWriter, req *http.Request) {

	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}
	if found {
		switch {
		case req.Method == http.MethodPost:

			Email := req.Header.Get("Email")

			if len(Email) > 0 {

				re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

				if !re.MatchString(Email) {
					shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
					return
				}

				// Авторизация под ролью пользователя
				err := setup.ServerSettings.SQL.Connect("admin_role_CRUD")

				if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
					return
				}
				defer setup.ServerSettings.SQL.Disconnect()

				mailexist, err := databases.PostgreSQLCheckUserMailExists(Email)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				if mailexist {
					messages.SendEmailConfirmationLetter(&setup.ServerSettings.SQL, Email, shared.CurrentPrefix+req.Host)

					shared.HandleSuccessMessage(w, "Письмо отправлено")

				} else {
					shared.HandleOtherError(w, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
				}

			} else {
				shared.HandleOtherError(w, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
			}

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}

// ConfirmEmail - обработчик вызывающий отправку подтверждения почты, принимает POST запрос
//
// POST
//
// 	ожидается параметр key с API ключом
//	ожидается заголовок Token с токеном для доступа
func ConfirmEmail(w http.ResponseWriter, req *http.Request) {
	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		switch {
		case req.Method == http.MethodPost:

			Token := req.Header.Get("Token")

			Token, err := url.QueryUnescape(Token)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Авторизация под ролью пользователя
			err = setup.ServerSettings.SQL.Connect("admin_role_CRUD")

			if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
				return
			}
			defer setup.ServerSettings.SQL.Disconnect()

			// Если токен не протух и существует в базе записали подтверждение пользователя
			err = databases.PostgreSQLGetTokenConfirmEmail(Token)

			if err != nil {
				if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
					return
				}
			}

			shared.HandleSuccessMessage(w, "Электронная почта подтверждена")

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}

// RequestResetEmail - отправляет письмо со ссылкой для сброса пароля
//
// POST
//
// 	ожидается параметр key с API ключом
// 	ожидается заголовок Email с электронной почтой
func RequestResetEmail(w http.ResponseWriter, req *http.Request) {
	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		switch {
		case req.Method == http.MethodPost:

			Email := req.Header.Get("Email")

			if len(Email) > 0 {

				re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

				if !re.MatchString(Email) {
					shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
					return
				}

				// Авторизация под ролью пользователя
				err := setup.ServerSettings.SQL.Connect("admin_role_CRUD")

				if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
					return
				}
				defer setup.ServerSettings.SQL.Disconnect()

				mailexist, err := databases.PostgreSQLCheckUserMailExists(Email)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				if mailexist {
					messages.SendEmailPasswordReset(&setup.ServerSettings.SQL, Email, shared.CurrentPrefix+req.Host)

					shared.HandleSuccessMessage(w, "Письмо отправлено")
				} else {
					shared.HandleOtherError(w, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
				}

			} else {
				shared.HandleOtherError(w, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
			}

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}

// ResetPassword - сбрасывает пароль для заданного пользователя
//
// POST
//
// 	ожидается параметр key с API ключом
//	ожидается заголовок Token с токеном для доступа
//  ожидается заголовок NewPassword c новым паролем
func ResetPassword(w http.ResponseWriter, req *http.Request) {
	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}
	if found {
		switch {
		case req.Method == http.MethodPost:

			Token := req.Header.Get("Token")

			Token, err := url.QueryUnescape(Token)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Разбираем и декодируем зашифрованный base64 пароль
			NewPassword := req.Header.Get("NewPassword")
			resbytepas, err := base64.StdEncoding.DecodeString(NewPassword)
			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}
			NewPassword = string(resbytepas)
			NewPassword, err = url.QueryUnescape(NewPassword)

			// Авторизация под ролью пользователя
			err = setup.ServerSettings.SQL.Connect("admin_role_CRUD")

			if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
				return
			}
			defer setup.ServerSettings.SQL.Disconnect()

			Hash, err := authentication.Argon2GenerateHash(NewPassword, &authentication.HashParams)

			if shared.HandleOtherError(w, "Ошибка при расчете хеша", err, http.StatusInternalServerError) {
				return
			}

			// Если токен не протух и существует в базе обновляем пароль пользователя
			err = databases.PostgreSQLGetTokenResetPassword(Token, Hash)

			if err != nil {
				if shared.HandleOtherError(w, err.Error(), err, http.StatusUnauthorized) {
					return
				}
			}

			shared.HandleSuccessMessage(w, "Пароль обновлён.")

		default:
			shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}

// HandleUsers - обработчик для работы с пользователями, принимает http запросы GET, POST и DELETE
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
// 	ожидается заголовок Page с номером страницы
// 	ожидается заголовок Limit с максимумом элементов на странице
//
// POST
//
// 	тело запроса должно быть заполнено JSON объектом
// 	идентичным по структуре UserDB при этом пароль
//	пропущен через через encodeURIComponent и btoa
//
// DELETE
//
// 	ожидается заголовок UserID с UUID пользователя пропущенным через encodeURIComponent и btoa (закодированным base64)
func HandleUsers(w http.ResponseWriter, req *http.Request) {

	found, err := CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		// Проверка токена и получение роли
		issued, role := TwoWayAuthentication(w, req)

		if issued {

			if setup.ServerSettings.CheckRoleForRead(role, "HandleUsers") {
				switch {
				case req.Method == http.MethodGet:

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

						if err != nil {
							if errors.Is(err, databases.ErrLimitOffsetInvalid) {
								shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
								return
							}

							if shared.HandleInternalServerError(w, err) {
								return
							}
						}

						shared.WriteObjectToJSON(w, usersresp)

					} else {
						shared.HandleOtherError(w, ErrHeadersNotFilled.Error(), ErrHeadersNotFilled, http.StatusBadRequest)
						return
					}

				case req.Method == http.MethodPost:
					// Создание и изменение пользователя
					var User databases.User

					err := json.NewDecoder(req.Body).Decode(&User)

					if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
						return
					}

					remai := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

					if !remai.MatchString(User.Email) {
						shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
						return
					}

					if len(User.Phone) > 0 {

						repho := regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3,4}\)?[\- ]?)?[\d\- ]{5,10}$`)

						if !repho.MatchString(User.Phone) {
							shared.HandleOtherError(w, "Некорректный телефонный номер", ErrBadPhone, http.StatusBadRequest)
							return
						}
					}

					if User.Role != "guest_role_read_only" && User.Role != "admin_role_CRUD" {
						shared.HandleOtherError(w, "Указана некорректная роль", ErrBadRole, http.StatusBadRequest)
						return
					}

					err = setup.ServerSettings.SQL.Connect(role)

					if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
						return
					}
					defer setup.ServerSettings.SQL.Disconnect()

					// Значение по умолчанию для хеша и пароля
					Hash := ""
					UpdatePassword := false

					NewPassword := req.Header.Get("NewPassword")

					if len(NewPassword) > 0 {

						// Разбираем и декодируем зашифрованный base64 пароль
						resbytepas, err := base64.StdEncoding.DecodeString(NewPassword)

						if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
							return
						}

						NewPassword = string(resbytepas)
						NewPassword, err = url.QueryUnescape(NewPassword)

						if len(NewPassword) < 6 {
							shared.HandleOtherError(w, "Пароль должен быть более шести символов", ErrPasswordTooShort, http.StatusBadRequest)
							return
						}
						Hash, err = authentication.Argon2GenerateHash(NewPassword, &authentication.HashParams)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						UpdatePassword = true
					}
					// Получаем обновлённого юзера (если новый с GUID)
					User, err = databases.PostgreSQLUsersInsertUpdate(User, Hash, UpdatePassword, true)

					if err != nil {
						if errors.Is(err, databases.ErrEmailIsOccupied) {
							shared.HandleOtherError(w, "Указанный адрес электронной почты уже занят", err, http.StatusInternalServerError)
							return
						}
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}

					// Пишем в тело ответа
					shared.WriteObjectToJSON(w, User)

				case req.Method == http.MethodDelete:
					// Удаление пользователя
					UserIDtoDelStr := req.Header.Get("UserID")

					if len(UserIDtoDelStr) > 0 {

						// Разбираем и декодируем зашифрованный base64 идентификатор
						resbytepas, err := base64.StdEncoding.DecodeString(UserIDtoDelStr)

						if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
							return
						}

						UserIDtoDelStr = string(resbytepas)
						UserIDtoDelStr, err = url.QueryUnescape(UserIDtoDelStr)

						err = setup.ServerSettings.SQL.Connect(role)

						if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
							return
						}
						defer setup.ServerSettings.SQL.Disconnect()

						UserID, err := uuid.FromString(UserIDtoDelStr)

						if shared.HandleOtherError(w, "Некорректный идентификатор пользователя", err, http.StatusBadRequest) {
							return
						}

						err = databases.PostgreSQLUsersDelete(UserID)

						if err != nil {
							if errors.Is(err, databases.ErrUserNotFound) {
								shared.HandleOtherError(w, "Пользователь не найден, невозможно удалить", err, http.StatusBadRequest)
								return
							}
						}

						if shared.HandleInternalServerError(w, err) {
							return
						}

						shared.HandleSuccessMessage(w, "Запись удалена")

					} else {
						shared.HandleOtherError(w, "Bad request", ErrHeadersNotFilled, http.StatusBadRequest)
					}

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

		shared.WriteObjectToJSON(w, AuthResponse)

	} else {
		shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
	}
}
