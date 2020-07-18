// Package main - Сервер книги рецептов и списка покупок
package main

import (
	"fmt"
	"io"
	"log"
	"myprojects/Shopping-lists-and-recipes/packages/files"
	"myprojects/Shopping-lists-and-recipes/packages/recipes"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shoppinglist"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
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
	if len(runargs) > 1 {
		first := runargs[1]
		if first == "-set" {
			forcesetup = true
		}
	}

	// Проверяем и запускаем мастер настройки если нужно
	// Иначе просто читаем данные из файла settings.json
	setup.InitialSettings(forcesetup)

	// Устанавливаем пути, по которым будут происходить http запросы

	// Раздаём файл сервером фронтенд и загрузки
	http.Handle("/", http.FileServer(http.Dir("./public/frontend")))
	http.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("./public/uploads"))))

	// Перенаправляем все запросы по разделам на индекс
	http.HandleFunc("/recipes/", RedirectToIndex)
	http.HandleFunc("/shopping-list/", RedirectToIndex)
	http.HandleFunc("/admin/", RedirectToIndex)

	// REST API
	http.HandleFunc("/api/Recipes", recipes.HandleRecipes)
	http.HandleFunc("/api/Recipes/Search", recipes.HandleRecipesSearch)
	http.HandleFunc("/api/SaveRecipePhoto", files.UploadFile)
	http.HandleFunc("/api/ShoppingList", shoppinglist.HandleShoppingList)

	// Запускаем либо http либо https сервер, в зависимости от наличия сертификата в папке с сервером

	if setup.СheckExists("cert.pem") && setup.СheckExists("key.pem") {
		//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
		log.Println("Запущен SSL веб сервер")
		log.Fatalln(http.ListenAndServeTLS(fmt.Sprintf(":%v", setup.ServerSettings.HTTPS), "cert.pem", "key.pem", nil))
	} else {
		log.Println("Запущен веб сервер без шифрования")
		log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", setup.ServerSettings.HTTP), nil))
	}

}

// RedirectToIndex - перенаправляет на файл index.html
func RedirectToIndex(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./public/frontend/index.html")
}
