package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	targetURL := "https://check.toproject.org"
	torProxy := "socks5://localhost:9050"
	torProxyURL, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ". ", err)
	}
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyURL)}
	client := &http.Client{
		Transport: torTransport,
		Timeout:   time.Second * 5,
	}
	response, err := client.Get(targetURL)
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
