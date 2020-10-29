package secondfactor

import (
	"bytes"
	"errors"
	"image/png"
	"shopping-lists-and-recipes/packages/aesencryptor"
	"shopping-lists-and-recipes/packages/databases"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// Список типовых ошибок
var (
	ErrSecretNotSaved = errors.New("Секретный ключ не сохранён, неверный код")
)

// UserSecondFactor - тип агрегирующий в себе данные о пользователе
type UserSecondFactor struct {
	URL  string
	User databases.User
	key  *otp.Key
}

// GenerateUserKey - создаёт новый ключ пользователя
func (usf *UserSecondFactor) GenerateUserKey() error {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      usf.URL,
		AccountName: usf.User.Email,
	})

	if err != nil {
		return err
	}

	usf.key = key

	es, encrkey, err := aesencryptor.GetStringEncrypted(key.Secret())

	if err != nil {
		return err
	}

	totps := databases.TOTPSecret{
		UserID:    usf.User.GUID,
		Secret:    es,
		EncKey:    encrkey,
		Confirmed: false,
	}

	databases.PostgreSQLChangeSecondFactorSecret(totps)

	return nil
}

// GetQR - получает буфер из байтов содержащий данные QR кода для приложения аутентификатора
func (usf *UserSecondFactor) GetQR(width int, height int) (bytes.Buffer, error) {

	usf.GenerateUserKey()

	// Convert TOTP key into a PNG
	var b bytes.Buffer

	img, err := usf.key.Image(width, height)

	if err != nil {
		return b, err
	}

	png.Encode(&b, img)

	return b, err
}

// EnableTOTP - проверяет правильность кода и сохраняет секрет если он верный
func EnableTOTP(Passcode string, u databases.User) error {

	result, err := databases.PostgreSQLGetSecretByUserID(u.GUID)

	if err != nil {
		return err
	}

	// Расшифровываем строку
	var encr aesencryptor.AESencryptor

	encr.SetKey(result.EncKey)

	result.Secret, err = encr.Decrypt(result.Secret)

	if err != nil {
		return err
	}

	valid := totp.Validate(Passcode, result.Secret)

	if valid {
		result.Confirmed = true
		err = databases.PostgreSQLChangeSecondFactorSecret(result)

		if err != nil {
			return err
		}

		u.SecondFactor = true
		_, err = databases.PostgreSQLUsersInsertUpdate(u, "", false, true)

		return nil
	}

	return ErrSecretNotSaved
}

// Validate - проверяет код токена против секрета из базы
func Validate(Passcode string, u databases.User) (bool, error) {

	result, err := databases.PostgreSQLGetSecretByUserID(u.GUID)

	if err != nil {
		return false, err
	}

	// Расшифровываем строку
	var encr aesencryptor.AESencryptor

	encr.SetKey(result.EncKey)

	result.Secret, err = encr.Decrypt(result.Secret)

	if err != nil {
		return false, err
	}

	valid := totp.Validate(Passcode, result.Secret)

	if valid {
		return true, nil
	}

	return false, nil
}
