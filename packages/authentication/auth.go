package authentication

import (
	"encoding/hex"
	"errors"
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"net/http"
)

// Список типовых ошибок
var (
	ErrNotAllowedMethod = errors.New("Запрошен недопустимый метод при авторизации")
)

// SignIn - обработчик для авторизации пользователя
func SignIn(w http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == http.MethodPost:
		w.Header().Set("Content-Type", "application/json")

		tokenb, err := GenerateRandomBytes(32)

		if shared.HandleInternalServerError(w, err) {
			return
		}

		log.Println(hex.EncodeToString(tokenb))
	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}
}

// SignUp - обработчик для регистрации пользователя
func SignUp(w http.ResponseWriter, req *http.Request) {

}
