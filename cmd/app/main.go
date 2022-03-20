// Package main - Сервер книги рецептов и списка покупок
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"shopping-lists-and-recipes/internal/recipes"
	"shopping-lists-and-recipes/internal/setup"
	"shopping-lists-and-recipes/internal/shoppinglist"
	"shopping-lists-and-recipes/packages/files"
	"shopping-lists-and-recipes/packages/gzipwrap"
	"shopping-lists-and-recipes/packages/secondfactor"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
	"strings"
)

// Список типовых ошибок
var (
	ErrWrongArgumentFormat = errors.New("неверный формат данных для логина: ожидатся -admincred:example@example.ru@@password")
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

	initpar := GetRunArgs()

	// Проверяем и запускаем мастер настройки если нужно
	// Иначе просто читаем данные из файла settings.json
	setup.InitialSettings(initpar)

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
	http.HandleFunc("/profile/", RedirectToIndex)
	http.HandleFunc("/totp/", RedirectToIndex)

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

	// Второй фактор
	http.HandleFunc("/api/TOTP/Check", secondfactor.CheckSecondFactor)
	http.HandleFunc("/api/TOTP/Settings", secondfactor.SecondFactor)
	http.HandleFunc("/api/TOTP/Qr.png", secondfactor.GetQRCode)

	// Админка
	http.HandleFunc("/api/Users", signinupout.HandleUsers)
	http.HandleFunc("/api/Users/Current", signinupout.CurrentUser)
	http.HandleFunc("/api/Sessions", signinupout.HandleSessions)

	// Сервис
	http.HandleFunc("/api/ConfirmEmail", signinupout.ConfirmEmail)
	http.HandleFunc("/api/ConfirmEmail/Send", signinupout.ResendEmail)
	http.HandleFunc("/api/PasswordReset", signinupout.ResetPassword)
	http.HandleFunc("/api/PasswordReset/Send", signinupout.RequestResetEmail)

	if initpar.CleanTokens {
		go signinupout.RegularConfirmTokensCleanup()
	}

	// Создаём пул соединений
	if setup.ServerSettings.SQL.Connected == false {
		setup.ServerSettings.SQL.Connect(false)
	}
	defer setup.ServerSettings.SQL.Disconnect()

	// Запускаем либо http либо https сервер, в зависимости от наличия сертификата в папке с сервером
	if setup.СheckExists("cert.pem") && setup.СheckExists("key.pem") {
		//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
		shared.CurrentPrefix = "https://"
		log.Println("Запущен SSL веб сервер")
		log.Fatalln(http.ListenAndServeTLS(fmt.Sprintf(":%v", setup.ServerSettings.HTTPS), "cert.pem", "key.pem", gzipwrap.MakeGzipHandler(http.DefaultServeMux)))
	} else {
		shared.CurrentPrefix = "http://"
		log.Println("Запущен веб сервер без шифрования")
		log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", setup.ServerSettings.HTTP), gzipwrap.MakeGzipHandler(http.DefaultServeMux)))
	}

}

// RedirectToIndex - перенаправляет на файл index.html
func RedirectToIndex(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./public/frontend/index.html")
}

// GetRunArgs - читает параметры запуска программы
func GetRunArgs() setup.InitParams {
	// Если первый аргумент -set - запускаем конфигуратор
	// вне зависимости от наличия файла settings.json
	runargs := os.Args

	// Инкапсулируем параметры установки в объекте
	var initpar setup.InitParams
	initpar.CreateAdmin = true

	if len(runargs) > 1 {
		for _, runarg := range runargs {
			// Принудительно запускает мастер настройки
			// если файл settings.json существует и перезаписывает его
			if runarg == "-set" {
				initpar.ForceSetup = true
			}

			// Запускает процесс автоочистки токенов и сессий
			// с интервалом пять минут
			if runarg == "-clean" {
				initpar.CleanTokens = true
			}

			// Запускает удаление базы без запуска мастера настройки
			// при существующем файле settings.json
			if runarg == "-dropdb" {
				initpar.DropDb = true
			}

			// Запускает создание базы без запуска мастера настройки
			// при существующем файле settings.json
			if runarg == "-makedb" {
				initpar.CreateDb = true
			}

			// Работает при -makedb (только для отладки)
			// Отключает создание администратора
			// после завершения создания базы и ролей
			if runarg == "-noadmin" {
				initpar.CreateAdmin = false
			}

			// Для пакетного режима
			// Работает при -makedb
			// Позволяет передать логин и пароль начального администратора через параметры командной строки
			if strings.HasPrefix(runarg, "-admincred:") {
				SetAdminCredentials(runarg, &initpar)
			}

			// Для пакетного режима
			// Работает при -makedb
			// Позволяет передать адрес вебсайта для формирования ссылок в почте через параметры командной строки
			if strings.HasPrefix(runarg, "-url:") {
				initpar.WebsiteURL = strings.ReplaceAll(runarg, "-url:", "")
			}

		}
	}

	// Для работы через докер
	admincred := os.Getenv("ADMIN_CRED")
	if len(admincred) > 0 {
		SetAdminCredentials(admincred, &initpar)
	}

	WebUrl := os.Getenv("URL")
	if len(WebUrl) > 0 {
		initpar.WebsiteURL = WebUrl
	}

	return initpar
}

func SetAdminCredentials(runarg string, initpar *setup.InitParams) {
	basestring := strings.ReplaceAll(runarg, "-admincred:", "")
	lp := strings.Split(basestring, "@@")

	if len(lp) == 2 {
		initpar.AdminLogin = lp[0]
		initpar.AdminPass = lp[1]
	} else {
		log.Println(lp)
		shared.WriteErrToLog(ErrWrongArgumentFormat)
	}
}
