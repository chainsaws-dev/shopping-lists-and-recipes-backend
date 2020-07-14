package settings

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

// TRule - Права для конкретной таблицы
type TRule struct {
	TName      string
	SELECT     bool
	INSERT     bool
	UPDATE     bool
	DELETE     bool
	REFERENCES bool
}
