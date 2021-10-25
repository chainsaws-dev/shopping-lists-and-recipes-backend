package settings

// WServerSettings - настройки веб сервера
type WServerSettings struct {
	HTTP  int
	HTTPS int
	SMTP  CredSMTP
	SQL   SQLServer
	TFO   bool
}

// CredSMTP - данные для входа SMTP
type CredSMTP struct {
	Use      bool
	SMTP     string
	SMTPPort int
	Login    string
	Pass     string
}

// SQLServer - данные для подключения к SQL серверу
type SQLServer struct {
	Type   string
	DbName string
	Addr   string
	Login  string
	Pass   string
	Roles  SQLRoles
}

// SQLRoles - список ролей которые должны быть созданы на сервере
type SQLRoles []SQLRole

// SQLRole - роль, которая должна быть создана на сервере
type SQLRole struct {
	Name    string
	Desc    string
	Login   string
	Pass    string
	TRules  SQLTRules
	Default bool
	Admin   bool
}

// SQLTRules - список прав на отдельные таблицы
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
