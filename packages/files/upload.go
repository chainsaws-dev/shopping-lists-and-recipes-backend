// Package files - содержит различные функции для работы с загрузкой файлов
package files

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

// UploadFile - Обработчик для загрузки файлов
func UploadFile(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		log.Println("Начинаем получение файла...")
		var furesp FileUploadResponse

		f, fh, err := req.FormFile("image")

		if shared.HandleInternalServerError(w, err) {
			return
		}
		defer f.Close()

		// TODO
		// Роль для поиска должна назначаться аутентификацией
		err = setup.ServerSettings.SQL.Connect("admin_role_CRUD")

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
		shared.HandleOtherError(w, "Method is not allowed", errors.New("Запрошен недопустимый метод для рецептов"), http.StatusMethodNotAllowed)
	}
}
