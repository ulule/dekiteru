package services

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

// Redis service
type Redis struct{}

// Run implements Service interface.
func (Redis) Run(parameters map[string]interface{}) error {
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
		return &SoftError{err}
	}

	defer c.Close()

	return nil
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
