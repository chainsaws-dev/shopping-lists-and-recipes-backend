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
	"strings"
)

// fileUpload - выполняет загрузку файла на сервер и сохранение в файловой системе и информации в базе данных
func fileUpload(w http.ResponseWriter, req *http.Request, role string) (databases.File, error) {

	log.Println("Начинаем получение файла...")

	var NewFile databases.File

	f, fh, err := req.FormFile("file")

	if shared.HandleInternalServerError(w, err) {
		return NewFile, err
	}
	defer f.Close()

	err = setup.ServerSettings.SQL.Connect(role)

	if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
		return NewFile, err
	}
	defer setup.ServerSettings.SQL.Disconnect()

	// Проверяем тип файла
	buff := make([]byte, 512)
	_, err = f.Read(buff)

	if shared.HandleInternalServerError(w, err) {
		return NewFile, err
	}

	filetype := http.DetectContentType(buff)

	if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/gif" ||
		filetype == "image/png" || filetype == "application/pdf" {

		ext := strings.Split(fh.Filename, ".")[1]

		fn := sha1.New()

		io.Copy(fn, f)

		filename := fmt.Sprintf("%x", fn.Sum(nil)) + "." + ext

		linktofile := strings.Join([]string{"uploads", filename}, "/")

		path := filepath.Join(".", "public", "uploads", filename)

		nf, err := os.Create(path)

		if shared.HandleInternalServerError(w, err) {
			return NewFile, err
		}

		defer nf.Close()

		_, err = f.Seek(0, 0)

		if shared.HandleInternalServerError(w, err) {
			return NewFile, err
		}

		_, err = io.Copy(nf, f)

		if shared.HandleInternalServerError(w, err) {
			return NewFile, err
		}

		log.Printf("Файл получен и сохранён под именем %s", filename)

		NewFile.Filename = fh.Filename
		NewFile.Filesize = int(fh.Size)
		NewFile.Filetype = ext
		NewFile.FileID = filename

		NewFile.ID, err = databases.PostgreSQLFileChange(NewFile)

		if err != nil {
			if errors.Is(databases.ErrFirstNotUpdate, err) {
				shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
				return NewFile, err
			}

			if shared.HandleInternalServerError(w, err) {
				return NewFile, err
			}
		}

		NewFile.FileID = linktofile

	} else {

		shared.HandleOtherError(w, ErrUnsupportedFileType.Error(), ErrUnsupportedFileType, http.StatusBadRequest)
		return NewFile, ErrUnsupportedFileType
	}

	return NewFile, nil

}
