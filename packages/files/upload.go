// Package files - содержит функции обработчики запросов для загрузки файлов
package files

import (
	"crypto/sha1"
	"encoding/json"
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

// UploadFile - обработчик для загрузки файлов POST запросом
//
// POST
//
// 	ожидается параметр key с API ключом
// 	тело запроса должно быть заполнено двоичными данными файла,
//	переданными через поле формы image
func UploadFile(w http.ResponseWriter, req *http.Request) {
	// Проверяем API ключ
	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		shared.HandleOtherError(w, ErrNoKeyInParams.Error(), ErrNoKeyInParams, http.StatusBadRequest)
		return
	}

	key := keys[0]

	_, found := shared.FindInStringSlice(setup.APIkeys, key)

	if found {
		// Проверка токена и получение роли
		issued, role := signinupout.CheckTokenIssued(*req)

		if issued {
			if req.Method == http.MethodPost {

				if role == "admin_role_CRUD" {

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

					js, err := json.Marshal(furesp)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					w.Header().Set("Content-Type", "application/json")
					w.Write(js)
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
