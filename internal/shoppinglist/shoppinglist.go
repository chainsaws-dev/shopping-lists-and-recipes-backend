// Package shoppinglist - содержит функции обработчики запросов для списка покупок
package shoppinglist

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"shopping-lists-and-recipes/internal/databases"
	"shopping-lists-and-recipes/internal/setup"
	"shopping-lists-and-recipes/packages/shared"
	"shopping-lists-and-recipes/packages/signinupout"
	"strconv"
)

var (
	ErrNoEntryFoundByName = errors.New("no shopping list entry found with specified name")
)

var (
	MesEntrySaved        = "entry saved"
	MesEntryDeleted      = "entry deleted"
	MesAllEntriesDeleted = "all entries deleted"
)

// HandleShoppingList - обрабатывает GET, POST и DELETE запросы для списка покупок
//
// Аутентификация
//
//  Куки
//  Session - шифрованная сессия
//	Email - шифрованный электронный адрес пользователя
//
//  или
//
//	Заголовки:
//  Auth - Токен доступа
//  Lang - Язык (ru или en)
//
//	и
//
//	ApiKey - Постоянный ключ доступа к API *
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

			if len(PageStr) > 0 && len(LimitStr) > 0 {

				Page, err := strconv.Atoi(PageStr)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				resp, err = databases.PostgreSQLShoppingListSelect(Page, Limit, setup.ServerSettings.SQL.ConnPool)

			} else {
				resp, err = databases.PostgreSQLShoppingListSelect(0, 0, setup.ServerSettings.SQL.ConnPool)
			}

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, resp)

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:
		// Обработка записи отдельного пункта списка покупок в базу данных

		if setup.ServerSettings.CheckRoleForChange(role, "HandleShoppingList") {
			var Ingredient databases.IngredientDB

			err = json.NewDecoder(req.Body).Decode(&Ingredient)

			if shared.HandleBadRequestError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			err = databases.PostgreSQLShoppingListInsertUpdate(Ingredient, setup.ServerSettings.SQL.ConnPool)

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesEntrySaved)

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:
		// Обработка удаления отдельного пункта списка покупок из базы данных

		if setup.ServerSettings.CheckRoleForDelete(role, "HandleShoppingList") {

			IngName := req.Header.Get("IngName")

			if IngName != "" {
				IngName, err := url.QueryUnescape(IngName)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				err = databases.PostgreSQLShoppingListDelete(IngName, setup.ServerSettings.SQL.ConnPool)

				if err != nil {
					if errors.Is(err, databases.ErrShoppingListNotFound) {
						shared.HandleOtherError(setup.ServerSettings.Lang, w, req, ErrNoEntryFoundByName.Error(), err, http.StatusBadRequest)
						return
					}
				}

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesEntryDeleted)

			} else {

				err = databases.PostgreSQLShoppingListDeleteAll(setup.ServerSettings.SQL.ConnPool)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, MesAllEntriesDeleted)
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
