// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Список типовых ошибок
var (
	ErrFirstNotDelete          = errors.New("первая запись в списке файлов техническая и не подлежит удалению")
	ErrFirstNotUpdate          = errors.New("первая запись в списке файлов техническая и не подлежит изменению")
	ErrRecipeNotFound          = errors.New("в таблице рецептов не найден указанный id")
	ErrShoppingListNotFound    = errors.New("не найдено ни одной записи в списке покупок с указанным названием")
	ErrSessionsNotFoundEmail   = errors.New("не найдено ни одной записи в списке сессий для указанной почты")
	ErrSessionsNotFoundExpired = errors.New("не найдено ни одной истекшей записи в списке сессий")
	ErrSessionsNotFoundSession = errors.New("не найдено ни одной записи в списке сессий")
	ErrEmptyPassword           = errors.New("не допустимо использование паролей с длинной менее шести символов")
	ErrNoUserWithEmail         = errors.New("электронная почта не найдена")
	ErrNoHashForUser           = errors.New("хеш пароля не найден")
	ErrEmailIsOccupied         = errors.New("указанный адрес электронной почты уже занят")
	ErrUserNotFound            = errors.New("в таблице пользователей не найден указанный id")
	ErrUserTOTPNotFound        = errors.New("в таблице секретов для двухфакторной авторизации не найден указанный id")
	ErrTOTPConfirmed           = errors.New("в таблице секретов для двухфакторной авторизации указанный id уже привязан")
	ErrEmailNotConfirmed       = errors.New("подтвердите адрес электронной почты")
	ErrTokenExpired            = errors.New("токен истёк или не существует")
	ErrLimitOffsetInvalid      = errors.New("limit и offset приняли недопустимое значение")
	ErrDatabaseDoesntExist     = errors.New("базы данных с указанным именем не найдено")
	ErrRoleDoesntExist         = errors.New("роли с указанным именем не найдено")
	ErrNoConnection            = errors.New("база данных недоступна")
	ErrDatabaseAlreadyExists   = errors.New("база данных с указанным именем уже существует")
	ErrTablesAlreadyExist      = errors.New("база данных содержит таблицы")
)

// PostgreSQLGetConnString - получаем строку соединения для PostgreSQL
// При начальной настройке строка возвращается без базы данных (она создаётся в процессе)
// При начальной настройке указывается пароль суперпользователя при штатной работе пароль соответствуещей роли
func PostgreSQLGetConnString(Login string, Password string, Addr string, DbName string, initialsetup bool, SSLmode bool, MaxPoolConns int) string {
	var strSSLmode string

	if SSLmode {
		strSSLmode = "enable"
	} else {
		strSSLmode = "disable"
	}

	if initialsetup {
		return fmt.Sprintf("postgres://%v:%v@%v/?sslmode=%v&pool_max_conns=%v", Login, Password, Addr, strSSLmode, MaxPoolConns)
	}

	return fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=%v&pool_max_conns=%v", Login, Password, Addr, DbName, strSSLmode, MaxPoolConns)

}

// PostgreSQLRollbackIfError - откатываем изменения транзакции если происходит ошибка и пишем её в лог и выходим
func PostgreSQLRollbackIfError(err error, critical bool, dbc *pgxpool.Pool) error {
	if err != nil {
		dbc.Exec(context.Background(), "ROLLBACK")

		if critical {
			log.Fatalln(err)
		} else {
			log.Println(err)
		}

		return err
	}

	return nil
}

// PostgreSQLConnect - подключаемся к базе данных (создаём пул соединений)
func PostgreSQLConnect(ConnectionString string) (dbc *pgxpool.Pool) {
	dbc, err := pgxpool.Connect(context.Background(), ConnectionString)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	log.Println("Установлено соединение с СУБД")
	return dbc
}

// PostgreSQLDisconnect - отключаемся от базы данных (убиваем пул соединений)
func PostgreSQLDisconnect(dbc *pgxpool.Pool) {
	dbc.Close()
	dbc = nil
	log.Println("Соединение с СУБД разорвано")
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
