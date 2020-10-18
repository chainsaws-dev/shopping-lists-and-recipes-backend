// Package settings - реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

import (
	"errors"
	"log"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/randompassword"
)

// Список типовых ошибок
var (
	ErrRoleNotFound         = errors.New("Роль с указанным именем не найдена")
	ErrDatabaseNotSupported = errors.New("Не реализована поддержка базы данных")
)

// AutoFillRoles - автозаполняет список ролей для SQL сервера
func (SQLsrv *SQLServer) AutoFillRoles() {

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "guest_role_read_only",
		Desc:    "Гостевая роль",
		Login:   "recipes_guest",
		Pass:    randompassword.NewRandomPassword(20),
		TRules:  GetTRulesForGuest(),
		Default: true,
		Admin:   false,
	})

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "admin_role_CRUD",
		Desc:    "Администратор",
		Login:   "recipes_admin",
		Pass:    randompassword.NewRandomPassword(20),
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
		TRule{
			TName:      "secret.\"confirmations\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"password_resets\"",
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
		TRule{
			TName:      "secret.\"confirmations\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
		TRule{
			TName:      "secret.\"password_resets\"",
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

		placeholder := databases.File{
			Filename: "placeholder.jpg",
			Filesize: 0,
			Filetype: "jpg",
			FileID:   "",
		}
		databases.PostgreSQLFileInsert(placeholder)

		for _, currole := range SQLsrv.Roles {

			databases.PostgreSQLCreateRole(currole.Login, currole.Pass, SQLsrv.DbName)

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

	switch {
	case SQLsrv.Type == "PostgreSQL":

		return databases.PostgreSQLConnect(
			databases.PostgreSQLGetConnString(
				ActiveRole.Login,
				ActiveRole.Pass,
				SQLsrv.Addr,
				SQLsrv.DbName,
				false))
	default:
		return ErrDatabaseNotSupported
	}
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

// Disconnect - Разрываем соединение с базой данных
func (SQLsrv *SQLServer) Disconnect() {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		databases.PostgreSQLCloseConn()
	default:
		log.Fatalln(ErrDatabaseNotSupported)
	}
}

// CheckRoleForRead - проверяет роль для разрешения доступа к разделу системы
func (ss WServerSettings) CheckRoleForRead(RoleName string, AppPart string) bool {

	switch {
	case AppPart == "HandleRecipes":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleRecipesSearch":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleShoppingList":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleFiles":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleUsers":
		return checkAdmin(RoleName)
	case AppPart == "HandleSessions":
		return checkAdmin(RoleName)
	default:
		return false
	}
}

// CheckRoleForChange - проверяет роль для разрешения изменений в разделе системы
func (ss WServerSettings) CheckRoleForChange(RoleName string, AppPart string) bool {
	switch {
	case AppPart == "HandleRecipes":
		return checkAdmin(RoleName)
	case AppPart == "HandleRecipesSearch":
		return checkAdmin(RoleName)
	case AppPart == "HandleShoppingList":
		return checkAdmin(RoleName)
	case AppPart == "HandleFiles":
		return checkAdmin(RoleName)
	case AppPart == "HandleUsers":
		return checkAdmin(RoleName)
	case AppPart == "HandleSessions":
		return checkAdmin(RoleName)
	default:
		return false
	}
}

// CheckRoleForDelete - проверяет роль для разрешения доступа к удалению элементов раздела системы
func (ss WServerSettings) CheckRoleForDelete(RoleName string, AppPart string) bool {
	switch {
	case AppPart == "HandleRecipes":
		return checkAdmin(RoleName)
	case AppPart == "HandleRecipesSearch":
		return checkAdmin(RoleName)
	case AppPart == "HandleShoppingList":
		return checkAdmin(RoleName)
	case AppPart == "HandleFiles":
		return checkAdmin(RoleName)
	case AppPart == "HandleUsers":
		return RoleName == "admin_role_CRUD"
	case AppPart == "HandleSessions":
		return checkAdmin(RoleName)
	default:
		return false
	}
}

func checkAdmin(RoleName string) bool {
	return RoleName == "admin_role_CRUD"
}

// CheckExistingRole - проверяет что роль это существующая роль
func (ss WServerSettings) CheckExistingRole(RoleName string) bool {

	for _, role := range ss.SQL.Roles {
		if role.Name == RoleName {
			return true
		}
	}

	return false

}
