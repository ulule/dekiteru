package services

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type Redis struct{}

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
		log.Printf("Error: \"%s\"\n", err)
		return 10, err
	}
	defer c.Close()
	return 0, nil
}

func (Redis) Name() string {
	return "redis"
}

func (Redis) Parameters() []string {
	return []string{
		"url",
	}
}
