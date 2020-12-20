package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You requested: "+request.URL.Path)
}

func customMiddlewareHandler(r http.ResponseWriter, q *http.Request, next http.HandlerFunc) {
	log.Println("Incoming request: " + q.URL.Path)
	log.Println("User agent: " + q.UserAgent())
	next(r, q)
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", indexHandler)
	negroniHandler := negroni.New()
	negroniHandler.Use(negroni.HandlerFunc(customMiddlewareHandler))
	http.ListenAndServe("localhost:3000", negroniHandler)
}
