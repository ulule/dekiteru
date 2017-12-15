package services

import (
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

// ElasticSearch service
type ElasticSearch struct{}

// Run implements Service interface.
func (ElasticSearch) Run(parameters map[string]interface{}) (int, error) {
	var (
		url string
		ok  bool
		err error
	)

	url, ok = parameters["url"].(string)
	if !ok || url == "" {
		url = "http://localhost:9201"
	}

	log.Printf(`url: "%s"`, url)

	_, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Printf(`Error: "%s"`, err)
		return 2, err
	}

	return 0, nil
}

// Name implements Service interface.
func (ElasticSearch) Name() string {
	return "elasticsearch"
}

// Parameters implements Service interface.
func (ElasticSearch) Parameters() []string {
	return []string{
		"url",
	}
}
