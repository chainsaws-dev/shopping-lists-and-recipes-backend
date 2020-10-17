package databases

import (
	"fmt"
	"math"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

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

	if PostgreSQLCheckLimitOffset(limit, offset) {

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
		return result, ErrLimitOffsetInvalid
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
