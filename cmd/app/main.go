// Package main - Сервер книги рецептов и списка покупок
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shopping-lists-and-recipes/internal/recipes"
	"shopping-lists-and-recipes/internal/setup"
	"shopping-lists-and-recipes/internal/shoppinglist"
	"shopping-lists-and-recipes/packages/files"
	"shopping-lists-and-recipes/packages/gzipwrap"
	"shopping-lists-and-recipes/packages/multilangtranslator"
	"shopping-lists-and-recipes/packages/secondfactor"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// Список типовых ошибок
var (
	ErrWrongArgumentFormat = errors.New("incorrect data format for login: expected -admincred:example@example.ru@@password")
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

	if initpar.CleanTokens {
		go signinupout.RegularConfirmTokensCleanup()
	}

	// Создаём пул соединений c СУБД
	if setup.ServerSettings.SQL.Connected == false {
		setup.ServerSettings.SQL.Connect(false)
	}

	ServerSetup()
}

func ServerSetup() {

	InitFrontendHandlers()

	http.Handle("/api/v1/", InitAPIHandlersV1())

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Хорошая практика устанавливать таймауты для избежания атак Slowloris
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      gzipwrap.MakeGzipHandler(http.DefaultServeMux),
	}

	// Запускаем сервер в отдельной горутине, чтобы он не блокировал исполнение
	go func() {
		// Запускаем либо http либо https сервер, в зависимости от наличия сертификата в папке с сервером
		if setup.СheckExists("cert.pem") && setup.СheckExists("key.pem") {
			//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
			shared.CurrentPrefix = "https://"
			log.Println(multilangtranslator.TranslateString("encrypted webserver is up", setup.ServerSettings.Lang))
			srv.Addr = fmt.Sprintf(":%v", setup.ServerSettings.HTTPS)

			err := srv.ListenAndServeTLS("cert.pem", "key.pem")

			if !errors.Is(err, http.ErrServerClosed) {
				setup.ServerSettings.SQL.Disconnect()
				log.Fatalln(err)
			}

		} else {
			shared.CurrentPrefix = "http://"
			log.Println("")
			log.Println(multilangtranslator.TranslateString("unencrypted webserver is up", setup.ServerSettings.Lang))
			srv.Addr = fmt.Sprintf(":%v", setup.ServerSettings.HTTP)
			err := srv.ListenAndServe()

			if !errors.Is(err, http.ErrServerClosed) {
				setup.ServerSettings.SQL.Disconnect()
				log.Fatalln(err)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	// Завершение работы сервера вызывается только по сигналу SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) не будут обработаны.
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Блокируем выполнение до получения сигнала
	<-c

	// Создаём срок ожидания завершения
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := srv.Shutdown(ctx)

	if err != nil {
		log.Println(err)
	}

	setup.ServerSettings.SQL.Disconnect()

	log.Println(multilangtranslator.TranslateString("server is shutting down...", setup.ServerSettings.Lang))

	os.Exit(0)
}

// InitFrontendHandlers - инициализирует сервер по умолчанию для раздачи фронтенда и файлов
func InitFrontendHandlers() {

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
}

// InitHandlers - инициализирует guerilla mux handler
func InitAPIHandlersV1() http.Handler {

	r := mux.NewRouter()

	// Устанавливаем пути, по которым будут происходить http запросы

	// REST API

	// Рецепты
	r.HandleFunc("/api/v1/Recipes", recipes.HandleRecipes).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/api/v1/Recipes/Search", recipes.HandleRecipesSearch).Methods("GET")

	// Файлы
	r.HandleFunc("/api/v1/Files", files.HandleFiles).Methods("GET", "POST", "DELETE")

	// Список покупок
	r.HandleFunc("/api/v1/ShoppingList", shoppinglist.HandleShoppingList).Methods("GET", "POST", "DELETE")

	// Авторизация и регистрация
	r.HandleFunc("/api/v1/Accounts/SignUp", signinupout.SignUp).Methods("POST")
	r.HandleFunc("/api/v1/Accounts/SignIn", signinupout.SignIn).Methods("POST")

	// Второй фактор
	r.HandleFunc("/api/v1/TOTP/Check", secondfactor.CheckSecondFactor).Methods("POST")
	r.HandleFunc("/api/v1/TOTP/Settings", secondfactor.SecondFactor).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/api/v1/TOTP/Qr.png", secondfactor.GetQRCode).Methods("GET")

	// Админка
	r.HandleFunc("/api/v1/Users", signinupout.HandleUsers).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/api/v1/Users/Current", signinupout.CurrentUser).Methods("GET", "POST")
	r.HandleFunc("/api/v1/Sessions", signinupout.HandleSessions).Methods("GET", "DELETE")

	// Сервис
	r.HandleFunc("/api/v1/ConfirmEmail", signinupout.ConfirmEmail).Methods("POST")
	r.HandleFunc("/api/v1/ConfirmEmail/Send", signinupout.ResendEmail).Methods("POST")
	r.HandleFunc("/api/v1/PasswordReset", signinupout.ResetPassword).Methods("POST")
	r.HandleFunc("/api/v1/PasswordReset/Send", signinupout.RequestResetEmail).Methods("POST")

	return r
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
		shared.WriteErrToLog(ErrWrongArgumentFormat, setup.ServerSettings.Lang)
	}
}
