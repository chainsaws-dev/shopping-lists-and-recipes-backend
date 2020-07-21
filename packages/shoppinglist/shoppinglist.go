// Package shoppinglist - содержит функции обработчики запросов для списка покупок
package shoppinglist

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
	ErrNotAllowedMethod = errors.New("Запрошен недопустимый метод для списка покупок")
	ErrNoKeyInParams    = errors.New("API ключ не указан в параметрах")
	ErrWrongKeyInParams = errors.New("API ключ не зарегистрирован")
	ErrNotAuthorized    = errors.New("Пройдите авторизацию")
)

// HandleShoppingList - обрабатывает POST, GET и DELETE запросы для изменения списка покупок
func HandleShoppingList(w http.ResponseWriter, req *http.Request) {
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
				// Обработка получения списка покупок с поддержкой постраничных порций
				w.Header().Set("Content-Type", "application/json")

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

				js, err := json.Marshal(resp)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				_, err = w.Write(js)

				if shared.HandleInternalServerError(w, err) {
					return
				}

			case req.Method == http.MethodPost:
				// Обработка записи отдельного пункта списка покупок в базу данных
				w.Header().Set("Content-Type", "application/json")

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

				w.WriteHeader(http.StatusOK)
				resulttext := fmt.Sprintf(`{"Error":{"Code":%v, "Message":"%v"}}`, http.StatusOK, "Запись сохранена")
				fmt.Fprintln(w, resulttext)

			case req.Method == http.MethodDelete:
				// Обработка удаления отдельного пункта списка покупок из базы данных
				w.Header().Set("Content-Type", "application/json")

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
						if err.Error() == "Не найдено ни одной записи в списке покупок с указанным названием" {
							shared.HandleOtherError(w, "Shopping list item not found and cannot be deleted", err, http.StatusBadRequest)
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

					err := setup.ServerSettings.SQL.Connect(role)

					if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
						return
					}
					defer setup.ServerSettings.SQL.Disconnect()

					err = databases.PostgreSQLShoppingListDeleteAll()

					if shared.HandleInternalServerError(w, err) {
						return
					}

					w.WriteHeader(http.StatusOK)
					resulttext := fmt.Sprintf(`{"Error":{"Code":%v, "Message":"%v"}}`, http.StatusOK, "Все записи удалены")
					fmt.Fprintln(w, resulttext)
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
