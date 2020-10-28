// Package shared - реализует весь функционал используемый в нескольких пакетах
package shared

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// CurrentPrefix - префикс URL
var CurrentPrefix = ""

// RequestResult - тип для хранения результата запроса
type RequestResult struct {
	Error ErrorResponse
}

// ErrorResponse - тип для хранения всяких ошибок и сообщений
type ErrorResponse struct {
	Code    int
	Message string
}

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

// HandleInternalServerError - обработчик внутренних ошибок сервера
func HandleInternalServerError(w http.ResponseWriter, err error) bool {

	if err != nil {

		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		Response := RequestResult{
			Error: ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
			},
		}

		w.WriteHeader(Response.Error.Code)

		WriteObjectToJSON(w, Response)

		return true
	}

	return false
}

// HandleBadRequestError - обработчик ошибки кривого запроса
func HandleBadRequestError(w http.ResponseWriter, err error) bool {

	if err != nil {

		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		Response := RequestResult{
			Error: ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Bad request",
			},
		}

		w.WriteHeader(Response.Error.Code)

		WriteObjectToJSON(w, Response)

		return true
	}

	return false
}

// HandleOtherError - обработчик прочих ошибок
func HandleOtherError(w http.ResponseWriter, message string, err error, statuscode int) bool {

	if err != nil {

		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		Response := RequestResult{
			Error: ErrorResponse{
				Code:    statuscode,
				Message: message,
			},
		}

		w.WriteHeader(Response.Error.Code)

		WriteObjectToJSON(w, Response)

		return true
	}

	return false
}

// HandleSuccessMessage - возвращает сообщение об успехе
func HandleSuccessMessage(w http.ResponseWriter, message string) {

	w.Header().Set("Content-Type", "application/json")

	Response := RequestResult{
		Error: ErrorResponse{
			Code:    200,
			Message: message,
		},
	}

	log.Println(message)

	WriteObjectToJSON(w, Response)

}

// FindInStringSlice - ищет в слайсе строк заданную строку
func FindInStringSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// WriteObjectToJSON - записывает в ответ произвольный объект
func WriteObjectToJSON(w http.ResponseWriter, v interface{}) {

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(v)

	if HandleInternalServerError(w, err) {
		return
	}

	_, err = w.Write(js)

	if HandleInternalServerError(w, err) {
		return
	}
}

// WriteBufferToPNG - записывает двоичный буффер в поток ответа для формата png
func WriteBufferToPNG(w http.ResponseWriter, b bytes.Buffer) {

	w.Header().Set("Content-Type", "image/png")

	_, err := w.Write(b.Bytes())

	if HandleInternalServerError(w, err) {
		return
	}

}
