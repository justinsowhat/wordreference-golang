package WordReferenceClient

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"

	s "github.com/justinsowhat/wordreference-golang/structs"
	utils "github.com/justinsowhat/wordreference-golang/utils"
)

const clientUrl = "https://www.wordreference.com/"

type WordReferenceClient struct {
	Dict string
}

// context path for available bidirectional dictionaries
const (
	ENGLISH_SPANISH   string = "es/translation.asp?tranword="
	SPANISH_ENGLISH          = "es/en/translation.asp?spen="
	SPANISH_FRENCH           = "esfr/"
	SPANISH_ITALIAN          = "esit/"
	SPANISH_GERMAN           = "esde/"
	SPANISH_PORTUGESE        = "espt/"
	ENGLISH_FRENCH           = "enfr/"
	FRENCH_ENGLISH           = "fren/"
	FRENCH_SPANISH           = "fres/"
	ITALIAN_ENGLISH          = "iten/"
	ENGLISH_ITALIAN          = "enit/"
	ITALIAN_SPANISH          = "ites/"
	ENGLISH_GERMAN           = "ende/"
	GERMAN_ENGLISH           = "deen/"
	GERMAN_SPANISH           = "dees/"
	ENGLISH_DUTCH            = "enhl/"
	DUTCH_ENGLISH            = "hlen/"
	ENGLISH_SWEDISH          = "ensv/"
	SWEDISH_ENGLISH          = "sven/"
	ENGLISH_RUSSIAN          = "enru/"
	RUSSIAN_ENGLISH          = "ruen/"
	ENGLISH_PORTUGESE        = "enpt/"
	PORTUGESE_ENGLISH        = "pten/"
	PORTUGESE_SPANISH        = "ptes/"
	ENGLISH_POLISH           = "enpl/"
	POLISH_ENGLISH           = "plen/"
	ENGLISH_ROMANIAN         = "enro/"
	ROMANIAN_ENGLISH         = "roen/"
	ENGLISH_CZECH            = "encz/"
	CZECH_ENGLISH            = "czen/"
	ENGLISH_GREEK            = "engr/"
	GREEK_ENGLISH            = "gren/"
	ENGLISH_TURKISH          = "entr/"
	TURKISH_ENGLISH          = "tren/"
	ENGLISH_CHINESE          = "enzh/"
	CHINESE_ENGLISH          = "zhen/"
	ENGLISH_JAPANESE         = "enja/"
	JAPANESE_ENGLISH         = "jaen/"
	ENGLISH_KOREAN           = "enko/"
	KOREAN_ENGLISH           = "koen/"
	ENGLISH_ARABIC           = "enar/"
	ARABIC_ENGLISH           = "aren/"
	ENGLISH_ICELANDIC        = "enis/"
)

func (client WordReferenceClient) sendGetRequest(entry string) io.ReadCloser {
	url := clientUrl + client.Dict + client.escapeSpaces(entry)

	responseBody := bytes.NewBufferString("")

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Body
}

func (client WordReferenceClient) LookUpWord(entry string) s.SearchResult {
	html := client.sendGetRequest(entry)
	parser := utils.Parser{}
	return parser.Parse(html)
}

func (client WordReferenceClient) escapeSpaces(entry string) string {
	pattern := regexp.MustCompile(`\+`)
	return pattern.ReplaceAllString(url.QueryEscape(entry), "%20")
}
