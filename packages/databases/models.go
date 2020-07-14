package databases

import uuid "github.com/satori/go.uuid"

// RecipeDB - тип для хранения информации о рецепте в базе данных
type RecipeDB struct {
	ID          int
	Name        string
	Description string
	ImagePath   string
	ImageDbID   int
	Ingredients IngredientsDB
}

// RecipesDB - тип для хранения массива рецептов
type RecipesDB []RecipeDB

// IngredientDB - тип для хранения информации о ингредиенте
type IngredientDB struct {
	Name   string
	Amount int
}

// IngredientsDB - тип для хранения слайсов ингредиентов
type IngredientsDB []IngredientDB

// RecipesResponse  - тип для возврата с ответом,
// описывающий список рецептов с обложками для постраничной разбивки
type RecipesResponse struct {
	Recipes RecipesDB
	Total   int
	Offset  int
	Limit   int
}

// FileDB - тип для хранения информации о файле в базе данных
type FileDB struct {
	ID       int
	Filename string
	Filesize int
	Filetype string
	FileID   string
}

// FilesList - тип для хранения списка файлов
type FilesList []FileDB

// FilesResponse - тип для возврата с ответом,
// описывающий список файлов для постраничной разбивки
type FilesResponse struct {
	Files  FilesList
	Total  int
	Offset int
	Limit  int
}

// ShoppingListDB - тип для хранения информации
// о списке покупок в базе данных
type ShoppingListDB struct {
	Items []IngredientDB
}

// ShoppingListResponse  - тип для возврата с ответом,
// описывающий список покупок для постраничной разбивки
type ShoppingListResponse struct {
	Items  ShoppingListDB
	Total  int
	Offset int
	Limit  int
}

// UserInfoDB - Информация о пользователе в базе данных
type UserInfoDB struct {
	GUID    uuid.UUID
	Role    string
	Email   string
	Phone   string
	Name    string
	IsAdmin bool
}

// UsersListDB - тип для хранения списка пользователей
type UsersListDB []UserInfoDB

// UserListResponse  - тип для возврата с ответом,
// описывающий список пользователей для постраничной разбивки
type UserListResponse struct {
	Items  UsersListDB
	Total  int
	Offset int
	Limit  int
}
