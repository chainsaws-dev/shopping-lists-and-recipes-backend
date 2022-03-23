package multilangtranslator

type Translation struct {
	SearchKey  string
	Translated string
}

func (Tr Translations) SearchTranslation(SearchKey string) string {
	for _, t := range Tr {
		if SearchKey == t.SearchKey {
			return t.Translated
		}
	}

	return SearchKey
}

type Translations []Translation
