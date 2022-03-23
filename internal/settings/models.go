package settings

import "github.com/jackc/pgx/v4/pgxpool"

// WServerSettings - настройки веб сервера
type WServerSettings struct {
	HTTP  int
	HTTPS int
	SMTP  CredSMTP
	SQL   SQLServer
	TFO   bool
	Lang  string
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
	Type        string
	DbName      string
	Addr        string
	Login       string
	Pass        string
	Roles       SQLRoles
	SSL         bool
	MaxConnPool int
	Connected   bool          `json:"-"`
	ConnPool    *pgxpool.Pool `json:"-"`
}

// SQLRoles - список ролей которые должны быть созданы на сервере
type SQLRoles []SQLRole

// SQLRole - роль, которая должна быть создана на сервере
type SQLRole struct {
	Name    string
	Desc    string
	Default bool
	Admin   bool
}
