// Package setup - выполняет начальную настройку и создание структуры папок при первом запуске сервера для рецептов и списка покупок
package setup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/settings"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"os"
	"strconv"
	"strings"
)

// ServerSettings - общие настройки сервера
var ServerSettings settings.WServerSettings

// InitialSettings - интерактивно спрашивает у пользователя параметры настроек
func InitialSettings(forcesetup bool) {

	if !СheckExists("settings.json") || forcesetup {

		if !forcesetup {
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

		// SQL
		ServerSettings.SQL.AutoFillRoles()

		AskString("Укажите тип сервера баз данных (поддерживается PostgreSQL): ", &ServerSettings.SQL.Type)

		AskString("Укажите адрес сервера баз данных: ", &ServerSettings.SQL.Addr)

		AskString("Укажите желаемое имя базы данных: ", &ServerSettings.SQL.DbName)

		AskString("Укажите имя суперпользователя базы данных: ", &ServerSettings.SQL.Login)

		AskString("Укажите пароль суперпользователя базы данных: ", &ServerSettings.SQL.Pass)

		var CreateDB string
		AskString("Создать базу данных с таблицами и ролями (Да или Нет): ", &CreateDB)
		CreateDB = strings.ToLower(CreateDB)

		if CreateDB == "да" || CreateDB == "д" {
			donech := make(chan bool)
			go ServerSettings.SQL.CreateDatabase(donech)

			if <-donech {
				log.Println("Процедура создания базы данных завершена")
			}
		}

		var Email string
		var LoginAdmin string
		var PasswordAdmin string

		AskString("Укажите e-mail администратора вебсайта: ", &Email)

		AskString("Укажите логин администратора вебсайта: ", &LoginAdmin)

		AskString("Укажите пароль администратора вебсайта: ", &PasswordAdmin)

		go ServerSettings.SQL.CreateAdmin(LoginAdmin, Email, PasswordAdmin)

		log.Println("Администратор сайта создан")

		bytes, err := json.Marshal(ServerSettings)

		shared.WriteErrToLog(err)

		setfile, err := os.Create("settings.json")
		defer setfile.Close()

		shared.WriteErrToLog(err)

		_, err = setfile.Write(bytes)

		shared.WriteErrToLog(err)

		log.Println("Файл настроек settings.json успешно создан")

	} else {
		log.Println("Читаем файл настроек settings.json...")

		bytes, err := ioutil.ReadFile("settings.json")

		shared.WriteErrToLog(err)

		err = json.Unmarshal(bytes, &ServerSettings)

		shared.WriteErrToLog(err)

		log.Println("Файл настроек settings.json успешно прочитан")
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
