package databases

import (
	"fmt"
	"math"
	"shopping-lists-and-recipes/packages/shared"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

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

	if PostgreSQLCheckLimitOffset(limit, offset) {

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
		return result, ErrLimitOffsetInvalid
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

	if PostgreSQLCheckLimitOffset(limit, offset) {

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
		return result, ErrLimitOffsetInvalid
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
