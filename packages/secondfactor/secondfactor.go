package secondfactor

import (
	"bytes"
	"errors"
	"image/png"
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

// generateUserKey - создаёт новый ключ пользователя
func (usf *UserSecondFactor) generateUserKey() error {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      usf.URL,
		AccountName: usf.User.Email,
	})

	if err != nil {
		return err
	}

	usf.key = key

	//databases.PostgreSQLSetSecret(usf.User.ID, key.Secret(), false)

	return nil
}

// GetQR - получает буфер из байтов содержащий данные QR кода для приложения аутентификатора
func (usf *UserSecondFactor) GetQR(width int, height int, new bool) (bytes.Buffer, error) {

	if new {
		usf.generateUserKey()
	}

	// Convert TOTP key into a PNG
	var b bytes.Buffer

	img, err := usf.key.Image(width, height)

	if err != nil {
		return b, err
	}

	png.Encode(&b, img)

	return b, err
}

// SaveSecret - проверяет правильность кода и сохраняет секрет если он верный
func SaveSecret(Passcode string) error {

	// TODO

	secret, err := "", errors.New("Not implemented") //databases.PostgreSQLGetSecret(usf.User.ID)

	if err != nil {
		return err
	}

	valid := totp.Validate(Passcode, secret)

	if valid {
		//databases.PostgreSQLSetSecret(UserID, secret, true)
		return nil
	}

	return ErrSecretNotSaved
}

// Validate - проверяет код токена против секрета из базы
func Validate(Passcode string) (bool, error) {

	// TODO

	secret, err := "", errors.New("Not implemented") //databases.PostgreSQLGetSecret(usf.User.ID)

	if err != nil {
		return false, err
	}

	valid := totp.Validate(Passcode, secret)

	if valid {
		return true, nil
	}

	return false, nil
}
