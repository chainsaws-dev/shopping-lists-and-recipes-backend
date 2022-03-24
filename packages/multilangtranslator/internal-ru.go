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
		Translation{
			SearchKey:  "unencrypted webserver is up",
			Translated: "запущен веб сервер без шифрования",
		},
		Translation{
			SearchKey:  "encrypted webserver is up",
			Translated: "запущен веб сервер с шифрованием",
		},
		Translation{
			SearchKey:  "file with index %v deleted",
			Translated: "файл с индексом %v удалён",
		},
		Translation{
			SearchKey:  "http request does not contain required parameter: FileID",
			Translated: "не заполнен обязательный параметр для удаления файла: FileID",
		},
		Translation{
			SearchKey:  "unsupported file type",
			Translated: "неподдерживаемый тип файла",
		},
		Translation{
			SearchKey:  "логин или пароль",
			Translated: "выбран слишком короткий пароль",
		},
		Translation{
			SearchKey:  "wrong login or password",
			Translated: "неверный логин или пароль",
		},
		Translation{
			SearchKey:  "invalid email specified",
			Translated: "указана некорректная электронная почта",
		},
		Translation{
			SearchKey:  "invalid phone number specified",
			Translated: "указан некорректный телефонный номер",
		},
		Translation{
			SearchKey:  "invalid role specified",
			Translated: "указана некорректная роль",
		},
		Translation{
			SearchKey:  "http request does not contain required parameters",
			Translated: "не заполнены обязательные параметры запроса",
		},
		Translation{
			SearchKey:  "invalid http request parameters Limit and Offset",
			Translated: "параметры Limit и Offset приняли недопустимое значение",
		},
		Translation{
			SearchKey:  "session is not found for specified email",
			Translated: "сессия не найдена для данной электронной почты",
		},
		Translation{
			SearchKey:  "session is not found for specified token",
			Translated: "сессия не найдена для данного токена",
		},
		Translation{
			SearchKey:  "sessions deleted",
			Translated: "сессии удалены",
		},
		Translation{
			SearchKey:  "session deleted",
			Translated: "сессия удалена",
		},
		Translation{
			SearchKey:  "email was sent",
			Translated: "письмо отправлено",
		},
		Translation{
			SearchKey:  "email successfully confirmed",
			Translated: "электронная почта подтверждена",
		},
		Translation{
			SearchKey:  "error calculating password hash",
			Translated: "ошибка при расчете хеша",
		},
		Translation{
			SearchKey:  "password changed",
			Translated: "пароль обновлён",
		},
		Translation{
			SearchKey:  "you are denied access to the resource",
			Translated: "вам закрыт доступ на ресурс",
		},
		Translation{
			SearchKey:  "password must be more than six characters",
			Translated: "пароль должен быть более шести символов",
		},
		Translation{
			SearchKey:  "password of the new user must be set",
			Translated: "пароль нового пользователя должен быть задан",
		},
		Translation{
			SearchKey:  "specified email address is already taken",
			Translated: "указанный адрес электронной почты уже занят",
		},
		Translation{
			SearchKey:  "invalid user id specified",
			Translated: "некорректный идентификатор пользователя",
		},
		Translation{
			SearchKey:  "user not found, unable to delete",
			Translated: "пользователь не найден, невозможно удалить",
		},
		Translation{
			SearchKey:  "user deleted",
			Translated: "пользователь удалён",
		},
		Translation{
			SearchKey:  "recipe not found, unable to delete",
			Translated: "рецепт не найден, невозможно удалить",
		},
		Translation{
			SearchKey:  "recipe deleted",
			Translated: "рецепт удалён",
		},
		Translation{
			SearchKey:  "entry saved",
			Translated: "запись сохранена",
		},
		Translation{
			SearchKey:  "no shopping list entry found with specified name",
			Translated: "не найдено ни одной записи в списке покупок с указанным названием",
		},
		Translation{
			SearchKey:  "entry deleted",
			Translated: "запись удалена",
		},
		Translation{
			SearchKey:  "all entries deleted",
			Translated: "все записи удалены",
		},
	}

}
