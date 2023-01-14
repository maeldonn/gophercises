package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/maeldonn/gophercises/url-shortener/urlshort"
)

var (
	yamlFilename string
	jsonFilename string
)

func init() {
	flag.StringVar(&yamlFilename, "yaml", "", "a yaml file in a format path, url")
	flag.StringVar(&jsonFilename, "json", "", "a json file in a format path, url")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	yamlHandler, err := urlshort.YAMLHandler(yamlFilename, mux)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler(jsonFilename, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Invalid path")
}
