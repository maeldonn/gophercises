package urlshort

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	yml, err := readFile(filename, fallback)
	if err != nil {
		return fallbackHandler(fallback), nil
	}

	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	yml, err := readFile(filename, fallback)
	if err != nil {
		return fallbackHandler(fallback), nil
	}

	parsedYaml, err := parseJSON(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func readFile(filename string, fallback http.Handler) ([]byte, error) {
	if filename == "" {
		return nil, errors.New("No filename specified")
	}

	return ioutil.ReadFile(filename)
}

func fallbackHandler(fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fallback.ServeHTTP(w, r)
	}
}
