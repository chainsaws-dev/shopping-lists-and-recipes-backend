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

		// Проверка прохождения двухфакторной авторизации
		sf := signinupout.SecondFactorAuthenticationCheck(w, req)

		if issued {
			if sf {
				switch {
				case req.Method == http.MethodGet:

					if setup.ServerSettings.CheckRoleForRead(role, "HandleRecipes") {
						// Обработка получения списка рецептов с поддержкой постраничных порций

						PageStr := req.Header.Get("Page")
						LimitStr := req.Header.Get("Limit")

						var recipesresp databases.RecipesResponse
						var err error

						err = setup.ServerSettings.SQL.Connect(role)

						if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
							return
						}
						defer setup.ServerSettings.SQL.Disconnect()

						if len(PageStr) > 0 && len(LimitStr) > 0 {

							Page, err := strconv.Atoi(PageStr)

							if shared.HandleInternalServerError(w, err) {
								return
							}

							Limit, err := strconv.Atoi(LimitStr)

							if shared.HandleInternalServerError(w, err) {
								return
							}

							recipesresp, err = databases.PostgreSQLRecipesSelect(Page, Limit)

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

						err := json.NewDecoder(req.Body).Decode(&Recipe)

						if shared.HandleOtherError(w, "Bad request", err, http.StatusBadRequest) {
							return
						}

						err = setup.ServerSettings.SQL.Connect(role)

						if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
							return
						}
						defer setup.ServerSettings.SQL.Disconnect()

						recipesresp, err := databases.PostgreSQLRecipesInsertUpdate(Recipe)

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

							err = setup.ServerSettings.SQL.Connect(role)

							if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
								return
							}
							defer setup.ServerSettings.SQL.Disconnect()

							err = databases.PostgreSQLRecipesDelete(RecipeID)

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
			} else {
				shared.HandleOtherError(w, shared.ErrNotAuthorizedTwoFactor.Error(), shared.ErrNotAuthorizedTwoFactor, http.StatusUnauthorized)
			}
		} else {
			shared.HandleOtherError(w, shared.ErrNotAuthorized.Error(), shared.ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", shared.ErrWrongKeyInParams, http.StatusBadRequest)
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

		// Проверка прохождения двухфакторной авторизации
		sf := signinupout.SecondFactorAuthenticationCheck(w, req)

		if issued {
			if sf {
				switch {
				case req.Method == http.MethodGet:

					if setup.ServerSettings.CheckRoleForRead(role, "HandleRecipesSearch") {
						// Обработка получения списка рецептов с поддержкой постраничных порций

						PageStr := req.Header.Get("Page")
						LimitStr := req.Header.Get("Limit")
						SearchStr := req.Header.Get("Search")

						var recipesresp databases.RecipesResponse
						var err error

						err = setup.ServerSettings.SQL.Connect(role)

						if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
							return
						}
						defer setup.ServerSettings.SQL.Disconnect()

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

							recipesresp, err = databases.PostgreSQLRecipesSelectSearch(Page, Limit, SearchStr)

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
			} else {
				shared.HandleOtherError(w, shared.ErrNotAuthorizedTwoFactor.Error(), shared.ErrNotAuthorizedTwoFactor, http.StatusUnauthorized)
			}
		} else {
			shared.HandleOtherError(w, shared.ErrNotAuthorized.Error(), shared.ErrNotAuthorized, http.StatusUnauthorized)
		}
	} else {
		shared.HandleOtherError(w, "Bad request", shared.ErrWrongKeyInParams, http.StatusBadRequest)
	}
}
