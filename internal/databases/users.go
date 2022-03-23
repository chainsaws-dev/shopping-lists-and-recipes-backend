// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"
	"fmt"
	"math"
	"shopping-lists-and-recipes/packages/shared"

	uuid "github.com/gofrs/uuid"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLUsersSelect - получает список пользователей в админке
func PostgreSQLUsersSelect(page int, limit int, dbc *pgxpool.Pool) (UsersResponse, error) {

	var result UsersResponse
	result.Users = Users{}

	sqlreq := `SELECT 
				COUNT(*) 
			FROM 
				secret.users;`

	row := dbc.QueryRow(context.Background(), sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {

		sqlreq = fmt.Sprintf(`SELECT 
								users.id,
								users.role,
								users.email,
								users.phone,
								users.name,
								users.lang,
								users.isadmin,
								users.confirmed,
								users.disabled,
								users.totp_active
							FROM 
								secret.users
							ORDER BY 
								email
							OFFSET %v LIMIT %v`, offset, limit)

	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(context.Background(), sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur User
		rows.Scan(&cur.GUID, &cur.Role, &cur.Email, &cur.Phone, &cur.Name, &cur.Lang, &cur.IsAdmin, &cur.Confirmed, &cur.Disabled, &cur.SecondFactor)
		result.Users = append(result.Users, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

// PostgreSQLUsersInsertUpdate - создаёт или обновляет существующего пользователя
func PostgreSQLUsersInsertUpdate(NewUserInfo User, Hash string, UpdatePassword bool, OverWrite bool, dbc *pgxpool.Pool) (User, error) {

	if NewUserInfo.Role == "admin_role_CRUD" {
		NewUserInfo.IsAdmin = true
	} else {
		NewUserInfo.IsAdmin = false
	}

	// Проверяем что почта уникальна
	var EmailCount int

	sqlreq := `SELECT COUNT(*) FROM secret.users WHERE email=$1;`

	EmailCountRow := dbc.QueryRow(context.Background(), sqlreq, NewUserInfo.Email)

	err := EmailCountRow.Scan(&EmailCount)

	if err != nil {
		return NewUserInfo, err
	}

	if EmailCount > 0 && !OverWrite || EmailCount > 0 && NewUserInfo.GUID == uuid.Nil {
		return NewUserInfo, ErrEmailIsOccupied
	}

	// Проверяем что пользователь с ID существует
	var UserCount int

	sqlreq = `SELECT COUNT(*) FROM secret.users WHERE id=$1;`

	UserCountRow := dbc.QueryRow(context.Background(), sqlreq, NewUserInfo.GUID)

	err = UserCountRow.Scan(&UserCount)

	if err != nil {
		return NewUserInfo, err
	}

	dbc.Exec(context.Background(), "BEGIN")

	if UserCount > 0 && OverWrite {

		// Обновляем существующего

		sqlreq = `UPDATE secret.users SET (role, email, phone, name, lang, isadmin, confirmed, disabled, totp_active) = ($1,$2,$3,$4,$5,$6,$7,$8,$9) WHERE id=$10;`

		_, err = dbc.Exec(context.Background(), sqlreq, NewUserInfo.Role, NewUserInfo.Email, NewUserInfo.Phone, NewUserInfo.Name,
			NewUserInfo.Lang, NewUserInfo.IsAdmin, NewUserInfo.Confirmed, NewUserInfo.Disabled, NewUserInfo.SecondFactor, NewUserInfo.GUID)

		if err != nil {
			return NewUserInfo, PostgreSQLRollbackIfError(err, false, dbc)
		}

		if UpdatePassword {
			if len(Hash) > 0 {

				sqlreq = `UPDATE secret.hashes SET value=$2 WHERE user_id=$1;`

				_, err = dbc.Exec(context.Background(), sqlreq, NewUserInfo.GUID, Hash)

				if err != nil {
					return NewUserInfo, PostgreSQLRollbackIfError(err, false, dbc)
				}
			} else {
				return NewUserInfo, PostgreSQLRollbackIfError(ErrEmptyPassword, false, dbc)
			}
		}

	} else {

		// Создаём нового

		// Генерируем новый уникальный идентификатор
		NewUserInfo.GUID, err = uuid.NewV4()

		if err != nil {
			return NewUserInfo, PostgreSQLRollbackIfError(err, false, dbc)
		}

		sqlreq = `INSERT INTO secret.users (id, role, email, phone, name, lang, isadmin, confirmed, disabled, totp_active) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`

		_, err = dbc.Exec(context.Background(), sqlreq, NewUserInfo.GUID, NewUserInfo.Role, NewUserInfo.Email, NewUserInfo.Phone,
			NewUserInfo.Name, NewUserInfo.Lang, NewUserInfo.IsAdmin, NewUserInfo.Confirmed, NewUserInfo.Disabled, NewUserInfo.SecondFactor)

		if err != nil {
			return NewUserInfo, PostgreSQLRollbackIfError(err, false, dbc)
		}

		if len(Hash) > 0 {

			sqlreq = `INSERT INTO secret.hashes (user_id, value) VALUES ($1,$2);`

			_, err = dbc.Exec(context.Background(), sqlreq, NewUserInfo.GUID, Hash)

			if err != nil {
				return NewUserInfo, PostgreSQLRollbackIfError(err, false, dbc)
			}
		} else {
			return NewUserInfo, PostgreSQLRollbackIfError(ErrEmptyPassword, false, dbc)
		}

	}

	dbc.Exec(context.Background(), "COMMIT")

	return NewUserInfo, nil
}

// PostgreSQLUsersDelete - удаляет пользователя с указанным GUID
func PostgreSQLUsersDelete(UserID uuid.UUID, dbc *pgxpool.Pool, locale string) error {

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.users
			WHERE 
				id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, UserID)

	var usercount int
	err := row.Scan(&usercount)

	shared.WriteErrToLog(err, locale)

	if usercount <= 0 {
		return ErrUserNotFound
	}

	dbc.Exec(context.Background(), "BEGIN")

	// Удаляем связанные хеши
	sqlreq = `DELETE FROM secret.hashes WHERE user_id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	// Удаляем подтверждения почты
	sqlreq = `DELETE FROM secret.confirmations WHERE user_id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	// Удаляем сбросы паролей
	sqlreq = `DELETE FROM secret.password_resets WHERE user_id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	// Удаляем привязки временных ключей
	sqlreq = `DELETE FROM secret.totp WHERE user_id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	// Удаляем пользователя
	sqlreq = `DELETE FROM secret.users WHERE id=$1;`

	_, err = dbc.Exec(context.Background(), sqlreq, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}

// PostgreSQLCheckUserMailExists - проверяет что почтовый ящик живого пользователя
func PostgreSQLCheckUserMailExists(Email string, dbc *pgxpool.Pool) (bool, error) {

	sqlreq := `SELECT COUNT(*) FROM secret.users WHERE email=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, Email)

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
func PostgreSQLGetUserByEmail(Email string, dbc *pgxpool.Pool) (User, error) {

	var result User

	if len(Email) > 0 {

		sqlreq := `SELECT 
					COUNT(*) 
				FROM 
					secret.users
				WHERE 
					users.email=$1;`

		row := dbc.QueryRow(context.Background(), sqlreq, Email)

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
							users.lang,
							users.isadmin,
							users.confirmed,
							users.disabled,
							users.totp_active
						FROM 
							secret.users
						WHERE 
							users.email=$1
						LIMIT 1`

			rows, err := dbc.Query(context.Background(), sqlreq, Email)

			if err != nil {
				return result, err
			}

			for rows.Next() {

				err = rows.Scan(&result.GUID, &result.Role, &result.Email, &result.Phone, &result.Name,
					&result.Lang, &result.IsAdmin, &result.Confirmed, &result.Disabled, &result.SecondFactor)

				if err != nil {
					return result, err
				}

			}

		} else {
			return result, ErrNoUserWithEmail
		}

	} else {
		return result, ErrNoUserWithEmail
	}

	return result, nil
}

// PostgreSQLCurrentUserUpdate - функция позволяющая менять поля, которые пользователь может менять
func PostgreSQLCurrentUserUpdate(NewUserInfo User, Hash string, UpdatePassword bool, dbc *pgxpool.Pool) (User, error) {

	// Проверяем что пользователь с ID существует
	var UserCount int

	sqlreq := `SELECT COUNT(*) FROM secret.users WHERE id=$1;`

	UserCountRow := dbc.QueryRow(context.Background(), sqlreq, NewUserInfo.GUID)

	err := UserCountRow.Scan(&UserCount)

	if err != nil {
		return NewUserInfo, err
	}

	dbc.Exec(context.Background(), "BEGIN")

	if UserCount > 0 {
		// Обновляем существующего
		sqlreq = `UPDATE secret.users SET (phone, name, lang) = ($1,$2,$3) WHERE id=$4;`

		_, err = dbc.Exec(context.Background(), sqlreq, NewUserInfo.Phone, NewUserInfo.Name, NewUserInfo.Lang, NewUserInfo.GUID)

		if err != nil {
			return NewUserInfo, PostgreSQLRollbackIfError(err, false, dbc)
		}

		if UpdatePassword {
			if len(Hash) > 0 {

				sqlreq = `UPDATE secret.hashes SET value=$2 WHERE user_id=$1;`

				_, err = dbc.Exec(context.Background(), sqlreq, NewUserInfo.GUID, Hash)

				if err != nil {
					return NewUserInfo, PostgreSQLRollbackIfError(err, false, dbc)
				}
			} else {
				return NewUserInfo, PostgreSQLRollbackIfError(ErrEmptyPassword, false, dbc)
			}
		}
	} else {
		return NewUserInfo, ErrNoUserWithEmail
	}

	dbc.Exec(context.Background(), "COMMIT")

	return NewUserInfo, nil
}
