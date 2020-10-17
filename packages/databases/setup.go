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
		createsql := fmt.Sprintf(`CREATE DATABASE "%s"
									WITH
									OWNER = postgres
									ENCODING = 'UTF8'
									LC_COLLATE = 'C.UTF-8'
									LC_CTYPE = 'C.UTF-8'
									TABLESPACE = pg_default
									CONNECTION LIMIT = -1;`, dbName)

		_, err = dbc.Exec(createsql)

		shared.WriteErrToLog(err)

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

	shared.WriteErrToLog(err)

	var tbq int

	for rows.Next() {
		rows.Scan(&tbq)
	}

	if tbq > 0 {
		log.Println("В базе найдены таблицы, дубликаты не создаём")
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

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу Files")

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

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу Recipes")

	sql = `ALTER TABLE public."Recipes"
				ADD CONSTRAINT "Recipes_image_id_fkey" FOREIGN KEY (image_id)
				REFERENCES public."Files" (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE CASCADE;
			CREATE INDEX "fki_Recipes_image_id_fkey"
				ON public."Recipes"(image_id);`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

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

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу RecipesIngredients")

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

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу Ingredients")

	sql = `ALTER TABLE public."RecipesIngredients"
			ADD CONSTRAINT "RecipesIngredients_recipe_id_fkey" FOREIGN KEY (recipe_id)
			REFERENCES public."Recipes" (id) MATCH FULL
			ON UPDATE RESTRICT
			ON DELETE SET NULL;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	sql = `ALTER TABLE public."RecipesIngredients"
			ADD CONSTRAINT "RecipesIngredients_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
			REFERENCES public."Ingredients" (id) MATCH FULL
			ON UPDATE RESTRICT
			ON DELETE SET NULL;
		CREATE INDEX "fki_RecipesIngredients_ingredient_id_fkey"
			ON public."RecipesIngredients"(ingredient_id);`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

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

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу ShoppingList")

	sql = `ALTER TABLE public."ShoppingList"
				ADD CONSTRAINT "ShoppingList_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
				REFERENCES public."Ingredients" (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE SET NULL;
			CREATE INDEX "fki_ShoppingList_ingredient_id_fkey"
				ON public."ShoppingList"(ingredient_id);`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	sql = `CREATE SCHEMA secret
			AUTHORIZATION postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	sql = `CREATE TABLE secret.users
			(
				id uuid NOT NULL,
				role character varying(50) NOT NULL,
				email character varying(50) NOT NULL,
				phone character varying(50),
				name character varying(150),
				isadmin boolean,
				confirmed boolean,
				PRIMARY KEY (id)
			);
			
			ALTER TABLE secret.users
				OWNER to postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу users")

	sql = `CREATE TABLE secret.hashes
		(
			id bigserial NOT NULL,
			user_id uuid NOT NULL,
			value bytea NOT NULL,
			PRIMARY KEY (id)
		);

		ALTER TABLE secret.hashes
			OWNER to postgres;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу hashes")

	sql = `ALTER TABLE secret.hashes
				ADD CONSTRAINT hashes_user_id_fkey FOREIGN KEY (user_id)
				REFERENCES secret.users (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE CASCADE;
			CREATE INDEX fki_hashes_user_id_fkey
				ON secret.hashes(user_id);`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	sql = `CREATE TABLE secret.confirmations
			(
				user_id uuid,
				token character varying(200) COLLATE pg_catalog."default" NOT NULL,
				"Created" timestamp with time zone NOT NULL,
				"Expires" timestamp with time zone NOT NULL,
				CONSTRAINT confirmations_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE NO ACTION
					ON DELETE RESTRICT
			);
			ALTER TABLE secret.confirmations
				OWNER to postgres;
			CREATE INDEX fki_confirmations_user_id_fkey
				ON secret.confirmations USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу confirmations")

	sql = `CREATE TABLE secret.password_resets
			(
				user_id uuid,
				token character varying(200) COLLATE pg_catalog."default" NOT NULL,
				"Created" timestamp with time zone NOT NULL,
				"Expires" timestamp with time zone NOT NULL,
				CONSTRAINT password_resets_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE NO ACTION
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;

			ALTER TABLE secret.password_resets
				OWNER to postgres;

			CREATE INDEX fki_password_resets_user_id_fkey
				ON secret.password_resets USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`

	_, err = dbc.Exec(sql)

	PostgreSQLRollbackIfError(err, true)

	log.Println("Создали таблицу password_resets")

	dbc.Exec("COMMIT")

	log.Println("Таблицы созданы")

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

	rolecreatesql := fmt.Sprintf(`CREATE USER %s WITH LOGIN ENCRYPTED PASSWORD 'md5%x';`, roleName, h.Sum(nil))

	_, err = dbc.Exec(rolecreatesql)

	PostgreSQLRollbackIfError(err, true)

	grantsql := fmt.Sprintf(`GRANT CONNECT ON DATABASE "%s" TO %s;`, dbName, roleName)

	_, err = dbc.Exec(grantsql)

	PostgreSQLRollbackIfError(err, true)

	grantsql = fmt.Sprintf(`GRANT USAGE ON SCHEMA %s TO %s;`, "public, secret", roleName)

	_, err = dbc.Exec(grantsql)

	PostgreSQLRollbackIfError(err, true)

	grantsql = fmt.Sprintf(`GRANT USAGE ON ALL SEQUENCES IN SCHEMA %s TO %s;`, "public, secret", roleName)

	_, err = dbc.Exec(grantsql)

	PostgreSQLRollbackIfError(err, true)

	grantsql = fmt.Sprintf(`REVOKE CREATE ON SCHEMA %s FROM %s;`, "public, secret", roleName)

	_, err = dbc.Exec(grantsql)

	PostgreSQLRollbackIfError(err, true)

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

	PostgreSQLRollbackIfError(err, true)

	dbc.Exec("COMMIT")

	log.Println("Права выданы")

}
