// Package settings - реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

import (
	"errors"
	"log"
	"net/http"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/packages/randompassword"
	"shopping-lists-and-recipes/packages/shared"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Список типовых ошибок
var (
	ErrRoleNotFound         = errors.New("Роль с указанным именем не найдена")
	ErrDatabaseNotSupported = errors.New("Не реализована поддержка базы данных")
	ErrDatabaseOffline      = errors.New("База данных недоступна")
	ErrUsupportedDBType     = errors.New("Указан неподдерживаемый тип базы данных")
)

// AutoFillRoles - автозаполняет список ролей для SQL сервера
func (SQLsrv *SQLServer) AutoFillRoles() {

	SQLsrv.Roles = SQLRoles{}

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
			UPDATE:     true,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"hashes\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     true,
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
		TRule{
			TName:      "secret.\"totp\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
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
		TRule{
			TName:      "secret.\"totp\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
	}
}

// DropDatabase - автоматизировано удаляет базу и роли
func (SQLsrv *SQLServer) DropDatabase(donech chan bool) {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		// Удаляем базу данных

		dbc, err := databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, "", true))
		if err != nil {
			log.Fatalln(err)
		}

		databases.PostgreSQLDropDatabase(SQLsrv.DbName, dbc)

		for _, currole := range SQLsrv.Roles {

			databases.PostgreSQLDropRole(currole.Login, dbc)
		}

		dbc.Close()

		donech <- true

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}
}

// CreateDatabase - Создаёт базу данных если её нет
func (SQLsrv *SQLServer) CreateDatabase(donech chan bool, CreateRoles bool) {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		// Создаём базу данных
		cs := databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, "", true)
		dbc, err := databases.PostgreSQLConnect(cs)
		if err != nil {
			log.Fatalln(err)
		}
		databases.PostgreSQLCreateDatabase(SQLsrv.DbName, dbc)
		dbc.Close()

		// Заполняем базу данных
		cs = databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, SQLsrv.DbName, false)
		dbc, err = databases.PostgreSQLConnect(cs)
		if err != nil {
			log.Fatalln(err)
		}

		err = databases.PostgreSQLCreateTables(dbc)

		if err != nil {
			if errors.Is(databases.ErrTablesAlreadyExist, err) {
				donech <- false
				return
			}
		}

		placeholder := databases.File{
			Filename: "placeholder.jpg",
			Filesize: 0,
			Filetype: "jpg",
			FileID:   "",
		}
		databases.PostgreSQLFileChange(placeholder, dbc)

		if CreateRoles {
			for _, currole := range SQLsrv.Roles {

				databases.PostgreSQLCreateRole(currole.Login, currole.Pass, SQLsrv.DbName, dbc)

				for _, tablerule := range currole.TRules {

					databases.PostgreSQLGrantRightsToRole(currole.Login, tablerule.TName, formRightsArray(tablerule), dbc)
				}
			}
		}

		dbc.Close()

		donech <- true

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
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

// GetConnectionString - Формируем строку соединения
func GetConnectionString(SQLsrv *SQLServer, Role string) (string, error) {

	ActiveRole, err := FindRoleInRoles(Role, SQLsrv.Roles)

	if err != nil {
		return "", err
	}

	return databases.PostgreSQLGetConnString(
		ActiveRole.Login,
		ActiveRole.Pass,
		SQLsrv.Addr,
		SQLsrv.DbName,
		false), nil
}

// Connect - открывает соединение с базой данных Postgresql
func (SQLsrv *SQLServer) Connect(w http.ResponseWriter, role string) *pgxpool.Pool {

	switch {
	case SQLsrv.Type == "PostgreSQL":
		cs, err := GetConnectionString(SQLsrv, role)

		if shared.HandleOtherError(w, "Роль не найдена", err, http.StatusServiceUnavailable) {
			return nil
		}

		dbc, err := databases.PostgreSQLConnect(cs)

		if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
			return nil
		}

		return dbc

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}

	return nil
}

// ConnectAsAdmin - подключаемся к базе с ролью администратора
func (SQLsrv *SQLServer) ConnectAsAdmin() *pgxpool.Pool {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		cs, err := GetConnectionString(SQLsrv, "admin_role_CRUD")

		if err != nil {
			log.Println(err)
			return nil
		}

		dbc, err := databases.PostgreSQLConnect(cs)

		if err != nil {
			log.Println(err)
			return nil
		}

		return dbc

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}

	return nil
}

// ConnectAsGuest - подключаемся к базе с ролью гостя
func (SQLsrv *SQLServer) ConnectAsGuest() *pgxpool.Pool {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		cs, err := GetConnectionString(SQLsrv, "guest_role_read_only")

		if err != nil {
			log.Println(err)
			return nil
		}

		dbc, err := databases.PostgreSQLConnect(cs)

		if err != nil {
			log.Println(err)
			return nil
		}

		return dbc

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}

	return nil
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

// CheckRoleForRead - проверяет роль для разрешения доступа к разделу системы
func (ss WServerSettings) CheckRoleForRead(RoleName string, AppPart string) bool {

	switch {
	case AppPart == "CurrentUser":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "CheckSecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "SecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "GetQRCode":
		return ss.CheckExistingRole(RoleName)
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
	case AppPart == "CurrentUser":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "CheckSecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "SecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "GetQRCode":
		return ss.CheckExistingRole(RoleName)
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
	case AppPart == "CurrentUser":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "CheckSecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "SecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "GetQRCode":
		return ss.CheckExistingRole(RoleName)
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
