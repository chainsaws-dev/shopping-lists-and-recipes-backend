package secondfactor

import (
	"bytes"
	"database/sql"
	"errors"
	"image/png"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/packages/aesencryptor"

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
func (usf *UserSecondFactor) GenerateUserKey(dbc *sql.DB) error {

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

	err = databases.PostgreSQLChangeSecondFactorSecret(totps, dbc)

	if err != nil {
		return err
	}

	return nil
}

// GetQR - получает буфер из байтов содержащий данные QR кода для приложения аутентификатора
func (usf *UserSecondFactor) GetQR(width int, height int, dbc *sql.DB) (bytes.Buffer, error) {

	var b bytes.Buffer

	err := usf.GenerateUserKey(dbc)

	if err != nil {
		return b, err
	}

	img, err := usf.key.Image(width, height)

	if err != nil {
		return b, err
	}

	err = png.Encode(&b, img)

	return b, err
}

// EnableTOTP - проверяет правильность кода и сохраняет секрет если он верный
func EnableTOTP(Passcode string, u databases.User, dbc *sql.DB) error {

	result, err := databases.PostgreSQLGetSecretByUserID(u.GUID, dbc)

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

		err = databases.PostgreSQLUpdateSecondFactorConfirmed(true, result.UserID, dbc)

		if err != nil {
			return err
		}

		return nil
	}

	return ErrSecretNotSaved
}

// Validate - проверяет код токена против секрета из базы
func Validate(Passcode string, u databases.User, dbc *sql.DB) (bool, error) {

	result, err := databases.PostgreSQLGetSecretByUserID(u.GUID, dbc)

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
