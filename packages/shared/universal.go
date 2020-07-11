// Package shared - реализует весь функционал используемый в нескольких пакетах
package shared

import (
	"database/sql"
	"fmt"
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

// WriteErrToLog - пишем критическую ошибку в лог
func WriteErrToLog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// HandleInternalServerError - Обработчик внутренних ошибок сервера
func HandleInternalServerError(w http.ResponseWriter, err error) bool {

	if err != nil {

		errortext := fmt.Sprintf(`{"Error":{"code":%v, "message":"%v"}}`, http.StatusInternalServerError, "Internal server error")
		ReturnJSONError(w, errortext, http.StatusInternalServerError)
		log.Println(err)
		return true
	}

	return false
}

// HandleForbiddenError - Обработчик ошибок нарушения прав доступа
func HandleForbiddenError(w http.ResponseWriter, err error) bool {

	if err != nil {

		errortext := fmt.Sprintf(`{"Error":{"code":%v, "message":"%v"}}`, http.StatusForbidden, "Access forbidden")
		ReturnJSONError(w, errortext, http.StatusInternalServerError)
		log.Println(err)
		return true
	}

	return false
}

// HandleOtherError - Обработчик прочих ошибок
func HandleOtherError(w http.ResponseWriter, err string, statuscode int) {

	errortext := fmt.Sprintf(`{"Error":{"code":%v, "message":"%v"}}`, statuscode, err)
	ReturnJSONError(w, errortext, statuscode)
	log.Println(err)

}

// ReturnJSONError - возвращает ошибку в виде JSON
func ReturnJSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, err)
}
