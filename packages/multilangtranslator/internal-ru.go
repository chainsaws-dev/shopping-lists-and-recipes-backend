package multilangtranslator

func GetRussianTranslations() Translations {
	return Translations{
		Translation{
			SearchKey:  "internal server error",
			Translated: "внутренняя ошибка сервера",
		},
		Translation{
			SearchKey:  "bad http request",
			Translated: "ошибка в http запросе",
		},
		Translation{
			SearchKey:  "http request to delete recipe does not contain required header RecipeID",
			Translated: "не заполнен обязательный заголовок RecipeID в запросе на удаление рецепта",
		},
		Translation{
			SearchKey:  "http request to search recipes does not contain required parameters: Page, Limit, Search",
			Translated: "не заполнены обязательные параметры поискового запроса: Page, Limit, Search",
		},
		Translation{
			SearchKey:  "http request method is not allowed",
			Translated: "запрошен недопустимый метод",
		},
		Translation{
			SearchKey:  "http request does not contain api key parameter",
			Translated: "api ключ не указан в параметрах",
		},
		Translation{
			SearchKey:  "api key is not registered",
			Translated: "api ключ не зарегистрирован",
		},
		Translation{
			SearchKey:  "authorization required",
			Translated: "пройдите авторизацию",
		},
		Translation{
			SearchKey:  "two factor authorization required",
			Translated: "пройдите авторизацию по второму фактору",
		},
		Translation{
			SearchKey:  "access forbidden",
			Translated: "доступ запрещён",
		},
		Translation{
			SearchKey:  "http request does not contain required parameters: Page, Limit",
			Translated: "не заполнены обязательные параметры запроса: Page, Limit",
		},
		Translation{
			SearchKey:  "cannot delete record referenced from other tables",
			Translated: "нельзя удалять записи, на которые имеются ссылки",
		},
		Translation{
			SearchKey:  "no rows in result set",
			Translated: "запрос вернул пустую таблицу",
		},
		Translation{
			SearchKey:  "invalid log level",
			Translated: "некорректный уровень журнала",
		},
		Translation{
			SearchKey:  "nothing to delete",
			Translated: "не найдено ни одной записи для удаления",
		},
		Translation{
			SearchKey:  "incorrect data format for login: expected -admincred:example@example.ru@@password",
			Translated: "неверный формат данных для логина: ожидатся -admincred:example@example.ru@@password",
		},
		Translation{
			SearchKey:  "server is shutting down...",
			Translated: "завершение работы сервера...",
		},
	}

}
