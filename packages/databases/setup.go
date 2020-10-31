package databases

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"shopping-lists-and-recipes/packages/shared"
	"strings"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLDropDatabase - удаляет базу данных с заданным именем
func PostgreSQLDropDatabase(dbName string) {

	if dbc != nil {

		log.Println("Идёт удаление базы данных...")

		// Считаем количество баз данных с заданным именем
		rows, err := dbc.Query(`SELECT COUNT(datname) FROM pg_catalog.pg_database WHERE datname = $1;`, dbName)

		shared.WriteErrToLog(err)

		var dbq int

		for rows.Next() {
			rows.Scan(&dbq)
		}

		// Если баз данных больше нуля, тогда ничего не делаем
		if dbq <= 0 {
			log.Printf("Не найдена база данных с именем %s\n", dbName)
			return
		}

		_, err = dbc.Exec(`SELECT pg_terminate_backend(pg_stat_activity.pid)
							FROM pg_stat_activity
							WHERE pg_stat_activity.datname = $1
							AND pid <> pg_backend_pid();`, dbName)

		if err != nil {
			log.Println(err)
			return
		}

		sqlreq := fmt.Sprintf(`DROP DATABASE "%s";`, dbName)

		_, err = dbc.Exec(sqlreq)

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
func PostgreSQLDropRole(rolename string) {

	if dbc != nil {
		var rq int
		// Считаем количество ролей с заданным именем
		rows, err := dbc.Query(`SELECT COUNT(*) FROM pg_catalog.pg_roles WHERE	rolname = $1;`, rolename)

		shared.WriteErrToLog(err)

		for rows.Next() {
			rows.Scan(&rq)
		}

		// Если ролей больше нуля, тогда ничего не делаем
		if rq <= 0 {
			log.Printf("Не найдена роль с именем %s\n", rolename)
			return
		}

		sqlreq := fmt.Sprintf(`DROP ROLE "%s";`, rolename)

		_, err = dbc.Exec(sqlreq)

		if err != nil {
			log.Printf("Не удалось удалить роль с именем %s\n", rolename)
			log.Println(err)
			return
		}

	} else {
		log.Println(ErrNoConnection)
	}
}

// PostgreSQLCreateDatabase - создаём базу данных для СУБД PostgreSQL
func PostgreSQLCreateDatabase(dbName string) {

	if dbc != nil {
		log.Println("Идёт создание базы данных...")

		// Считаем количество баз данных с заданным именем
		rows, err := dbc.Query(`SELECT COUNT(datname) FROM pg_catalog.pg_database WHERE datname = $1;`, dbName)

		shared.WriteErrToLog(err)

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
									CONNECTION LIMIT = -1;`, dbName)

		_, err = dbc.Exec(sqlreq)

		shared.WriteErrToLog(err)

		log.Println("База данных успешно создана")
	}

}

// PostgreSQLCreateTables - Создаём таблицы в базе данных
func PostgreSQLCreateTables() error {

	log.Println("Проверяем, что база пустая")

	// Проверяем что таблиц нет
	sqlreq := `SELECT 
				count(*)
			FROM 
				information_schema.tables
			WHERE 
				table_schema = 'public';`

	rows, err := dbc.Query(sqlreq)

	shared.WriteErrToLog(err)

	var tbq int

	for rows.Next() {
		rows.Scan(&tbq)
	}

	if tbq > 0 {
		log.Println("В базе найдены таблицы, дубликаты не создаём")
		return ErrTablesAlreadyExist
	}

	log.Println("Начинаем создание таблиц")

	dbc.Exec("BEGIN")

	sqlreq = `CREATE TABLE public."Files"
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

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу Files")

	sqlreq = `CREATE TABLE public."Recipes"
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

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу Recipes")

	sqlreq = `ALTER TABLE public."Recipes"
				ADD CONSTRAINT "Recipes_image_id_fkey" FOREIGN KEY (image_id)
				REFERENCES public."Files" (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE RESTRICT;
			CREATE INDEX "fki_Recipes_image_id_fkey"
				ON public."Recipes"(image_id);`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = `CREATE TABLE public."RecipesIngredients"
			(
				recipe_id bigserial NOT NULL,
				ingredient_id bigint NOT NULL,
				quantity bigint NOT NULL
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."RecipesIngredients"
				OWNER to postgres;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу RecipesIngredients")

	sqlreq = `CREATE TABLE public."Ingredients"
			(
				id bigserial NOT NULL,
				name character varying(100) NOT NULL,
				CONSTRAINT "Ingredients_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;                  
			
			ALTER TABLE public."Ingredients"
				OWNER to postgres;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу Ingredients")

	sqlreq = `ALTER TABLE public."RecipesIngredients"
			ADD CONSTRAINT "RecipesIngredients_recipe_id_fkey" FOREIGN KEY (recipe_id)
			REFERENCES public."Recipes" (id) MATCH FULL
			ON UPDATE RESTRICT
			ON DELETE SET NULL;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = `ALTER TABLE public."RecipesIngredients"
			ADD CONSTRAINT "RecipesIngredients_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
			REFERENCES public."Ingredients" (id) MATCH FULL
			ON UPDATE RESTRICT
			ON DELETE SET NULL;
		CREATE INDEX "fki_RecipesIngredients_ingredient_id_fkey"
			ON public."RecipesIngredients"(ingredient_id);`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = `CREATE TABLE public."ShoppingList"
			(
				id bigserial NOT NULL,
				ingredient_id bigint NOT NULL,
				quantity bigint NOT NULL,
				CONSTRAINT "ShoppingList_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."ShoppingList"
				OWNER to postgres;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу ShoppingList")

	sqlreq = `ALTER TABLE public."ShoppingList"
				ADD CONSTRAINT "ShoppingList_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
				REFERENCES public."Ingredients" (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE SET NULL;
			CREATE INDEX "fki_ShoppingList_ingredient_id_fkey"
				ON public."ShoppingList"(ingredient_id);`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = `CREATE SCHEMA secret
			AUTHORIZATION postgres;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = `CREATE TABLE secret.users
			(
				id uuid NOT NULL,
				role character varying(50) NOT NULL,
				email character varying(50) NOT NULL,
				phone character varying(50),
				name character varying(150),
				isadmin boolean,
				confirmed boolean,
				totp_active boolean,
				PRIMARY KEY (id)
			);
			
			ALTER TABLE secret.users
				OWNER to postgres;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу users")

	sqlreq = `CREATE TABLE secret.hashes
		(
			id bigserial NOT NULL,
			user_id uuid NOT NULL,
			value bytea NOT NULL,
			PRIMARY KEY (id)
		);

		ALTER TABLE secret.hashes
			OWNER to postgres;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу hashes")

	sqlreq = `ALTER TABLE secret.hashes
				ADD CONSTRAINT hashes_user_id_fkey FOREIGN KEY (user_id)
				REFERENCES secret.users (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE RESTRICT;
			CREATE INDEX fki_hashes_user_id_fkey
				ON secret.hashes(user_id);`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = `CREATE TABLE secret.confirmations
			(
				user_id uuid,
				token character varying(200) COLLATE pg_catalog."default" NOT NULL,
				"Created" timestamp with time zone NOT NULL,
				"Expires" timestamp with time zone NOT NULL,
				CONSTRAINT confirmations_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			);
			ALTER TABLE secret.confirmations
				OWNER to postgres;
			CREATE INDEX fki_confirmations_user_id_fkey
				ON secret.confirmations USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу confirmations")

	sqlreq = `CREATE TABLE secret.password_resets
			(
				user_id uuid,
				token character varying(200) COLLATE pg_catalog."default" NOT NULL,
				"Created" timestamp with time zone NOT NULL,
				"Expires" timestamp with time zone NOT NULL,
				CONSTRAINT password_resets_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;

			ALTER TABLE secret.password_resets
				OWNER to postgres;

			CREATE INDEX fki_password_resets_user_id_fkey
				ON secret.password_resets USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу password_resets")

	sqlreq = `CREATE TABLE secret.totp
	(
		user_id uuid NOT NULL,
		secret text COLLATE pg_catalog."default",
		key bytea,
		confirmed boolean,
		CONSTRAINT user_id_fkey FOREIGN KEY (user_id)
			REFERENCES secret.users (id) MATCH FULL
			ON UPDATE RESTRICT
			ON DELETE RESTRICT
	)
	
	TABLESPACE pg_default;
	
	ALTER TABLE secret.totp
		OWNER to postgres;
	
	CREATE INDEX fki_user_id_fkey
		ON secret.totp USING btree
		(user_id ASC NULLS LAST)
		TABLESPACE pg_default;`

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу totp")

	dbc.Exec("COMMIT")

	log.Println("Таблицы созданы")

	return nil

}

// PostgreSQLCreateRole - создание отдельной роли для базы данных
func PostgreSQLCreateRole(roleName string, password string, dbName string) {

	rows, err := dbc.Query(`SELECT COUNT(*) FROM pg_catalog.pg_roles WHERE  rolname = $1`, roleName)

	shared.WriteErrToLog(err)

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

	dbc.Exec("BEGIN")

	sqlreq := fmt.Sprintf(`CREATE USER %s WITH LOGIN ENCRYPTED PASSWORD 'md5%x';`, roleName, h.Sum(nil))

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = fmt.Sprintf(`GRANT CONNECT ON DATABASE "%s" TO %s;`, dbName, roleName)

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = fmt.Sprintf(`GRANT USAGE ON SCHEMA %s TO %s;`, "public, secret", roleName)

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = fmt.Sprintf(`GRANT UPDATE, USAGE ON ALL SEQUENCES IN SCHEMA %s TO %s;`, "public, secret", roleName)

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	sqlreq = fmt.Sprintf(`REVOKE CREATE ON SCHEMA %s FROM %s;`, "public, secret", roleName)

	_, err = dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	dbc.Exec("COMMIT")

	log.Println("Роль создана")

}

// PostgreSQLGrantRightsToRole - предоставляем права заданной роли для заданной таблицы
func PostgreSQLGrantRightsToRole(roleName string, tableName string, rights []string) {

	dbc.Exec("BEGIN")

	reqrights := strings.Join(rights, ", ")

	log.Printf("Даём доступ %s к таблице %s c правами %s ", roleName, tableName, reqrights)

	sqlreq := fmt.Sprintf(`GRANT %s ON %s TO %s`, reqrights, tableName, roleName)

	_, err := dbc.Exec(sqlreq)

	PostgreSQLRollbackIfError(err, true)

	dbc.Exec("COMMIT")

	log.Println("Права выданы")

}
