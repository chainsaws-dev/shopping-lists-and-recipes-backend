package databases

import (
	"database/sql"
	"log"
)

// SQLConnect - соединиться с базой данных и выполнить команду
// Не забываем в точке вызова defer db.Close()
func SQLConnect(DbType string, ConStr string) *sql.DB {

	db, err := sql.Open(DbType, ConStr)

	if err != nil {
		log.Fatalln(err)
	}

	// Проверяем что база данных доступна
	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// SQLSelect - Выполняет выборку данных с параметрами или без
// Не забываем в точке вызова defer rows.Close()
func SQLSelect(Db *sql.DB, SelectStmnt string, Args ...interface{}) (*sql.Rows, error) {

	var rows *sql.Rows
	var err error

	if Args == nil {
		rows, err = Db.Query(SelectStmnt)
	} else {
		rows, err = Db.Query(SelectStmnt, Args)
	}

	return rows, err

}

// SQLExecute - Выполняет TSQL выражение с параметрами
func SQLExecute(db *sql.DB, SelectStmnt string, Args ...interface{}) (sql.Result, error) {

	result, err := db.Exec(SelectStmnt, Args)

	return result, err

}
