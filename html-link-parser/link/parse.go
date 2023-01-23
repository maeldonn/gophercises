package link

import (
	"io"
	"strings"

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
	nodes := linkNodes(doc)
	links := make([]Link, len(nodes))
	for i, node := range nodes {
		links[i] = buildLink(node)
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var l Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			l.Href = attr.Val
			break
		}
	}
	l.Text = getText(n)
	return l
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var builder strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		builder.WriteString(getText(c))
	}
	return strings.Join(strings.Fields(builder.String()), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}
	return nodes
}
