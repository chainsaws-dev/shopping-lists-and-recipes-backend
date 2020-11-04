// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLFileChange - определяет существует ли такой же файл в базе
// и изменяет или создаёт новый в зависимости от результата проверки
func PostgreSQLFileChange(f File, dbc *sql.DB) (int, error) {

	sqlreq := `SELECT 
					COUNT(*)
				FROM
					public."Files"
				WHERE 
					file_id=$1`

	row := dbc.QueryRow(sqlreq, f.FileID)

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
func PostgreSQLFileInsert(f File, dbc *sql.DB) (int, error) {

	dbc.Exec("BEGIN")

	sqlreq := `INSERT INTO 
			public."Files"
			(filename, filesize, filetype, file_id) 
		  VALUES 
			($1, $2, $3, $4) RETURNING id;`

	row := dbc.QueryRow(sqlreq, f.Filename, f.Filesize, f.Filetype, f.FileID)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return curid, PostgreSQLRollbackIfError(err, true, dbc)
	}

	log.Printf("Данные о файле сохранены в базу данных под индексом %v", curid)

	dbc.Exec("COMMIT")

	return curid, nil
}

// PostgreSQLFileUpdate - перезаписывает данные в базе о уже существующем файле
//
// Параметры:
//
// f - тип файл, содержащий данные о файле (имя, размер, тип, имя на сервере)
//
func PostgreSQLFileUpdate(f File, dbc *sql.DB) (int, error) {

	sqlreq := `SELECT 
					id
				FROM
					public."Files"
				WHERE 
					file_id=$1`

	row := dbc.QueryRow(sqlreq, f.FileID)

	var DbID int
	err := row.Scan(&DbID)

	if err != nil {
		return -1, err
	}

	f.ID = DbID

	dbc.Exec("BEGIN")

	sqlreq = `UPDATE 
				public."Files"
				SET (filename, filesize, filetype, file_id) = ($1, $2, $3, $4)
				WHERE
					file_id=$4;`

	_, err = dbc.Exec(sqlreq, f.Filename, f.Filesize, f.Filetype, f.FileID)

	if err != nil {
		return f.ID, PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return f.ID, nil
}

// PostgreSQLFileDelete - удаляет запись в базе данных о загруженном файле
func PostgreSQLFileDelete(fileid int, dbc *sql.DB) error {

	if fileid == 1 {
		return ErrFirstNotDelete
	}

	dbc.Exec("BEGIN")

	sqlreq := `SELECT 
				file_id
			FROM 
				public."Files"
			WHERE id=$1;`

	row := dbc.QueryRow(sqlreq, fileid)

	var filename string
	err := row.Scan(&filename)

	if err != nil {
		return err
	}

	sqlreq = `DELETE FROM 
				public."Files"
			WHERE id=$1;`

	_, err = dbc.Exec(sqlreq, fileid)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	path := strings.Join([]string{".", "public", "uploads", filename}, "/")

	if СheckExists(path) {
		err = os.Remove(path)
	}

	if err != nil {
		return err
	}

	sqlreq = `select setval('"public"."Files_id_seq"',(select COALESCE(max("id"),1) from "public"."Files")::bigint);`

	_, err = dbc.Exec(sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLFilesSelect - получает информацию о файлах
//
// Параметры:
//
// page - номер страницы результата для вывода
// limit - количество строк на странице
//
func PostgreSQLFilesSelect(page int, limit int, dbc *sql.DB) (FilesResponse, error) {

	var result FilesResponse
	result.Files = FilesList{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Files"
			WHERE
				id<>1;`

	row := dbc.QueryRow(sqlreq)

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
							files.file_id
						FROM 
							public."Files" AS files
						WHERE
							id<>1
						ORDER BY files.id
						OFFSET %v LIMIT %v;`, offset, limit)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur File
		err = rows.Scan(&cur.ID, &cur.Filename, &cur.Filesize, &cur.Filetype, &cur.FileID)
		if err != nil {
			return result, err
		}

		cur.FileID = strings.Join([]string{"", "uploads", cur.FileID}, "/")
		result.Files = append(result.Files, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}
