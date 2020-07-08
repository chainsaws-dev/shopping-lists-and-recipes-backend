// Package setup - Выполняет начальную настройку и создание структуры папок при первом запуске сервера для рецептов и списка покупок
package setup

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// WServerSettings - Настройки веб сервера
type WServerSettings struct {
	http  int
	https int
	SQL   SQLServer
}

// SQLServer - Данные для подключения к SQL серверу
type SQLServer struct {
	Type   string
	DbName string
	Addr   string
	Login  string
	Pass   string
	Roles  []SQLRole
}

// SQLRole - Роль, которая должна быть создана на сервере
type SQLRole struct {
	Name    string
	Desc    string
	Login   string
	Pass    string
	TRules  []TRule
	Default bool
	Admin   bool
}

// TRule - права для конкретной таблицы
type TRule struct {
	TName      string
	SELECT     bool
	INSERT     bool
	UPDATE     bool
	DELETE     bool
	REFERENCES bool
}

// InitialSettings - интерактивно спрашивает у пользователя параметры настроек
func InitialSettings(forcesetup bool) {

	if !СheckExists("settings.json") || forcesetup {

		log.Println("Не найден файл settings.json. Запущена процедура начальной настройки.")

		var inputstring string
		var settings WServerSettings

		fmt.Println("Добро пожаловать в мастер настройки сервера рецептов и покупок!")
		fmt.Println("Для начальной настройки сервера необходимо создать конфигурационный файл.")
		fmt.Println("Он будет создан автоматически по результатам ваших ответов на вопросы.")
		fmt.Println("ВНИМАНИЕ! Для некоторых значений настроек потребуются права суперпользователя!")

		fmt.Println("Укажите порт для http соединений: ")
		fmt.Scanln(&inputstring)
		value, err := strconv.Atoi(inputstring)
		WriteErrToConsole(err)
		settings.http = value

		inputstring = ""
		value = 0

		fmt.Println("Укажите порт для https соединений: ")
		fmt.Scanln(&inputstring)
		value, err = strconv.Atoi(inputstring)
		WriteErrToConsole(err)
		settings.https = value

	}
}

// FolderCreate - создаёт всю необходимую структуру папок на сервере
func FolderCreate() {

	if !СheckExists("public") {
		WriteErrToConsole(os.Mkdir("public", 0700))
	}

	if !СheckExists("public/frontend") {
		WriteErrToConsole(os.Mkdir("public/frontend", 0700))
	}

	if !СheckExists("public/uploads") {
		WriteErrToConsole(os.Mkdir("public/uploads", 0700))
	}

	if !СheckExists("logs") {
		WriteErrToConsole(os.Mkdir("logs", 0700))
	}
}

// СheckExists - проверяем что файл существует
func СheckExists(filename string) bool {

	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}

// WriteErrToConsole - до создания лога пишем ошибку в консоль
func WriteErrToConsole(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func WriteErrToLog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
