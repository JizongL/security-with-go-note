package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Search for HTML comments in a URL")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	url := os.Args[1]
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading HTTP body. ", err)
	}
	re := regexp.MustCompile("<!--(.|\n)*?-->")
	matches := re.FindAllString(string(body), -1)
	if matches == nil {
		fmt.Println("No HTML comments found.")
		os.Exit(0)
	}
	for _, match := range matches {
		fmt.Println(match)
	}
}
