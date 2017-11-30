package elasticsearch

import (
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

const (
	defaultURL = "http://localhost:9201"
)

// Check checks Elasticsearch service.
func Check(parameters map[string]interface{}) (int, error) {
	var (
		url string
		ok  bool
		err error
	)

	url, ok = parameters["url"].(string)
	if !ok || url == "" {
		url = defaultURL
	}

	log.Printf(`url: "%s"\n`, url)

	_, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Printf(`Error: "%s"\n`, err)
		return 10, err
	}

	return 0, nil
}
