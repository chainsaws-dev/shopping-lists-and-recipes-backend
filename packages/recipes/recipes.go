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

		PageStr := req.Header.Get("Page")

		if PageStr == "" {
			shared.HandleInternalServerError(w, errors.New("При GET запросе рецептов не указана страница в заголовке Page"))
			return
		}

		Page, err := strconv.Atoi(PageStr)

		shared.HandleInternalServerError(w, err)

		// TODO
		// Должна назначаться аутентификацией
		ActiveRole := setup.ServerSettings.SQL.Roles[1]
		databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(ActiveRole.Login, ActiveRole.Pass,
			setup.ServerSettings.SQL.Addr, setup.ServerSettings.SQL.DbName, false))
		defer databases.PostgreSQLCloseConn()

		recipesresp := databases.PostgreSQLRecipesSelect(Page)

		js, err := json.Marshal(recipesresp)

		if shared.HandleInternalServerError(w, err) {
			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(js)

		if shared.HandleInternalServerError(w, err) {
			return
		}

	case req.Method == http.MethodPost:
		//TODO
		http.Error(w, "Method is not implemented", http.StatusNotImplemented)
	case req.Method == http.MethodDelete:
		//TODO
		http.Error(w, "Method is not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}
}
