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
func SQLConnect(DbType string, ConStr string) (*sql.DB, error) {

	db, err := sql.Open(DbType, ConStr)

	if err != nil {
		return db, err
	}

	// Проверяем что база данных доступна
	err = db.Ping()

	if err != nil {
		return db, err
	}

	return db, nil
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

		errortext := fmt.Sprintf(`{"Error":{"Code":%v, "Message":"%v"}}`, http.StatusInternalServerError, "Internal server error")
		ReturnJSONError(w, errortext, http.StatusInternalServerError)

		log.Println(err)

		return true
	}

	return false
}

// HandleForbiddenError - Обработчик ошибок нарушения прав доступа
func HandleForbiddenError(w http.ResponseWriter, err error) bool {

	if err != nil {

		errortext := fmt.Sprintf(`{"Error":{"Code":%v, "Message":"%v"}}`, http.StatusForbidden, "Access forbidden")
		ReturnJSONError(w, errortext, http.StatusInternalServerError)
		log.Println(err)
		return true
	}

	return false
}

// HandleOtherError - Обработчик прочих ошибок
func HandleOtherError(w http.ResponseWriter, message string, err error, statuscode int) bool {

	if err != nil {
		errortext := fmt.Sprintf(`{"Error":{"Code":%v, "Message":"%v"}}`, statuscode, message)
		ReturnJSONError(w, errortext, statuscode)
		log.Println(err)
		return true
	}

	return false
}

// ReturnJSONError - возвращает ошибку в виде JSON
func ReturnJSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, err)
}
