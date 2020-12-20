package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: transport}
	_, err = client.Get("https://example.com")

	if err != nil {
		log.Println("Error making request. ", err)
	}
}
