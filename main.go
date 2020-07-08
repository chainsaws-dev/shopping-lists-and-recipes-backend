package main

import (
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
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

	log.SetOutput(file)

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

	// Проверяем и запускаем установщик если нужно
	setup.InitialSettings(forcesetup)

	// Устанавливаем пути, по которым будут происходить http запросы
	/*
		http.Handle("/", http.FileServer(http.Dir("./public/frontend")))

		// Запускаем либо http либо https сервер, в зависимости от наличия сертификата в папке с сервером

		if setup.СheckExists("cert.pem") && setup.СheckExists("key.pem") {
			//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
			fmt.Println("SSL web server up")
			log.Println("Started SSL web server")
			log.Fatalln(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
		} else {
			fmt.Println("Plain web server up")
			log.Println("Started plain web server")
			log.Fatalln(http.ListenAndServe(":8080", nil))
		}
	*/
}
