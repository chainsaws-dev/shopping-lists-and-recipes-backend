package files

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/internal/setup"
	"shopping-lists-and-recipes/packages/shared"
	"strings"

	"github.com/bakape/thumbnailer/v2"
)

// fileUpload - выполняет загрузку файла на сервер и сохранение в файловой системе и информации в базе данных
func fileUpload(w http.ResponseWriter, req *http.Request, role string) (databases.File, error) {

	log.Println("Начинаем получение файла...")

	var NewFile databases.File
	// Читаем из формы двоичные данные
	f, fh, err := req.FormFile("file")

	if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
		return NewFile, err
	}
	defer f.Close()

	// Проверяем тип файла
	buff := make([]byte, 512)
	_, err = f.Read(buff)

	if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
		return NewFile, err
	}

	filetype := http.DetectContentType(buff)

	if filetype == "image/jpeg" || filetype == "image/jpg" ||
		filetype == "image/gif" || filetype == "image/png" {

		// На всякий случай сохраняем расширение
		ext := strings.Split(fh.Filename, ".")[1]

		fn := sha1.New()

		io.Copy(fn, f)

		basename := fmt.Sprintf("%x", fn.Sum(nil))

		filename := basename + "." + ext

		linktofile := strings.Join([]string{"uploads", filename}, "/")

		pathtofile := filepath.Join(".", "public", "uploads", filename)

		nf, err := os.Create(pathtofile)

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return NewFile, err
		}

		defer nf.Close()

		_, err = f.Seek(0, 0)

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return NewFile, err
		}

		_, err = io.Copy(nf, f)

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return NewFile, err
		}

		log.Printf("Файл получен и сохранён под именем %s", filename)

		NewFile.Filename = fh.Filename
		NewFile.Filesize = int(fh.Size)
		NewFile.Filetype = ext
		NewFile.FileID = filename

		// Создаем превьюшку
		_, thumb, err := thumbnailer.Process(f, thumbnailer.Options{
			ThumbDims: thumbnailer.Dims{
				Width:  440,
				Height: 248,
			},
		})

		var previewname string

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return NewFile, err
		}

		if strings.ToLower(ext) == "png" {
			previewname = basename + "pv.png"
		} else {
			previewname = basename + "pv.jpg"
		}

		linktopreview := strings.Join([]string{"uploads", previewname}, "/")

		pathtopreview := filepath.Join(".", "public", "uploads", previewname)

		nfp, err := os.Create(pathtopreview)

		if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
			return NewFile, err
		}

		defer nfp.Close()

		if strings.ToLower(ext) == "png" {
			png.Encode(nfp, thumb)
		} else {
			jpeg.Encode(nfp, thumb, &jpeg.Options{Quality: jpeg.DefaultQuality})
		}

		log.Printf("Превью файла сохранено под именем %s", previewname)

		NewFile.PreviewID = previewname

		NewFile.ID, err = databases.PostgreSQLFileChange(NewFile, setup.ServerSettings.SQL.ConnPool)

		if err != nil {
			if errors.Is(databases.ErrFirstNotUpdate, err) {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
				return NewFile, err
			}

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return NewFile, err
			}
		}

		NewFile.FileID = linktofile
		NewFile.PreviewID = linktopreview

	} else {

		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrUnsupportedFileType.Error(), ErrUnsupportedFileType, http.StatusBadRequest)
		return NewFile, ErrUnsupportedFileType
	}

	return NewFile, nil

}
