package databases

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgreSQLCleanAccessToken - Удаляет заданный токен доступа
func PostgreSQLCleanAccessToken(Token string, TokenStorageTableName string) error {

	_, err := dbc.Exec(fmt.Sprintf(`DELETE FROM %v WHERE token=$1;`, TokenStorageTableName), Token)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	return nil
}

// PostgreSQLCleanAccessTokens - Удаляет все истекшие токены доступа
func PostgreSQLCleanAccessTokens() error {

	dbc.Exec("BEGIN")

	_, err := dbc.Exec(`DELETE FROM secret.confirmations WHERE "Expires" < now();`)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	_, err = dbc.Exec(`DELETE FROM secret.password_resets WHERE "Expires" < now();`)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLGetTokenConfirmEmail - Ищем токен из запроса и устанавливаем у пользователя подтверждение если он существует
func PostgreSQLGetTokenConfirmEmail(Token string) error {

	sql := `SELECT 
				COUNT(*) 
			FROM 
				secret.confirmations 
			WHERE 
				token=$1 
				AND "Expires" >= now()
			LIMIT 1;`

	row := dbc.QueryRow(sql, Token)

	var TokenCount int

	err := row.Scan(&TokenCount)

	if err != nil {
		return err
	}

	if TokenCount > 0 {

		sql = `SELECT 
					confirmations.user_id
				FROM
					secret.confirmations
				WHERE
					token=$1 
					AND "Expires" >= now()
				LIMIT 1;`

		row := dbc.QueryRow(sql, Token)

		var UID uuid.UUID

		err := row.Scan(&UID)

		if err != nil {
			return err
		}

		dbc.Exec("BEGIN")

		sql = "UPDATE secret.users SET confirmed=true WHERE id=$1;"

		_, err = dbc.Exec(sql, UID)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		err = PostgreSQLCleanAccessToken(Token, "secret.confirmations")

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		dbc.Exec("COMMIT")

	} else {
		return ErrTokenExpired
	}

	return nil

}

// PostgreSQLGetTokenResetPassword - ищем токен среди выданных и не протухших и обновляем хеш пароля для пользователя
func PostgreSQLGetTokenResetPassword(Token string, Hash string) error {

	sql := `SELECT 
				COUNT(*) 
			FROM 
				secret.password_resets 
			WHERE 
				token=$1 
				AND "Expires" >= now()
			LIMIT 1;`

	row := dbc.QueryRow(sql, Token)

	var TokenCount int

	err := row.Scan(&TokenCount)

	if err != nil {
		return err
	}

	if TokenCount > 0 {

		sql = `SELECT 
					password_resets.user_id
				FROM
					secret.password_resets
				WHERE
					token=$1 
					AND "Expires" >= now()
				LIMIT 1;`

		row := dbc.QueryRow(sql, Token)

		var UID uuid.UUID

		err := row.Scan(&UID)

		if err != nil {
			return err
		}

		dbc.Exec("BEGIN")

		if len(Hash) > 0 {

			sql = `UPDATE secret.hashes SET value=$2 WHERE user_id=$1;`

			_, err = dbc.Exec(sql, UID, Hash)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}
		} else {
			return PostgreSQLRollbackIfError(ErrEmptyPassword, false)
		}

		err = PostgreSQLCleanAccessToken(Token, "secret.password_resets")

		if err != nil {
			return PostgreSQLRollbackIfError(err, false)
		}

		dbc.Exec("COMMIT")

	} else {
		return ErrTokenExpired
	}

	return nil

}

// PostgreSQLSaveAccessToken - сохраняем токен для подтверждения почты
func PostgreSQLSaveAccessToken(Token string, Email string, TokenTableName string) error {

	if len(Token) > 0 && len(Email) > 0 {

		sql := `SELECT COUNT(*) FROM secret.users WHERE email=$1 LIMIT 1;`

		row := dbc.QueryRow(sql, Email)

		var UsCount int

		err := row.Scan(&UsCount)

		if err != nil {
			return err
		}

		if UsCount > 0 {

			sql = `SELECT id FROM secret.users WHERE email=$1 LIMIT 1;`

			row = dbc.QueryRow(sql, Email)

			var CurUID uuid.UUID

			err = row.Scan(&CurUID)

			if err != nil {
				return err
			}

			dbc.Exec("BEGIN")

			sql = fmt.Sprintf(`DELETE FROM %v WHERE user_id=$1;`, TokenTableName)

			_, err = dbc.Exec(sql, CurUID)

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}

			sql = fmt.Sprintf(`INSERT INTO %v (user_id, token, "Created", "Expires") VALUES ($1,$2,$3,$4);`, TokenTableName)

			cd := time.Now()

			_, err = dbc.Exec(sql, CurUID, Token, cd, cd.Add(time.Minute*10))

			if err != nil {
				return PostgreSQLRollbackIfError(err, false)
			}

			dbc.Exec("COMMIT")
		}

	}
	return nil
}

// PostgreSQLGetTokenForUser - получает токен для проверки при авторизации
func PostgreSQLGetTokenForUser(email string) (string, string, error) {

	var UserCount int
	var UserID uuid.UUID
	var HashesCount int
	var UserRole string
	var Confirmed bool

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
				role, 
				confirmed
			FROM 
				secret.users
			WHERE
				email=$1 
			LIMIT 1`

	UserIDRow := dbc.QueryRow(sql, email)

	err = UserIDRow.Scan(&UserID, &UserRole, &Confirmed)

	if err != nil {
		return "", "", err
	}

	if !Confirmed {
		return "", "", ErrEmailNotConfirmed
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
