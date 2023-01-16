package story

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Options
}

type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func ParseStory(r io.Reader) (Story, error) {
	var story Story

	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
