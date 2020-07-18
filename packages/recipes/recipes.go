// Package recipes - содержит функции обработчики запросов для рецептов
package recipes

import (
	"encoding/json"
	"errors"
	"fmt"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"myprojects/Shopping-lists-and-recipes/packages/signinupout"
	"net/http"
	"net/url"
	"strconv"
)

// Список типовых ошибок
var (
	ErrNotAllowedMethod       = errors.New("Запрошен недопустимый метод для рецептов")
	ErrRecipeIDNotFilled      = errors.New("Не заполнен обязательный заголовок RecipeID в запросе на удаление рецепта")
	ErrHeadersSearchNotFilled = errors.New("Не заполнены обязательные параметры поискового запроса: Page, Limit, Search")
	ErrNoKeyInParams          = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams       = errors.New("API ключ не зарегистрирован")
	ErrNotAuthorized          = errors.New("Пройдите авторизацию")
)

// HandleRecipes - обрабатывает POST, GET и DELETE запросы для изменения рецептов
func HandleRecipes(w http.ResponseWriter, req *http.Request) {

	// Проверяем API ключ
	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		shared.HandleOtherError(w, ErrNoKeyInParams.Error(), ErrNoKeyInParams, http.StatusBadRequest)
		return
	}

	key := keys[0]

	_, found := shared.FindInStringSlice(setup.APIkeys, key)

	if found {
		// Проверка токена и получение роли
		Auth := req.Header.Get("Auth")

		issued, role := signinupout.CheckTokenIssued(Auth)

		if issued {

			switch {
			case req.Method == http.MethodGet:
				// Обработка получения списка рецептов с поддержкой постраничных порций
				w.Header().Set("Content-Type", "application/json")

				PageStr := req.Header.Get("Page")
				LimitStr := req.Header.Get("Limit")

				var recipesresp databases.RecipesResponse
				var err error

				// TODO
				// Роль для поиска должна назначаться аутентификацией
				err = setup.ServerSettings.SQL.Connect(role)

				if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
					return
				}
				defer setup.ServerSettings.SQL.Disconnect()

				if PageStr != "" && LimitStr != "" {

					Page, err := strconv.Atoi(PageStr)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					Limit, err := strconv.Atoi(LimitStr)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					recipesresp, err = databases.PostgreSQLRecipesSelect(Page, Limit)

				} else {
					recipesresp, err = databases.PostgreSQLRecipesSelect(0, 0)
				}

				if shared.HandleInternalServerError(w, err) {
					return
				}

				js, err := json.Marshal(recipesresp)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				_, err = w.Write(js)

				if shared.HandleInternalServerError(w, err) {
					return
				}

			case req.Method == http.MethodPost:
				// Обработка записи отдельного рецепта в базу данных
				w.Header().Set("Content-Type", "application/json")

				var Recipe databases.RecipeDB

				err := json.NewDecoder(req.Body).Decode(&Recipe)

				if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
					return
				}

				// TODO
				// Роль для поиска должна назначаться аутентификацией
				err = setup.ServerSettings.SQL.Connect("admin_role_CRUD")

				if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
					return
				}
				defer setup.ServerSettings.SQL.Disconnect()

				recipesresp, err := databases.PostgreSQLRecipesInsertUpdate(Recipe)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				js, err := json.Marshal(recipesresp)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				_, err = w.Write(js)

				if shared.HandleInternalServerError(w, err) {
					return
				}

			case req.Method == http.MethodDelete:
				// Обработка удаления отдельного рецепта из базы данных и его обложки с фаловой системы
				w.Header().Set("Content-Type", "application/json")

				RecipeIDToDelStr := req.Header.Get("RecipeID")

				if RecipeIDToDelStr != "" {

					RecipeID, err := strconv.Atoi(RecipeIDToDelStr)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					// TODO
					// Роль для поиска должна назначаться аутентификацией
					err = setup.ServerSettings.SQL.Connect("admin_role_CRUD")

					if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
						return
					}
					defer setup.ServerSettings.SQL.Disconnect()

					err = databases.PostgreSQLRecipesDelete(RecipeID)

					if err != nil {
						if err.Error() == "В таблице рецептов не найден указанный id" {
							shared.HandleOtherError(w, "Recipe not found and cannot be deleted", err, http.StatusBadRequest)
							return
						}
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}

					w.WriteHeader(http.StatusOK)
					resulttext := fmt.Sprintf(`{"Error":{"Code":%v, "Message":"%v"}}`, http.StatusOK, "Запись удалена")
					fmt.Fprintln(w, resulttext)

				} else {
					shared.HandleOtherError(w, "Bad request", ErrRecipeIDNotFilled, http.StatusBadRequest)
				}
			default:
				shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
			}
		} else {
			shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}

// HandleRecipesSearch - обрабатывает GET запросы для поиска рецептов
func HandleRecipesSearch(w http.ResponseWriter, req *http.Request) {
	// Проверяем API ключ
	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		shared.HandleOtherError(w, ErrNoKeyInParams.Error(), ErrNoKeyInParams, http.StatusBadRequest)
		return
	}

	key := keys[0]

	_, found := shared.FindInStringSlice(setup.APIkeys, key)

	if found {
		// Проверка токена и получение роли
		Auth := req.Header.Get("Auth")
		issued, role := signinupout.CheckTokenIssued(Auth)

		if issued {
			switch {
			case req.Method == http.MethodGet:
				// Обработка получения списка рецептов с поддержкой постраничных порций
				w.Header().Set("Content-Type", "application/json")

				PageStr := req.Header.Get("Page")
				LimitStr := req.Header.Get("Limit")
				SearchStr := req.Header.Get("Search")

				var recipesresp databases.RecipesResponse
				var err error

				// TODO
				// Роль для поиска должна назначаться аутентификацией
				err = setup.ServerSettings.SQL.Connect(role)

				if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
					return
				}
				defer setup.ServerSettings.SQL.Disconnect()

				if PageStr != "" && LimitStr != "" && SearchStr != "" {

					Page, err := strconv.Atoi(PageStr)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					Limit, err := strconv.Atoi(LimitStr)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					SearchStr, err := url.QueryUnescape(SearchStr)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					recipesresp, err = databases.PostgreSQLRecipesSelectSearch(Page, Limit, SearchStr)

				} else {
					shared.HandleOtherError(w, ErrHeadersSearchNotFilled.Error(), ErrHeadersSearchNotFilled, http.StatusBadRequest)
					return
				}

				if shared.HandleInternalServerError(w, err) {
					return
				}

				js, err := json.Marshal(recipesresp)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				_, err = w.Write(js)

				if shared.HandleInternalServerError(w, err) {
					return
				}

			default:
				shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
			}
		} else {
			shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}
}
