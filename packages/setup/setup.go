// Package setup - Выполняет начальную настройку и создание структуры папок при первом запуске сервера для рецептов и списка покупок
package setup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/settings"
	"os"
	"strconv"
)

// ServerSettings - Общие настройки сервера
var ServerSettings settings.WServerSettings

// InitialSettings - интерактивно спрашивает у пользователя параметры настроек
func InitialSettings(forcesetup bool) *settings.WServerSettings {

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

		ServerSettings.SQL.Type = "PostgreSQL"

		AskString("Укажите адрес сервера баз данных PostgreSQL: ", &ServerSettings.SQL.Addr)

		AskString("Укажите желаемое имя базы данных PostgreSQL: ", &ServerSettings.SQL.DbName)

		AskString("Укажите имя суперпользователя PostgreSQL: ", &ServerSettings.SQL.Login)

		AskString("Укажите пароль суперпользователя PostgreSQL: ", &ServerSettings.SQL.Pass)

		ServerSettings.SQL.AutoFillRoles()

		bytes, err := json.Marshal(ServerSettings)

		WriteErrToLog(err)

		setfile, err := os.Create("settings.json")
		defer setfile.Close()

		WriteErrToLog(err)

		setfile.Write(bytes)

	} else {
		bytes, err := ioutil.ReadFile("settings.json")

		WriteErrToLog(err)

		err = json.Unmarshal(bytes, &ServerSettings)

		WriteErrToLog(err)
	}

	return &ServerSettings
}

// AskString - Спрашивает вопрос и сохраняет ответ как строку в заданное поле
func AskString(Question string, fieldToWriteIn *string) {

	var inputstring string

	fmt.Println("")
	fmt.Println(Question)
	fmt.Scanln(&inputstring)
	*fieldToWriteIn = inputstring

}

// AskInt - Спрашивает вопрос и сохраняет ответ как int в заданное поле
func AskInt(Question string, fieldToWriteIn *int) {

	var inputstring string

	fmt.Println("")
	fmt.Println(Question)
	fmt.Scanln(&inputstring)
	value, err := strconv.Atoi(inputstring)
	WriteErrToLog(err)
	*fieldToWriteIn = value

}

// FolderCreate - создаёт всю необходимую структуру папок на сервере
func FolderCreate() {

	if !СheckExists("public") {
		WriteErrToLog(os.Mkdir("public", 0700))
	}

	if !СheckExists("public/frontend") {
		WriteErrToLog(os.Mkdir("public/frontend", 0700))
	}

	if !СheckExists("public/uploads") {
		WriteErrToLog(os.Mkdir("public/uploads", 0700))
	}

	if !СheckExists("logs") {
		WriteErrToLog(os.Mkdir("logs", 0700))
	}
}

// СheckExists - проверяем что файл или папка существует
func СheckExists(filename string) bool {

	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}

// WriteErrToLog - пишем ошибку в лог
func WriteErrToLog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
