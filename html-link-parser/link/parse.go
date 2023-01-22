package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represent a link (<a href"...">) in a HTML document
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	depthFirstSearch(doc, "")
	return nil, nil
}

func depthFirstSearch(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		depthFirstSearch(c, padding+"  ")
	}
}
