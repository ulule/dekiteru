package services

import (
	"github.com/ulule/dekiteru/services/elasticsearch"
	"github.com/ulule/dekiteru/services/postgresql"
	"github.com/ulule/dekiteru/services/redis"
)

// Checker is a service checker.
type Checker func(parameters map[string]interface{}) (int, error)

// Services are built-in services.
var Services = map[string]Checker{
	"postgresql":    postgresql.Check,
	"redis":         redis.Check,
	"elasticsearch": elasticsearch.Check,
}
