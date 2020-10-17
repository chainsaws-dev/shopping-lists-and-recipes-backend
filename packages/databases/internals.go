package databases

import (
	"log"
	"shopping-lists-and-recipes/packages/shared"
)

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
