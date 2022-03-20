// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"
	"fmt"
	"math"
	"shopping-lists-and-recipes/packages/shared"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLRecipesSelect - получает информацию о рецептах и связанном файле обложки
func PostgreSQLRecipesSelect(page int, limit int, dbc *pgxpool.Pool) (RecipesResponse, error) {

	var result RecipesResponse
	result.Recipes = RecipesDB{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"`

	row := dbc.QueryRow(context.Background(), sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {

		sqlreq = fmt.Sprintf(`SELECT 
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

	rows, err := dbc.Query(context.Background(), sqlreq)

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

		sqlreq = `SELECT 	
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

		ings, err := dbc.Query(context.Background(), sqlreq, cur.ID)

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
func PostgreSQLRecipesSelectSearch(page int, limit int, search string, dbc *pgxpool.Pool) (RecipesResponse, error) {

	var result RecipesResponse
	result.Recipes = RecipesDB{}

	search = "%" + search + "%"

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"
			WHERE 
				"Recipes".name LIKE $1
				OR "Recipes".description LIKE $1`

	row := dbc.QueryRow(context.Background(), sqlreq, search)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {

		sqlreq = fmt.Sprintf(`SELECT 
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

	rows, err := dbc.Query(context.Background(), sqlreq, search)

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

		sqlreq = `SELECT 	
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

		ings, err := dbc.Query(context.Background(), sqlreq, cur.ID)

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
func PostgreSQLRecipesInsertUpdate(RecipeUpd RecipeDB, dbc *pgxpool.Pool) (RecipeDB, error) {

	if RecipeUpd.ImageDbID == 0 {
		RecipeUpd.ImageDbID = 1
	}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"
			WHERE 
				id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, RecipeUpd.ID)

	var recipecount int
	err := row.Scan(&recipecount)

	if err != nil {
		return RecipeUpd, err
	}

	dbc.Exec(context.Background(), "BEGIN")

	if recipecount > 0 && RecipeUpd.ID != 0 {
		// Если запись найдена по индексу и индекс не равен нулю (случай новой записи)
		// Обновляем существующую запись
		sqlreq = `UPDATE 
					public."Recipes" 
				SET 
					(name, description, image_id) = ($1, $2, $3) 
				WHERE 
					id=$4;`

		_, err = dbc.Exec(context.Background(), sqlreq, RecipeUpd.Name, RecipeUpd.Description, RecipeUpd.ImageDbID, RecipeUpd.ID)

	} else {
		// Иначе вставляем новую запись
		sqlreq = `INSERT INTO public."Recipes" (name, description, image_id) VALUES ($1, $2, $3) RETURNING id;`

		row := dbc.QueryRow(context.Background(), sqlreq, RecipeUpd.Name, RecipeUpd.Description, RecipeUpd.ImageDbID)

		err = row.Scan(&RecipeUpd.ID)
	}

	if err != nil {
		return RecipeUpd, PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `DELETE FROM public."RecipesIngredients" WHERE recipe_id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, RecipeUpd.ID)

	if err != nil {
		return RecipeUpd, PostgreSQLRollbackIfError(err, false, dbc)
	}

	for _, OneRecipe := range RecipeUpd.Ingredients {

		sqlreq = `SELECT 
					COUNT(*)
				FROM 
					public."Ingredients" 
				WHERE 
					name = $1
				LIMIT 1;`

		row := dbc.QueryRow(context.Background(), sqlreq, OneRecipe.Name)

		var count int
		err := row.Scan(&count)

		if err != nil {
			return RecipeUpd, PostgreSQLRollbackIfError(err, false, dbc)
		}

		var curid int

		if count > 0 {

			sqlreq = `SELECT 
						id						 
					FROM 
						public."Ingredients" 
					WHERE 
						name = $1
					LIMIT 1;`

			row := dbc.QueryRow(context.Background(), sqlreq, OneRecipe.Name)

			err := row.Scan(&curid)

			if err != nil {
				return RecipeUpd, PostgreSQLRollbackIfError(err, false, dbc)
			}

		} else {
			sqlreq = `INSERT INTO public."Ingredients" (name) VALUES ($1) RETURNING id;`

			row := dbc.QueryRow(context.Background(), sqlreq, OneRecipe.Name)

			err := row.Scan(&curid)

			if err != nil {
				return RecipeUpd, PostgreSQLRollbackIfError(err, false, dbc)
			}
		}

		sqlreq = `INSERT INTO public."RecipesIngredients" (recipe_id, ingredient_id, quantity) VALUES ($1,$2,$3);`

		_, err = dbc.Exec(context.Background(), sqlreq, RecipeUpd.ID, curid, OneRecipe.Amount)

		if err != nil {
			return RecipeUpd, PostgreSQLRollbackIfError(err, false, dbc)
		}

	}

	dbc.Exec(context.Background(), "COMMIT")

	return RecipeUpd, nil

}

// PostgreSQLRecipesDelete - удаляет рецепт и связанные с ним ингредиенты по индексу рецепта
func PostgreSQLRecipesDelete(ID int, dbc *pgxpool.Pool) error {

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Recipes"
			WHERE 
				id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, ID)

	var recipecount int
	err := row.Scan(&recipecount)

	shared.WriteErrToLog(err)

	if recipecount <= 0 {
		return ErrRecipeNotFound
	}

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq = `DELETE FROM public."RecipesIngredients" WHERE recipe_id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `DELETE FROM public."Recipes" WHERE id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, ID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	sqlreq = `select setval('"public"."Recipes_id_seq"',(select COALESCE(max("id"),1) from "public"."Recipes")::bigint);`

	_, err = dbc.Exec(context.Background(), sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}
