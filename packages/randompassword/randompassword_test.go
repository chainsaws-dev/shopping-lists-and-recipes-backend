// Package randompassword содержит методы и структуры для создания случайного пароля из чисел и букв английского алфавита нижнего и верхнего регистра
package randompassword

import (
	"fmt"
	"strings"
	"testing"
)

func TestInt64NumberRange_GetNumberDiff(t *testing.T) {
	type fields struct {
		LowerBorder int64
		UpperBorder int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Test-001",
			fields: fields{
				LowerBorder: 100,
				UpperBorder: 200,
			},
			want: 100,
		},
		{
			name: "Test-002",
			fields: fields{
				LowerBorder: 200,
				UpperBorder: 100,
			},
			want: -100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CharacterRange := &Int64NumberRange{
				LowerBorder: tt.fields.LowerBorder,
				UpperBorder: tt.fields.UpperBorder,
			}
			if got := CharacterRange.GetNumberDiff(); got != tt.want {
				t.Errorf("Int64NumberRange.GetNumberDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

//!+example

func ExampleInt64NumberRange_GetNumberDiff() {
	CharacterRange1 := &Int64NumberRange{
		LowerBorder: 140,
		UpperBorder: 240,
	}

	CharacterRange2 := &Int64NumberRange{
		LowerBorder: 240,
		UpperBorder: 140,
	}

	fmt.Println(CharacterRange1.GetNumberDiff())
	fmt.Println(CharacterRange2.GetNumberDiff())
	// Output:
	// 100
	// -100
}

//!-example

func BenchmarkInt64NumberRange_GetNumberDiff(t *testing.B) {

	t.ResetTimer()

	for i := 0; i <= t.N; i++ {

		CharacterRange := &Int64NumberRange{
			LowerBorder: 140,
			UpperBorder: 240,
		}

		CharacterRange.GetNumberDiff()

	}

}

func TestInt64NumberRange_GenerateRandomChar(t *testing.T) {

	var password strings.Builder

	type args struct {
		password *strings.Builder
	}
	tests := []struct {
		name   string
		fields Int64NumberRange
		args   args
		want   int
	}{
		{
			name:   "Test-001",
			fields: Numbers,
			args: args{
				password: &password,
			},
			want: 1,
		},
		{
			name:   "Test-001",
			fields: UpperCase,
			args: args{
				password: &password,
			},
			want: 2,
		},
		{
			name:   "Test-001",
			fields: LowerCase,
			args: args{
				password: &password,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			CharacterRange := &tt.fields
			CharacterRange.GenerateRandomChar(tt.args.password)

			reschar := tt.args.password.String()

			got := len(reschar)

			if got != tt.want {
				t.Errorf("Int64NumberRange.GenerateRandomChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInt64NumberRange_GenerateRandomChar(t *testing.B) {

	t.ResetTimer()

	for i := 0; i <= t.N; i++ {

		var password strings.Builder
		LowerCase.GenerateRandomChar(&password)

	}

}

func TestGenerateRandomPassword(t *testing.T) {

	var tp strings.Builder

	type args struct {
		password    *strings.Builder
		PasswordLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test-001",
			args: args{
				password:    &tp,
				PasswordLen: 10,
			},
			want: 10,
		},
		{
			name: "Test-002",
			args: args{
				password:    &tp,
				PasswordLen: 20,
			},
			want: 30,
		},
		{
			name: "Test-003",
			args: args{
				password:    &tp,
				PasswordLen: 30,
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateRandomPassword(tt.args.password, tt.args.PasswordLen)

			reschar := tt.args.password.String()

			got := len(reschar)

			if got != tt.want {
				t.Errorf("GenerateRandomPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGenerateRandomPassword(t *testing.B) {

	t.ResetTimer()

	for i := 0; i <= t.N; i++ {

		var password strings.Builder
		GenerateRandomPassword(&password, 60)

	}

}

func TestNewRandomPassword(t *testing.T) {
	type args struct {
		PasswordLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test-01",
			args: args{
				PasswordLen: 30,
			},
			want: 30,
		},
		{
			name: "Test-02",
			args: args{
				PasswordLen: 60,
			},
			want: 60,
		},
		{
			name: "Test-03",
			args: args{
				PasswordLen: 90,
			},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(NewRandomPassword(tt.args.PasswordLen)); got != tt.want {
				t.Errorf("NewRandomPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewRandomPassword(t *testing.B) {

	t.ResetTimer()

	for i := 0; i <= t.N; i++ {

		NewRandomPassword(60)

	}

}
