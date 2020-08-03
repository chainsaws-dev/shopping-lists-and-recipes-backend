// Package admin - содержит функции создания пользователей разных уровней доступа
package admin

import (
	"errors"
	"shopping-lists-and-recipes/packages/authentication"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/settings"
)

// Список типовых ошибок
var (
	ErrBasicFieldsNotFilled = errors.New("Не заполнены обязательные поля, невозможно создать пользователя")
)

// CreateAdmin - создаём пользователя для администратора
func CreateAdmin(SQL *settings.SQLServer, Login string, Email string, Password string) error {

	if len(Login) == 0 || len(Password) == 0 || len(Email) == 0 {
		return ErrBasicFieldsNotFilled
	}

	err := SQL.Connect("admin_role_CRUD")

	if err != nil {
		return err
	}

	defer SQL.Disconnect()

	Hash, err := authentication.Argon2GenerateHash(Password, &authentication.HashParams)

	if err != nil {
		return err
	}

	var UserInfo = databases.UserDB{
		Role:    "admin_role_CRUD",
		Email:   Email,
		Phone:   "",
		Name:    Login,
		IsAdmin: true,
	}

	err = databases.PostgreSQLUsersCreateUpdate(UserInfo, Hash, true, false)

	if err != nil {
		return err
	}

	return nil
}

// CreateUser - создаём пользователя для гостя
func CreateUser(SQL *settings.SQLServer, Login string, Email string, Password string) error {

	if len(Login) == 0 || len(Password) == 0 || len(Email) == 0 {
		return ErrBasicFieldsNotFilled
	}

	err := SQL.Connect("admin_role_CRUD")

	if err != nil {
		return err
	}

	defer SQL.Disconnect()

	Hash, err := authentication.Argon2GenerateHash(Password, &authentication.HashParams)

	if err != nil {
		return err
	}

	var UserInfo = databases.UserDB{
		Role:    "guest_role_read_only",
		Email:   Email,
		Phone:   "",
		Name:    Login,
		IsAdmin: false,
	}

	err = databases.PostgreSQLUsersCreateUpdate(UserInfo, Hash, true, false)

	if err != nil {
		return err
	}

	return nil
}
