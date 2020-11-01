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

// File - тип для хранения информации о файле в базе данных
type File struct {
	ID       int
	Filename string
	Filesize int
	Filetype string
	FileID   string
}

// FilesList - тип для хранения списка файлов
type FilesList []File

// FilesResponse - тип для возврата с ответом,
// описывающий список файлов для постраничной разбивки
type FilesResponse struct {
	Files  FilesList
	Total  int
	Offset int
	Limit  int
}

// ShoppingListResponse  - тип для возврата с ответом,
// описывающий список покупок для постраничной разбивки
type ShoppingListResponse struct {
	Items  IngredientsDB
	Total  int
	Offset int
	Limit  int
}

// User - тип для хранения данных о пользователе в базе данных
type User struct {
	GUID         uuid.UUID
	Role         string
	Email        string
	Phone        string
	Name         string
	IsAdmin      bool
	Confirmed    bool
	Disabled     bool
	SecondFactor bool
}

// Users - тип для хранения списка пользователей
type Users []User

// UsersResponse  - тип для возврата с ответом,
// описывающий список пользователей для постраничной разбивки
type UsersResponse struct {
	Users  Users
	Total  int
	Offset int
	Limit  int
}

// TOTPSecret - секрет для Time Based One Time Password
type TOTPSecret struct {
	UserID    uuid.UUID
	Secret    string
	EncKey    []byte
	Confirmed bool
}

// TOTPResponse - тип для подтверждения наличия секрета
type TOTPResponse struct {
	UserID    uuid.UUID
	Confirmed bool
}
