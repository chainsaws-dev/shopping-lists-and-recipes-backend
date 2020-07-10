// Package shared - реализует весь функционал используемый в нескольких пакетах
package shared

import (
	"database/sql"
	"log"
	"net/http"
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

// WriteErrToLog - пишем ошибку в лог
func WriteErrToLog(err error) {
	if err != nil {
		log.Fatalln(err)
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
