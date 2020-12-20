package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	foundPaths  []string
	startingURL *url.URL
	timeout     = time.Duration(8 * time.Second)
)

func crawlURL(path string) {
	var targetURL url.URL
	targetURL.Scheme = startingURL.Scheme
	targetURL.Host = startingURL.Host
	targetURL.Path = path
	httpClient := http.Client{Timeout: timeout}
	response, err := httpClient.Get(targetURL.String())
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		fmt.Println(href, "href test")
		parsedURL, err := url.Parse(href)
		if err != nil {
			return
		}
		if urlIsInScope(parsedURL) {
			foundPaths = append(foundPaths, parsedURL.Path)
			log.Println("Found new path to crawl: " + parsedURL.String())
			crawlURL(parsedURL.Path)
		}
	})

}

func urlIsInScope(tempURL *url.URL) bool {
	if tempURL.Host != "" && tempURL.Host != startingURL.Host {
		return false
	}
	if tempURL.Path == "" {
		return false
	}
	for _, existingPath := range foundPaths {
		if existingPath == tempURL.Path {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Crawl a website, depth-first")
		fmt.Println("Usage: " + os.Args[0] + "<startingUrl>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	foundPaths = make([]string, 0)
	startingURL = new(url.URL)
	startingURL, err := url.Parse(os.Args[1])
	fmt.Println(startingURL, "starting url 81")
	if err != nil {
		log.Fatal("Error parsing starting URL. ", err)
	}
	crawlURL(startingURL.Path)
	for _, path := range foundPaths {
		fmt.Println(path)
	}
	log.Printf("Total unique paths crawled: %d\n", len(foundPaths))
}
