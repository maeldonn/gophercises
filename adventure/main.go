package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maeldonn/gophercises/adventure/story"
)

var (
	filename string
	port     int
)

func init() {
	flag.StringVar(&filename, "file", "gopher.json", "The JSON file with the story")
	flag.IntVar(&port, "port", 3000, "The port to start the web server")
	flag.Parse()
}

func main() {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	s, err := story.ParseStory(f)
	if err != nil {
		panic(err)
	}

	handler := story.NewHandler(s)

	fmt.Printf("Starting the server on port :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
