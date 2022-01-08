package structs

type SearchResult struct {
	IPA               string             `json:"ipa"`
	TranslationGroups []TranslationGroup `json:"translationGroups"`
}

type TranslationGroup struct {
	Title        string             `json:"title"`
	Translations []TranslationEntry `json:"translations"`
}

type TranslationEntry struct {
	Id          string   `json:"id"`
	FromWord    string   `json:"fromWord"`
	FromType    string   `json:"fromType"`
	ToWord      string   `json:"toWord"`
	ToType      string   `json:"toType"`
	FromExample []string `json:"fromExample"`
	ToExample   []string `json:"toExample"`
}

func (entry *TranslationEntry) AddFromExample(example string) {
	entry.FromExample = append(entry.FromExample, example)
}

func (entry *TranslationEntry) AddToExample(example string) {
	entry.ToExample = append(entry.ToExample, example)
}
