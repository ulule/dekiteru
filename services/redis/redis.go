package redis

import (
	"log"

	redigo "github.com/garyburd/redigo/redis"
)

const (
	defaultURL = "redis://localhost:6379/0"
)

// Run checks Redis service.
func Run(parameters map[string]interface{}) (int, error) {
	var (
		c   redigo.Conn
		url string
		ok  bool
		err error
	)

	url, ok = parameters["url"].(string)
	if !ok || url == "" {
		url = defaultURL
	}

	log.Printf("url: \"%s\"\n", url)

	c, err = redigo.DialURL(url)
	if err != nil {
		log.Printf("Error: \"%s\"\n", err)
		return 10, err
	}
	defer c.Close()

	return 0, nil
}
