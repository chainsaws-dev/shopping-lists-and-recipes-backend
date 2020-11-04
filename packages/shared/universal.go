// Package shared - реализует весь функционал используемый в нескольких пакетах
package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Список общих типовых ошибок
var (
	ErrNotAllowedMethod       = errors.New("Запрошен недопустимый метод для файлов")
	ErrNoKeyInParams          = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams       = errors.New("API ключ не зарегистрирован")
	ErrNotAuthorized          = errors.New("Пройдите авторизацию")
	ErrNotAuthorizedTwoFactor = errors.New("Пройдите авторизацию по второму фактору")
	ErrForbidden              = errors.New("Доступ запрещён")
	ErrHeadersFetchNotFilled  = errors.New("Не заполнены обязательные параметры запроса списка файлов: Page, Limit")
	ErrFkeyViolation          = errors.New("Нельзя удалять записи, на которые имеются ссылки")
	ErroNoRowsInResult        = errors.New("Не найдено ни одной записи для удаления")
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
