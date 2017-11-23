package contrib

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

func CheckRedis() error {
	url := os.Getenv("DEKITERU_REDIS_URL")
	if url == "" {
		url = "redis://localhost:6379/0"
	}
	log.Printf("REDIS_URL: \"%s\"\n", url)
	c, err := redis.DialURL(url)
	if err != nil {
		log.Printf("Error: \"%s\"\n", err)
		return err
	}
	defer c.Close()
	return nil
}
