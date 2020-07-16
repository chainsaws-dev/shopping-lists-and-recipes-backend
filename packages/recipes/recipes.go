// Package recipes - содержит функции обработчики для рецептов
package recipes

import (
	"encoding/json"
	"errors"
	"fmt"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"net/http"
	"net/url"
	"strconv"
)

// HandleRecipes - обрабатывает POST, GET и DELETE запросы для изменения рецептов
func HandleRecipes(w http.ResponseWriter, req *http.Request) {

	switch {
	case req.Method == http.MethodGet:
		// Обработка получения списка рецептов с поддержкой постраничных порций
		w.Header().Set("Content-Type", "application/json")

		PageStr := req.Header.Get("Page")
		LimitStr := req.Header.Get("Limit")

		var recipesresp databases.RecipesResponse
		var err error

		// TODO
		// Должна назначаться аутентификацией
		ActiveRole := setup.ServerSettings.SQL.Roles[1]

		err = databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
			setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
		if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
			return
		}
		defer databases.PostgreSQLCloseConn()

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
		// Должна назначаться аутентификацией
		ActiveRole := setup.ServerSettings.SQL.Roles[1]

		err = databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
			setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
		if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
			return
		}
		defer databases.PostgreSQLCloseConn()

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
			// Должна назначаться аутентификацией
			ActiveRole := setup.ServerSettings.SQL.Roles[1]

			err = databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
				setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
			if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
				return
			}
			defer databases.PostgreSQLCloseConn()

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
			shared.HandleOtherError(w, "Bad request", errors.New("Не заполнен обязательный заголовок RecipeID в запросе на удаление рецепта"), http.StatusBadRequest)
		}
	default:
		shared.HandleOtherError(w, "Method is not allowed", errors.New("Запрошен недопустимый метод для рецептов"), http.StatusMethodNotAllowed)
	}
}

// HandleRecipesSearch - обрабатывает GET запросы для поиска рецептов
func HandleRecipesSearch(w http.ResponseWriter, req *http.Request) {
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
		// Должна назначаться аутентификацией
		ActiveRole := setup.ServerSettings.SQL.Roles[1]

		err = databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
			setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
		if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
			return
		}
		defer databases.PostgreSQLCloseConn()

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
			errtext := "Не заполнены обязательные параметры поискового запроса: Page, Limit, Search"
			shared.HandleOtherError(w, errtext, errors.New(errtext), http.StatusBadRequest)
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
		shared.HandleOtherError(w, "Method is not allowed", errors.New("Запрошен недопустимый метод для поиска рецептов"), http.StatusMethodNotAllowed)
	}
}
