package services

import (
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

type ElasticSearchService struct{}

func (ElasticSearchService) Run(parameters map[string]interface{}) (int, error) {
	var (
		url string
		ok  bool
		err error
	)
	url, ok = parameters["url"].(string)
	if !ok || url == "" {
		url = "http://localhost:9201"
	}
	log.Printf("url: \"%s\"\n", url)

	_, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Printf("Error: \"%s\"\n", err)
		return 10, err
	}
	return 0, nil
}

func (ElasticSearchService) Name() string {
	return "elasticsearch"
}

func (ElasticSearchService) Parameters() []string {
	return []string{
		"url",
	}
}
