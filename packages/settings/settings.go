// Package settings - Реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

import (
	"math/rand"
	"strings"
	"time"
)

// WServerSettings - Настройки веб сервера
type WServerSettings struct {
	HTTP  int
	HTTPS int
	SQL   SQLServer
}

// SQLServer - Данные для подключения к SQL серверу
type SQLServer struct {
	Type   string
	DbName string
	Addr   string
	Login  string
	Pass   string
	Roles  SQLRoles
}

// SQLRoles - Список ролей которые должны быть созданы на сервере
type SQLRoles []SQLRole

// SQLRole - Роль, которая должна быть создана на сервере
type SQLRole struct {
	Name    string
	Desc    string
	Login   string
	Pass    string
	TRules  SQLTRules
	Default bool
	Admin   bool
}

// SQLTRules - Список прав на отдельные таблицы
type SQLTRules []TRule

// TRule - права для конкретной таблицы
type TRule struct {
	TName      string
	SELECT     bool
	INSERT     bool
	UPDATE     bool
	DELETE     bool
	REFERENCES bool
}

// AutoFillRoles - автозаполняет список ролей для SQL сервера
func (SQLsrv SQLServer) AutoFillRoles() {

	rand.Seed(time.Now().Unix())

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "guest_role_read_only",
		Desc:    "Гостевая роль",
		Login:   "recipes_guest",
		Pass:    GeneratePassword(20, 5, 5, 5),
		TRules:  GetTRulesForGuest(),
		Default: true,
		Admin:   false,
	})

	rand.Seed(time.Now().Unix())

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "admin_role_CRUD",
		Desc:    "Администратор",
		Login:   "recipes_admin",
		Pass:    GeneratePassword(20, 5, 5, 5),
		Default: true,
		Admin:   false,
	})

}

// GetTRulesForGuest - Возвращает заполненный список ролей по всем таблицам будущей базы данных
func GetTRulesForGuest() SQLTRules {
	return SQLTRules{}
}

// GeneratePassword - генерирует случайный пароль
func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {

	var (
		lowerCharSet   = "abcdedfghijklmnopqrst"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "!@#$%&*"
		numberSet      = "0123456789"
		allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	)

	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
