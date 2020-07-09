// Package settings - Реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

import (
	"crypto/rand"
	"log"
	"math/big"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"strings"
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
func (SQLsrv *SQLServer) AutoFillRoles() {

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "guest_role_read_only",
		Desc:    "Гостевая роль",
		Login:   "recipes_guest",
		Pass:    GeneratePassword(20, 5, 5, 5),
		TRules:  GetTRulesForGuest(),
		Default: true,
		Admin:   false,
	})

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "admin_role_CRUD",
		Desc:    "Администратор",
		Login:   "recipes_admin",
		Pass:    GeneratePassword(20, 5, 5, 5),
		TRules:  GetTRulesForAdmin(),
		Default: false,
		Admin:   true,
	})
}

// GetTRulesForGuest - Возвращает заполненный список ролей по всем таблицам будущей базы данных для гостя
func GetTRulesForGuest() SQLTRules {
	return SQLTRules{
		TRule{
			TName:      "public.\"Files\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"Recipes\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"RecipesIngredients\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"Ingredients\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"ShoppingList\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: true,
		},
	}
}

// GetTRulesForAdmin - Возвращает заполненный список ролей по всем таблицам будущей базы данных для админа
func GetTRulesForAdmin() SQLTRules {
	return SQLTRules{
		TRule{
			TName:      "public.\"Files\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"Recipes\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"RecipesIngredients\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"Ingredients\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
		TRule{
			TName:      "public.\"ShoppingList\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
	}
}

// CreateDatabase - Создаёт базу данных если её нет
func (SQLsrv *SQLServer) CreateDatabase() {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		// Создаём базу данных
		databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, "", true))
		databases.PostgreSQLCreateDatabase(SQLsrv.DbName)
		databases.PostgreSQLCloseConn()

		// Заполняем базу данных
		databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, SQLsrv.DbName, false))
		databases.PostgreSQLCreateTables()
		for _, currole := range SQLsrv.Roles {

			databases.PostgreSQLCreateRole(currole.Login, currole.Pass, SQLsrv.DbName)

			for _, tablerule := range currole.TRules {

				databases.PostgreSQLGrantRightsToRole(currole.Login, tablerule.TName, formRightsArray(tablerule))
			}
		}
		databases.PostgreSQLCloseConn()

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}
}

// formRightsArray - формирует массив прав для таблицы
func formRightsArray(rule TRule) []string {
	var result []string

	if rule.SELECT {
		result = append(result, "SELECT")
		result = append(result, "REFERENCES")
	}

	if rule.INSERT {
		result = append(result, "INSERT")
	}

	if rule.UPDATE {
		result = append(result, "UPDATE")
	}

	if rule.DELETE {
		result = append(result, "DELETE")
	}
	return result
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

		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialCharSet))))
		password.WriteString(string(specialCharSet[random.Int64()]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(numberSet))))
		password.WriteString(string(numberSet[random.Int64()]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(upperCharSet))))
		password.WriteString(string(upperCharSet[random.Int64()]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(lowerCharSet))))
		password.WriteString(string(allCharSet[random.Int64()]))
	}

	return password.String()
}
