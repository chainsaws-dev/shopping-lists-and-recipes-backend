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
	// Проверяем API ключ
	found, err := signinupout.CheckAPIKey(w, req)

	if err != nil {
		if shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest) {
			return
		}
	}

	if found {
		// Проверка токена и получение роли
		issued, role := signinupout.TwoWayAuthentication(w, req)

		if issued {
			switch {
			case req.Method == http.MethodGet:

				if setup.ServerSettings.CheckRoleForRead(role, "HandleShoppingList") {

					// Обработка получения списка покупок с поддержкой постраничных порций

					PageStr := req.Header.Get("Page")
					LimitStr := req.Header.Get("Limit")

					var resp databases.ShoppingListResponse
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

						resp, err = databases.PostgreSQLShoppingListSelect(Page, Limit)

					} else {
						resp, err = databases.PostgreSQLShoppingListSelect(0, 0)
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

					err := json.NewDecoder(req.Body).Decode(&Ingredient)

					if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
						return
					}

					err = setup.ServerSettings.SQL.Connect(role)

					if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
						return
					}
					defer setup.ServerSettings.SQL.Disconnect()

					err = databases.PostgreSQLShoppingListInsertUpdate(Ingredient)

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

						err = setup.ServerSettings.SQL.Connect(role)

						if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
							return
						}
						defer setup.ServerSettings.SQL.Disconnect()

						err = databases.PostgreSQLShoppingListDelete(IngName)

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

						err := setup.ServerSettings.SQL.Connect(role)

						if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
							return
						}
						defer setup.ServerSettings.SQL.Disconnect()

						err = databases.PostgreSQLShoppingListDeleteAll()

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
		} else {
			shared.HandleOtherError(w, ErrNotAuthorized.Error(), ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", ErrWrongKeyInParams, http.StatusBadRequest)
	}

}
