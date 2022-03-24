package multilangtranslator

import "errors"

func TranslateString(SearchKey string, Locale string) string {

	var ts Translations

	switch Locale {
	case "ru":
		ts = GetRussianTranslations()
	case "en":
		return SearchKey
	default:
		return SearchKey
	}

	return ts.SearchTranslation(SearchKey)
}

func TranslateError(err error, Locale string) error {
	return errors.New(TranslateString(err.Error(), Locale))
}
