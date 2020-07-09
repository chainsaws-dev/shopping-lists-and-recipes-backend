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
