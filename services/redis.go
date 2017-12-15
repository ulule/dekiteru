package services

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

// Redis service
type Redis struct{}

// Run implements Service interface.
func (Redis) Run(parameters map[string]interface{}) (int, error) {
	var (
		c   redis.Conn
		url string
		ok  bool
		err error
	)

	url, ok = parameters["url"].(string)
	if !ok || url == "" {
		url = "redis://localhost:6379/0"
	}

	log.Printf("url: \"%s\"\n", url)

	c, err = redis.DialURL(url)
	if err != nil {
		log.Printf(`Error: "%s"`, err)
		return 10, err
	}

	defer c.Close()

	return 0, nil
}

// Name implements Service interface
func (Redis) Name() string {
	return "redis"
}

// Parameters implements Service interface
func (Redis) Parameters() []string {
	return []string{
		"url",
	}
}
