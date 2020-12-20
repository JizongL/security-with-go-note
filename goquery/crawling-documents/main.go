package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var documentExtensions = []string{"doc", "docx", "pdf", "csv", "xls", "xlsx", "zip", "gz", "tar"}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Find all links in a web page")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	url := os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if exists && linkContainsDocument(href) {
			fmt.Println(href)
		}
	})
}

func linkContainsDocument(url string) bool {
	urlPieces := strings.Split(url, ".")

	if len(urlPieces) < 2 {
		return false
	}
	for _, extension := range documentExtensions {
		if urlPieces[len(urlPieces)-1] == extension {
			return true
		}
	}
	return false
}
