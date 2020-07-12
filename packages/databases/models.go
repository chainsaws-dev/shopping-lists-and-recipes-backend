package databases

// RecipeDB - тип для хранения информации о рецепте в базе данных
type RecipeDB struct {
	ID          int
	Name        string
	Description string
	ImagePath   string
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

// FilesResponse - тип для возврата с ответом,
// описывающий список файлов для постраничной разбивки
type FilesResponse struct {
	Files  []FileDB
	Total  int
	Offset int
	Limit  int
}
