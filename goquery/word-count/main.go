package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var total int
	if len(os.Args) != 2 {
		fmt.Println("List all words by frequency from a web page")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	url := os.Args[1]
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	wordCountMap := make(map[string]int)
	doc.Find("p").Each(func(i int, body *goquery.Selection) {
		fmt.Println(body.Text())
		words := strings.Split(body.Text(), " ")
		for _, word := range words {
			trimmedWord := strings.Trim(word, "\t\n\r,.?!")
			if trimmedWord == "" {
				continue
			}
			total++
			wordCountMap[strings.ToLower(trimmedWord)]++
		}
	})
	for word, count := range wordCountMap {
		fmt.Printf("%d | %s\n", count, word)
	}
	fmt.Printf("Total word count is: %d\n", total)
}
