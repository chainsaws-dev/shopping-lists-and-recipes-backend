// Package files - содержит функции обработчики запросов для загрузки файлов
package files

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
	"strings"
)

// Список типовых ошибок
var (
	ErrNotAllowedMethod = errors.New("Запрошен недопустимый метод для файлов")
	ErrNoKeyInParams    = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams = errors.New("API ключ не зарегистрирован")
	ErrNotAuthorized    = errors.New("Пройдите авторизацию")
	ErrForbidden        = errors.New("Доступ запрещён")
)

// FileUploadResponse - тип для ответа на запрос
type FileUploadResponse struct {
	FileName string
	FileSize int64
	FileType string
	DbID     int
	FileID   string
	Error    string
}

// HandleFiles - обработчик для загрузки файлов POST запросом
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
// POST
//
// 	тело запроса должно быть заполнено двоичными данными файла,
//	переданными через поле формы image
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
			if req.Method == http.MethodPost {

				if setup.ServerSettings.CheckRoleForChange(role, "HandleFiles") {

					log.Println("Начинаем получение файла...")
					var furesp FileUploadResponse

					f, fh, err := req.FormFile("image")

					if shared.HandleInternalServerError(w, err) {
						return
					}
					defer f.Close()

					err = setup.ServerSettings.SQL.Connect(role)

					if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
						return
					}
					defer setup.ServerSettings.SQL.Disconnect()

					// Проверяем тип файла
					buff := make([]byte, 512)
					_, err = f.Read(buff)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					filetype := http.DetectContentType(buff)

					if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/gif" ||
						filetype == "image/png" || filetype == "application/pdf" {

						ext := strings.Split(fh.Filename, ".")[1]

						fn := sha1.New()

						io.Copy(fn, f)

						filename := fmt.Sprintf("%x", fn.Sum(nil)) + "." + ext

						filename = strings.Join([]string{"uploads", filename}, "/")

						path := filepath.Join(".", "public", "uploads", filename)

						nf, err := os.Create(path)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						defer nf.Close()

						_, err = f.Seek(0, 0)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						_, err = io.Copy(nf, f)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						log.Printf("Файл получен и сохранён под именем %s", filename)

						fileid, err := databases.PostgreSQLFileInsert(fh.Filename, fh.Size, ext, filename)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						furesp = FileUploadResponse{
							FileName: fh.Filename,
							FileID:   filename,
							FileType: ext,
							DbID:     fileid,
							FileSize: fh.Size,
							Error:    "",
						}

					} else {
						furesp = FileUploadResponse{
							FileName: fh.Filename,
							FileID:   "",
							FileType: "",
							DbID:     -1,
							FileSize: fh.Size,
							Error:    "Unsupported file type",
						}
					}

					shared.WriteObjectToJSON(w, furesp)

				} else {
					shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
				}

			} else {
				shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
			}
		} else {
			shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}
