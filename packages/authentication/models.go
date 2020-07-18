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

// AuthRequestData - запрос при авторизации и регистрации
type AuthRequestData struct {
	Email             string
	Password          string
	ReturnSecureToken bool
}

// AuthResponseData - ответ при авторизации и регистрации
type AuthResponseData struct {
	Token      string
	Email      string
	ExpiresIn  int
	Registered bool
	Role       string
}

// ActiveToken - тип для хранения в списке активных токенов
type ActiveToken struct {
	Token   string
	IssDate time.Time
	ExpDate time.Time
	Role    string
}
