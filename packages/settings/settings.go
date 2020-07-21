// Package settings - реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"strings"
)

// Список типовых ошибок
var (
	ErrRoleNotFound = errors.New("Роль с указанным именем не найдена")
)

// AutoFillRoles - автозаполняет список ролей для SQL сервера
func (SQLsrv *SQLServer) AutoFillRoles() {

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "guest_role_read_only",
		Desc:    "Гостевая роль",
		Login:   "recipes_guest",
		Pass:    GeneratePassword(20, 5, 5),
		TRules:  GetTRulesForGuest(),
		Default: true,
		Admin:   false,
	})

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "admin_role_CRUD",
		Desc:    "Администратор",
		Login:   "recipes_admin",
		Pass:    GeneratePassword(20, 5, 5),
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
			REFERENCES: false,
		},
		TRule{
			TName:      "public.\"Recipes\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "public.\"RecipesIngredients\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "public.\"Ingredients\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "public.\"ShoppingList\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"users\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"hashes\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
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
		TRule{
			TName:      "secret.\"users\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
		TRule{
			TName:      "secret.\"hashes\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
	}
}

// CreateDatabase - Создаёт базу данных если её нет
func (SQLsrv *SQLServer) CreateDatabase(donech chan bool) {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		// Создаём базу данных
		err := databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, "", true))
		if err != nil {
			log.Fatalln(err)
		}
		databases.PostgreSQLCreateDatabase(SQLsrv.DbName)
		databases.PostgreSQLCloseConn()

		// Заполняем базу данных
		err = databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, SQLsrv.DbName, false))
		if err != nil {
			log.Fatalln(err)
		}
		databases.PostgreSQLCreateTables()
		databases.PostgreSQLFileInsert("placeholder.jpg", 0, "jpg", "")

		for _, currole := range SQLsrv.Roles {

			databases.PostgreSQLCreateRole(currole.Login, currole.Pass, SQLsrv.DbName, currole.Admin)

			for _, tablerule := range currole.TRules {

				databases.PostgreSQLGrantRightsToRole(currole.Login, tablerule.TName, formRightsArray(tablerule))
			}
		}
		databases.PostgreSQLCloseConn()

		donech <- true

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}
}

// Connect - Соединяемся с базой данных
func (SQLsrv *SQLServer) Connect(RoleName string) error {

	ActiveRole, err := FindRoleInRoles(RoleName, SQLsrv.Roles)

	if err != nil {
		return err
	}

	return databases.PostgreSQLConnect(
		databases.PostgreSQLGetConnString(
			ActiveRole.Login,
			ActiveRole.Pass,
			SQLsrv.Addr,
			SQLsrv.DbName,
			false))
}

// FindRoleInRoles - Ищем роль в списке ролей по имени
func FindRoleInRoles(RoleName string, Roles SQLRoles) (SQLRole, error) {
	for _, si := range Roles {
		if si.Name == RoleName {
			return si, nil
		}
	}
	return SQLRole{}, ErrRoleNotFound
}

// Disconnect - Разрываем соединение с базой данных
func (SQLsrv *SQLServer) Disconnect() {
	databases.PostgreSQLCloseConn()
}

// formRightsArray - формирует массив прав для таблицы
func formRightsArray(rule TRule) []string {
	var result []string

	if rule.SELECT {
		result = append(result, "SELECT")
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

	if rule.REFERENCES {
		result = append(result, "REFERENCES")
	}

	return result
}

// GeneratePassword - генерирует случайный пароль
func GeneratePassword(passwordLength, minNum, minUpperCase int) string {

	var (
		lowerCharSet = "abcdedfghijklmnopqrst"
		upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numberSet    = "0123456789"
		allCharSet   = lowerCharSet + upperCharSet + numberSet
	)

	var password strings.Builder

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

	//Set lowercase
	remainingLength := passwordLength - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(lowerCharSet))))
		password.WriteString(string(allCharSet[random.Int64()]))
	}

	return password.String()
}
