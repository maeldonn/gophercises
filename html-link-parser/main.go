package main

import (
	"fmt"
	"io"
	"os"

	"github.com/maeldonn/gophercises/html-link-parser/link"
)

func main() {
    r := mustReadFile("index.html")

    links, err := link.Parse(r)
    if err != nil {
        // TODO: Do smthg
    }

    fmt.Println(links)
}

func mustReadFile(filename string) io.Reader {
    buffer, err := os.Open("templates/" + filename)
    if err != nil {
        panic(err)
    }

    return buffer
}
