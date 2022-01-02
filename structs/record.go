package structs

type SearchResult struct {
	IPA                    string
	PrincipalTranslations  []TranslationEntry
	AdditionalTranslations []TranslationEntry
	CompoundForms          []TranslationEntry
}

type TranslationEntry struct {
	FromWord    string
	FromType    string
	ToWord      string
	ToType      string
	FromExample []string
	ToExample   []string
}

func (entry *TranslationEntry) AddFromExample(example string) {
	entry.FromExample = append(entry.FromExample, example)
}

func (entry *TranslationEntry) AddToExample(example string) {
	entry.ToExample = append(entry.ToExample, example)
}
