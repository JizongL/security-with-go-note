package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	proxyURLString := "http://52.149.152.236:80"
	proxyURL, err := url.Parse(proxyURLString)
	if err != nil {
		log.Fatal("Error parsing URL. ", err)
	}
	customTransport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	httpClient := &http.Client{
		Transport: customTransport,
		Timeout:   time.Second * 5,
	}
	response, err := httpClient.Get("http://www.example.com")
	if err != nil {
		log.Fatal("Error making GET request. ", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading body of response. ", err)
	}
	log.Println(string(body))
}
