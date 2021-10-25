package setup

// InitParams - параметры начальной настройки
type InitParams struct {
	ForceSetup     bool
	CreateDb       bool
	DropDb         bool
	ResetRolesPass bool
	CreateRoles    bool
	CreateAdmin    bool
	CleanTokens    bool
	AdminLogin     string
	AdminPass      string
	WebsiteURL     string
}
