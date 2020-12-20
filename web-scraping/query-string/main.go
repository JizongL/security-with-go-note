package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Search for a keyword in the contents of a URL")
		fmt.Println("Usage: " + os.Args[0] + " <url><keyword>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com/ NanoDano")
		os.Exit(1)
	}
	url := os.Args[1]
	needle := os.Args[2]
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := client.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}
	fmt.Println(string(body))
	if strings.Contains(string(body), needle) {
		fmt.Println("Match found for " + needle + " in URL " + url)
	} else {
		fmt.Println("No match found for " + needle + " in URL " + url)
	}
}
