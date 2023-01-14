package urlshort

import (
	"database/sql"
	"encoding/json"

	"github.com/go-sql-driver/mysql"
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

func getAllUrls() (map[string]string, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
	}
	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM path")
	if err != nil {
		return nil, err
	}

	paths := make(map[string]string)
	for rows.Next() {
		var path pathUrl
		if err = rows.Scan(&path.Path, &path.URL); err != nil {
			return nil, err
		}
		paths[path.Path] = path.URL
	}

	return paths, nil
}
