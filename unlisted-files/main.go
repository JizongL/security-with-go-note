package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func checkIfURLExists(baseURL, filePath string, doneChannel chan bool) {
	targetURL, err := url.Parse(baseURL)
	if err != nil {
		log.Println("Error parsing base URL. ", err)
	}
	targetURL.Path = filePath
	response, err := http.Head(targetURL.String())
	file, err := os.OpenFile("record.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal("Writting error.")
	}
	// recordWriter := bufio.NewWriter(file)
	if err != nil {
		log.Println("Error fetching ", targetURL.String())
	}
	if response.StatusCode == 200 {
		// written, err := recordWriter.Write([]byte(targetURL.String()))
		// if err != nil {
		// log.Fatal(err)
		// }
		// log.Printf("wrote %d bytes.\n", written)
		log.Println(targetURL.String())
	}
	doneChannel <- true
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println(os.Args[0] + " - Perform an HTTP HEAD request to a URL")
		fmt.Println("Usage: " + os.Args[0] + " <wordlist_file><url><maxThreads>")
		fmt.Println("Example: " + os.Args[0] + " wordlist.txt https://devdungeon.com 10")
		os.Exit(1)
	}
	wordlistFilename := os.Args[1]
	baseURL := os.Args[2]
	maxThreads, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("Error converting maxThread value to integer. ", err)
	}
	activeThreads := 0
	doneChannel := make(chan bool)
	wordlistFile, err := os.Open(wordlistFilename)
	if err != nil {
		log.Fatal("Error opening wordlist file. ", err)
	}
	scanner := bufio.NewScanner(wordlistFile)
	for scanner.Scan() {
		go checkIfURLExists(baseURL, scanner.Text(), doneChannel)
		activeThreads++
		if activeThreads >= maxThreads {
			<-doneChannel
			activeThreads = activeThreads - 1
		}
	}
	for activeThreads > 0 {
		<-doneChannel
		activeThreads = activeThreads - 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading wordlist file. ", err)
	}

}
