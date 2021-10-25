// Package admin - содержит функции создания пользователей разных уровней доступа
package admin

import (
	"database/sql"
	"errors"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/internal/settings"
	"shopping-lists-and-recipes/packages/authentication"
)

// Список типовых ошибок
var (
	ErrBasicFieldsNotFilled = errors.New("не заполнены обязательные поля, невозможно создать пользователя")
)

// CreateAdmin - создаём пользователя для администратора
func CreateAdmin(SQL *settings.SQLServer, Login string, Email string, Password string, ConfirmEnabled bool, dbc *sql.DB) error {

	if len(Login) == 0 || len(Password) == 0 || len(Email) == 0 {
		return ErrBasicFieldsNotFilled
	}

	Hash, err := authentication.Argon2GenerateHash(Password, &authentication.HashParams)

	if err != nil {
		return err
	}

	var UserInfo = databases.User{
		Role:      "admin_role_CRUD",
		Email:     Email,
		Phone:     "",
		Name:      Login,
		IsAdmin:   true,
		Confirmed: !ConfirmEnabled,
	}

	_, err = databases.PostgreSQLUsersInsertUpdate(UserInfo, Hash, true, false, dbc)

	if err != nil {
		return err
	}

	return nil
}

// CreateUser - создаём пользователя для гостя
func CreateUser(SQL *settings.SQLServer, Login string, Email string, Password string, ConfirmEnabled bool, dbc *sql.DB) error {

	if len(Login) == 0 || len(Password) == 0 || len(Email) == 0 {
		return ErrBasicFieldsNotFilled
	}

	Hash, err := authentication.Argon2GenerateHash(Password, &authentication.HashParams)

	if err != nil {
		return err
	}

	var UserInfo = databases.User{
		Role:      "guest_role_read_only",
		Email:     Email,
		Phone:     "",
		Name:      Login,
		IsAdmin:   false,
		Confirmed: !ConfirmEnabled,
	}

	_, err = databases.PostgreSQLUsersInsertUpdate(UserInfo, Hash, true, false, dbc)

	if err != nil {
		return err
	}

	return nil
}
