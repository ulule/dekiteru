package contrib

import (
	"log"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

func CheckElasticSearch() error {
	url := os.Getenv("DEKITERU_ELASTICSEARCH_URL")
	if url == "" {
		url = "http://localhost:9201"
	}
	log.Printf("ELASTICSEARCH_URL: \"%s\"\n", url)

	_, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Printf("Error: \"%s\"\n", err)
		return err
	}
	return nil
}
