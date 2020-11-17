// Package shoppinglist - содержит функции обработчики запросов для списка покупок
package shoppinglist

import (
	"encoding/json"
	"errors"
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
	ErrNotAllowedMethod = errors.New("Запрошен недопустимый метод для списка покупок")
	ErrNoKeyInParams    = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams = errors.New("API ключ не зарегистрирован")
	ErrNotAuthorized    = errors.New("Пройдите авторизацию")
	ErrForbidden        = errors.New("Доступ запрещён")
)

// HandleShoppingList - обрабатывает GET, POST и DELETE запросы для списка покупок
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
// 	идентичным по структуре IngredientDB
//
// DELETE
//
// 	ожидается параметр key с API ключом
// 	ожидается заголовок IngName с названием продукта из списка покупок
func HandleShoppingList(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "HandleShoppingList") {

			// Обработка получения списка покупок с поддержкой постраничных порций

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")

			var resp databases.ShoppingListResponse

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
				return
			}
			defer dbc.Close()

			if len(PageStr) > 0 && len(LimitStr) > 0 {

				Page, err := strconv.Atoi(PageStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				resp, err = databases.PostgreSQLShoppingListSelect(Page, Limit, dbc)

			} else {
				resp, err = databases.PostgreSQLShoppingListSelect(0, 0, dbc)
			}

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.WriteObjectToJSON(w, resp)

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:
		// Обработка записи отдельного пункта списка покупок в базу данных

		if setup.ServerSettings.CheckRoleForChange(role, "HandleShoppingList") {
			var Ingredient databases.IngredientDB

			err = json.NewDecoder(req.Body).Decode(&Ingredient)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
				return
			}
			defer dbc.Close()

			err = databases.PostgreSQLShoppingListInsertUpdate(Ingredient, dbc)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.HandleSuccessMessage(w, "Запись сохранена")

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:
		// Обработка удаления отдельного пункта списка покупок из базы данных

		if setup.ServerSettings.CheckRoleForDelete(role, "HandleShoppingList") {

			IngName := req.Header.Get("IngName")

			if IngName != "" {
				IngName, err := url.QueryUnescape(IngName)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				dbc := setup.ServerSettings.SQL.Connect(w, role)
				if dbc == nil {
					shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
					return
				}
				defer dbc.Close()

				err = databases.PostgreSQLShoppingListDelete(IngName, dbc)

				if err != nil {
					if errors.Is(err, databases.ErrShoppingListNotFound) {
						shared.HandleOtherError(w, "Не найдено ни одной записи в списке покупок с указанным названием", err, http.StatusBadRequest)
						return
					}
				}

				if shared.HandleInternalServerError(w, err) {
					return
				}

				shared.HandleSuccessMessage(w, "Запись удалена")

			} else {

				dbc := setup.ServerSettings.SQL.Connect(w, role)
				if dbc == nil {
					shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
					return
				}

				defer dbc.Close()

				err = databases.PostgreSQLShoppingListDeleteAll(dbc)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				shared.HandleSuccessMessage(w, "Все записи удалены")
			}

		} else {
			shared.HandleOtherError(w, ErrForbidden.Error(), ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
