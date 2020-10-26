package databases

import (
	"fmt"
	"math"
	"shopping-lists-and-recipes/packages/shared"

	uuid "github.com/satori/go.uuid"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLUsersSelect - получает список пользователей в админке
func PostgreSQLUsersSelect(page int, limit int) (UsersResponse, error) {

	var result UsersResponse
	result.Users = Users{}

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

	if PostgreSQLCheckLimitOffset(limit, offset) {

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
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(sql)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur User
		rows.Scan(&cur.GUID, &cur.Role, &cur.Email, &cur.Phone, &cur.Name, &cur.IsAdmin, &cur.Confirmed)
		result.Users = append(result.Users, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLUsersInsertUpdate - создаёт или обновляет существующего пользователя
func PostgreSQLUsersInsertUpdate(NewUserInfo User, Hash string, UpdatePassword bool, OverWrite bool) (User, error) {

	if NewUserInfo.Role == "admin_role_CRUD" {
		NewUserInfo.IsAdmin = true
	} else {
		NewUserInfo.IsAdmin = false
	}

	// Проверяем что почта уникальна
	var EmailCount int

	sql := `SELECT COUNT(*) FROM secret.users WHERE email=$1;`

	EmailCountRow := dbc.QueryRow(sql, NewUserInfo.Email)

	err := EmailCountRow.Scan(&EmailCount)

	if err != nil {
		return NewUserInfo, err
	}

	if EmailCount > 0 && !OverWrite || EmailCount > 0 && NewUserInfo.GUID == uuid.Nil {
		return NewUserInfo, ErrEmailIsOccupied
	}

	// Проверяем что пользователь с ID существует
	var UserCount int

	sql = `SELECT COUNT(*) FROM secret.users WHERE id=$1;`

	UserCountRow := dbc.QueryRow(sql, NewUserInfo.GUID)

	err = UserCountRow.Scan(&UserCount)

	if err != nil {
		return NewUserInfo, err
	}

	dbc.Exec("BEGIN")

	if UserCount > 0 && OverWrite {

		// Обновляем существующего

		sql = `UPDATE secret.users SET (role, email, phone, name, isadmin, confirmed) = ($1,$2,$3,$4,$5,$6) WHERE id=$7;`

		_, err = dbc.Exec(sql, NewUserInfo.Role, NewUserInfo.Email, NewUserInfo.Phone, NewUserInfo.Name, NewUserInfo.IsAdmin, NewUserInfo.Confirmed, NewUserInfo.GUID)

		if err != nil {
			return NewUserInfo, PostgreSQLRollbackIfError(err, false)
		}

		if UpdatePassword {
			if len(Hash) > 0 {

				sql = `UPDATE secret.hashes SET value=$2 WHERE user_id=$1;`

				_, err = dbc.Exec(sql, NewUserInfo.GUID, Hash)

				if err != nil {
					return NewUserInfo, PostgreSQLRollbackIfError(err, false)
				}
			} else {
				return NewUserInfo, PostgreSQLRollbackIfError(ErrEmptyPassword, false)
			}
		}

	} else {

		// Создаём нового

		// Генерируем новый уникальный идентификатор
		NewUserInfo.GUID = uuid.NewV4()

		sql = `INSERT INTO secret.users (id, role, email, phone, name, isadmin, confirmed) VALUES ($1,$2,$3,$4,$5,$6,$7);`

		_, err = dbc.Exec(sql, NewUserInfo.GUID, NewUserInfo.Role, NewUserInfo.Email, NewUserInfo.Phone, NewUserInfo.Name, NewUserInfo.IsAdmin, NewUserInfo.Confirmed)

		if err != nil {
			return NewUserInfo, PostgreSQLRollbackIfError(err, false)
		}

		if len(Hash) > 0 {

			sql = `INSERT INTO secret.hashes (user_id, value) VALUES ($1,$2);`

			_, err = dbc.Exec(sql, NewUserInfo.GUID, Hash)

			if err != nil {
				return NewUserInfo, PostgreSQLRollbackIfError(err, false)
			}
		} else {
			return NewUserInfo, PostgreSQLRollbackIfError(ErrEmptyPassword, false)
		}

	}

	dbc.Exec("COMMIT")

	return NewUserInfo, nil
}

// PostgreSQLUsersDelete - удаляет пользователя с указанным GUID
func PostgreSQLUsersDelete(UserID uuid.UUID) error {

	sql := `SELECT 
				COUNT(*)
			FROM 
				secret.users
			WHERE 
				id=$1;`

	row := dbc.QueryRow(sql, UserID)

	var usercount int
	err := row.Scan(&usercount)

	shared.WriteErrToLog(err)

	if usercount <= 0 {
		return ErrUserNotFound
	}

	dbc.Exec("BEGIN")

	// Удаляем связанные хеши
	sql = `DELETE FROM secret.hashes WHERE user_id=$1;`

	_, err = dbc.Exec(sql, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	// Удаляем подтверждения если есть
	sql = `DELETE FROM secret.confirmations WHERE user_id=$1;`

	_, err = dbc.Exec(sql, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	// Удаляем пользователя
	sql = `DELETE FROM secret.users WHERE id=$1;`

	_, err = dbc.Exec(sql, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLCheckUserMailExists - проверяет что почтовый ящик живого пользователя
func PostgreSQLCheckUserMailExists(Email string) (bool, error) {

	sql := `SELECT COUNT(*) FROM secret.users WHERE email=$1;`

	row := dbc.QueryRow(sql, Email)

	var UsersCount int

	err := row.Scan(&UsersCount)

	if err != nil {
		return false, err
	}

	if UsersCount > 0 {
		return true, nil
	}

	return false, nil
}

// PostgreSQLGetUserByEmail - получает данные о пользователе по электронной почте
func PostgreSQLGetUserByEmail(Email string) (User, error) {

	var result User

	if len(Email) > 0 {

		sqlreq := `SELECT 
					COUNT(*) 
				FROM 
					secret.users
				WHERE 
					users.email=$1;`

		row := dbc.QueryRow(sqlreq, Email)

		var countRows int

		err := row.Scan(&countRows)

		if err != nil {
			return result, err
		}

		if countRows > 0 {

			sqlreq := `SELECT 
							users.id,
							users.role,
							users.email,
							users.phone,
							users.name,
							users.isadmin,
							users.confirmed
						FROM 
							secret.users
						WHERE 
							users.email=$1
						LIMIT 1`

			rows, err := dbc.Query(sqlreq, Email)

			if err != nil {
				return result, err
			}

			for rows.Next() {

				err = rows.Scan(&result.GUID, &result.Role, &result.Email, &result.Phone, &result.Name, &result.IsAdmin, &result.Confirmed)

				if err != nil {
					return result, err
				}

			}

		} else {
			return result, ErrNoUserWithEmail
		}

	}

	return result, nil
}
