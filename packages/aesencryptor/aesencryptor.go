package aesencryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// GetStringEncrypted - шифрует строку и возвращает ключ шифрования
func GetStringEncrypted(s string) (dataenc string, key []byte, e error) {

	var enc AESencryptor

	err := enc.MakeSecretKey()

	if err != nil {
		return "", []byte{}, err
	}

	es, err := enc.Encrypt(s)

	if err != nil {
		return "", []byte{}, err
	}

	enckey := enc.GetKey()

	return es, enckey, nil
}

// AESencryptor - тип для хранения информации о шифровании
type AESencryptor struct {
	key []byte
}

// GetKey - возвращает слайс байтов ключа шифрования
func (enc *AESencryptor) GetKey() []byte {
	return enc.key
}

// SetKey - устанавливает слайс байтов ключа шифрования
func (enc *AESencryptor) SetKey(newkey []byte) {
	enc.key = newkey
}

// MakeSecretKey - генерит случайный ключ шифрования для AES 256
func (enc *AESencryptor) MakeSecretKey() error {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)

	if err != nil {
		return err
	}

	enc.key = bytes

	return nil
}

// Encrypt - шифрует входящий текст
func (enc *AESencryptor) Encrypt(stringToEncrypt string) (encryptedString string, err error) {

	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(enc.key)
	if err != nil {
		return "", err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt - расшифровывает входящий текст
func (enc *AESencryptor) Decrypt(encryptedString string) (decryptedString string, err error) {

	stren, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(enc.key)
	if err != nil {
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := stren[:nonceSize], stren[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}
