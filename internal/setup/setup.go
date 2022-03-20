// Package setup - выполняет начальную настройку и создание структуры папок при первом запуске сервера для рецептов и списка покупок
package setup

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"shopping-lists-and-recipes/internal/settings"
	"shopping-lists-and-recipes/packages/admin"
	"shopping-lists-and-recipes/packages/messages"
	"shopping-lists-and-recipes/packages/shared"
	"strconv"
	"strings"
)

// Список типовых ошибок
var (
	ErrDeleteInterrupted   = errors.New("удаление базы данных, таблиц и ролей завершилось с ошибкой")
	ErrCreationInterrupted = errors.New("создание базы данных, таблиц и ролей завершилось с ошибкой")
)

// ServerSettings - общие настройки сервера
var ServerSettings settings.WServerSettings

// APIkeys - API ключи которым разрешено работать с API
var APIkeys = []string{
	"AIzaSyB3Jr8tp5wotjeS-re9iBSgX2b1zbM0Fx4",
}

// InitialSettings - интерактивно спрашивает у пользователя параметры настроек
func InitialSettings(initpar InitParams) {

	if !СheckExists("settings.json") || initpar.ForceSetup {

		if !initpar.ForceSetup {
			log.Println("Не найден файл settings.json. Запущена процедура начальной настройки.")
		} else {
			log.Println("Принудительно запущена процедура начальной настройки.")
		}

		fmt.Println("")
		fmt.Println("*******************************************************************")
		fmt.Println("* Добро пожаловать в мастер настройки сервера рецептов и покупок! *")
		fmt.Println("* Для настройки сервера необходимо создать конфигурационный файл. *")
		fmt.Println("* Он будет создан по результатам ваших ответов на вопросы.        *")
		fmt.Println("*                           ВНИМАНИЕ!                             *")
		fmt.Println("*  Некоторые значения настроек потребуют запуска сервера от sudo  *")
		fmt.Println("*******************************************************************")

		// WEB

		AskInt("Укажите порт для http соединений (например 80): ", &ServerSettings.HTTP)

		AskInt("Укажите порт для https соединений (например 443): ", &ServerSettings.HTTPS)

		// 2FA
		var TwoFactor string
		AskString("Обязательная двухфакторная авторизация (Да или Нет): ", &TwoFactor)
		TwoFactor = strings.ToLower(TwoFactor)

		if TwoFactor == "да" || TwoFactor == "д" {
			ServerSettings.TFO = true
		}

		var ConfirmEmails string
		AskString("Подтверждать электронную почту письмом со ссылкой (Да или Нет): ", &ConfirmEmails)
		ConfirmEmails = strings.ToLower(ConfirmEmails)

		if ConfirmEmails == "да" || ConfirmEmails == "д" {
			// SMTP

			AskString("Укажите адрес SMTP сервера для отправки почты: ", &ServerSettings.SMTP.SMTP)

			AskInt("Укажите порт для соединения с SMTP: ", &ServerSettings.SMTP.SMTPPort)

			AskString("Укажите логин пользователя SMTP сервера: ", &ServerSettings.SMTP.Login)

			AskString("Укажите пароль пользователя SMTP сервера: ", &ServerSettings.SMTP.Pass)

			ServerSettings.SMTP.Use = true

			messages.SetCredentials(ServerSettings.SMTP)
		} else {
			ServerSettings.SMTP = settings.CredSMTP{}
		}

		// SQL

		AskString("Укажите тип сервера баз данных (поддерживается PostgreSQL): ", &ServerSettings.SQL.Type)

		AskString("Укажите адрес сервера баз данных: ", &ServerSettings.SQL.Addr)

		AskString("Укажите желаемое имя базы данных: ", &ServerSettings.SQL.DbName)

		AskString("Укажите имя суперпользователя базы данных: ", &ServerSettings.SQL.Login)

		AskString("Укажите пароль суперпользователя базы данных: ", &ServerSettings.SQL.Pass)

		var CreateDB string
		AskString("Создать базу данных с таблицами и ролями (Да или Нет): ", &CreateDB)
		CreateDB = strings.ToLower(CreateDB)

		if CreateDB == "да" || CreateDB == "д" {

			// Создаём базу данных
			ServerSettings.SQL.AutoFillRoles()
			err := StartCreateDatabase()
			if err == nil {
				if initpar.CreateAdmin {
					SetDefaultAdmin(initpar.AdminLogin, initpar.AdminPass, initpar.WebsiteURL)
				}
			}
		}

		WriteToJSON()

	} else {

		log.Println("Читаем файл настроек settings.json...")

		bytes, err := ioutil.ReadFile("settings.json")

		shared.WriteErrToLog(err)

		err = json.Unmarshal(bytes, &ServerSettings)

		shared.WriteErrToLog(err)

		messages.SetCredentials(ServerSettings.SMTP)

		DbHost := os.Getenv("DATABASE_HOST")

		if len(DbHost) > 0 {
			ServerSettings.SQL.Addr = DbHost
		} else {
			ServerSettings.SQL.Addr = "localhost"
		}

		log.Println("Файл настроек settings.json успешно прочитан")

		// Удаляем базу данных и роли
		if initpar.DropDb {
			err = StartDropDatabase()
			if err != nil {
				log.Println(err)
			}
		}

		// Пересоздаём базу данных без перенастройки
		if initpar.CreateDb {

			ServerSettings.SQL.AutoFillRoles()
			err := StartCreateDatabase()
			if err == nil {
				if initpar.CreateAdmin {
					SetDefaultAdmin(initpar.AdminLogin, initpar.AdminPass, initpar.WebsiteURL)
				}
			}
			WriteToJSON()
		}

	}

}

// AskString - спрашивает вопрос и сохраняет ответ как строку в заданное поле
func AskString(Question string, fieldToWriteIn *string) {

	var inputstring string

	fmt.Println("")
	fmt.Println(Question)
	fmt.Scanln(&inputstring)
	*fieldToWriteIn = inputstring

}

// AskInt - спрашивает вопрос и сохраняет ответ как int в заданное поле
func AskInt(Question string, fieldToWriteIn *int) {

	var inputstring string

	fmt.Println("")
	fmt.Println(Question)
	fmt.Scanln(&inputstring)
	value, err := strconv.Atoi(inputstring)
	shared.WriteErrToLog(err)
	*fieldToWriteIn = value

}

// FolderCreate - создаёт всю необходимую структуру папок на сервере
func FolderCreate() {

	if !СheckExists("public") {
		shared.WriteErrToLog(os.Mkdir("public", 0700))
	}

	if !СheckExists("public/frontend") {
		shared.WriteErrToLog(os.Mkdir("public/frontend", 0700))
	}

	if !СheckExists("public/templates") {
		shared.WriteErrToLog(os.Mkdir("public/templates", 0700))
	}

	if !СheckExists("public/uploads") {
		shared.WriteErrToLog(os.Mkdir("public/uploads", 0700))
	}

	if !СheckExists("logs") {
		shared.WriteErrToLog(os.Mkdir("logs", 0700))
	}
}

// СheckExists - проверяем что файл или папка существует
func СheckExists(filename string) bool {

	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}

// WriteToJSON - записывает объект настроек в JSON файл
func WriteToJSON() {
	bytes, err := json.Marshal(ServerSettings)

	shared.WriteErrToLog(err)

	setfile, err := os.Create("settings.json")

	shared.WriteErrToLog(err)

	_, err = setfile.Write(bytes)

	shared.WriteErrToLog(err)

	log.Println("Файл настроек settings.json успешно создан")
}

// SetDefaultAdmin - позволяет настроить администратора по умолчанию
func SetDefaultAdmin(login string, password string, websiteurl string) string {

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	var Email string
	var LoginAdmin string
	var PasswordAdmin string
	var URI string

	if len(login) > 1 && re.MatchString(login) {
		Email = login
		LoginAdmin = login
	}

	if len(password) >= 6 {
		PasswordAdmin = password
	}

	for len(Email) < 1 || !re.MatchString(Email) {

		AskString("Укажите e-mail администратора вебсайта: ", &Email)

		if len(Email) < 1 || !re.MatchString(Email) {
			fmt.Println("Некорректный e-mail!")
		}
	}

	for len(LoginAdmin) < 1 {
		AskString("Укажите логин администратора вебсайта: ", &LoginAdmin)

		if len(LoginAdmin) < 1 {
			fmt.Println("Некорректный логин!")
		}
	}

	for len(PasswordAdmin) < 6 {
		AskString("Укажите пароль администратора вебсайта: ", &PasswordAdmin)
		if len(PasswordAdmin) < 6 {
			fmt.Println("Пароль должен быть больше шести символов!")
		}
	}

	err := admin.CreateAdmin(&ServerSettings.SQL, LoginAdmin, Email, PasswordAdmin, ServerSettings.SMTP.Use, ServerSettings.SQL.ConnPool)

	shared.WriteErrToLog(err)

	if ServerSettings.SMTP.Use {
		if len(websiteurl) < 1 {
			for len(URI) < 1 {
				AskString("Укажите адрес вебсайта с портом (например: http://localhost:8080): ", &URI)
			}
		} else {
			URI = websiteurl
		}

		messages.SendEmailConfirmationLetter(&ServerSettings.SQL, Email, URI, ServerSettings.SQL.ConnPool)
	}

	log.Println("Администратор сайта создан")

	return Email

}

// StartCreateDatabase - запускает в фоне процесс создания базы данных
func StartCreateDatabase() error {

	donech := make(chan bool)
	go ServerSettings.SQL.CreateDatabase(donech)

	if <-donech {
		log.Println("Процедура создания базы данных завершена")
		return nil
	}

	return ErrCreationInterrupted

}

// StartDropDatabase - запускает в фоне процесс удаления базы данных и ролей
func StartDropDatabase() error {

	donech := make(chan bool)
	go ServerSettings.SQL.DropDatabase(donech)

	if <-donech {
		log.Println("Процедура удаления базы данных завершена")
		return nil
	}

	return ErrDeleteInterrupted

}
