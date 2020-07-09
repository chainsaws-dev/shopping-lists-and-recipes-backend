package databases

import (
	"fmt"
	"log"
)

// PostgreSQLGetConnString - получаем строку соединения для PostgreSQL
// При начальной настройке строка возвращается без базы данных (она создаётся в процессе)
// При начальной настройке указывается пароль суперпользователя при штатной работе пароль соответствуещей роли
func PostgreSQLGetConnString(Login string, Password string, Addr string, DbName string, initialsetup bool) string {

	if initialsetup {
		return fmt.Sprintf("postgres://%v:%v@%v/", Login, Password, Addr)
	}

	return fmt.Sprintf("postgres://%v:%v@%v/%v", Login, Password, Addr, DbName)

}

// PostgreSQLCreateDatabase - создаём базу данных для СУБД PostgreSQL
func PostgreSQLCreateDatabase(dbName string, ConnString string) {

	log.Println("Идёт создание базы данных...")

	// Подключаемся к базе данных
	dbc := SQLConnect("postgres", ConnString)
	defer dbc.Close()

	// Считаем количество баз данных с заданным именем
	rows, err := dbc.Query(`SELECT COUNT(datname) FROM pg_catalog.pg_database WHERE datname = $1;`, dbName)

	WriteErrToLog(err)

	var dbq int

	for rows.Next() {
		rows.Scan(&dbq)
	}

	// Если баз данных больше нуля, тогда ничего не делаем
	if dbq > 0 {
		return
	}

	// Иначе создаём базу данных с заданным именем
	// Параметром не подставляется не кртично ибо не используется в обычной работе
	// а только при установке, а так то это место для SQL инъекций
	createsql := `CREATE DATABASE "` + dbName + `"
	WITH
	OWNER = postgres
	ENCODING = 'UTF8'
	LC_COLLATE = 'C.UTF-8'
	LC_CTYPE = 'C.UTF-8'
	TABLESPACE = pg_default
	CONNECTION LIMIT = -1;`

	_, err = dbc.Exec(createsql)

	WriteErrToLog(err)

	log.Println("База данных успешно создана")

}
