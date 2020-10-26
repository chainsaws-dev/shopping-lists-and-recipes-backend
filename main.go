// Package main - Сервер книги рецептов и списка покупок
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"shopping-lists-and-recipes/packages/files"
	"shopping-lists-and-recipes/packages/recipes"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/shoppinglist"
	"shopping-lists-and-recipes/packages/signinupout"
)

// main - главная точка входа в программу
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
	var cleantokens bool
	var createdb bool

	if len(runargs) > 1 {
		for _, runarg := range runargs {
			if runarg == "-set" {
				forcesetup = true
			}

			if runarg == "-clean" {
				cleantokens = true
			}

			if runarg == "-makedb" {
				createdb = true
			}
		}
	}

	// Проверяем и запускаем мастер настройки если нужно
	// Иначе просто читаем данные из файла settings.json
	setup.InitialSettings(forcesetup, createdb)

	// Устанавливаем пути, по которым будут происходить http запросы

	// Раздаём файл сервером фронтенд и загрузки
	http.Handle("/", http.FileServer(http.Dir("./public/frontend")))
	http.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("./public/uploads"))))

	// Перенаправляем все запросы по разделам на индекс
	http.HandleFunc("/recipes/", RedirectToIndex)
	http.HandleFunc("/shopping-list/", RedirectToIndex)
	http.HandleFunc("/admin/", RedirectToIndex)
	http.HandleFunc("/auth/", RedirectToIndex)
	http.HandleFunc("/confirm-email/", RedirectToIndex)
	http.HandleFunc("/reset-password/", RedirectToIndex)

	// REST API

	// Рецепты
	http.HandleFunc("/api/Recipes", recipes.HandleRecipes)
	http.HandleFunc("/api/Recipes/Search", recipes.HandleRecipesSearch)

	// Файлы
	http.HandleFunc("/api/Files", files.HandleFiles)

	// Список покупок
	http.HandleFunc("/api/ShoppingList", shoppinglist.HandleShoppingList)

	// Авторизация и регистрация
	http.HandleFunc("/api/Accounts/SignUp", signinupout.SignUp)
	http.HandleFunc("/api/Accounts/SignIn", signinupout.SignIn)

	// Админка
	http.HandleFunc("/api/Users", signinupout.HandleUsers)
	http.HandleFunc("/api/Users/Current", signinupout.GetCurrentUser)
	http.HandleFunc("/api/Sessions", signinupout.HandleSessions)

	// Сервис
	http.HandleFunc("/api/ConfirmEmail", signinupout.ConfirmEmail)
	http.HandleFunc("/api/ConfirmEmail/Send", signinupout.ResendEmail)
	http.HandleFunc("/api/PasswordReset", signinupout.ResetPassword)
	http.HandleFunc("/api/PasswordReset/Send", signinupout.RequestResetEmail)

	if cleantokens {
		go signinupout.RegularConfirmTokensCleanup()
	}

	// Запускаем либо http либо https сервер, в зависимости от наличия сертификата в папке с сервером
	if setup.СheckExists("cert.pem") && setup.СheckExists("key.pem") {
		//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
		shared.CurrentPrefix = "https://"
		log.Println("Запущен SSL веб сервер")
		log.Fatalln(http.ListenAndServeTLS(fmt.Sprintf(":%v", setup.ServerSettings.HTTPS), "cert.pem", "key.pem", nil))
	} else {
		shared.CurrentPrefix = "http://"
		log.Println("Запущен веб сервер без шифрования")
		log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", setup.ServerSettings.HTTP), nil))
	}

}

// RedirectToIndex - перенаправляет на файл index.html
func RedirectToIndex(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./public/frontend/index.html")
}
