package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("List all Javascript files in a webpage")
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
	fmt.Println("Scripts found in", url)
	fmt.Println("===================")
	doc.Find("script").Each(func(i int, script *goquery.Selection) {
		src, exists := script.Attr("src")
		if exists {
			fmt.Println(src)
		}
	})
}
