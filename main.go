// Package main - Сервер книги рецептов и списка покупок
package main

import (
	"io"
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"net/http"
	"os"
)

func main() {

	// Создаём папки если их не существует
	setup.FolderCreate()

	// Задаём куда писать логи сервера
	file, err := os.OpenFile("./logs/MainServer.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)

	log.SetOutput(mw)

	// Если первый аргумент -set - запускаем конфигуратор
	// вне зависимости от наличия файла settings.json
	runargs := os.Args
	var forcesetup bool
	if len(runargs) > 1 {
		first := runargs[1]
		if first == "-set" {
			forcesetup = true
		}
	}

	// Проверяем и запускаем мастер настройки если нужно
	// Иначе просто читаем данные из файла settings.json
	ServerSettings := setup.InitialSettings(forcesetup)

	// Устанавливаем пути, по которым будут происходить http запросы
	http.Handle("/", http.FileServer(http.Dir("./public/frontend")))

	// Запускаем либо http либо https сервер, в зависимости от наличия сертификата в папке с сервером

	if setup.СheckExists("cert.pem") && setup.СheckExists("key.pem") {
		//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
		log.Println("Started SSL web server")
		log.Fatalln(http.ListenAndServeTLS(":"+string(ServerSettings.HTTPS), "cert.pem", "key.pem", nil))
	} else {
		log.Println("Started plain web server")
		log.Fatalln(http.ListenAndServe(":"+string(ServerSettings.HTTP), nil))
	}

}
