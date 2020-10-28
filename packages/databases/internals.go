package databases

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"shopping-lists-and-recipes/packages/shared"
)

// Список типовых ошибок
var (
	ErrFirstNotDelete       = errors.New("Первая запись в списке файлов техническая и не подлежит удалению")
	ErrFirstNotUpdate       = errors.New("Первая запись в списке файлов техническая и не подлежит изменению")
	ErrRecipeNotFound       = errors.New("В таблице рецептов не найден указанный id")
	ErrShoppingListNotFound = errors.New("Не найдено ни одной записи в списке покупок с указанным названием")
	ErrEmptyPassword        = errors.New("Не допустимо использование паролей с длинной менее шести символов")
	ErrNoUserWithEmail      = errors.New("Электронная почта не найдена")
	ErrNoHashForUser        = errors.New("Хеш пароля не найден")
	ErrEmailIsOccupied      = errors.New("Указанный адрес электронной почты уже занят")
	ErrUserNotFound         = errors.New("В таблице пользователей не найден указанный id")
	ErrUserTOTPNotFound     = errors.New("В таблице секретов для двухфакторной авторизации не найден указанный id")
	ErrTOTPConfirmed        = errors.New("В таблице секретов для двухфакторной авторизации указанный id уже привязан")
	ErrEmailNotConfirmed    = errors.New("Подтвердите адрес электронной почты")
	ErrTokenExpired         = errors.New("Токен истёк или не существует")
	ErrLimitOffsetInvalid   = errors.New("Limit и Offset приняли недопустимое значение")
)

var dbc *sql.DB

// PostgreSQLGetConnString - получаем строку соединения для PostgreSQL
// При начальной настройке строка возвращается без базы данных (она создаётся в процессе)
// При начальной настройке указывается пароль суперпользователя при штатной работе пароль соответствуещей роли
func PostgreSQLGetConnString(Login string, Password string, Addr string, DbName string, initialsetup bool) string {

	if initialsetup {
		return fmt.Sprintf("postgres://%v:%v@%v/", Login, Password, Addr)
	}

	return fmt.Sprintf("postgres://%v:%v@%v/%v", Login, Password, Addr, DbName)

}

// PostgreSQLRollbackIfError - откатываем изменения транзакции если происходит ошибка и пишем её в лог и выходим
func PostgreSQLRollbackIfError(err error, critical bool) error {
	if err != nil {
		dbc.Exec("ROLLBACK")

		if critical {
			log.Fatalln(err)
		} else {
			log.Println(err)
		}

		return err
	}

	return nil
}

// PostgreSQLCloseConn - закрываем соединение с базой данных
func PostgreSQLCloseConn() {
	dbc.Close()
}

// PostgreSQLConnect - подключаемся к базе данных
func PostgreSQLConnect(ConnString string) error {

	var err error

	dbc, err = shared.SQLConnect("postgres", ConnString)

	return err

}

// PostgreSQLCheckLimitOffset - проверяем значение лимита и сдвига
//
// Параметры:
//
// Limit - максимум строк на странице (должно быть меньше 50)
// Offset - сдвиг относительно первой строки
//
func PostgreSQLCheckLimitOffset(Limit int, Offset int) bool {
	return Offset >= 0 && Limit > 0 && Limit <= 50
}

// СheckExists - проверяем что файл или папка существует
func СheckExists(filename string) bool {

	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}
