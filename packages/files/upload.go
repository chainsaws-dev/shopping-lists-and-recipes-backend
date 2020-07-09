// Package files - содержит различные функции для работы с загрузкой файлов
package files

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	DbID     int64
	FileID   string
	Error    string
}

// UploadFile - Обработчик для загрузки файлов
func UploadFile(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		log.Println("Начинаем получение файла...")
		var furesp FileUploadResponse

		f, fh, err := req.FormFile("image")

		if HandleInternalServerError(w, err) {
			return
		}
		defer f.Close()

		// Проверяем тип файла
		buff := make([]byte, 512)
		_, err = f.Read(buff)

		if HandleInternalServerError(w, err) {
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

			if HandleInternalServerError(w, err) {
				return
			}

			defer nf.Close()
			f.Seek(0, 0)

			io.Copy(nf, f)

			log.Printf("Файл получен и сохранён под именем %s", filename)

			furesp = FileUploadResponse{
				FileName: fh.Filename,
				FileID:   filename,
				FileType: ext,
				FileSize: fh.Size,
				Error:    "",
			}
		} else {
			furesp = FileUploadResponse{
				FileName: fh.Filename,
				FileID:   "",
				FileType: "",
				FileSize: fh.Size,
				Error:    "Unsupported file type",
			}
		}

		js, err := json.Marshal(furesp)

		if HandleInternalServerError(w, err) {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	} else {
		http.Error(w, "Request method is not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleInternalServerError - Обработчик внутренних ошибок сервера
func HandleInternalServerError(w http.ResponseWriter, err error) bool {

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
		return true
	}

	return false
}
