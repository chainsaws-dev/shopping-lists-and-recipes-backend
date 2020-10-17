// Package securecookies - содержит функции и процедуры для работы с куками
package securecookies

import (
	"errors"
	"net/http"
	"shopping-lists-and-recipes/packages/authentication"
	"time"

	"github.com/gorilla/securecookie"
)

// Список типовых ошибок
var (
	ErrPairNotFound = errors.New("Не найдена запись по параметрам запроса")
)

var seccookie *securecookie.SecureCookie

func init() {
	seccookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

// MyCookies - тип для хранения данных из Кук
type MyCookies struct {
	Email   string
	Session []byte
}

// GetCookies - получает безопасные куки для почты, идентификатора письма и сессии
func GetCookies(w http.ResponseWriter, req *http.Request) (MyCookies, error) {
	// Читаем куки
	cookie, err := req.Cookie("Email")
	if err != nil {
		return MyCookies{}, err
	}
	var Email string
	err = seccookie.Decode("Email", cookie.Value, &Email)
	if err != nil {
		return MyCookies{}, err
	}

	cookie, err = req.Cookie("Session")
	if err != nil {
		return MyCookies{}, err
	}
	var Session []byte
	err = seccookie.Decode("Session", cookie.Value, &Session)
	if err != nil {
		return MyCookies{}, err
	}

	return MyCookies{
		Email:   Email,
		Session: Session,
	}, nil
}

// SetCookies - записывает безопасные куки в клиент
func SetCookies(expires time.Time, activetoken authentication.ActiveToken, w http.ResponseWriter) error {
	// Пишем куки
	encoded, err := seccookie.Encode("Email", activetoken.Email)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Email",
		Value:   encoded,
		Path:    "/",
		Expires: expires,
	})

	encoded, err = seccookie.Encode("Session", activetoken.Session)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Session",
		Value:   encoded,
		Path:    "/",
		Expires: expires,
	})

	return nil
}
