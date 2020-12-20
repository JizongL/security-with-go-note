package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	url := "http://www.devdungeon.com/archive"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Downloading %d byte file.\n", numBytesWritten)
}
