// Package settings - реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

import (
	"errors"
	"log"
	"shopping-lists-and-recipes/internal/databases"
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
		Default: true,
		Admin:   false,
	})

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "admin_role_CRUD",
		Desc:    "Администратор",
		Default: false,
		Admin:   true,
	})
}

// DropDatabase - автоматизировано удаляет базу и роли
func (SQLsrv *SQLServer) DropDatabase(donech chan bool) {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		// Подключаемся без контекста базы данных
		SQLsrv.Connect(true)

		// Удаляем базу данных
		databases.PostgreSQLDropDatabase(SQLsrv.DbName, SQLsrv.ConnPool)
		donech <- true

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}
}

// CreateDatabase - Создаёт базу данных если её нет
func (SQLsrv *SQLServer) CreateDatabase(donech chan bool) {
	switch {
	case SQLsrv.Type == "PostgreSQL":

		// Подключаемся без контекста базы данных
		SQLsrv.Connect(true)

		// Создаём базу данных
		databases.PostgreSQLCreateDatabase(SQLsrv.DbName, SQLsrv.ConnPool)

		// Подключаемся под базу данных
		SQLsrv.Connect(false)

		// Заполняем базу данных
		err := databases.PostgreSQLCreateTables(SQLsrv.ConnPool)

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
		databases.PostgreSQLFileChange(placeholder, SQLsrv.ConnPool)

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
func GetConnectionString(SQLsrv *SQLServer, Init bool) string {

	return databases.PostgreSQLGetConnString(
		SQLsrv.Login,
		SQLsrv.Pass,
		SQLsrv.Addr,
		SQLsrv.DbName,
		Init)
}

// Connect - открывает соединение с базой данных
func (SQLsrv *SQLServer) Connect(Init bool) {

	switch {
	case SQLsrv.Type == "PostgreSQL":

		if SQLsrv.Connected {
			SQLsrv.Disconnect()
		}

		SQLsrv.ConnPool = databases.PostgreSQLConnect(GetConnectionString(SQLsrv, Init))

		if SQLsrv.ConnPool == nil {
			SQLsrv.Connected = false
		} else {
			SQLsrv.Connected = true
		}

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
		SQLsrv.ConnPool = nil
	}

}

// Disconnect - разрывает соединение с базой данных
func (SQLsrv *SQLServer) Disconnect() {
	switch {
	case SQLsrv.Type == "PostgreSQL":

		databases.PostgreSQLDisconnect(SQLsrv.ConnPool)
		SQLsrv.Connected = false

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}

	SQLsrv.ConnPool = nil
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
