package services

import (
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

// Elasticsearch service
type Elasticsearch struct{}

// Run implements Service interface.
func (Elasticsearch) Run(parameters map[string]interface{}) (int, error) {
	var (
		url string
		ok  bool
		err error
	)

	url, ok = parameters["url"].(string)
	if !ok || url == "" {
		url = "http://localhost:9201"
	}

	log.Printf(`url: "%s"\n`, url)

	_, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Printf(`Error: "%s"\n`, err)
		return 10, err
	}

	return 0, nil
}

// Name implements Service interface.
func (Elasticsearch) Name() string {
	return "elasticsearch"
}

// Parameters implements Service interface.
func (Elasticsearch) Parameters() []string {
	return []string{
		"url",
	}
}
