// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var dbc *sql.DB

// Список типовых ошибок
var (
	ErrFirstNotDelete       = errors.New("Первая запись в списке файлов техническая и не подлежит удалению")
	ErrRecipeNotFound       = errors.New("В таблице рецептов не найден указанный id")
	ErrShoppingListNotFound = errors.New("Не найдено ни одной записи в списке покупок с указанным названием")
	ErrEmptyPassword        = errors.New("Не допустимо использование паролей с длинной менее шести символов")
	ErrNoUserWithEmail      = errors.New("Электронная почта не найдена")
	ErrNoHashForUser        = errors.New("Хеш пароля не найден")
	ErrEmailIsOccupied      = errors.New("Указанный адрес электронной почты уже занят")
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

// PostgreSQLFileInsert - создаёт записи в базе данных для хранения информации о загруженном файле
func PostgreSQLFileInsert(filename string, filesize int64, filetype string, fileid string) (int, error) {

	dbc.Exec("BEGIN")

	sql := `INSERT INTO 
			public."Files" 
			(filename, filesize, filetype, file_id) 
		  VALUES 
			($1, $2, $3, $4) RETURNING id;`

	row := dbc.QueryRow(sql, filename, filesize, filetype, fileid)

	var curid int
	err := row.Scan(&curid)

	if err != nil {
		return curid, PostgreSQLRollbackIfError(err, true)
	}

	log.Printf("Данные о файле сохранены в базу данных под индексом %v", curid)

	dbc.Exec("COMMIT")

	return curid, nil
}

// PostgreSQLFileUpdate - обновляет записи в базе данных для хранения информации о загруженном файле
func PostgreSQLFileUpdate(filename string, filesize int64, filetype string, fileid string, id string) (int, error) {

	dbc.Exec("BEGIN")

	sql := `UPDATE
				public."Files" 
			SET 
				(filename, filesize, filetype, file_id) = ($1, $2, $3, $4) 
			WHERE 
				id = $5
			RETURNING id;`

	row := dbc.QueryRow(sql, filename, filesize, filetype, fileid, id)

	var curid int
	err := row.Scan(&curid)

	log.Printf("Данные о файле сохранены в базу данных под индексом %v", curid)

	if err != nil {
		return 1, PostgreSQLRollbackIfError(err, true)
	}

	dbc.Exec("COMMIT")

	return curid, nil
}

// PostgreSQLFileDelete - удаляет запись в базе данных о загруженном файле
func PostgreSQLFileDelete(fileid int) error {

	if fileid == 1 {
		return ErrFirstNotDelete
	}

	dbc.Exec("BEGIN")

	sql := `SELECT 
				file_id
			FROM 
				public."Files" 
			WHERE id=$1;`

	row := dbc.QueryRow(sql, fileid)

	var filename string
	err := row.Scan(&filename)

	if err != nil {
		return err
	}

	sql = `DELETE FROM 
				public."Files" 
			WHERE id=$1;`

	_, err = dbc.Exec(sql, fileid)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	path := strings.Join([]string{".", "public", "uploads", filename}, "/")
	err = os.Remove(path)

	if err != nil {
		return err
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLFilesSelect - получает информацию о файлах
func PostgreSQLFilesSelect(offset int, limit int) (FilesResponse, error) {

	var result FilesResponse

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Files"`

	row := dbc.QueryRow(sql)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if offset > 0 && limit > 0 {
		sql = fmt.Sprintf(`SELECT 
							"Files".id,
							"Files".filename,
							"Files".filesize,
							"Files".filetype,
							"Files".file_id
						FROM 
							public."Files"
						ORDER BY "Files".id
						OFFSET %v LIMIT %v`, offset, limit)
	} else {
		offset = 0
		limit = 0
		sql = fmt.Sprintln(`SELECT 
							"Files".id,
							"Files".filename,
							"Files".filesize,
							"Files".filetype,
							"Files".file_id
						FROM 
							public."Files"
						ORDER BY "Files".id`)
	}

	rows, err := dbc.Query(sql)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur FileDB
		rows.Scan(&cur.ID, &cur.Filename, &cur.Filesize, &cur.Filetype, &cur.FileID)
		result.Files = append(result.Files, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLRecipesSelect - получает информацию о рецептах и связанном файле обложки
func PostgreSQLRecipesSelect(page int, limit int) (RecipesResponse, error) {

	var result RecipesResponse
	result.Recipes = RecipesDB{}

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"`

	row := dbc.QueryRow(sql)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if offset >= 0 && limit > 0 {

		sql = fmt.Sprintf(`SELECT 
							"Recipes".id, 
							"Recipes".name, 
							"Recipes".description,
							"Recipes".image_id,
							"Files".file_id	
						FROM 
							public."Recipes"
							LEFT JOIN 
							public."Files"
							ON "Recipes".image_id="Files".id
						ORDER BY "Recipes".id
						OFFSET %v LIMIT %v`, offset, limit)

	} else {
		offset = 0
		limit = 0
		sql = fmt.Sprintln(`SELECT 
							"Recipes".id, 
							"Recipes".name, 
							"Recipes".description,
							"Recipes".image_id,
							"Files".file_id	
						FROM 
							public."Recipes"
							LEFT JOIN 
							public."Files"
							ON "Recipes".image_id="Files".id
						ORDER BY "Recipes".id`)
	}
	rows, err := dbc.Query(sql)

	if err != nil {
		return result, err
	}

	for rows.Next() {

		var cur RecipeDB
		cur.Ingredients = IngredientsDB{}
		rows.Scan(&cur.ID, &cur.Name, &cur.Description, &cur.ImageDbID, &cur.ImagePath)
		if cur.ImagePath != "" {
			cur.ImagePath = "/uploads/" + cur.ImagePath
		} else {
			cur.ImageDbID = 1
		}

		sql = `SELECT 	
				Ing.name,
				RecIng.quantity
			FROM 
				(SELECT 
					recipe_id,
					ingredient_id,
					quantity
				FROM 
					public."RecipesIngredients"
				WHERE
					recipe_id=$1) AS RecIng
			LEFT JOIN 
				public."Ingredients" AS Ing
			ON Ing.id = RecIng.ingredient_id`

		ings, err := dbc.Query(sql, cur.ID)

		if err != nil {
			return result, err
		}

		for ings.Next() {
			var ing IngredientDB
			ings.Scan(&ing.Name, &ing.Amount)
			cur.Ingredients = append(cur.Ingredients, ing)
		}

		result.Recipes = append(result.Recipes, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLRecipesSelectSearch - получает информацию о рецептах и связанном файле обложки для поискового запроса
func PostgreSQLRecipesSelectSearch(page int, limit int, search string) (RecipesResponse, error) {

	var result RecipesResponse
	result.Recipes = RecipesDB{}

	search = "%" + search + "%"

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"
			WHERE 
				"Recipes".name LIKE $1
				OR "Recipes".description LIKE $1`

	row := dbc.QueryRow(sql, search)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if offset >= 0 && limit > 0 {

		sql = fmt.Sprintf(`SELECT 
							"Recipes".id, 
							"Recipes".name, 
							"Recipes".description,
							"Recipes".image_id,
							"Files".file_id	
						FROM 
							public."Recipes"
							LEFT JOIN 
							public."Files"
							ON "Recipes".image_id="Files".id
						WHERE 
							"Recipes".name LIKE $1
							OR "Recipes".description LIKE $1
						ORDER BY "Recipes".id
						OFFSET %v LIMIT %v`, offset, limit)

	} else {
		offset = 0
		limit = 0
		sql = fmt.Sprintln(`SELECT 
							"Recipes".id, 
							"Recipes".name, 
							"Recipes".description,
							"Recipes".image_id,
							"Files".file_id	
						FROM 
							public."Recipes"
							LEFT JOIN 
							public."Files"
							ON "Recipes".image_id="Files".id
						WHERE 
							"Recipes".name LIKE $1
							OR "Recipes".description LIKE $1
						ORDER BY "Recipes".id`)
	}
	rows, err := dbc.Query(sql, search)

	if err != nil {
		return result, err
	}

	for rows.Next() {

		var cur RecipeDB
		cur.Ingredients = IngredientsDB{}
		rows.Scan(&cur.ID, &cur.Name, &cur.Description, &cur.ImageDbID, &cur.ImagePath)
		if cur.ImagePath != "" {
			cur.ImagePath = "/uploads/" + cur.ImagePath
		} else {
			cur.ImageDbID = 1
		}

		sql = `SELECT 	
				Ing.name,
				RecIng.quantity
			FROM 
				(SELECT 
					recipe_id,
					ingredient_id,
					quantity
				FROM 
					public."RecipesIngredients"
				WHERE
					recipe_id=$1) AS RecIng
			LEFT JOIN 
				public."Ingredients" AS Ing
			ON Ing.id = RecIng.ingredient_id`

		ings, err := dbc.Query(sql, cur.ID)

		if err != nil {
			return result, err
		}

		for ings.Next() {
			var ing IngredientDB
			ings.Scan(&ing.Name, &ing.Amount)
			cur.Ingredients = append(cur.Ingredients, ing)
		}

		result.Recipes = append(result.Recipes, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLRecipesInsertUpdate - обновляет существующий рецепт или вставляет новый рецепт в базу данных
func PostgreSQLRecipesInsertUpdate(RecipeUpd RecipeDB) (RecipeDB, error) {

	if RecipeUpd.ImageDbID == 0 {
		RecipeUpd.ImageDbID = 1
	}

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"
			WHERE 
				id=$1;`

	row := dbc.QueryRow(sql, RecipeUpd.ID)

	var recipecount int
	err := row.Scan(&recipecount)

	if err != nil {
		return RecipeUpd, err
	}

	dbc.Exec("BEGIN")

	if recipecount > 0 && RecipeUpd.ID != 0 {
		// Если запись найдена по индексу и индекс не равен нулю (случай новой записи)
		// Обновляем существующую запись
		sql = `UPDATE 
					public."Recipes" 
				SET 
					(name, description, image_id) = ($1, $2, $3) 
				WHERE 
					id=$4;`

		_, err = dbc.Exec(sql, RecipeUpd.Name, RecipeUpd.Description, RecipeUpd.ImageDbID, RecipeUpd.ID)

	} else {
		// Иначе вставляем новую запись
		sql = `INSERT INTO public."Recipes" (name, description, image_id) VALUES ($1, $2, $3) RETURNING id;`

		row := dbc.QueryRow(sql, RecipeUpd.Name, RecipeUpd.Description, RecipeUpd.ImageDbID)

		err = row.Scan(&RecipeUpd.ID)
	}

	if err != nil {
		return RecipeUpd, PostgreSQLRollbackIfError(err, false)
	}

	sql = `DELETE FROM public."RecipesIngredients" WHERE recipe_id=$1;`

	_, err = dbc.Exec(sql, RecipeUpd.ID)

	if err != nil {
		return RecipeUpd, PostgreSQLRollbackIfError(err, false)
	}

	for _, OneRecipe := range RecipeUpd.Ingredients {

		sql = `SELECT 
					COUNT(*)
				FROM 
					public."Ingredients" 
				WHERE 
					name = $1
				LIMIT 1;`

		row := dbc.QueryRow(sql, OneRecipe.Name)

		var count int
		err := row.Scan(&count)

		if err != nil {
			return RecipeUpd, PostgreSQLRollbackIfError(err, false)
		}

		var curid int

		if count > 0 {

			sql = `SELECT 
						id						 
					FROM 
						public."Ingredients" 
					WHERE 
						name = $1
					LIMIT 1;`

			row := dbc.QueryRow(sql, OneRecipe.Name)

			err := row.Scan(&curid)

			if err != nil {
				return RecipeUpd, PostgreSQLRollbackIfError(err, false)
			}

		} else {
			sql = `INSERT INTO public."Ingredients" (name) VALUES ($1) RETURNING id;`

			row := dbc.QueryRow(sql, OneRecipe.Name)

			err := row.Scan(&curid)

			if err != nil {
				return RecipeUpd, PostgreSQLRollbackIfError(err, false)
			}
		}

		sql = `INSERT INTO public."RecipesIngredients" (recipe_id, ingredient_id, quantity) VALUES ($1,$2,$3);`

		_, err = dbc.Exec(sql, RecipeUpd.ID, curid, OneRecipe.Amount)

		if err != nil {
			return RecipeUpd, PostgreSQLRollbackIfError(err, false)
		}

	}

	dbc.Exec("COMMIT")

	return RecipeUpd, nil

}

// PostgreSQLRecipesDelete - удаляет рецепт и связанные с ним ингредиенты по индексу рецепта
func PostgreSQLRecipesDelete(ID int) error {

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"
			WHERE 
				id=$1;`

	row := dbc.QueryRow(sql, ID)

	var recipecount int
	err := row.Scan(&recipecount)

	shared.WriteErrToLog(err)

	if recipecount <= 0 {
		return ErrRecipeNotFound
	}

	dbc.Exec("BEGIN")

	sql = `DELETE FROM public."RecipesIngredients" WHERE recipe_id=$1;`

	_, err = dbc.Exec(sql, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	sql = `DELETE FROM public."Recipes" WHERE id=$1;`

	_, err = dbc.Exec(sql, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLShoppingListSelect - получает информацию о списке покупок
func PostgreSQLShoppingListSelect(page int, limit int) (ShoppingListResponse, error) {

	var result ShoppingListResponse

	result.Items = IngredientsDB{}

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."ShoppingList"`

	row := dbc.QueryRow(sql)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if offset >= 0 && limit > 0 {
		sql = fmt.Sprintf(`SELECT 
								"Ingredients"."name",
								"ShoppingList".quantity 
							FROM 
								public."ShoppingList"
								LEFT JOIN
								public."Ingredients"
								ON "ShoppingList".ingredient_id = "Ingredients".id
							ORDER BY
								"Ingredients"."name"							
							OFFSET %v LIMIT %v;`, offset, limit)
	} else {
		offset = 0
		limit = 0
		sql = fmt.Sprintln(`SELECT 
								"Ingredients"."name",
								"ShoppingList".quantity 
							FROM 
								public."ShoppingList"
								LEFT JOIN
								public."Ingredients"
								ON "ShoppingList".ingredient_id = "Ingredients".id
							ORDER BY
								"Ingredients"."name"`)
	}

	rows, err := dbc.Query(sql)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur IngredientDB
		rows.Scan(&cur.Name, &cur.Amount)
		result.Items = append(result.Items, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLShoppingListInsertUpdate - обновляет существующую запись в списке покупок или вставляет новую в базу данных
func PostgreSQLShoppingListInsertUpdate(ShoppingItem IngredientDB) error {

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Ingredients"
			WHERE 
				"Ingredients".name=$1;`

	row := dbc.QueryRow(sql, ShoppingItem.Name)

	var IngCount int

	err := row.Scan(&IngCount)

	if err != nil {
		return err
	}

	var countRows int

	var ingID int

	dbc.Exec("BEGIN")

	if IngCount > 0 {

		sql := `SELECT 
					id
				FROM 
					public."Ingredients"
				WHERE 
					"Ingredients".name=$1;`

		ingrow := dbc.QueryRow(sql, ShoppingItem.Name)

		err = ingrow.Scan(&ingID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		sql = `SELECT 
					COUNT(*)
				FROM
					public."ShoppingList"
				WHERE ingredient_id =$1;`

		row := dbc.QueryRow(sql, ingID)

		err := row.Scan(&countRows)

		if err != nil {
			return err
		}

		if countRows == 0 {
			// Добавляем новую
			sql := `INSERT INTO public."ShoppingList" (ingredient_id, quantity) VALUES ($1,$2);`

			_, err = dbc.Exec(sql, ingID, ShoppingItem.Amount)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		} else {
			// Обновляем существующую
			sql = `UPDATE public."ShoppingList" SET quantity = $1 WHERE ingredient_id=$2;`

			_, err = dbc.Exec(sql, ShoppingItem.Amount, ingID)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		}
	} else {
		// Добавляем ингредиент в справочник
		sql = `INSERT INTO public."Ingredients" (name) VALUES ($1) RETURNING id;`

		ingrow := dbc.QueryRow(sql, ShoppingItem.Name)

		err := ingrow.Scan(&ingID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		sql = `SELECT 
					COUNT(*)
				FROM
					public."ShoppingList"
				WHERE ingredient_id =$1;`

		row := dbc.QueryRow(sql, ingID)

		err = row.Scan(&countRows)

		if err != nil {
			return err
		}

		if countRows == 0 {
			// Добавляем новую запись в список покупок
			sql = `INSERT INTO public."ShoppingList" (ingredient_id, quantity) VALUES ($1,$2);`

			_, err = dbc.Exec(sql, ingID, ShoppingItem.Amount)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		} else {
			// Обновляем существующую
			sql = `UPDATE public."ShoppingList" SET quantity = $1 WHERE ingredient_id=$2;`

			_, err = dbc.Exec(sql, ShoppingItem.Amount, ingID)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		}
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLShoppingListDelete - удаляет запись из списка покупок по имени
func PostgreSQLShoppingListDelete(IngName string) error {

	sql := `SELECT 
				COUNT(*)
			FROM 
				public."Ingredients"
			WHERE 
				"Ingredients".name=$1;`

	row := dbc.QueryRow(sql, IngName)

	var IngCount int

	err := row.Scan(&IngCount)

	if err != nil {
		return err
	}

	var ingID int
	var CountRows int

	if IngCount > 0 {

		sql := `SELECT 
					id
				FROM 
					public."Ingredients"
				WHERE 
					"Ingredients".name=$1;`

		ingrow := dbc.QueryRow(sql, IngName)

		err = ingrow.Scan(&ingID)

		if err != nil {
			return err
		}

		sql = `SELECT 
					COUNT(*)
				FROM 
					public."ShoppingList"
				WHERE 
					ingredient_id=$1;`

		slrow := dbc.QueryRow(sql, ingID)

		err = slrow.Scan(&CountRows)

		if err != nil {
			return err
		}

		if CountRows <= 0 {
			return ErrShoppingListNotFound
		}

		dbc.Exec("BEGIN")

		sql = `DELETE FROM public."ShoppingList" WHERE ingredient_id=$1;`
		_, err = dbc.Exec(sql, ingID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		dbc.Exec("COMMIT")

	} else {
		return ErrShoppingListNotFound
	}

	return nil

}

// PostgreSQLGetTokenForUser - получает токен для проверки при авторизации
func PostgreSQLGetTokenForUser(email string) (string, string, error) {

	var UserCount int
	var UserID uuid.UUID
	var HashesCount int
	var UserRole string

	var Hash string

	sql := `SELECT 
				COUNT(*) 
			FROM 
				secret.users 
			WHERE 
				email=$1;`

	UserCountRow := dbc.QueryRow(sql, email)

	err := UserCountRow.Scan(&UserCount)

	if err != nil {
		return "", "", err
	}

	if UserCount <= 0 {
		return "", "", ErrNoUserWithEmail
	}

	sql = `SELECT 
				id,
				role
			FROM 
				secret.users
			WHERE
				email=$1 
			LIMIT 1`

	UserIDRow := dbc.QueryRow(sql, email)

	err = UserIDRow.Scan(&UserID, &UserRole)

	if err != nil {
		return "", "", err
	}

	sql = `SELECT 
				COUNT(*)
			FROM 
				secret.hashes
			WHERE 
				user_id = $1;`

	HashesRow := dbc.QueryRow(sql, UserID)

	err = HashesRow.Scan(&HashesCount)

	if err != nil {
		return "", "", err
	}

	if HashesCount <= 0 {
		return "", "", ErrNoHashForUser
	}

	sql = `SELECT 
				value
			FROM 
				secret.hashes
			WHERE 
				user_id = $1;`

	HashValueRow := dbc.QueryRow(sql, UserID)

	err = HashValueRow.Scan(&Hash)

	if err != nil {
		return "", "", err
	}

	return Hash, UserRole, nil
}

// PostgreSQLUsersCreateUpdate - создаёт или обновляет существующего пользователя
func PostgreSQLUsersCreateUpdate(NewUserInfo UserDB, Hash string, UpdatePassword bool, OverWrite bool) error {

	// Проверяем что почта уникальна
	var EmailCount int

	sql := `SELECT COUNT(*) FROM secret.users WHERE email=$1;`

	EmailCountRow := dbc.QueryRow(sql, NewUserInfo.Email)

	err := EmailCountRow.Scan(&EmailCount)

	if err != nil {
		return err
	}

	if EmailCount > 0 && !OverWrite {
		return ErrEmailIsOccupied
	}

	// Проверяем что пользователь с ID существует
	var UserCount int

	sql = `SELECT COUNT(*) FROM secret.users WHERE id=$1;`

	UserCountRow := dbc.QueryRow(sql, NewUserInfo.GUID)

	err = UserCountRow.Scan(&UserCount)

	if err != nil {
		return err
	}

	dbc.Exec("BEGIN")

	if UserCount > 0 && OverWrite {

		// Обновляем существующего

		sql = `UPDATE secret.users SET (role, email, phone, name, isadmin) = ($1,$2,$3,$4,$5) WHERE id=$6;`

		_, err = dbc.Exec(sql, NewUserInfo.Role, NewUserInfo.Email, NewUserInfo.Phone, NewUserInfo.Name, NewUserInfo.IsAdmin, NewUserInfo.GUID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		if UpdatePassword {
			if len(Hash) > 0 {

				sql = `UPDATE secret.hashes SET value=$2 WHERE user_id=$1;`

				_, err = dbc.Exec(sql, NewUserInfo.GUID, Hash)

				if err != nil {
					return PostgreSQLRollbackIfError(err, false)
				}
			} else {
				return PostgreSQLRollbackIfError(ErrEmptyPassword, false)
			}
		}

	} else {

		// Создаём нового

		// Генерируем новый уникальный идентификатор
		NewUserInfo.GUID, err = uuid.NewV4()
		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		sql = `INSERT INTO secret.users (id, role, email, phone, name, isadmin) VALUES ($1,$2,$3,$4,$5,$6);`

		_, err = dbc.Exec(sql, NewUserInfo.GUID, NewUserInfo.Role, NewUserInfo.Email, NewUserInfo.Phone, NewUserInfo.Name, NewUserInfo.IsAdmin)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		if len(Hash) > 0 {

			sql = `INSERT INTO secret.hashes (user_id, value) VALUES ($1,$2);`

			_, err = dbc.Exec(sql, NewUserInfo.GUID, Hash)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		} else {
			return PostgreSQLRollbackIfError(ErrEmptyPassword, false)
		}

	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLUsersSelect - получает список пользователей в админке
func PostgreSQLUsersSelect(page int, limit int) (UsersResponse, error) {

	var result UsersResponse
	result.Users = UsersDB{}

	sql := `SELECT 
				COUNT(*) 
			FROM 
				secret.users;`

	row := dbc.QueryRow(sql)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if offset >= 0 && limit > 0 {

		sql = fmt.Sprintf(`SELECT 
								users.id,
								users.role,
								users.email,
								users.phone,
								users.name,
								users.isadmin,
								users.confirmed
							FROM 
								secret.users
							ORDER BY 
								email
							OFFSET %v LIMIT %v`, offset, limit)

	} else {
		offset = 0
		limit = 0
		sql = fmt.Sprintln(`SELECT 
								users.id,
								users.role,
								users.email,
								users.phone,
								users.name,
								users.isadmin,
								users.confirmed
							FROM 
								secret.users
							ORDER BY 
								email`)
	}
	rows, err := dbc.Query(sql)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur UserDB
		rows.Scan(&cur.GUID, &cur.Role, &cur.Email, &cur.Phone, &cur.Name, &cur.IsAdmin, &cur.Confirmed)
		result.Users = append(result.Users, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLShoppingListDeleteAll - удаляет все записи из списка покупок
func PostgreSQLShoppingListDeleteAll() error {

	dbc.Exec("BEGIN")

	sql := `DELETE FROM public."ShoppingList";`

	_, err := dbc.Exec(sql)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
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
