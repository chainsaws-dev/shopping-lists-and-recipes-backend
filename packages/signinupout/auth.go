// Package signinupout - отвечает за авторизацию, регистрацию, админку и функции для восстановления пароля и подтверждение почты
package signinupout

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/internal/setup"
	"shopping-lists-and-recipes/packages/admin"
	"shopping-lists-and-recipes/packages/authentication"
	"shopping-lists-and-recipes/packages/messages"
	"shopping-lists-and-recipes/packages/shared"
	"strconv"

	uuid "github.com/gofrs/uuid"
)

// TokenList - список активных токенов
var TokenList Sessions

// SignIn - обработчик для авторизации пользователя POST запросом
//
// POST
//
// 	ожидается заголовок ApiKey с API ключом
//  ожидается заголовок Lang - Язык (ru или en)
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

		if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		if !re.MatchString(AuthRequest.Email) {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
			return
		}

		AuthRequest.Password, err = url.QueryUnescape(AuthRequest.Password)

		if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		// Вызываем проверку пароля и выдачу токена
		secretauth(w, req, AuthRequest)

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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

		if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		if !re.MatchString(SignUpRequest.Email) {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
			return
		}

		SignUpRequest.Password, err = url.QueryUnescape(SignUpRequest.Password)

		if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		SignUpRequest.Name, err = url.QueryUnescape(SignUpRequest.Name)

		if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		// Создаём пользователя
		err = admin.CreateUser(&setup.ServerSettings.SQL, SignUpRequest.Name, SignUpRequest.Email, SignUpRequest.Password, setup.ServerSettings.SMTP.Use, setup.ServerSettings.SQL.ConnPool)

		if err != nil {

			if errors.Is(err, databases.ErrEmailIsOccupied) || errors.Is(err, admin.ErrBasicFieldsNotFilled) {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
				return
			}

		}

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		// Авторизация пользователя
		secretauth(w, req, ConvertToSignInRequest(SignUpRequest))

		// Отправляем письмо-подтверждение
		messages.SendEmailConfirmationLetter(&setup.ServerSettings.SQL, SignUpRequest.Email, shared.CurrentPrefix+req.Host, setup.ServerSettings.SQL.ConnPool)

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
				return
			}

			mailexist, err := databases.PostgreSQLCheckUserMailExists(Email, setup.ServerSettings.SQL.ConnPool)

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			if mailexist {
				messages.SendEmailConfirmationLetter(&setup.ServerSettings.SQL, Email, shared.CurrentPrefix+req.Host, setup.ServerSettings.SQL.ConnPool)

				shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesEmailSent)

			} else {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
		}

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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

		if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		// Если токен не протух и существует в базе записали подтверждение пользователя
		err = databases.PostgreSQLGetTokenConfirmEmail(Token, setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			if shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest) {
				return
			}
		}

		shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesEmailConfirmed)

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
				return
			}

			mailexist, err := databases.PostgreSQLCheckUserMailExists(Email, setup.ServerSettings.SQL.ConnPool)

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			if mailexist {
				messages.SendEmailPasswordReset(&setup.ServerSettings.SQL, Email, shared.CurrentPrefix+req.Host, setup.ServerSettings.SQL.ConnPool)

				shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesEmailSent)
			} else {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
		}

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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

		if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
			return
		}

		// Разбираем и декодируем зашифрованный base64 пароль
		NewPassword := req.Header.Get("NewPassword")
		NewPassword, err = url.QueryUnescape(NewPassword)

		Hash, err := authentication.Argon2GenerateHash(NewPassword, &authentication.HashParams)

		if shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrPasswdHashCalc.Error(), err, http.StatusInternalServerError) {
			return
		}

		// Если токен не протух и существует в базе обновляем пароль пользователя
		err = databases.PostgreSQLGetTokenResetPassword(Token, Hash, setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			if shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusUnauthorized) {
				return
			}
		}

		shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesPasswdChanged)

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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

			if PageStr != "" && LimitStr != "" {

				Page, err := strconv.Atoi(PageStr)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				usersresp, err = databases.PostgreSQLUsersSelect(Page, Limit, setup.ServerSettings.SQL.ConnPool)

				if err != nil {
					if errors.Is(err, databases.ErrLimitOffsetInvalid) {
						shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
						return
					}
				}

				shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, usersresp)

			} else {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrHeadersNotFilled.Error(), shared.ErrHeadersNotFilled, http.StatusBadRequest)
				return
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:

		if setup.ServerSettings.CheckRoleForChange(role, "HandleUsers") {
			// Создание и изменение пользователя
			var User databases.User

			err := json.NewDecoder(req.Body).Decode(&User)

			if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			remai := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !remai.MatchString(User.Email) {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
				return
			}

			if len(User.Phone) > 0 {

				repho := regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3,4}\)?[\- ]?)?[\d\- ]{5,10}$`)

				if !repho.MatchString(User.Phone) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadPhone.Error(), ErrBadPhone, http.StatusBadRequest)
					return
				}
			}

			if !setup.ServerSettings.CheckExistingRole(User.Role) {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadRole.Error(), ErrBadRole, http.StatusBadRequest)
				return
			}

			// Значение по умолчанию для хеша и пароля
			Hash := ""
			UpdatePassword := false

			NewPassword := req.Header.Get("NewPassword")

			if len(NewPassword) > 0 {

				NewPassword, err = url.QueryUnescape(NewPassword)

				if len(NewPassword) < 6 {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrPasswordLength.Error(), ErrPasswordLength, http.StatusBadRequest)
					return
				}
				Hash, err = authentication.Argon2GenerateHash(NewPassword, &authentication.HashParams)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				UpdatePassword = true
			}

			if len(NewPassword) == 0 && uuid.Nil.String() == User.GUID.String() {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrPasswordNewUserNotFilled.Error(), ErrPasswordNewUserNotFilled, http.StatusBadRequest)
				return
			}

			// Получаем обновлённого юзера (если новый с GUID)
			User, err = databases.PostgreSQLUsersInsertUpdate(User, Hash, UpdatePassword, true, setup.ServerSettings.SQL.ConnPool)

			if err != nil {
				if errors.Is(err, databases.ErrEmailIsOccupied) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrEmailRegistered.Error(), err, http.StatusBadRequest)
					return
				}
			}

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			// Пишем в тело ответа
			shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, User)

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "HandleUsers") {
			// Удаление пользователя
			UserIDtoDelStr := req.Header.Get("UserID")

			if len(UserIDtoDelStr) > 0 {

				UserIDtoDelStr, err := url.QueryUnescape(UserIDtoDelStr)

				UserID, err := uuid.FromString(UserIDtoDelStr)

				if shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrIncorrectUserID.Error(), err, http.StatusBadRequest) {
					return
				}

				err = databases.PostgreSQLUsersDelete(UserID, setup.ServerSettings.SQL.ConnPool, setup.ServerSettings.Lang)

				if err != nil {
					if errors.Is(err, databases.ErrUserNotFound) {
						shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrUnableToDeleteAbsent.Error(), err, http.StatusBadRequest)
						return
					}
				}

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesUserDeleted)

			} else {
				shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, shared.ErrHeadersNotFilled)
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				sessionsresp, err = GetSessionsList(Page, Limit)

				if err != nil {
					if errors.Is(err, databases.ErrLimitOffsetInvalid) {
						shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
						return
					}
				}

				shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, sessionsresp)

			} else {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrHeadersNotFilled.Error(), shared.ErrHeadersNotFilled, http.StatusBadRequest)
				return
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
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
							shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrSessionNotFoundByEmail.Error(), err, http.StatusBadRequest)
							return
						}
					}

					shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesSessionsDeleted)
				}

				if len(Token) > 0 {
					err = DeleteSessionByToken(Token)

					if err != nil {
						if errors.Is(err, ErrSessionNotFoundByToken) {
							shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrSessionNotFoundByToken.Error(), err, http.StatusBadRequest)
							return
						}
					}

					shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesSessionDeleted)
				}

			} else {
				shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, shared.ErrHeadersNotFilled)
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
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

			FoundUser, err := databases.PostgreSQLGetUserByEmail(Email, setup.ServerSettings.SQL.ConnPool)

			if err != nil {
				if errors.Is(databases.ErrNoUserWithEmail, err) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
					return
				}
			}

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, FoundUser)
		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}
	case req.Method == http.MethodPost:
		if setup.ServerSettings.CheckRoleForChange(role, "CurrentUser") {
			// Создание и изменение пользователя
			var User databases.User

			err = json.NewDecoder(req.Body).Decode(&User)

			if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			// Получаем данные текущего пользователя
			User.Email = GetCurrentUserEmail(w, req)

			FoundUser, err := databases.PostgreSQLGetUserByEmail(User.Email, setup.ServerSettings.SQL.ConnPool)

			if err != nil {
				if errors.Is(databases.ErrNoUserWithEmail, err) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
					return
				}
			}

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			User.GUID = FoundUser.GUID
			User.IsAdmin = FoundUser.IsAdmin
			User.Role = FoundUser.Role
			User.Confirmed = FoundUser.Confirmed
			User.SecondFactor = FoundUser.SecondFactor

			remai := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !remai.MatchString(User.Email) {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadEmail.Error(), ErrBadEmail, http.StatusBadRequest)
				return
			}

			if len(User.Phone) > 0 {

				repho := regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3,4}\)?[\- ]?)?[\d\- ]{5,10}$`)

				if !repho.MatchString(User.Phone) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadPhone.Error(), ErrBadPhone, http.StatusBadRequest)
					return
				}
			}

			if !setup.ServerSettings.CheckExistingRole(User.Role) {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrBadRole.Error(), ErrBadRole, http.StatusBadRequest)
				return
			}

			// Значение по умолчанию для хеша и пароля
			Hash := ""
			UpdatePassword := false

			NewPassword := req.Header.Get("NewPassword")

			if len(NewPassword) > 0 {

				NewPassword, err = url.QueryUnescape(NewPassword)

				if len(NewPassword) < 6 {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrPasswordLength.Error(), ErrPasswordLength, http.StatusBadRequest)
					return
				}
				Hash, err = authentication.Argon2GenerateHash(NewPassword, &authentication.HashParams)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				UpdatePassword = true
			}

			if len(NewPassword) == 0 && uuid.Nil.String() == User.GUID.String() {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrPasswordNewUserNotFilled.Error(), ErrPasswordNewUserNotFilled, http.StatusBadRequest)
				return
			}

			// Получаем обновлённого юзера
			User, err = databases.PostgreSQLCurrentUserUpdate(User, Hash, UpdatePassword, setup.ServerSettings.SQL.ConnPool)

			if err != nil {
				if errors.Is(err, databases.ErrEmailIsOccupied) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrEmailRegistered.Error(), err, http.StatusInternalServerError)
					return
				}
			}

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			// Пишем в тело ответа
			shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, User)

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}
	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
