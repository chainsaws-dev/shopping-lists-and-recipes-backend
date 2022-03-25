package databases

import (
	"context"
	"fmt"
	"math"
	"shopping-lists-and-recipes/packages/authentication"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLSessionsSelect - получает информацию о списке сессий
func PostgreSQLSessionsSelect(page int, limit int, dbc *pgxpool.Pool) (SessionsResponse, error) {

	var result SessionsResponse

	result.Sessions = Sessions{}

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.sessions`

	row := dbc.QueryRow(context.Background(), sqlreq)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	offset := int(math.RoundToEven(float64((page - 1) * limit)))

	if PostgreSQLCheckLimitOffset(limit, offset) {

		sqlreq = fmt.Sprintf(`SELECT 
									sessions.email, 
									sessions.token, 
									sessions.session, 
									sessions.iss_date, 
									sessions.exp_date, 
									sessions.role, 
									sessions.ip, 
									sessions.user_agent, 
									sessions.second_factor_enabled, 
									sessions.second_factor_check_result
								FROM 
									secret.sessions
							ORDER BY
								sessions.email,
								sessions.exp_date						
							OFFSET %v LIMIT %v;`, offset, limit)

	} else {
		return result, ErrLimitOffsetInvalid
	}

	rows, err := dbc.Query(context.Background(), sqlreq)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var cur authentication.ActiveToken
		var SecFacEnabled bool
		var SecFacCheckRes bool

		err = rows.Scan(&cur.Email, &cur.Token, &cur.Session, &cur.IssDate, &cur.ExpDate,
			&cur.Role, &cur.IP, &cur.UserAgent, &SecFacEnabled, &SecFacCheckRes)

		cur.SecondFactor.Enabled = SecFacEnabled
		cur.SecondFactor.CheckResult = SecFacCheckRes

		result.Sessions = append(result.Sessions, cur)
	}

	result.Total = countRows
	result.Limit = limit
	result.Offset = offset

	return result, nil
}

func PostgreSQLCountTokensByEmail(email string, dbc *pgxpool.Pool) (int, error) {

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.sessions
			WHERE				
				email = $1`

	row := dbc.QueryRow(context.Background(), sqlreq, email)

	var countRows int

	err := row.Scan(&countRows)

	return countRows, err
}

func PostgreSQLGetActiveTokenByToken(token string, dbc *pgxpool.Pool) (authentication.ActiveToken, error) {

	ct := time.Now()

	var result authentication.ActiveToken

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.sessions
			WHERE				
				exp_date > $1
				AND token = $2`

	row := dbc.QueryRow(context.Background(), sqlreq, ct, token)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows > 0 {
		sqlreq := `SELECT 
						sessions.email, 
						sessions.token, 
						sessions.session, 
						sessions.iss_date, 
						sessions.exp_date, 
						sessions.role, 
						sessions.ip, 
						sessions.user_agent, 
						sessions.second_factor_enabled, 
						sessions.second_factor_check_result
					FROM 
						secret.sessions
					WHERE 						
						sessions.exp_date > $1
						AND sessions.token = $2
					LIMIT 1`

		rows, err := dbc.Query(context.Background(), sqlreq, ct, token)

		if err != nil {
			return result, err
		}

		for rows.Next() {

			var SecFacEnabled bool
			var SecFacCheckRes bool

			err = rows.Scan(&result.Email, &result.Token, &result.Session, &result.IssDate, &result.ExpDate,
				&result.Role, &result.IP, &result.UserAgent, &SecFacEnabled, &SecFacCheckRes)

			result.SecondFactor.Enabled = SecFacEnabled
			result.SecondFactor.CheckResult = SecFacCheckRes
		}
	} else {
		return result, ErrSessionsNotFoundSession
	}

	return result, err
}

// PostgreSQLGetActiveTokenBySession - получает объект активного токена по электронной почте и сессии
func PostgreSQLGetActiveTokenBySession(email string, session []byte, dbc *pgxpool.Pool) (authentication.ActiveToken, error) {

	ct := time.Now()

	var result authentication.ActiveToken

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.sessions
			WHERE				
				email = $1
				AND sessions.exp_date > $2
				AND session = $3`

	row := dbc.QueryRow(context.Background(), sqlreq, email, ct, session)

	var countRows int

	err := row.Scan(&countRows)

	if err != nil {
		return result, err
	}

	if countRows > 0 {
		sqlreq := `SELECT 
						sessions.email, 
						sessions.token, 
						sessions.session, 
						sessions.iss_date, 
						sessions.exp_date, 
						sessions.role, 
						sessions.ip, 
						sessions.user_agent, 
						sessions.second_factor_enabled, 
						sessions.second_factor_check_result
					FROM 
						secret.sessions
					WHERE 
						email = $1
						AND sessions.exp_date > $2
						AND session = $3
					LIMIT 1`

		rows, err := dbc.Query(context.Background(), sqlreq, email, ct, session)

		if err != nil {
			return result, err
		}

		for rows.Next() {

			var SecFacEnabled bool
			var SecFacCheckRes bool

			err = rows.Scan(&result.Email, &result.Token, &result.Session, &result.IssDate, &result.ExpDate,
				&result.Role, &result.IP, &result.UserAgent, &SecFacEnabled, &SecFacCheckRes)

			result.SecondFactor.Enabled = SecFacEnabled
			result.SecondFactor.CheckResult = SecFacCheckRes
		}
	} else {
		return result, ErrSessionsNotFoundSession
	}

	return result, err
}

// PostgreSQLSessionsInsert - Добавляет новую активную сессию в список сессий
func PostgreSQLSessionsInsert(NewActiveToken authentication.ActiveToken, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	// Добавляем новую
	sqlreq := `INSERT INTO secret.sessions (
		email, token, session, iss_date, exp_date, 
		role, ip, user_agent, second_factor_enabled, 
		second_factor_check_result
	  ) 
	  VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	_, err := dbc.Exec(context.Background(), sqlreq,
		NewActiveToken.Email,
		NewActiveToken.Token,
		NewActiveToken.Session,
		NewActiveToken.IssDate,
		NewActiveToken.ExpDate,
		NewActiveToken.Role,
		NewActiveToken.IP,
		NewActiveToken.UserAgent,
		NewActiveToken.SecondFactor.Enabled,
		NewActiveToken.SecondFactor.CheckResult)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil

}

func PostgreSQLSessionsUpdate(NewActiveToken authentication.ActiveToken, dbc *pgxpool.Pool) error {

	dbc.Exec(context.Background(), "BEGIN")

	sqlreq := `UPDATE 
					secret.sessions 
				SET 
					(email, token, session, iss_date, exp_date, 
					role, ip, user_agent, second_factor_enabled, 
					second_factor_check_result) = 
					($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
				WHERE 
					session=$3;`

	_, err := dbc.Exec(context.Background(), sqlreq,
		NewActiveToken.Email,
		NewActiveToken.Token,
		NewActiveToken.Session,
		NewActiveToken.IssDate,
		NewActiveToken.ExpDate,
		NewActiveToken.Role,
		NewActiveToken.IP,
		NewActiveToken.UserAgent,
		NewActiveToken.SecondFactor.Enabled,
		NewActiveToken.SecondFactor.CheckResult)

	if err != nil {
		return PostgreSQLRollbackIfError(err, false, dbc)
	}

	dbc.Exec(context.Background(), "COMMIT")

	return nil
}

// PostgreSQLDeleteSessionsByToken - удаляет сессию по токену
func PostgreSQLDeleteSessionsByToken(Token string, dbc *pgxpool.Pool) error {

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.sessions
			WHERE 
				sessions.token=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, Token)

	var SessionsCount int

	err := row.Scan(&SessionsCount)

	if err != nil {
		return err
	}

	if SessionsCount > 0 {

		dbc.Exec(context.Background(), "BEGIN")

		sqlreq = `DELETE FROM secret.sessions WHERE token=$1;`
		_, err = dbc.Exec(context.Background(), sqlreq, Token)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		sqlreq = `select setval('"secret"."sessions_id_seq"',(select COALESCE(max("id"),1) from "secret"."sessions")::bigint);`

		_, err = dbc.Exec(context.Background(), sqlreq)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		dbc.Exec(context.Background(), "COMMIT")

	} else {
		return ErrSessionsNotFoundEmail
	}

	return nil

}

// PostgreSQLDeleteSessionsByEmail - удаляет сессию по электронной почте
func PostgreSQLDeleteSessionsByEmail(Email string, dbc *pgxpool.Pool) error {

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.sessions
			WHERE 
				sessions.email=$1;`

	row := dbc.QueryRow(context.Background(), sqlreq, Email)

	var SessionsCount int

	err := row.Scan(&SessionsCount)

	if err != nil {
		return err
	}

	if SessionsCount > 0 {

		dbc.Exec(context.Background(), "BEGIN")

		sqlreq = `DELETE FROM secret.sessions WHERE email=$1;`
		_, err = dbc.Exec(context.Background(), sqlreq, Email)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		sqlreq = `select setval('"secret"."sessions_id_seq"',(select COALESCE(max("id"),1) from "secret"."sessions")::bigint);`

		_, err = dbc.Exec(context.Background(), sqlreq)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		dbc.Exec(context.Background(), "COMMIT")

	} else {
		return ErrSessionsNotFoundEmail
	}

	return nil
}

// PostgreSQLDeleteExpiredSessions - удаляет устаревшие сессии
func PostgreSQLDeleteExpiredSessions(dbc *pgxpool.Pool) error {

	ct := time.Now()

	sqlreq := `SELECT 
				COUNT(*)
			FROM 
				secret.sessions
			WHERE 
				sessions.exp_date <= $1;`

	row := dbc.QueryRow(context.Background(), sqlreq, ct)

	var SessionsCount int

	err := row.Scan(&SessionsCount)

	if err != nil {
		return err
	}

	if SessionsCount > 0 {

		dbc.Exec(context.Background(), "BEGIN")

		sqlreq = `DELETE FROM secret.sessions WHERE exp_date <= $1;`
		_, err = dbc.Exec(context.Background(), sqlreq, ct)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		sqlreq = `select setval('"secret"."sessions_id_seq"',(select COALESCE(max("id"),1) from "secret"."sessions")::bigint);`

		_, err = dbc.Exec(context.Background(), sqlreq)

		if err != nil {
			return PostgreSQLRollbackIfError(err, false, dbc)
		}

		dbc.Exec(context.Background(), "COMMIT")

	} else {
		return ErrSessionsNotFoundExpired
	}

	return nil

}
