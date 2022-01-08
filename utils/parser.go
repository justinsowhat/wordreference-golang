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

type Parser struct {
}

func (p *Parser) Parse(responseBody io.Reader) structs.SearchResult {

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
		IPA:               ipa,
		TranslationGroups: make([]structs.TranslationGroup, 0),
	}

	header := ""
	var translations []structs.TranslationEntry

	doc.Find("table.WRD").Each(func(i int, s *goquery.Selection) {

		prev_id := ""

		entry := structs.TranslationEntry{}

		s.Find("tr").Each(func(i int, node *goquery.Selection) {

			if p.isHeaderItem(node) {
				header, _ = node.Find("td").Attr("title")
				translations = []structs.TranslationEntry{}
			}

			id, isEntry := p.isTranslationEntry(node)

			if prev_id != "" && prev_id != id && id != "" {
				translations = append(translations, entry)
				entry = structs.TranslationEntry{
					Id: id,
				}
			}
			if prev_id != id && id != "" {
				prev_id = id
			}

			if isEntry {
				entry = p.parseTranslationEntry(entry, node)
			}
			if p.isExampleSentence(node) {
				entry = p.appendExample(entry, node)
			}

		})
		translations = append(translations, entry)

		result.TranslationGroups = append(result.TranslationGroups, structs.TranslationGroup{
			Title:        header,
			Translations: translations,
		})
	})

	return result
}

func (p *Parser) parseTranslationEntry(entry structs.TranslationEntry, node *goquery.Selection) structs.TranslationEntry {
	entry.FromWord = node.Find("strong").Text()

	node.Find(".ToWrd em span").Remove()
	node.Find(".FrWrd em span").Remove()

	entry.FromType = node.Find(".FrWrd em").Text()
	entry.ToType = node.Find(".ToWrd em").Text()

	node.Find(".ToWrd em").Remove()
	entry.ToWord = node.Find(".ToWrd").Text()

	return entry
}

func (p *Parser) isHeaderItem(element *goquery.Selection) bool {
	val, exists := element.Attr("class")
	if exists {
		return val == "wrtopsection"
	}
	return exists
}

func (p *Parser) isTranslationEntry(element *goquery.Selection) (string, bool) {
	id, id_exists := element.Attr("id")
	cls, _ := element.Attr("class")

	return id, id_exists && (cls == "even" || cls == "odd")
}

func (p *Parser) isExampleSentence(element *goquery.Selection) bool {
	_, id_exists := element.Attr("id")
	cls, _ := element.Attr("class")

	return !id_exists && (cls == "even" || cls == "odd")
}

func (p *Parser) appendExample(entry structs.TranslationEntry, element *goquery.Selection) structs.TranslationEntry {

	if element.Find(".FrEx").Text() != "" {
		entry.AddFromExample(element.Find(".FrEx").Text())
	}

	if element.Find(".ToEx").Text() != "" {
		entry.AddToExample(element.Find(".ToEx").Text())
	}

	return entry
}
