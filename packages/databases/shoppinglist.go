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

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."ShoppingList"`

	row := dbc.QueryRow(sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {

		sqlreq = fmt.Sprintf(`SELECT 
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

	rows, err := dbc.Query(sqlreq)

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

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Ingredients"
			WHERE 
				"Ingredients".name=$1;`

	row := dbc.QueryRow(sqlreq, ShoppingItem.Name)

	var IngCount int

	err := row.Scan(&IngCount)

	if err != nil {
		return err
	}

	var countRows int

	var ingID int

	dbc.Exec("BEGIN")

	if IngCount > 0 {

		sqlreq := `SELECT 
					id
				FROM 
					public."Ingredients"
				WHERE 
					"Ingredients".name=$1;`

		ingrow := dbc.QueryRow(sqlreq, ShoppingItem.Name)

		err = ingrow.Scan(&ingID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		sqlreq = `SELECT 
					COUNT(*)
				FROM
					public."ShoppingList"
				WHERE ingredient_id =$1;`

		row := dbc.QueryRow(sqlreq, ingID)

		err := row.Scan(&countRows)

		if err != nil {
			return err
		}

		if countRows == 0 {
			// Добавляем новую
			sqlreq := `INSERT INTO public."ShoppingList" (ingredient_id, quantity) VALUES ($1,$2);`

			_, err = dbc.Exec(sqlreq, ingID, ShoppingItem.Amount)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		} else {
			// Обновляем существующую
			sqlreq = `UPDATE public."ShoppingList" SET quantity = $1 WHERE ingredient_id=$2;`

			_, err = dbc.Exec(sqlreq, ShoppingItem.Amount, ingID)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		}
	} else {
		// Добавляем ингредиент в справочник
		sqlreq = `INSERT INTO public."Ingredients" (name) VALUES ($1) RETURNING id;`

		ingrow := dbc.QueryRow(sqlreq, ShoppingItem.Name)

		err := ingrow.Scan(&ingID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		sqlreq = `SELECT 
					COUNT(*)
				FROM
					public."ShoppingList"
				WHERE ingredient_id =$1;`

		row := dbc.QueryRow(sqlreq, ingID)

		err = row.Scan(&countRows)

		if err != nil {
			return err
		}

		if countRows == 0 {
			// Добавляем новую запись в список покупок
			sqlreq = `INSERT INTO public."ShoppingList" (ingredient_id, quantity) VALUES ($1,$2);`

			_, err = dbc.Exec(sqlreq, ingID, ShoppingItem.Amount)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		} else {
			// Обновляем существующую
			sqlreq = `UPDATE public."ShoppingList" SET quantity = $1 WHERE ingredient_id=$2;`

			_, err = dbc.Exec(sqlreq, ShoppingItem.Amount, ingID)

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

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				public."Ingredients"
			WHERE 
				"Ingredients".name=$1;`

	row := dbc.QueryRow(sqlreq, IngName)

	var IngCount int

	err := row.Scan(&IngCount)

	if err != nil {
		return err
	}

	var ingID int
	var CountRows int

	if IngCount > 0 {

		sqlreq := `SELECT 
					id
				FROM 
					public."Ingredients"
				WHERE 
					"Ingredients".name=$1;`

		ingrow := dbc.QueryRow(sqlreq, IngName)

		err = ingrow.Scan(&ingID)

		if err != nil {
			return err
		}

		sqlreq = `SELECT 
					COUNT(*)
				FROM 
					public."ShoppingList"
				WHERE 
					ingredient_id=$1;`

		slrow := dbc.QueryRow(sqlreq, ingID)

		err = slrow.Scan(&CountRows)

		if err != nil {
			return err
		}

		if CountRows <= 0 {
			return ErrShoppingListNotFound
		}

		dbc.Exec("BEGIN")

		sqlreq = `DELETE FROM public."ShoppingList" WHERE ingredient_id=$1;`
		_, err = dbc.Exec(sqlreq, ingID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		sqlreq = `select setval('"public"."ShoppingList_id_seq"',(select COALESCE(max("id"),1) from "public"."ShoppingList")::bigint);`

		_, err = dbc.Exec(sqlreq)

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

	sqlreq := `DELETE FROM public."ShoppingList";`

	_, err := dbc.Exec(sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	sqlreq = `select setval('"public"."ShoppingList_id_seq"',(select COALESCE(max("id"),1) from "public"."ShoppingList")::bigint);`

	_, err = dbc.Exec(sqlreq)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
}
