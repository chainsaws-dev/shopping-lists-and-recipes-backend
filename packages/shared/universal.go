// Package shared - реализует весь функционал используемый в нескольких пакетах
package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"shopping-lists-and-recipes/packages/multilangtranslator"
)

// Список общих типовых ошибок
var (
	ErrNotAllowedMethod       = errors.New("http request method is not allowed")
	ErrNoKeyInParams          = errors.New("http request does not contain api key parameter")
	ErrWrongKeyInParams       = errors.New("api key is not registered")
	ErrNotAuthorized          = errors.New("authorization required")
	ErrNotAuthorizedTwoFactor = errors.New("two factor authorization required")
	ErrForbidden              = errors.New("access forbidden")
	ErrBadHttpRequest         = errors.New("bad http request")
	ErrHeadersFetchNotFilled  = errors.New("http request does not contain required parameters: Page, Limit")
	ErrHeadersNotFilled       = errors.New("http request does not contain required parameters")
	ErrLimitOffsetInvalid     = errors.New("invalid http request parameters Limit and Offset")
	ErrFkeyViolation          = errors.New("cannot delete record referenced from other tables")
	ErroNoRowsInResult        = errors.New("nothing to delete")
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
func HandleInternalServerError(ServerLanguage string, w http.ResponseWriter, r *http.Request, err error) bool {

	locale := r.Header.Get("Lang")

	if err != nil {

		log.Println(multilangtranslator.TranslateError(err, ServerLanguage))

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		Response := RequestResult{
			Error: ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: multilangtranslator.TranslateString("Internal server error", locale),
			},
		}

		w.WriteHeader(Response.Error.Code)

		WriteObjectToJSON(ServerLanguage, w, r, Response)

		return true
	}

	return false
}

// HandleBadRequestError - обработчик ошибки кривого запроса
func HandleBadRequestError(ServerLanguage string, w http.ResponseWriter, r *http.Request, err error) bool {

	locale := r.Header.Get("Lang")

	if err != nil {

		log.Println(multilangtranslator.TranslateError(err, ServerLanguage))

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		Response := RequestResult{
			Error: ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: multilangtranslator.TranslateString(ErrBadHttpRequest.Error(), locale),
			},
		}

		w.WriteHeader(Response.Error.Code)

		WriteObjectToJSON(ServerLanguage, w, r, Response)

		return true
	}

	return false
}

// HandleOtherError - обработчик прочих ошибок
func HandleOtherError(ServerLanguage string, w http.ResponseWriter, r *http.Request, message string, err error, statuscode int) bool {

	locale := r.Header.Get("Lang")

	if err != nil {

		log.Println(multilangtranslator.TranslateError(err, ServerLanguage))

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		Response := RequestResult{
			Error: ErrorResponse{
				Code:    statuscode,
				Message: multilangtranslator.TranslateString(message, locale),
			},
		}

		w.WriteHeader(Response.Error.Code)

		WriteObjectToJSON(ServerLanguage, w, r, Response)

		return true
	}

	return false
}

// HandleSuccessMessage - возвращает сообщение об успехе
func HandleSuccessMessage(ServerLanguage string, w http.ResponseWriter, r *http.Request, message string) {

	locale := r.Header.Get("Lang")

	w.Header().Set("Content-Type", "application/json")

	Response := RequestResult{
		Error: ErrorResponse{
			Code:    200,
			Message: multilangtranslator.TranslateString(message, locale),
		},
	}

	log.Println(multilangtranslator.TranslateString(message, ServerLanguage))

	WriteObjectToJSON(ServerLanguage, w, r, Response)

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
func WriteObjectToJSON(ServerLanguage string, w http.ResponseWriter, r *http.Request, v interface{}) {

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(v)

	if HandleInternalServerError(ServerLanguage, w, r, err) {
		return
	}

	_, err = w.Write(js)

	if HandleInternalServerError(ServerLanguage, w, r, err) {
		return
	}
}

// WriteBufferToPNG - записывает двоичный буффер в поток ответа для формата png
func WriteBufferToPNG(ServerLanguage string, w http.ResponseWriter, r *http.Request, b bytes.Buffer) {

	w.Header().Set("Content-Type", "image/png")

	_, err := w.Write(b.Bytes())

	if HandleInternalServerError(ServerLanguage, w, r, err) {
		return
	}

}
