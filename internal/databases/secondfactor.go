// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLGetSecretByUserID - получает секрет и расшифровывает его
func PostgreSQLGetSecretByUserID(UserID uuid.UUID, dbc *pgxpool.Pool) (TOTPSecret, error) {

	var result TOTPSecret

	sqlreq := `SELECT 
					COUNT(*) 
				FROM 
					secret.totp
				WHERE 
					user_id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, UserID)

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

		rows, err := dbc.Query(context.Background(), sqlreq, UserID)

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
func PostgreSQLChangeSecondFactorSecret(totps TOTPSecret, dbc *pgxpool.Pool) error {

	sqlreq := `SELECT 
				COUNT(*)
			FROM
				secret.totp
			WHERE
				user_id=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, totps.UserID)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return err
	}

	if countRows > 0 {

		result, err := PostgreSQLGetSecretByUserID(totps.UserID, dbc)

		if err != nil {
			return err
		}

		if result.Confirmed {
			return ErrTOTPConfirmed
		}

		return PostgreSQLUpdateSecondFactorSecret(totps, dbc)
	}

	return PostgreSQLInsertSecondFactorSecret(totps, dbc)

}

// PostgreSQLInsertSecondFactorSecret - вставляет новую запись в
// список серкетов для двухфакторной авторизации
func PostgreSQLInsertSecondFactorSecret(totps TOTPSecret, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `INSERT INTO secret.totp(user_id, secret, key, confirmed) VALUES ($1, $2, $3, $4);`

	_, err := dbc.Exec(context.Background(), sqlreq, totps.UserID, totps.Secret, totps.EncKey, totps.Confirmed)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}

// PostgreSQLUpdateSecondFactorSecret - обновляет существующую запись в
// списке серкетов для двухфакторной авторизации
func PostgreSQLUpdateSecondFactorSecret(totps TOTPSecret, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `UPDATE secret.totp SET (secret, key, confirmed) = ($1, $2, $3) WHERE user_id=$4;`

	_, err := dbc.Exec(context.Background(), sqlreq, totps.Secret, totps.EncKey, totps.Confirmed, totps.UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}

// PostgreSQLUpdateSecondFactorConfirmed - взводит флаг подтверждения для второго фактора
func PostgreSQLUpdateSecondFactorConfirmed(Confirmed bool, UserID uuid.UUID, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `UPDATE secret.totp SET confirmed = $1 WHERE user_id=$2;`

	_, err := dbc.Exec(context.Background(), sqlreq, Confirmed, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	err = PostgreSQLSetUserSecondFactorActive(true, UserID, dbc)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}

// PostgreSQLDeleteSecondFactorSecret - удаляет привязанные секреты для двухфакторной авторизации пользователя
func PostgreSQLDeleteSecondFactorSecret(UserID uuid.UUID, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `DELETE FROM secret.totp WHERE user_id=$1;`

	_, err := dbc.Exec(context.Background(), sqlreq, UserID)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	err = PostgreSQLSetUserSecondFactorActive(false, UserID, dbc)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}

// PostgreSQLSetUserSecondFactorActive - меняет статус флага двухфакторной авторизации в пользователе
func PostgreSQLSetUserSecondFactorActive(Activate bool, UserID uuid.UUID, dbc *pgxpool.Pool) error {

	sqlreq := `UPDATE secret.users SET totp_active = $1 WHERE id=$2;`

	_, err := dbc.Exec(context.Background(), sqlreq, Activate, UserID)

	return err

}
