// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLFileChange - определяет существует ли такой же файл в базе
// и изменяет или создаёт новый в зависимости от результата проверки
func PostgreSQLFileChange(f File, dbc *pgxpool.Pool) (int, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM
					public."Files"
				WHERE 
					file_id=$1`

	row := dbc.QueryRow(context.Background(), sqlreq, f.FileID)

	var CountRows int
	err := row.Scan(&CountRows)

	if err != nil {
		return -1, err
	}

	if CountRows > 0 {
		return PostgreSQLFileUpdate(f, dbc)
	}

	return PostgreSQLFileInsert(f, dbc)
}

// PostgreSQLFileInsert - создаёт записи в базе данных для хранения информации о загруженном файле
//
// Параметры:
//
// f - тип файл, содержащий данные о файле (имя, размер, тип, имя на сервере)
//
func PostgreSQLFileInsert(f File, dbc *pgxpool.Pool) (int, error) {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `INSERT INTO 
			public."Files"
			(filename, filesize, filetype, file_id, preview_id) 
		  VALUES 
			($1, $2, $3, $4, $5) RETURNING id;`

	row := dbc.QueryRow(context.Background(), sqlreq, f.Filename, f.Filesize, f.Filetype, f.FileID, f.PreviewID)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return curid, PostgreSQLRollbackIfError(err, true, dbc)
	}

	log.Printf("Данные о файле сохранены в базу данных под индексом %v", curid)

	dbc.Exec(context.Background(), "COMMIT")

	return curid, nil
}

// PostgreSQLFileUpdate - перезаписывает данные в базе о уже существующем файле
//
// Параметры:
//
// f - тип файл, содержащий данные о файле (имя, размер, тип, имя на сервере)
//
func PostgreSQLFileUpdate(f File, dbc *pgxpool.Pool) (int, error) {

	sqlreq := `SELECT 
					id
				FROM
					public."Files"
				WHERE 
					file_id=$1`

	row := dbc.QueryRow(context.Background(), sqlreq, f.FileID)

	var DbID int
	err := row.Scan(&DbID)

	if err != nil {
		return -1, err
	}

	f.ID = DbID

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq = `UPDATE 
				public."Files"
				SET (filename, filesize, filetype, file_id, preview_id) = ($1, $2, $3, $4, $5)
				WHERE
					file_id=$4;`

	_, err = dbc.Exec(context.Background(), sqlreq, f.Filename, f.Filesize, f.Filetype, f.FileID, f.PreviewID)

	if err != nil {
		return f.ID, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return f.ID, nil
}

// PostgreSQLFileDelete - удаляет запись в базе данных о загруженном файле
func PostgreSQLFileDelete(fileid int, dbc *pgxpool.Pool) error {

	if fileid == 1 {
		return ErrFirstNotDelete
	}

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `SELECT 
				file_id,
				preview_id
			FROM 
				public."Files"
			WHERE id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, fileid)

	var filename string
	var previewname string
	err := row.Scan(&filename, &previewname)

	if err != nil {
		return err
	}

	sqlreq = `DELETE FROM 
				public."Files"
			WHERE id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, fileid)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"public"."Files_id_seq"',(select COALESCE(max("id"),1) from "public"."Files")::bigint);`

	_, err = dbc.Exec(context.Background(), sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	// Удаляем сами файлы с диска
	err = DeleteFileFromDisk(filename)
	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	err = DeleteFileFromDisk(previewname)
	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}

// DeleteFileFromDisk - удаляет файл с жесткого диска сервера
func DeleteFileFromDisk(filename string) error {

	var err error

	path := strings.Join([]string{".", "public", "uploads", filename}, "/")

	if СheckExists(path) {
		err = os.Remove(path)
	}

	if err != nil {
		return err
	}

	return nil
}

// PostgreSQLFilesSelect - получает информацию о файлах
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLFilesSelect(page int, limit int, dbc *pgxpool.Pool) (FilesResponse, error) {

	var result FilesResponse
	result.Files = FilesList{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Files"
			WHERE
				id<>1;`

	row := dbc.QueryRow(context.Background(), sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows <= 0 {
		return result, nil
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {
		sqlreq = fmt.Sprintf(`SELECT 
							files.id,
							files.filename,
							files.filesize,
							files.filetype,
							files.file_id,
							files.preview_id
						FROM 
							public."Files" AS files
						WHERE
							id<>1
						ORDER BY files.id
						OFFSET %v LIMIT %v;`, offset, limit)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(context.Background(), sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur File
		err = rows.Scan(&cur.ID, &cur.Filename, &cur.Filesize, &cur.Filetype, &cur.FileID, &cur.PreviewID)
		if err != nil {
			return result, err
		}

		cur.FileID = strings.Join([]string{"", "uploads", cur.FileID}, "/")
		cur.PreviewID = strings.Join([]string{"", "uploads", cur.PreviewID}, "/")
		result.Files = append(result.Files, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}
