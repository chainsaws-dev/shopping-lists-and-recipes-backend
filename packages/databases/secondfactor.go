package databases

import (
	uuid "github.com/satori/go.uuid"
)

// PostgreSQLGetSecretByUserID - получает секрет и расшифровывает его
func PostgreSQLGetSecretByUserID(UserID uuid.UUID) (TOTPSecret, error) {

	var result TOTPSecret

	sqlreq := `SELECT 
					COUNT(*) 
				FROM 
					secret.totp
				WHERE 
					user_id=$1;`

	row := dbc.QueryRow(sqlreq, UserID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows > 0 {

		sqlreq := `SELECT 
						user_id,
						secret, 
						key,
						confirmed
					FROM 
						secret.totp
					WHERE 
						user_id=$1
					LIMIT 1`

		rows, err := dbc.Query(sqlreq, UserID)

		if err != nil {
			return result, err
		}

		for rows.Next() {

			err = rows.Scan(&result.UserID, &result.Secret, &result.EncKey, &result.Confirmed)

			if err != nil {
				return result, err
			}

		}

	} else {
		return result, ErrUserTOTPNotFound
	}

	return result, nil
}

// PostgreSQLChangeSecondFactorSecret - принимает решение и обновляет
//  или добавляет новую запись в список серкетов для двухфакторной авторизации
func PostgreSQLChangeSecondFactorSecret(totps TOTPSecret) error {

	sqlreq := `SELECT 
				COUNT(*)
			FROM
				secret.totp
			WHERE
				user_id=$1;`

	row := dbc.QueryRow(sqlreq, totps.UserID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return err
	}

	if countRows > 0 {

		result, err := PostgreSQLGetSecretByUserID(totps.UserID)

		if err != nil {
			return err
		}

		if result.Confirmed {
			return ErrTOTPConfirmed
		}

		return PostgreSQLUpdateSecondFactorSecret(totps)
	}

	return PostgreSQLInsertSecondFactorSecret(totps)

}

// PostgreSQLInsertSecondFactorSecret - вставляет новую запись в
// список серкетов для двухфакторной авторизации
func PostgreSQLInsertSecondFactorSecret(totps TOTPSecret) error {

	dbc.Exec("BEGIN")

	sqlreq := `INSERT INTO secret.totp(user_id, secret, key, confirmed) VALUES ($1, $2, $3, $4);`

	_, err := dbc.Exec(sqlreq, totps.UserID, totps.Secret, totps.EncKey, totps.Confirmed)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLUpdateSecondFactorSecret - обновляет существующую запись в
// списке серкетов для двухфакторной авторизации
func PostgreSQLUpdateSecondFactorSecret(totps TOTPSecret) error {

	dbc.Exec("BEGIN")

	sqlreq := `UPDATE secret.totp SET (secret, key, confirmed) = ($1, $2, $3) WHERE user_id=$4;`

	_, err := dbc.Exec(sqlreq, totps.Secret, totps.EncKey, totps.Confirmed, totps.UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
}

// PostgreSQLDeleteSecondFactorSecret - удаляет привязанные секреты для двухфакторной авторизации пользователя
func PostgreSQLDeleteSecondFactorSecret(UserID uuid.UUID) error {

	dbc.Exec("BEGIN")

	sqlreq := `DELETE FROM secret.totp WHERE user_id=$1;`

	_, err := dbc.Exec(sqlreq, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false)
	}

	dbc.Exec("COMMIT")

	return nil
}
