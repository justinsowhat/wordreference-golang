package client

import (
	"bytes"
	"fmt"
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

func (client WordReferenceClient) sendGetRequest(entry string) io.ReadCloser {
	url := clientUrl + client.Dict + "/" + escapeSpaces(entry)
	fmt.Println(url)

	responseBody := bytes.NewBufferString("")

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Body
}

func (client WordReferenceClient) LookUpWord(entry string) s.SearchResult {
	html := client.sendGetRequest(entry)

	return utils.Parse(html)
}

func escapeSpaces(entry string) string {
	pattern := regexp.MustCompile(`\+`)
	return pattern.ReplaceAllString(url.QueryEscape(entry), "%20")
}
