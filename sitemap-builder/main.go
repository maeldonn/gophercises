package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/maeldonn/gophercises/html-link-parser/link"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

var (
	urlFlag  string
	maxDepth int
)

type location struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []location `xml:"url"`
	Xmlns string     `xml:"xmlns,attr"`
}

func init() {
	flag.StringVar(&urlFlag, "url", "https://gophercises.com/", "The url that you want to build a sitemap for")
	flag.IntVar(&maxDepth, "depth", 3, "The maximum number of links deep to traverse")
	flag.Parse()
}

func main() {
	pages := bfs(urlFlag)

	err := saveToFile(pages)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The sitemap has been successfully generated for %s", urlFlag)
}

func bfs(urlStr string) []string {
	seen := make(map[string]struct{})
	var queue map[string]struct{}
	nextQueue := map[string]struct{}{
		urlStr: {},
	}

	for i := 0; i <= maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]struct{})
		if len(queue) == 0 {
			break
		}

		for url := range queue {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nextQueue[link] = struct{}{}
				}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		return []string{}
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
		return []string{}
	}

	return filter(pages, withPrefix(base))
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

func saveToFile(pages []string) error {
	buffer, err := encode(pages)
	if err != nil {
		return err
	}

	file, err := os.Create("sitemap.xml")
	if err != nil {
		return err
	}

	_, err = file.Write(buffer.Bytes())
	return err
}

func encode(pages []string) (*bytes.Buffer, error) {
	toXml := urlset{
		Urls:  make([]location, len(pages)),
		Xmlns: xmlns,
	}
	for i, page := range pages {
		toXml.Urls[i] = location{page}
	}

	buffer := bytes.NewBufferString(xml.Header)

	enc := xml.NewEncoder(buffer)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		return nil, err
	}
	if _, err := buffer.WriteString("\n"); err != nil {
		return nil, err
	}

	return buffer, nil
}
