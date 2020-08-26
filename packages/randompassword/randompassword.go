// Package randompassword содержит методы и структуры для создания случайного пароля из чисел и букв английского алфавита нижнего и верхнего регистра
package randompassword

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// ASCII числовые интервалы для подбора случайного пароля (границы последовательностей)
var (
	Numbers   = Int64NumberRange{LowerBorder: 48, UpperBorder: 57}
	UpperCase = Int64NumberRange{LowerBorder: 65, UpperBorder: 90}
	LowerCase = Int64NumberRange{LowerBorder: 97, UpperBorder: 122}
)

// Int64NumberRange - числовая последовательность int64 для хранения интервала
type Int64NumberRange struct {
	LowerBorder int64
	UpperBorder int64
}

// GetNumberDiff - считает длинну числового интервала между верхней и нижней границей
func (CharacterRange *Int64NumberRange) GetNumberDiff() int64 {
	return CharacterRange.UpperBorder - CharacterRange.LowerBorder
}

// GenerateRandomChar - генерируем рандомное число и выбираем из ASCII таблицы символ сохраняем в password
func (CharacterRange *Int64NumberRange) GenerateRandomChar(password *strings.Builder) {

	numberSet := CharacterRange.GetNumberDiff() + 1

	random, _ := rand.Int(rand.Reader, big.NewInt(numberSet))
	password.WriteString(fmt.Sprintf("%c", CharacterRange.LowerBorder+random.Int64()))

}

// NewRandomPassword - создаёт новый случайный пароль заданной длины
func NewRandomPassword(PasswordLen int) string {

	var password strings.Builder

	GenerateRandomPassword(&password, PasswordLen)

	return password.String()
}

// GenerateRandomPassword - генерируем рандомное число и выбираем из ASCII таблицы символ
// при этом сначала генерируется случайное число от 0 до 2 и потом для чисел или букв в построитель
// строки добавляется символ ASCII
func GenerateRandomPassword(password *strings.Builder, PasswordLen int) {

	for i := 0; i < PasswordLen; i++ {

		random, _ := rand.Int(rand.Reader, big.NewInt(3))

		SeqNo := random.Int64()

		switch {
		case SeqNo == 0:
			Numbers.GenerateRandomChar(password)
		case SeqNo == 1:
			UpperCase.GenerateRandomChar(password)
		case SeqNo == 2:
			LowerCase.GenerateRandomChar(password)
		}
	}
}
