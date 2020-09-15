// Package recipes - содержит функции обработчики запросов для рецептов
package recipes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"shopping-lists-and-recipes/packages/databases"
	"shopping-lists-and-recipes/packages/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
	"strconv"
)

// Список типовых ошибок
var (
	ErrNotAllowedMethod       = errors.New("Запрошен недопустимый метод для рецептов")
	ErrRecipeIDNotFilled      = errors.New("Не заполнен обязательный заголовок RecipeID в запросе на удаление рецепта")
	ErrHeadersSearchNotFilled = errors.New("Не заполнены обязательные параметры поискового запроса: Page, Limit, Search")
	ErrHeadersFetchNotFilled  = errors.New("Не заполнены обязательные параметры запроса списка рецептов: Page, Limit")
	ErrNoKeyInParams          = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams       = errors.New("API ключ не зарегистрирован")
	ErrNotAuthorized          = errors.New("Пройдите авторизацию")
	ErrForbidden              = errors.New("Доступ запрещён")
)

// HandleRecipes - обрабатывает POST, GET и DELETE запросы для изменения рецептов
//
// GET
//
// 	ожидается параметр key с API ключом
// 	ожидается заголовок Page с номером страницы
// 	ожидается заголовок Limit с максимумом элементов на странице
//
// POST
//
// 	ожидается параметр key с API ключом
// 	тело запроса должно быть заполнено JSON объектом
// 	идентичным по структуре RecipeDB
//
// DELETE
//
// 	ожидается параметр key с API ключом
// 	ожидается заголовок RecipeID с номером рецепта на удаление
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
		issued, role := signinupout.CheckTokenIssued(*req)

		if issued {

			switch {
			case req.Method == http.MethodGet:
				// Обработка получения списка рецептов с поддержкой постраничных порций
				w.Header().Set("Content-Type", "application/json")

				PageStr := req.Header.Get("Page")
				LimitStr := req.Header.Get("Limit")

				var recipesresp databases.RecipesResponse
				var err error

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
					shared.HandleOtherError(w, ErrHeadersFetchNotFilled.Error(), ErrHeadersFetchNotFilled, http.StatusBadRequest)
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

			case req.Method == http.MethodPost:
				// Обработка записи отдельного рецепта в базу данных
				w.Header().Set("Content-Type", "application/json")

				if role == "admin_role_CRUD" {

					var Recipe databases.RecipeDB

					err := json.NewDecoder(req.Body).Decode(&Recipe)

					if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
						return
					}

					err = setup.ServerSettings.SQL.Connect(role)

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

				} else {
					shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
				}

			case req.Method == http.MethodDelete:
				// Обработка удаления отдельного рецепта из базы данных и его обложки с фаловой системы
				w.Header().Set("Content-Type", "application/json")

				if role == "admin_role_CRUD" {

					RecipeIDToDelStr := req.Header.Get("RecipeID")

					if RecipeIDToDelStr != "" {

						RecipeID, err := strconv.Atoi(RecipeIDToDelStr)

						if shared.HandleInternalServerError(w, err) {
							return
						}

						err = setup.ServerSettings.SQL.Connect(role)

						if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
							return
						}
						defer setup.ServerSettings.SQL.Disconnect()

						err = databases.PostgreSQLRecipesDelete(RecipeID)

						if err != nil {
							if errors.Is(err, databases.ErrRecipeNotFound) {
								shared.HandleOtherError(w, "Рецепт не найден, невозможно удалить", err, http.StatusBadRequest)
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

				} else {
					shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
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
//
// GET
//
// 	ожидается параметр key с API ключом
// 	ожидается заголовок Page с номером страницы
// 	ожидается заголовок Limit с максимумом элементов на странице
// 	ожидается заголовок Search с поисковым запросом пропущенным через encodeURIComponent
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
		issued, role := signinupout.CheckTokenIssued(*req)

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
