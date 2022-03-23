// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"shopping-lists-and-recipes/packages/shared"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Schemas - список схем, которые должны быть созданы перед созданием таблиц базы
var Schemas = []string{
	`public`,
	`secret`,
}

// PostgreSQLCreateTables - Создаём таблицы в базе данных
func PostgreSQLCreateTables(dbc *pgxpool.Pool, locale string) error {

	// Начало транзакции
	dbc.Exec(context.Background(), "BEGIN")

	log.Println("Создаём базу и схемы...")

	err := PostgreSQLCreateSchemas(dbc, locale)

	if err != nil {
		return err
	}

	log.Println("Создаём таблицы ...")

	log.Println("\tДля схемы public ...")
	// Создание таблиц для списка покупок и рецептов
	PostgreSQLCreateTablesPublic(dbc)

	log.Println("\tДля схемы secret ...")
	// Создание таблиц для авторизации и админки
	PostgreSQLCreateTablesSecret(dbc)

	// Фиксация транзакции
	dbc.Exec(context.Background(), "COMMIT")

	log.Println("Таблицы созданы")

	return nil

}

// PostgreSQLCreateDatabase - создаём базу данных для СУБД PostgreSQL
func PostgreSQLCreateDatabase(dbName string, dbc *pgxpool.Pool, locale string) {

	if dbc != nil {
		log.Println("Идёт создание базы данных...")

		// Считаем количество баз данных с заданным именем
		rows, err := dbc.Query(context.Background(), `SELECT COUNT(datname) FROM pg_catalog.pg_database WHERE datname = $1;`, dbName)

		shared.WriteErrToLog(err, locale)

		var dbq int

		for rows.Next() {
			rows.Scan(&dbq)
		}

		// Если баз данных больше нуля, тогда ничего не делаем
		if dbq > 0 {
			log.Printf("Уже существует база данных с именем %s\n", dbName)
			return
		}

		// Иначе создаём базу данных с заданным именем
		// Параметром не подставляется не кртично ибо не используется в обычной работе
		// а только при установке, а так то это место для SQL инъекций
		sqlreq := fmt.Sprintf(`CREATE DATABASE "%s"
									WITH
									OWNER = postgres
									ENCODING = 'UTF8'
									LC_COLLATE = 'C.UTF-8'
									LC_CTYPE = 'C.UTF-8'
									TABLESPACE = pg_default
									CONNECTION LIMIT = -1
									TEMPLATE = template0;`, dbName)

		_, err = dbc.Exec(context.Background(), sqlreq)

		shared.WriteErrToLog(err, locale)

		log.Println("База данных успешно создана")
	}

}

// PostgreSQLDropDatabase - удаляет базу данных с заданным именем
func PostgreSQLDropDatabase(dbName string, dbc *pgxpool.Pool, locale string) {

	if dbc != nil {

		log.Println("Идёт удаление базы данных...")

		// Считаем количество баз данных с заданным именем
		rows, err := dbc.Query(context.Background(), `SELECT COUNT(datname) FROM pg_catalog.pg_database WHERE datname = $1;`, dbName)

		shared.WriteErrToLog(err, locale)

		var dbq int

		for rows.Next() {
			rows.Scan(&dbq)
		}

		// Если баз данных больше нуля, тогда ничего не делаем
		if dbq <= 0 {
			log.Printf("Не найдена база данных с именем %s\n", dbName)
			return
		}

		_, err = dbc.Exec(context.Background(), `SELECT pg_terminate_backend(pg_stat_activity.pid)
							FROM pg_stat_activity
							WHERE pg_stat_activity.datname = $1
							AND pid <> pg_backend_pid();`, dbName)

		if err != nil {
			log.Println(err)
			return
		}

		sqlreq := fmt.Sprintf(`DROP DATABASE "%s";`, dbName)

		_, err = dbc.Exec(context.Background(), sqlreq)

		if err != nil {
			log.Printf("Не удалось удалить базу данных с именем %s\n", dbName)
			log.Println(err)
			return
		}

	} else {
		log.Println(ErrNoConnection)
	}

}

// PostgreSQLDropRole - удаляет роль с заданным именем
func PostgreSQLDropRole(rolename string, dbc *pgxpool.Pool, locale string) {

	if dbc != nil {
		var rq int
		// Считаем количество ролей с заданным именем
		rows, err := dbc.Query(context.Background(), `SELECT COUNT(*) FROM pg_catalog.pg_roles WHERE	rolname = $1;`, rolename)

		shared.WriteErrToLog(err, locale)

		for rows.Next() {
			rows.Scan(&rq)
		}

		// Если ролей больше нуля, тогда ничего не делаем
		if rq <= 0 {
			log.Printf("Не найдена роль с именем %s\n", rolename)
			return
		}

		sqlreq := fmt.Sprintf(`DROP ROLE "%s";`, rolename)

		_, err = dbc.Exec(context.Background(), sqlreq)

		if err != nil {
			log.Printf("Не удалось удалить роль с именем %s\n", rolename)
			log.Println(err)
			return
		}

	} else {
		log.Println(ErrNoConnection)
	}
}

// PostgreSQLCreateSchemas - Создаём cхемы в базе данных
func PostgreSQLCreateSchemas(dbc *pgxpool.Pool, locale string) error {

	log.Println("Проверяем, что база пустая")

	// Проверяем что таблиц нет
	sqlreq := `SELECT 
				count(*)
			FROM 
				information_schema.tables
			WHERE 
				table_schema = ANY($1);`

	rows, err := dbc.Query(context.Background(), sqlreq, Schemas)

	shared.WriteErrToLog(err, locale)

	var tbq int

	for rows.Next() {
		err = rows.Scan(&tbq)
		shared.WriteErrToLog(err, locale)
	}

	if tbq > 0 {
		log.Println("В базе найдены таблицы, дубликаты не создаём")
		return ErrTablesAlreadyExist
	}

	// Создаём схемы
	for _, Schema := range Schemas {
		if Schema != "public" {
			PostgreSQLCreateSchema(Schema, dbc)
		}
	}

	return nil

}

// PostgreSQLCreateRole - создание отдельной роли для базы данных
func PostgreSQLCreateRole(roleName string, password string, dbName string, dbc *pgxpool.Pool, locale string) {

	rows, err := dbc.Query(context.Background(), `SELECT COUNT(*) FROM pg_catalog.pg_roles WHERE  rolname = $1`, roleName)

	shared.WriteErrToLog(err, locale)

	var rq int

	for rows.Next() {
		rows.Scan(&rq)
	}

	if rq > 0 {
		log.Printf("В базе данных найдена роль %s, дубликаты не создаём\n", roleName)
		return
	}

	log.Println("Создаём роль " + roleName)

	// Делаем MD5 хеш
	h := md5.New()
	io.WriteString(h, password+roleName)

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := fmt.Sprintf(`CREATE USER %s WITH LOGIN ENCRYPTED PASSWORD 'md5%x';`, roleName, h.Sum(nil))

	_, err = dbc.Exec(context.Background(), sqlreq)

	PostgreSQLRollbackIfError(err, true, dbc)

	sqlreq = fmt.Sprintf(`GRANT CONNECT ON DATABASE "%s" TO %s;`, dbName, roleName)

	_, err = dbc.Exec(context.Background(), sqlreq)

	PostgreSQLRollbackIfError(err, true, dbc)

	sqlreq = fmt.Sprintf(`GRANT USAGE ON SCHEMA %s TO %s;`, "public, secret", roleName)

	_, err = dbc.Exec(context.Background(), sqlreq)

	PostgreSQLRollbackIfError(err, true, dbc)

	sqlreq = fmt.Sprintf(`GRANT UPDATE, USAGE ON ALL SEQUENCES IN SCHEMA %s TO %s;`, "public, secret", roleName)

	_, err = dbc.Exec(context.Background(), sqlreq)

	PostgreSQLRollbackIfError(err, true, dbc)

	sqlreq = fmt.Sprintf(`REVOKE CREATE ON SCHEMA %s FROM %s;`, "public, secret", roleName)

	_, err = dbc.Exec(context.Background(), sqlreq)

	PostgreSQLRollbackIfError(err, true, dbc)

	dbc.Exec(context.Background(), "COMMIT")

	log.Println("Роль создана")

}

// PostgreSQLGrantRightsToRole - предоставляем права заданной роли для заданной таблицы
func PostgreSQLGrantRightsToRole(roleName string, tableName string, rights []string, dbc *pgxpool.Pool) {

	dbc.Exec(context.Background(), "BEGIN")

	reqrights := strings.Join(rights, ", ")

	log.Printf("Даём доступ %s к таблице %s c правами %s ", roleName, tableName, reqrights)

	sqlreq := fmt.Sprintf(`GRANT %s ON %s TO %s`, reqrights, tableName, roleName)

	_, err := dbc.Exec(context.Background(), sqlreq)

	PostgreSQLRollbackIfError(err, true, dbc)

	dbc.Exec(context.Background(), "COMMIT")

	log.Println("Права выданы")

}

// PostgreSQLCreateSchema - создаёт схему с заданным именем
//
// Параметры:
//
// SchemaName - имя схемы которую нужно создать
//
func PostgreSQLCreateSchema(SchemaName string, dbc *pgxpool.Pool) {

	log.Println("Создаём схему", SchemaName)

	sqlreq := fmt.Sprintf(`CREATE SCHEMA %v 
	AUTHORIZATION postgres;`, SchemaName)

	_, err := dbc.Exec(context.Background(), sqlreq)

	PostgreSQLRollbackIfError(err, true, dbc)

}

// PostgreSQLInsertStatus - записывает статус с заданным именем
//
// Параметры:
//
// StatusName - имя статуса который нужно создать
func PostgreSQLInsertStatus(StatusName string, dbc *pgxpool.Pool) {

	log.Println("Записываем статус", StatusName)

	sqlreq := `INSERT INTO "references".request_status(name) VALUES ($1);`

	_, err := dbc.Exec(context.Background(), sqlreq, StatusName)

	PostgreSQLRollbackIfError(err, true, dbc)

}

// PostgreSQLExecuteCreateStatement - выполняет sql запрос на создание таблицы и прекращает выполнение в случае ошибки
func PostgreSQLExecuteCreateStatement(dbc *pgxpool.Pool, ncs NamedCreateStatement) {

	log.Println("Создаём таблицу", ncs.TableName)

	_, err := dbc.Exec(context.Background(), ncs.CreateStatement)

	PostgreSQLRollbackIfError(err, true, dbc)

}
