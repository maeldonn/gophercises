package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/maeldonn/gophercises/html-link-parser/link"
)

var urlFlag string

func init() {
	flag.StringVar(&urlFlag, "url", "https://gophercises.com", "The url that you want to build a sitemap for")
	flag.Parse()
}

func main() {
	pages, err := get(urlFlag)
	if err != nil {
		panic("Impossible to get links from page")
	}

	for _, p := range pages {
		fmt.Println(p)
	}
}

func get(urlStr string) ([]string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic("Impossible to get website")
	}

	defer resp.Body.Close()
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	pages, err := getHrefs(resp.Body, base)
	if err != nil {
		return nil, err
	}

	return filter(pages, withPrefix(base)), nil
}

func getHrefs(body io.Reader, base string) ([]string, error) {
	links, err := link.Parse(body)
	if err != nil {
		return nil, err
	}

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	return hrefs, nil
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
