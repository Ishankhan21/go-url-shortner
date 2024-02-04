package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	urlshort "github.com/Ishankhan21/go-url-shortner/handler"
)

func main() {
	mux := defaultMux()

	// TODO: Add URLs in Map or JSON
	pathsToUrls := map[string]string{}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	jsonBytes, err := ioutil.ReadFile("URLS.JSON")
	if err != nil {
		fmt.Println("ERROR ++++", err)
	}

	JSONHandler, err := urlshort.JSONHandler(jsonBytes, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", JSONHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
