package databases

// Recipe - тип для хранения информации о рецепте в базе данных
type Recipe struct {
	id          int
	name        string
	description string
	imagePath   string
	ingredients []Ingredient
}

// Ingredient - тип для хранения информации о ингредиенте
type Ingredient struct {
	name   string
	amount int
}

// RecipesResponse  - тип для возврата с ответом,
// описывающий список рецептов с обложками для постраничной разбивки
type RecipesResponse struct {
	Recipes []Recipe
	Total   int
	Offset  int
	Limit   int
}

// FileDB - тип для хранения информации о файле в базе данных
type FileDB struct {
	id       int
	filename string
	filesize int64
	filetype string
	fileID   string
}

// FilesResponse - тип для возврата с ответом,
// описывающий список файлов для постраничной разбивки
type FilesResponse struct {
	Files  []FileDB
	Total  int64
	Offset int64
	Limit  int64
}
