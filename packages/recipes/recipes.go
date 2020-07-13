// Package recipes - содержит функции обработчики для рецептов
package recipes

import (
	"encoding/json"
	"errors"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"net/http"
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

		databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
			setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
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

		databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
			setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
		defer databases.PostgreSQLCloseConn()

		err = databases.PostgreSQLRecipesInsertUpdate(Recipe)

		if shared.HandleInternalServerError(w, err) {
			return
		}

		w.WriteHeader(http.StatusOK)

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

			databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
				setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
			defer databases.PostgreSQLCloseConn()

			err = databases.PostgreSQLRecipesDelete(RecipeID)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			w.WriteHeader(http.StatusOK)

		} else {
			shared.HandleOtherError(w, "Bad request", errors.New("Не заполнен обязательный заголовок RecipeID в запросе на удаление рецепта"), http.StatusBadRequest)
		}
	default:
		shared.HandleOtherError(w, "Method is not allowed", errors.New("Запрошен недопустимый метод для рецептов"), http.StatusMethodNotAllowed)
	}
}
