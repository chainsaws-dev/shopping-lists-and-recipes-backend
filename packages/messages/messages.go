// Package messages - отвечает за отправку писем
package messages

import (
	"bytes"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/internal/settings"
	"shopping-lists-and-recipes/packages/authentication"
	"text/template"

	send "github.com/go-mail/mail"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseGlob("public/templates/*")
	if err != nil {
		log.Fatalln(err)
	}
}

// SetCredentials - заполняет данные для входа на почту
func SetCredentials(sc settings.CredSMTP) {
	SendCred = sc
}

// SendCred - данные для отправки почты
var SendCred = settings.CredSMTP{}

// SendEmail - отправляет письмо по заданным адресам
func SendEmail(Recepients []string, LetterBodyHTML string, LetterSubject string) {

	m := send.NewMessage()
	m.SetHeader("From", SendCred.Login)
	m.SetHeader("To", Recepients...)
	m.SetHeader("Subject", LetterSubject)
	m.SetBody("text/html", LetterBodyHTML)

	d := send.NewDialer(SendCred.SMTP, SendCred.SMTPPort, SendCred.Login, SendCred.Pass)

	d.StartTLSPolicy = send.MandatoryStartTLS

	err := d.DialAndSend(m)

	if err != nil {
		log.Println("Письмо не отправлено")
		log.Println(err)
	} else {
		for _, OneRec := range Recepients {
			log.Printf("Отправили пользователю %v письмо", OneRec)
		}
	}
}

// GetStringTemplate - получаем шаблон в виде строки
func GetStringTemplate(TemplateName string, ObjectToInsert string) string {
	var Etpl bytes.Buffer
	err := tpl.ExecuteTemplate(&Etpl, TemplateName, ObjectToInsert)

	if err != nil {
		log.Fatalln(err)
	}

	return Etpl.String()
}

// SendEmailConfirmationLetter - отправляет письмо с ссылкой для подтверждения электронной почты
func SendEmailConfirmationLetter(SQL *settings.SQLServer, Recepient string, ReqHost string, dbc *sql.DB) {

	if SendCred.Use {

		log.Printf("Отправляем пользователю %v письмо...", Recepient)

		fn := sha1.New()

		tokenb, err := authentication.GenerateRandomBytes(32)

		if err != nil {
			log.Fatal(err)
		}

		strtoken := fmt.Sprintf("%x", fn.Sum(tokenb))

		prurl := fmt.Sprintf("%v/confirm-email?Token=%v", ReqHost, url.PathEscape(strtoken))

		go SendEmail([]string{Recepient}, GetStringTemplate("EmailConfirm.gohtml", prurl), "Подтвердите электронную почту")

		go SaveTokenForUser(SQL, strtoken, "secret.confirmations", Recepient, dbc)

	}
}

// SaveTokenForUser - сохраняем токен доступа в базу данных в заданную таблицу
func SaveTokenForUser(SQL *settings.SQLServer, strtoken string, TableName string, Recepient string, dbc *sql.DB) {

	log.Printf("Сохраняем токен для пользователя %v...", Recepient)

	err := databases.PostgreSQLSaveAccessToken(strtoken, Recepient, TableName, dbc)

	if err != nil {
		log.Fatal(err)
	}

	defer dbc.Close()

	log.Printf("Сохранили токен для пользователя %v", Recepient)

}

// SendEmailPasswordReset - отправляет письмо с ссылкой для сброса пароля
func SendEmailPasswordReset(SQL *settings.SQLServer, Recepient string, ReqHost string, dbc *sql.DB) {

	if SendCred.Use {

		log.Printf("Отправляем пользователю %v письмо...", Recepient)

		fn := sha1.New()

		tokenb, err := authentication.GenerateRandomBytes(32)

		if err != nil {
			log.Fatal(err)
		}

		strtoken := fmt.Sprintf("%x", fn.Sum(tokenb))

		prurl := fmt.Sprintf("%v/reset-password?Token=%v", ReqHost, url.PathEscape(strtoken))

		go SendEmail([]string{Recepient}, GetStringTemplate("EmailPasswordReset.gohtml", prurl), "Сброс пароля")

		go SaveTokenForUser(SQL, strtoken, "secret.password_resets", Recepient, dbc)

	}
}
