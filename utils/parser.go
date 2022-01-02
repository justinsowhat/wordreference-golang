package utils

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/justinsowhat/wordreference-golang/structs"
)

const PRINCIPAL_TRANSLATIONS = "Principal Translations"
const ADDITIONAL_TRANSLATION = "Additional Translations"
const COMPOUND_FORMS = "Compound Forms"

func Parse(responseBody io.Reader) structs.SearchResult {

	if responseBody == nil {
		return structs.SearchResult{}
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(responseBody)
	if err != nil {
		log.Fatal(err)
		return structs.SearchResult{}
	}

	if doc == nil {
		log.Fatal("document is nil!")
		return structs.SearchResult{}
	}

	ipa := doc.Find("span#pronWR").Text()

	result := structs.SearchResult{
		IPA: ipa,
	}

	header := ""
	var translations []structs.TranslationEntry

	doc.Find("table.WRD").Each(func(i int, s *goquery.Selection) {

		prev_id := ""

		entry := structs.TranslationEntry{}

		s.Find("tr").Each(func(i int, node *goquery.Selection) {

			if isHeaderItem(node) {
				header, _ = node.Find("td").Attr("title")
				translations = []structs.TranslationEntry{}
			}

			id, isEntry := isTranslationEntry(node)

			if prev_id != "" && prev_id != id && id != "" {
				translations = append(translations, entry)
				entry = structs.TranslationEntry{}
			}
			if prev_id != id && id != "" {
				prev_id = id
			}

			if isEntry {
				entry = parseTranslationEntry(entry, node)
			}
			if isExampleSentence(node) {
				entry = appendExample(entry, node)
			}

		})
		translations = append(translations, entry)

		result = setTranslationsByHeader(header, result, translations)
	})

	return result
}

func parseTranslationEntry(entry structs.TranslationEntry, node *goquery.Selection) structs.TranslationEntry {
	entry.FromWord = node.Find("strong").Text()

	node.Find(".ToWrd em span").Remove()
	node.Find(".FrWrd em span").Remove()

	entry.FromType = node.Find(".FrWrd em").Text()
	entry.ToType = node.Find(".ToWrd em").Text()

	node.Find(".ToWrd em").Remove()
	entry.ToWord = node.Find(".ToWrd").Text()

	return entry
}

func setTranslationsByHeader(header string, result structs.SearchResult, translations []structs.TranslationEntry) structs.SearchResult {
	if header == PRINCIPAL_TRANSLATIONS && result.PrincipalTranslations == nil {
		result.PrincipalTranslations = translations
	} else if header == ADDITIONAL_TRANSLATION && result.AdditionalTranslations == nil {
		result.AdditionalTranslations = translations
	} else if header == COMPOUND_FORMS && result.CompoundForms == nil {
		result.CompoundForms = translations
	}
	return result
}

func isHeaderItem(element *goquery.Selection) bool {
	val, exists := element.Attr("class")
	if exists {
		return val == "wrtopsection"
	}
	return exists
}

func isTranslationEntry(element *goquery.Selection) (string, bool) {
	id, id_exists := element.Attr("id")
	cls, _ := element.Attr("class")

	return id, id_exists && (cls == "even" || cls == "odd")
}

func isExampleSentence(element *goquery.Selection) bool {
	_, id_exists := element.Attr("id")
	cls, _ := element.Attr("class")

	return !id_exists && (cls == "even" || cls == "odd")
}

func appendExample(entry structs.TranslationEntry, element *goquery.Selection) structs.TranslationEntry {

	if element.Find(".FrEx").Text() != "" {
		entry.AddFromExample(element.Find(".FrEx").Text())
	}

	if element.Find(".ToEx").Text() != "" {
		entry.AddToExample(element.Find(".ToEx").Text())
	}

	return entry
}
