// Package signinupout - отвечает за авторизацию, регистрацию, админку и функции для восстановления пароля и подтверждение почты
package signinupout

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"shopping-lists-and-recipes/packages/admin"
	"shopping-lists-and-recipes/packages/authentication"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/messages"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

// TokenList - список активных токенов
var TokenList Sessions

// SignIn - обработчик для авторизации пользователя POST запросом
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
// 	в теле запроса JSON объект AuthRequestData
//	Email и пароль должны быть пропущены через через encodeURIComponent и btoa
func SignIn(w http.ResponseWriter, req *http.Request) {

	if !AuthBasic(w, req) {
		return
	}

	switch {
	case req.Method == http.MethodPost:

		// Читаем тело запроса в структуру
		var AuthRequest authentication.AuthRequestData

		err := json.NewDecoder(req.Body).Decode(&AuthRequest)

		if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
			return
		}

		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		if !re.MatchString(AuthRequest.Email) {
			shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
			return
		}

		AuthRequest.Password, err = url.QueryUnescape(AuthRequest.Password)

		if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
			return
		}

		// Вызываем проверку пароля и выдачу токена
		secretauth(w, req, AuthRequest)

	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}

// SignUp - обработчик для регистрации пользователя POST запросом
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
// 	в теле запроса JSON объект AuthSignUpRequestData
//	Email, имя пользователя и пароль должны быть
//	пропущены через через encodeURIComponent и btoa
func SignUp(w http.ResponseWriter, req *http.Request) {

	if !AuthBasic(w, req) {
		return
	}

	switch {
	case req.Method == http.MethodPost:

		// Читаем тело запроса в структуру
		var SignUpRequest authentication.AuthSignUpRequestData

		err := json.NewDecoder(req.Body).Decode(&SignUpRequest)

		if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
			return
		}

		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		if !re.MatchString(SignUpRequest.Email) {
			shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
			return
		}

		SignUpRequest.Password, err = url.QueryUnescape(SignUpRequest.Password)

		if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
			return
		}

		SignUpRequest.Name, err = url.QueryUnescape(SignUpRequest.Name)

		if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
			return
		}

		// Подключаемся под ролью админа
		dbc := setup.ServerSettings.SQL.ConnectAsAdmin()
		if dbc == nil {
			return
		}

		// Создаём пользователя
		err = admin.CreateUser(&setup.ServerSettings.SQL, SignUpRequest.Name, SignUpRequest.Email, SignUpRequest.Password, setup.ServerSettings.SMTP.Use, dbc)

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
		messages.SendEmailConfirmationLetter(&setup.ServerSettings.SQL, SignUpRequest.Email, shared.CurrentPrefix+req.Host, dbc)

	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}

// ResendEmail - отправляет письмо подтверждение повторно в ответ на POST запрос
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
// 	ожидается заголовок Email с электронной почтой
func ResendEmail(w http.ResponseWriter, req *http.Request) {

	if !AuthBasic(w, req) {
		return
	}

	switch {
	case req.Method == http.MethodPost:

		Email := req.Header.Get("Email")

		if len(Email) > 0 {

			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !re.MatchString(Email) {
				shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
				return
			}

			// Подключаемся под ролью админа
			dbc := setup.ServerSettings.SQL.ConnectAsAdmin()
			if dbc == nil {
				return
			}

			mailexist, err := databases.PostgreSQLCheckUserMailExists(Email, dbc)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			if mailexist {
				messages.SendEmailConfirmationLetter(&setup.ServerSettings.SQL, Email, shared.CurrentPrefix+req.Host, dbc)

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

}

// ConfirmEmail - обработчик вызывающий отправку подтверждения почты, принимает POST запрос
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
//	ожидается заголовок Token с токеном для доступа
func ConfirmEmail(w http.ResponseWriter, req *http.Request) {

	if !AuthBasic(w, req) {
		return
	}

	switch {
	case req.Method == http.MethodPost:

		Token := req.Header.Get("Token")

		Token, err := url.QueryUnescape(Token)

		if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
			return
		}

		// Подключаемся под ролью админа
		dbc := setup.ServerSettings.SQL.ConnectAsAdmin()
		if dbc == nil {
			return
		}
		defer dbc.Close()

		// Если токен не протух и существует в базе записали подтверждение пользователя
		err = databases.PostgreSQLGetTokenConfirmEmail(Token, dbc)

		if err != nil {
			if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
				return
			}
		}

		shared.HandleSuccessMessage(w, "Электронная почта подтверждена")

	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}

// RequestResetEmail - отправляет письмо со ссылкой для сброса пароля
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
// 	ожидается заголовок Email с электронной почтой
func RequestResetEmail(w http.ResponseWriter, req *http.Request) {

	if !AuthBasic(w, req) {
		return
	}

	switch {
	case req.Method == http.MethodPost:

		Email := req.Header.Get("Email")

		if len(Email) > 0 {

			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !re.MatchString(Email) {
				shared.HandleOtherError(w, "Некорректная электронная почта", ErrBadEmail, http.StatusBadRequest)
				return
			}

			// Подключаемся под ролью админа
			dbc := setup.ServerSettings.SQL.ConnectAsAdmin()
			if dbc == nil {
				return
			}

			mailexist, err := databases.PostgreSQLCheckUserMailExists(Email, dbc)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			if mailexist {
				messages.SendEmailPasswordReset(&setup.ServerSettings.SQL, Email, shared.CurrentPrefix+req.Host, dbc)

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

}

// ResetPassword - сбрасывает пароль для заданного пользователя
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
//	ожидается заголовок Token с токеном для доступа
//  ожидается заголовок NewPassword c новым паролем
func ResetPassword(w http.ResponseWriter, req *http.Request) {

	if !AuthBasic(w, req) {
		return
	}

	switch {
	case req.Method == http.MethodPost:

		Token := req.Header.Get("Token")

		Token, err := url.QueryUnescape(Token)

		if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
			return
		}

		// Разбираем и декодируем зашифрованный base64 пароль
		NewPassword := req.Header.Get("NewPassword")
		NewPassword, err = url.QueryUnescape(NewPassword)

		// Подключаемся под ролью админа
		dbc := setup.ServerSettings.SQL.ConnectAsAdmin()
		if dbc == nil {
			return
		}
		defer dbc.Close()

		Hash, err := authentication.Argon2GenerateHash(NewPassword, &authentication.HashParams)

		if shared.HandleOtherError(w, "Ошибка при расчете хеша", err, http.StatusInternalServerError) {
			return
		}

		// Если токен не протух и существует в базе обновляем пароль пользователя
		err = databases.PostgreSQLGetTokenResetPassword(Token, Hash, dbc)

		if err != nil {
			if shared.HandleOtherError(w, err.Error(), err, http.StatusUnauthorized) {
				return
			}
		}

		shared.HandleSuccessMessage(w, "Пароль обновлён.")

	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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
//	и записан в заголовке NewPassword
//
// DELETE
//
// 	ожидается заголовок UserID с UUID пользователя пропущенным через encodeURIComponent и btoa (закодированным base64)
func HandleUsers(w http.ResponseWriter, req *http.Request) {

	role, auth := AuthGeneral(w, req)

	if !auth {
		return
	}

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "HandleUsers") {

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")

			var usersresp databases.UsersResponse

			// Авторизация под ролью пользователя
			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			if PageStr != "" && LimitStr != "" {

				Page, err := strconv.Atoi(PageStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				usersresp, err = databases.PostgreSQLUsersSelect(Page, Limit, dbc)

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

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:

		if setup.ServerSettings.CheckRoleForChange(role, "HandleUsers") {
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

			if !setup.ServerSettings.CheckExistingRole(User.Role) {
				shared.HandleOtherError(w, "Указана некорректная роль", ErrBadRole, http.StatusBadRequest)
				return
			}

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			// Значение по умолчанию для хеша и пароля
			Hash := ""
			UpdatePassword := false

			NewPassword := req.Header.Get("NewPassword")

			if len(NewPassword) > 0 {

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

			if len(NewPassword) == 0 && uuid.Equal(uuid.Nil, User.GUID) {
				shared.HandleOtherError(w, "Пароль нового пользователя должен быть задан", ErrPasswordTooShort, http.StatusBadRequest)
				return
			}

			// Получаем обновлённого юзера (если новый с GUID)
			User, err = databases.PostgreSQLUsersInsertUpdate(User, Hash, UpdatePassword, true, dbc)

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

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "HandleUsers") {
			// Удаление пользователя
			UserIDtoDelStr := req.Header.Get("UserID")

			if len(UserIDtoDelStr) > 0 {

				UserIDtoDelStr, err := url.QueryUnescape(UserIDtoDelStr)

				dbc := setup.ServerSettings.SQL.Connect(w, role)
				if dbc == nil {
					return
				}
				defer dbc.Close()

				UserID, err := uuid.FromString(UserIDtoDelStr)

				if shared.HandleOtherError(w, "Некорректный идентификатор пользователя", err, http.StatusBadRequest) {
					return
				}

				err = databases.PostgreSQLUsersDelete(UserID, dbc)

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

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}

// HandleSessions - обработчик для работы с сессиями, принимает http запросы GET, POST и DELETE
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
// DELETE
//
// 	ожидается заголовок Email для удаления сессий из списка сессий по электронной почте
//
//  или
//
//	ожидается заголовок Token для удаления сессий из списка сессий по токену
func HandleSessions(w http.ResponseWriter, req *http.Request) {

	role, auth := AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "HandleSessions") {

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")

			var sessionsresp SessionsResponse

			if PageStr != "" && LimitStr != "" {

				Page, err := strconv.Atoi(PageStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				sessionsresp, err = GetSessionsList(Page, Limit)

				if err != nil {
					if errors.Is(err, databases.ErrLimitOffsetInvalid) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}
				}

				shared.WriteObjectToJSON(w, sessionsresp)

			} else {
				shared.HandleOtherError(w, ErrHeadersNotFilled.Error(), ErrHeadersNotFilled, http.StatusBadRequest)
				return
			}

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "HandleSessions") {
			// Удаление сессии по электронной почте
			Email := req.Header.Get("Email")
			Token := req.Header.Get("Token")

			if len(Email) > 0 || len(Token) > 0 {

				if len(Email) > 0 {
					err = DeleteSessionByEmail(Email)

					if err != nil {
						if errors.Is(err, ErrSessionNotFoundByEmail) {
							shared.HandleOtherError(w, "Сессии не найдены, невозможно удалить", err, http.StatusBadRequest)
							return
						}
					}

					shared.HandleSuccessMessage(w, "Сессии удалены")
				}

				if len(Token) > 0 {
					err = DeleteSessionByToken(Token)

					if err != nil {
						if errors.Is(err, ErrSessionNotFoundByToken) {
							shared.HandleOtherError(w, "Сессия не найдена, невозможно удалить", err, http.StatusBadRequest)
							return
						}
					}

					shared.HandleSuccessMessage(w, "Сессия удалена")
				}

			} else {
				shared.HandleOtherError(w, "Bad request", ErrHeadersNotFilled, http.StatusBadRequest)
			}

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}

// CurrentUser - обработчик для получения и сохранения данных текущего пользователя
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
// 	идентичным по структуре UserDB при этом пароль
//	пропущен через через encodeURIComponent и btoa
//	и записан в заголовке NewPassword
func CurrentUser(w http.ResponseWriter, req *http.Request) {

	role, auth := AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "CurrentUser") {
			// Получаем данные текущего пользователя
			Email := GetCurrentUserEmail(w, req)

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
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

			shared.WriteObjectToJSON(w, FoundUser)
		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}
	case req.Method == http.MethodPost:
		if setup.ServerSettings.CheckRoleForChange(role, "CurrentUser") {
			// Создание и изменение пользователя
			var User databases.User

			err = json.NewDecoder(req.Body).Decode(&User)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			// Получаем данные текущего пользователя
			User.Email = GetCurrentUserEmail(w, req)

			// Подключаемся к базе данных
			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			FoundUser, err := databases.PostgreSQLGetUserByEmail(User.Email, dbc)

			if err != nil {
				if errors.Is(databases.ErrNoUserWithEmail, err) {
					shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
					return
				}
			}

			if shared.HandleInternalServerError(w, err) {
				return
			}

			User.GUID = FoundUser.GUID
			User.IsAdmin = FoundUser.IsAdmin
			User.Role = FoundUser.Role
			User.Confirmed = FoundUser.Confirmed

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

			if !setup.ServerSettings.CheckExistingRole(User.Role) {
				shared.HandleOtherError(w, "Указана некорректная роль", ErrBadRole, http.StatusBadRequest)
				return
			}

			// Значение по умолчанию для хеша и пароля
			Hash := ""
			UpdatePassword := false

			NewPassword := req.Header.Get("NewPassword")

			if len(NewPassword) > 0 {

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

			if len(NewPassword) == 0 && uuid.Equal(uuid.Nil, User.GUID) {
				shared.HandleOtherError(w, "Пароль нового пользователя должен быть задан", ErrPasswordTooShort, http.StatusBadRequest)
				return
			}

			// Получаем обновлённого юзера
			User, err = databases.PostgreSQLCurrentUserUpdate(User, Hash, UpdatePassword, dbc)

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

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}
	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
