// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
	"log"
	"strings"
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

// PostgreSQLCreateDatabase - создаём базу данных для СУБД PostgreSQL
func PostgreSQLCreateDatabase(dbName string) {

	if dbc != nil {
		log.Println("Идёт создание базы данных...")

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
		createsql := fmt.Sprintf(`CREATE DATABASE "%s"
									WITH
									OWNER = postgres
									ENCODING = 'UTF8'
									LC_COLLATE = 'C.UTF-8'
									LC_CTYPE = 'C.UTF-8'
									TABLESPACE = pg_default
									CONNECTION LIMIT = -1;`, dbName)

		_, err = dbc.Exec(createsql)

		WriteErrToLog(err)

		log.Println("База данных успешно создана")
	}

}

// PostgreSQLCreateTables - Создаём таблицы в базе данных
func PostgreSQLCreateTables() {

	log.Println("Проверяем, что база пустая")

	// Проверяем что таблиц нет
	sql := `SELECT 
				count(*)
			FROM 
				information_schema.tables
			WHERE 
				table_schema = 'public';`

	rows, err := dbc.Query(sql)

	WriteErrToLog(err)

	var tbq int

	for rows.Next() {
		rows.Scan(&tbq)
	}

	if tbq > 0 {
		return
	}

	log.Println("Начинаем создание таблиц")

	dbc.Exec("BEGIN")

	sql = `CREATE TABLE public."Files"
			(
				id bigserial NOT NULL,
				filename character varying(255),
				filesize bigint,
				filetype character varying(50),
				file_id character varying(50),
				CONSTRAINT "Files_pkey" PRIMARY KEY (id)
			)

			TABLESPACE pg_default;

			ALTER TABLE public."Files"
				OWNER to postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `CREATE TABLE public."Recipes"
			(
				id bigserial NOT NULL,
				name character varying(100),
				description text,
				image_id bigint,
				CONSTRAINT "Recipes_pkey" PRIMARY KEY (id)
			)            

			TABLESPACE pg_default; 

			ALTER TABLE public."Recipes"
				OWNER to postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `ALTER TABLE public."Recipes"
				ADD CONSTRAINT "Recipes_image_id_fkey" FOREIGN KEY (image_id)
				REFERENCES public."Files" (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE SET NULL;
			CREATE INDEX "fki_Recipes_image_id_fkey"
				ON public."Recipes"(image_id);`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `CREATE TABLE public."RecipesIngredients"
			(
				recipe_id bigserial NOT NULL,
				ingredient_id bigint NOT NULL,
				quantity bigint NOT NULL
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."RecipesIngredients"
				OWNER to postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `CREATE TABLE public."Ingredients"
			(
				id bigserial NOT NULL,
				name character varying(100) NOT NULL,
				CONSTRAINT "Ingredients_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;                  
			
			ALTER TABLE public."Ingredients"
				OWNER to postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `ALTER TABLE public."RecipesIngredients"
			ADD CONSTRAINT "RecipesIngredients_recipe_id_fkey" FOREIGN KEY (recipe_id)
			REFERENCES public."Recipes" (id) MATCH FULL
			ON UPDATE RESTRICT
			ON DELETE SET NULL;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `ALTER TABLE public."RecipesIngredients"
			ADD CONSTRAINT "RecipesIngredients_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
			REFERENCES public."Ingredients" (id) MATCH FULL
			ON UPDATE RESTRICT
			ON DELETE SET NULL;
		CREATE INDEX "fki_RecipesIngredients_ingredient_id_fkey"
			ON public."RecipesIngredients"(ingredient_id);`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `CREATE TABLE public."ShoppingList"
			(
				id bigserial NOT NULL,
				ingredient_id bigint NOT NULL,
				quantity bigint NOT NULL,
				CONSTRAINT "ShoppingList_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."ShoppingList"
				OWNER to postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	sql = `ALTER TABLE public."ShoppingList"
				ADD CONSTRAINT "ShoppingList_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
				REFERENCES public."Ingredients" (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE SET NULL;
			CREATE INDEX "fki_ShoppingList_ingredient_id_fkey"
				ON public."ShoppingList"(ingredient_id);`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err)

	dbc.Exec("COMMIT")

}

// PostgreSQLCreateRole - Создание отдельной роли для базы данных
func PostgreSQLCreateRole(roleName string, password string, dbName string) {

	log.Println("Создаём роль " + roleName)

	// Делаем MD5 хеш
	h := md5.New()
	io.WriteString(h, password)

	dbc.Exec("BEGIN")

	rolecreatesql := fmt.Sprintf(`CREATE ROLE %s WITH LOGIN ENCRYPTED PASSWORD 'md5%x';`, roleName, h.Sum(nil))

	_, err := dbc.Exec(rolecreatesql)

	PostgreSQLRollbackIfError(err)

	grantsql := fmt.Sprintf(`GRANT CONNECT ON DATABASE "%s" TO %s;`, dbName, roleName)

	_, err = dbc.Exec(grantsql)

	PostgreSQLRollbackIfError(err)

	dbc.Exec("COMMIT")

	log.Println("Роль создана")

}

// PostgreSQLGrantRightsToRole - предоставляем права заданной роли для заданной таблицы
func PostgreSQLGrantRightsToRole(roleName string, tableName string, rights []string) {

	dbc.Exec("BEGIN")

	reqrights := strings.Join(rights, ", ")

	log.Printf("Даём доступ %s к таблице %s c правами %s ", roleName, tableName, reqrights)

	grantsql := fmt.Sprintf(`GRANT %s ON %s TO %s`, reqrights, tableName, roleName)

	_, err := dbc.Exec(grantsql)

	PostgreSQLRollbackIfError(err)

	dbc.Exec("COMMIT")

	log.Println("Права выданы")

}

// PostgreSQLRollbackIfError - откатываем изменения транзакции если происходит ошибка и пишем её в лог и выходим
func PostgreSQLRollbackIfError(err error) {
	if err != nil {
		dbc.Exec("ROLLBACK")
		WriteErrToLog(err)
	}
}

// PostgreSQLCloseConn - Закрываем соединение с базой данных
func PostgreSQLCloseConn() {
	dbc.Close()
}

// PostgreSQLConnect - Подключаемся к базе данных
func PostgreSQLConnect(ConnString string) {
	dbc = SQLConnect("postgres", ConnString)
}
