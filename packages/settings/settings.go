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
	Roles  []SQLRole
}

// SQLRole - Роль, которая должна быть создана на сервере
type SQLRole struct {
	Name    string
	Desc    string
	Login   string
	Pass    string
	TRules  []TRule
	Default bool
	Admin   bool
}

// TRule - права для конкретной таблицы
type TRule struct {
	TName      string
	SELECT     bool
	INSERT     bool
	UPDATE     bool
	DELETE     bool
	REFERENCES bool
}
