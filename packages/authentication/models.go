package authentication

import "time"

// Argon2Params - параметры хеширования Argon 2
type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// AuthRequestData - запрос при авторизации
type AuthRequestData struct {
	Email             string
	Password          string
	ReturnSecureToken bool
}

// AuthResponseData - ответ при авторизации и регистрации
type AuthResponseData struct {
	Token        string
	Email        string
	ExpiresIn    int
	Registered   bool
	Role         string
	SecondFactor TOTP
}

// AuthSignUpRequestData - запрос при регистрации
type AuthSignUpRequestData struct {
	Email             string
	Name              string
	Password          string
	ReturnSecureToken bool
}

// ActiveToken - тип для хранения в списке активных токенов
type ActiveToken struct {
	Email        string
	Token        string
	Session      []byte
	IssDate      time.Time
	ExpDate      time.Time
	Role         string
	IP           string
	UserAgent    string
	SecondFactor TOTP
}

// TOTP - содержит информацию о проверке
// Time Based One Time Password
// включена ли она и каков результат проверки
type TOTP struct {
	Enabled     bool
	CheckResult bool
}
