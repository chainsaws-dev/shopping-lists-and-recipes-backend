// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLFileInsert - создаёт записи в базе данных для хранения информации о загруженном файле
func PostgreSQLFileInsert(NewFile File) (int, error) {

	dbc.Exec("BEGIN")

	sql := `INSERT INTO 
			public."Files" 
			(filename, filesize, filetype, file_id) 
		  VALUES 
			($1, $2, $3, $4) RETURNING id;`

	row := dbc.QueryRow(sql, NewFile.Filename, NewFile.Filesize, NewFile.Filetype, NewFile.FileID)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return -1, PostgreSQLRollbackIfError(err, true)
	}

	log.Printf("Данные о файле сохранены в базу данных под индексом %v", curid)

	dbc.Exec("COMMIT")

	return curid, nil
}

// PostgreSQLFileUpdate - обновляет записи в базе данных для хранения информации о загруженном файле
func PostgreSQLFileUpdate(f File) error {

	if f.ID == 1 {
		return ErrFirstNotUpdate
	}

	dbc.Exec("BEGIN")

	sql := `UPDATE
				public."Files" 
			SET 
				(filename, filesize, filetype, file_id) = ($1, $2, $3, $4) 
			WHERE 
				id = $5;`

	_, err := dbc.Exec(sql, f.Filename, f.Filesize, f.Filetype, strings.ReplaceAll(f.FileID, "/uploads/", ""), f.ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, true)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLFileDelete - удаляет запись в базе данных о загруженном файле
func PostgreSQLFileDelete(fileid int) error {

	if fileid == 1 {
		return ErrFirstNotDelete
	}

	dbc.Exec("BEGIN")

	sql := `SELECT 
				file_id
			FROM 
				public."Files" 
			WHERE id=$1;`

	row := dbc.QueryRow(sql, fileid)

	var filename string
	err := row.Scan(&filename)

	if err != nil {
		return err
	}

	sql = `DELETE FROM 
				public."Files" 
			WHERE id=$1;`

	_, err = dbc.Exec(sql, fileid)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	path := strings.Join([]string{".", "public", "uploads", filename}, "/")
	err = os.Remove(path)

	if err != nil {
		return err
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
func PostgreSQLFilesSelect(page int, limit int) (FilesResponse, error) {

	var result FilesResponse
	result.Files = FilesList{}

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Files";`

	row := dbc.QueryRow(sql)

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
		sql = fmt.Sprintf(`SELECT 
							files.id,
							files.filename,
							files.filesize,
							files.filetype,
							files.file_id
						FROM 
							public."Files" AS files
						ORDER BY files.id
						OFFSET %v LIMIT %v;`, offset, limit)
	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(sql)

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
