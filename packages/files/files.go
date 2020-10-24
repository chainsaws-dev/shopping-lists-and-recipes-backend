// Package files - содержит функции обработчики запросов для загрузки файлов
package files

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
	"strconv"
	"strings"
)

// Список типовых ошибок
var (
	ErrNotAllowedMethod      = errors.New("Запрошен недопустимый метод для файлов")
	ErrNoKeyInParams         = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams      = errors.New("API ключ не зарегистрирован")
	ErrNotAuthorized         = errors.New("Пройдите авторизацию")
	ErrForbidden             = errors.New("Доступ запрещён")
	ErrHeadersFetchNotFilled = errors.New("Не заполнены обязательные параметры запроса списка файлов: Page, Limit")
	ErrHeaderDeleteNotFilled = errors.New("Не заполнен обязательный параметр для удаления файла: FileID")
	ErroNoRowsInResult       = errors.New("Не найдено ни одной записи для удаления")
	ErrFkeyViolation         = errors.New("Нельзя удалять записи, на которые имеются ссылки")
	ErrUnsupportedFileType   = errors.New("Неподдерживаемый тип файла")
)

// HandleFiles - обрабатывает POST, GET и DELETE запросы для работы с файлами
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
// 	тело запроса должно быть заполнено двоичными данными файла,
//	переданными через поле формы file
//
// DELETE
//
// 	ожидается заголовок FileID с номером файла на удаление
func HandleFiles(w http.ResponseWriter, req *http.Request) {
	// Проверяем API ключ
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

				// Обработка получения списка файлов с поддержкой постраничных порций

				PageStr := req.Header.Get("Page")
				LimitStr := req.Header.Get("Limit")

				var FilesResponse databases.FilesResponse
				var err error

				err = setup.ServerSettings.SQL.Connect(role)

				if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
					return
				}
				defer setup.ServerSettings.SQL.Disconnect()

				if PageStr != "" && LimitStr != "" {
					Page, err := strconv.Atoi(PageStr)

					if shared.HandleBadRequestError(w, err) {
						return
					}

					Limit, err := strconv.Atoi(LimitStr)

					if shared.HandleBadRequestError(w, err) {
						return
					}

					FilesResponse, err = databases.PostgreSQLFilesSelect(Page, Limit)

					if shared.HandleInternalServerError(w, err) {
						return
					}

				} else {
					shared.HandleOtherError(w, ErrHeadersFetchNotFilled.Error(), ErrHeadersFetchNotFilled, http.StatusBadRequest)
					return
				}

				shared.WriteObjectToJSON(w, FilesResponse)

			case req.Method == http.MethodPost:

				// Обработка добавления нового файла в список файлов

				if setup.ServerSettings.CheckRoleForChange(role, "HandleFiles") {

					furesp, err := fileUpload(w, req, role)

					if err != nil {
						return
					}

					shared.WriteObjectToJSON(w, furesp)

				} else {
					shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
				}

			case req.Method == http.MethodDelete:

				// Обработка удаления файла по номеру в базе

				FileID := req.Header.Get("FileID")

				var err error

				err = setup.ServerSettings.SQL.Connect(role)

				if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
					return
				}
				defer setup.ServerSettings.SQL.Disconnect()

				if len(FileID) > 0 {

					ID, err := strconv.Atoi(FileID)

					if shared.HandleBadRequestError(w, err) {
						return
					}

					if ID > 1 {
						err = databases.PostgreSQLFileDelete(ID)
					}

					if err != nil {
						if errors.Is(databases.ErrFirstNotDelete, err) {
							shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
							return
						}

						if errors.Is(sql.ErrNoRows, err) {
							shared.HandleOtherError(w, ErroNoRowsInResult.Error(), ErroNoRowsInResult, http.StatusBadRequest)
							return
						}

						if strings.Contains(err.Error(), "violates foreign key constraint") {
							shared.HandleOtherError(w, ErrFkeyViolation.Error(), ErrFkeyViolation, http.StatusBadRequest)
							return
						}

						if shared.HandleInternalServerError(w, err) {
							return
						}
					}

					shared.HandleSuccessMessage(w, fmt.Sprintf("Файл с индексом %v удалён", FileID))

				} else {
					shared.HandleOtherError(w, ErrHeaderDeleteNotFilled.Error(), ErrHeaderDeleteNotFilled, http.StatusBadRequest)
					return
				}

			default:
				shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
			}

		} else {
			shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}
