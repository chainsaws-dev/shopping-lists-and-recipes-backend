// Package recipes - содержит функции обработчики для рецептов
package recipes

import (
	"encoding/json"
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

		w.Header().Set("Content-Type", "application/json")

		PageStr := req.Header.Get("Page")
		LimitStr := req.Header.Get("Limit")

		var recipesresp databases.RecipesResponse

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

			recipesresp = databases.PostgreSQLRecipesSelect(Page, Limit)

		} else {
			recipesresp = databases.PostgreSQLRecipesSelect(0, 0)
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
		//TODO
		shared.HandleOtherError(w, "POST method is not implemented", http.StatusNotImplemented)
	case req.Method == http.MethodDelete:
		//TODO
		shared.HandleOtherError(w, "DELETE method is not implemented", http.StatusNotImplemented)
	default:
		shared.HandleOtherError(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}
}
