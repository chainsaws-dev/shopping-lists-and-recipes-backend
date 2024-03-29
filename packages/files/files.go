// Package files - содержит функции обработчики запросов для загрузки файлов
package files

import (
	"database/sql"
	"errors"
	"net/http"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/internal/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
	"strconv"
	"strings"
)

// Список типовых ошибок
var (
	ErrHeaderDeleteNotFilled = errors.New("http request does not contain required parameter: FileID")
	ErrUnsupportedFileType   = errors.New("unsupported file type")
)

var (
	MsgFileDeleted = "file deleted"
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

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	switch {
	case req.Method == http.MethodGet:

		// Обработка получения списка файлов с поддержкой постраничных порций

		PageStr := req.Header.Get("Page")
		LimitStr := req.Header.Get("Limit")

		var FilesResponse databases.FilesResponse

		if PageStr != "" && LimitStr != "" {
			Page, err := strconv.Atoi(PageStr)

			if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			Limit, err := strconv.Atoi(LimitStr)

			if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			FilesResponse, err = databases.PostgreSQLFilesSelect(Page, Limit, setup.ServerSettings.SQL.ConnPool)

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrHeadersFetchNotFilled.Error(), shared.ErrHeadersFetchNotFilled, http.StatusBadRequest)
			return
		}

		shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, FilesResponse)

	case req.Method == http.MethodPost:

		// Обработка добавления нового файла в список файлов

		if setup.ServerSettings.CheckRoleForChange(role, "HandleFiles") {

			furesp, err := fileUpload(w, req, role)

			if err != nil {
				return
			}

			shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, furesp)

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		// Обработка удаления файла по номеру в базе

		FileID := req.Header.Get("FileID")

		if len(FileID) > 0 {

			ID, err := strconv.Atoi(FileID)

			if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			if ID > 1 {
				err = databases.PostgreSQLFileDelete(ID, setup.ServerSettings.SQL.ConnPool)
			}

			if err != nil {
				if errors.Is(databases.ErrFirstNotDelete, err) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
					return
				}

				if errors.Is(sql.ErrNoRows, err) {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErroNoRowsInResult.Error(), shared.ErroNoRowsInResult, http.StatusBadRequest)
					return
				}

				if strings.Contains(err.Error(), "violates foreign key constraint") {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrFkeyViolation.Error(), shared.ErrFkeyViolation, http.StatusBadRequest)
					return
				}

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}
			}

			shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MsgFileDeleted)

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrHeaderDeleteNotFilled.Error(), ErrHeaderDeleteNotFilled, http.StatusBadRequest)
			return
		}

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
