package auth

// Argon2Params - параметры хеширования Argon 2
type Argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}
