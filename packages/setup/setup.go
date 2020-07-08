// Package setup - Выполняет начальную настройку и создание структуры папок при первом запуске сервера для рецептов и списка покупок
package setup

import (
	"fmt"
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/settings"
	"os"
	"strconv"
)

// InitialSettings - интерактивно спрашивает у пользователя параметры настроек
func InitialSettings(forcesetup bool) {

	if !СheckExists("settings.json") || forcesetup {

		if !forcesetup {
			log.Println("Не найден файл settings.json. Запущена процедура начальной настройки.")
		} else {
			log.Println("Принудительно запущена процедура начальной настройки.")
		}

		var ServerSettings settings.WServerSettings

		fmt.Println("")
		fmt.Println("*******************************************************************")
		fmt.Println("* Добро пожаловать в мастер настройки сервера рецептов и покупок! *")
		fmt.Println("* Для настройки сервера необходимо создать конфигурационный файл. *")
		fmt.Println("* Он будет создан по результатам ваших ответов на вопросы.        *")
		fmt.Println("*                           ВНИМАНИЕ!                             *")
		fmt.Println("*  Некоторые значения настроек потребуют запуска сервера от sudo  *")
		fmt.Println("*******************************************************************")

		AskInt("Укажите порт для http соединений (например 80): ", &ServerSettings.HTTP)

		AskInt("Укажите порт для https соединений (например 443): ", &ServerSettings.HTTPS)

		AskString("Укажите адрес сервера баз данных PostgreSQL: ", &ServerSettings.SQL.Addr)

	}
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

// СheckExists - проверяем что файл существует
func СheckExists(filename string) bool {

	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}

// WriteErrToLog - пишем ошибку в лог (лог файл должен быть уже задан)
func WriteErrToLog(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
