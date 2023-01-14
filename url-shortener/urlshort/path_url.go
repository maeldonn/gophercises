package urlshort

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(yml []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl

	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

func parseJSON(j []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl

	err := json.Unmarshal(j, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return pathsToUrls
}
