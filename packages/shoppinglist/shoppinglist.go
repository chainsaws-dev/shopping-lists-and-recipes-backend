package shoppinglist

import (
	"encoding/json"
	"errors"
	"myprojects/Shopping-lists-and-recipes/packages/databases"
	"myprojects/Shopping-lists-and-recipes/packages/setup"
	"myprojects/Shopping-lists-and-recipes/packages/shared"
	"net/http"
	"strconv"
)

// HandleShoppingList - обрабатывает POST, GET и DELETE запросы для изменения списка покупок
func HandleShoppingList(w http.ResponseWriter, req *http.Request) {

	switch {
	case req.Method == http.MethodGet:
		// Обработка получения списка покупок с поддержкой постраничных порций
		w.Header().Set("Content-Type", "application/json")

		PageStr := req.Header.Get("Page")
		LimitStr := req.Header.Get("Limit")

		var resp databases.ShoppingListResponse
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

			resp, err = databases.PostgreSQLShoppingListSelect(Page, Limit)

		} else {
			resp, err = databases.PostgreSQLShoppingListSelect(0, 0)
		}

		if shared.HandleInternalServerError(w, err) {
			return
		}

		js, err := json.Marshal(resp)

		if shared.HandleInternalServerError(w, err) {
			return
		}

		_, err = w.Write(js)

		if shared.HandleInternalServerError(w, err) {
			return
		}

	case req.Method == http.MethodPost:
		shared.HandleOtherError(w, "Method is not implemented", errors.New("Запрошен не реализованный метод для списка покупок"), http.StatusNotImplemented)
	case req.Method == http.MethodDelete:
		shared.HandleOtherError(w, "Method is not implemented", errors.New("Запрошен не реализованный метод для списка покупок"), http.StatusNotImplemented)
	default:
		shared.HandleOtherError(w, "Method is not allowed", errors.New("Запрошен недопустимый метод для списка покупок"), http.StatusMethodNotAllowed)
	}

}