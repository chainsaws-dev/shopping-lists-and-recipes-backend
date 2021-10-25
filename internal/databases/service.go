// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLCleanAccessToken - Удаляет заданный токен доступа
func PostgreSQLCleanAccessToken(Token string, TokenStorageTableName string, dbc *sql.DB) error {

	_, err := dbc.Exec(fmt.Sprintf(`DELETE FROM %v WHERE token=$1;`, TokenStorageTableName), Token)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	return nil
}

// PostgreSQLCleanAccessTokens - Удаляет все истекшие токены доступа
func PostgreSQLCleanAccessTokens(dbc *sql.DB) error {

	dbc.Exec("BEGIN")

	_, err := dbc.Exec(`DELETE FROM secret.confirmations WHERE expires < now();`)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	_, err = dbc.Exec(`DELETE FROM secret.password_resets WHERE expires < now();`)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLGetTokenConfirmEmail - Ищем токен из запроса и устанавливаем у пользователя подтверждение если он существует
func PostgreSQLGetTokenConfirmEmail(Token string, dbc *sql.DB) error {

	sqlreq := `SELECT 
				COUNT(*) 
			FROM 
				secret.confirmations 
			WHERE 
				token=$1 
				AND expires >= now()
			LIMIT 1;`

	row := dbc.QueryRow(sqlreq, Token)

	var TokenCount int

	err := row.Scan(&TokenCount)

	if err != nil {
		return err
	}

	if TokenCount > 0 {

		sqlreq = `SELECT 
					confirmations.user_id
				FROM
					secret.confirmations
				WHERE
					token=$1 
					AND expires >= now()
				LIMIT 1;`

		row := dbc.QueryRow(sqlreq, Token)

		var UID uuid.UUID

		err := row.Scan(&UID)

		if err != nil {
			return err
		}

		dbc.Exec("BEGIN")

		sqlreq = "UPDATE secret.users SET confirmed=true WHERE id=$1;"

		_, err = dbc.Exec(sqlreq, UID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		err = PostgreSQLCleanAccessToken(Token, "secret.confirmations", dbc)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		dbc.Exec("COMMIT")

	} else {
		return ErrTokenExpired
	}

	return nil

}

// PostgreSQLGetTokenResetPassword - ищем токен среди выданных и не протухших и обновляем хеш пароля для пользователя
func PostgreSQLGetTokenResetPassword(Token string, Hash string, dbc *sql.DB) error {

	sqlreq := `SELECT 
				COUNT(*) 
			FROM 
				secret.password_resets 
			WHERE 
				token=$1 
				AND expires >= now()
			LIMIT 1;`

	row := dbc.QueryRow(sqlreq, Token)

	var TokenCount int

	err := row.Scan(&TokenCount)

	if err != nil {
		return err
	}

	if TokenCount > 0 {

		sqlreq = `SELECT 
					password_resets.user_id
				FROM
					secret.password_resets
				WHERE
					token=$1 
					AND expires >= now()
				LIMIT 1;`

		row := dbc.QueryRow(sqlreq, Token)

		var UID uuid.UUID

		err := row.Scan(&UID)

		if err != nil {
			return err
		}

		dbc.Exec("BEGIN")

		if len(Hash) > 0 {

			sqlreq = `UPDATE secret.hashes SET value=$2 WHERE user_id=$1;`

			_, err = dbc.Exec(sqlreq, UID, Hash)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false, dbc)
			}
		} else {
			return PostgreSQLRollbackIfError(ErrEmptyPassword, false, dbc)
		}

		err = PostgreSQLCleanAccessToken(Token, "secret.password_resets", dbc)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		dbc.Exec("COMMIT")

	} else {
		return ErrTokenExpired
	}

	return nil

}

// PostgreSQLSaveAccessToken - сохраняем токен для подтверждения почты
func PostgreSQLSaveAccessToken(Token string, Email string, TokenTableName string, dbc *sql.DB) error {

	if len(Token) > 0 && len(Email) > 0 {

		sqlreq := `SELECT COUNT(*) FROM secret.users WHERE email=$1 LIMIT 1;`

		row := dbc.QueryRow(sqlreq, Email)

		var UsCount int

		err := row.Scan(&UsCount)

		if err != nil {
			return err
		}

		if UsCount > 0 {

			sqlreq = `SELECT id FROM secret.users WHERE email=$1 LIMIT 1;`

			row = dbc.QueryRow(sqlreq, Email)

			var CurUID uuid.UUID

			err = row.Scan(&CurUID)

			if err != nil {
				return err
			}

			dbc.Exec("BEGIN")

			sqlreq = fmt.Sprintf(`DELETE FROM %v WHERE user_id=$1;`, TokenTableName)

			_, err = dbc.Exec(sqlreq, CurUID)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false, dbc)
			}

			sqlreq = fmt.Sprintf(`INSERT INTO %v (user_id, token, created, expires) VALUES ($1,$2,$3,$4);`, TokenTableName)

			cd := time.Now()

			_, err = dbc.Exec(sqlreq, CurUID, Token, cd, cd.Add(time.Minute*10))

			if err != nil {
				return PostgreSQLRollbackIfError(err, false, dbc)
			}

			dbc.Exec("COMMIT")
		}

	}
	return nil
}

// PostgreSQLGetTokenForUser - получает токен для проверки при авторизации
func PostgreSQLGetTokenForUser(email string, dbc *sql.DB) (string, string, error) {

	var UserCount int
	var UserID uuid.UUID
	var HashesCount int
	var UserRole string
	var Confirmed bool

	var Hash string

	sqlreq := `SELECT 
				COUNT(*) 
			FROM 
				secret.users 
			WHERE 
				email=$1;`

	UserCountRow := dbc.QueryRow(sqlreq, email)

	err := UserCountRow.Scan(&UserCount)

	if err != nil {
		return "", "", err
	}

	if UserCount <= 0 {
		return "", "", ErrNoUserWithEmail
	}

	sqlreq = `SELECT 
				id,
				role, 
				confirmed
			FROM 
				secret.users
			WHERE
				email=$1 
			LIMIT 1`

	UserIDRow := dbc.QueryRow(sqlreq, email)

	err = UserIDRow.Scan(&UserID, &UserRole, &Confirmed)

	if err != nil {
		return "", "", err
	}

	if !Confirmed {
		return "", "", ErrEmailNotConfirmed
	}

	sqlreq = `SELECT 
				COUNT(*)
			FROM 
				secret.hashes
			WHERE 
				user_id = $1;`

	HashesRow := dbc.QueryRow(sqlreq, UserID)

	err = HashesRow.Scan(&HashesCount)

	if err != nil {
		return "", "", err
	}

	if HashesCount <= 0 {
		return "", "", ErrNoHashForUser
	}

	sqlreq = `SELECT 
				value
			FROM 
				secret.hashes
			WHERE 
				user_id = $1;`

	HashValueRow := dbc.QueryRow(sqlreq, UserID)

	err = HashValueRow.Scan(&Hash)

	if err != nil {
		return "", "", err
	}

	return Hash, UserRole, nil
}
