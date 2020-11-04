// Package recipes - содержит функции обработчики запросов для рецептов
package recipes

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
	ErrRecipeIDNotFilled      = errors.New("Не заполнен обязательный заголовок RecipeID в запросе на удаление рецепта")
	ErrHeadersSearchNotFilled = errors.New("Не заполнены обязательные параметры поискового запроса: Page, Limit, Search")
)

// HandleRecipes - обрабатывает POST, GET и DELETE запросы для изменения рецептов
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
//
//	и
//
//	ApiKey - Постоянный ключ доступа к API *
//
// GET
//
// 	ожидается заголовок Page с номером страницы
// 	ожидается заголовок Limit с максимумом элементов на странице
//
// POST
//
// 	тело запроса должно быть заполнено JSON объектом
// 	идентичным по структуре RecipeDB
//
// DELETE
//
// 	ожидается заголовок RecipeID с номером рецепта на удаление
func HandleRecipes(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "HandleRecipes") {
			// Обработка получения списка рецептов с поддержкой постраничных порций

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")

			var recipesresp databases.RecipesResponse

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
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

				recipesresp, err = databases.PostgreSQLRecipesSelect(Page, Limit, dbc)

				if shared.HandleInternalServerError(w, err) {
					return
				}

			} else {
				shared.HandleOtherError(w, shared.ErrHeadersFetchNotFilled.Error(), shared.ErrHeadersFetchNotFilled, http.StatusBadRequest)
				return
			}

			shared.WriteObjectToJSON(w, recipesresp)

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:
		// Обработка записи отдельного рецепта в базу данных

		if setup.ServerSettings.CheckRoleForChange(role, "HandleRecipes") {

			var Recipe databases.RecipeDB

			err = json.NewDecoder(req.Body).Decode(&Recipe)

			if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
				return
			}

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			recipesresp, err := databases.PostgreSQLRecipesInsertUpdate(Recipe, dbc)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.WriteObjectToJSON(w, recipesresp)

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "HandleRecipes") {
			// Обработка удаления отдельного рецепта из базы данных и его обложки с фаловой системы

			RecipeIDToDelStr := req.Header.Get("RecipeID")

			if RecipeIDToDelStr != "" {

				RecipeID, err := strconv.Atoi(RecipeIDToDelStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				dbc := setup.ServerSettings.SQL.Connect(w, role)
				if dbc == nil {
					return
				}
				defer dbc.Close()

				err = databases.PostgreSQLRecipesDelete(RecipeID, dbc)

				if err != nil {
					if errors.Is(err, databases.ErrRecipeNotFound) {
						shared.HandleOtherError(w, "Рецепт не найден, невозможно удалить", err, http.StatusBadRequest)
						return
					}
				}

				if shared.HandleInternalServerError(w, err) {
					return
				}

				shared.HandleSuccessMessage(w, "Запись удалена")

			} else {
				shared.HandleOtherError(w, "Bad request", ErrRecipeIDNotFilled, http.StatusBadRequest)
			}

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}

// HandleRecipesSearch - обрабатывает GET запросы для поиска рецептов
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
//
//	и
//
//	ApiKey - Постоянный ключ доступа к API *
//
// GET
//
// 	ожидается заголовок Page с номером страницы
// 	ожидается заголовок Limit с максимумом элементов на странице
// 	ожидается заголовок Search с поисковым запросом пропущенным через encodeURIComponent
func HandleRecipesSearch(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "HandleRecipesSearch") {
			// Обработка получения списка рецептов с поддержкой постраничных порций

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")
			SearchStr := req.Header.Get("Search")

			var recipesresp databases.RecipesResponse

			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			if len(PageStr) > 0 && len(LimitStr) > 0 && len(SearchStr) > 0 {

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

				recipesresp, err = databases.PostgreSQLRecipesSelectSearch(Page, Limit, SearchStr, dbc)

			} else {
				shared.HandleOtherError(w, ErrHeadersSearchNotFilled.Error(), ErrHeadersSearchNotFilled, http.StatusBadRequest)
				return
			}

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.WriteObjectToJSON(w, recipesresp)

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}
	default:
		shared.HandleOtherError(w, "Method is not allowed", shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
