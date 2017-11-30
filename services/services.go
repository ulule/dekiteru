package services

import (
	"github.com/ulule/dekiteru/services/elasticsearch"
	"github.com/ulule/dekiteru/services/postgres"
	"github.com/ulule/dekiteru/services/redis"
)

// Checker is a service checker.
type Checker func(parameters map[string]interface{}) (int, error)

// Services are built-in service checkers.
var Services = map[string]Checker{
	"postgres":      postgres.Run,
	"redis":         redis.Run,
	"elasticsearch": elasticsearch.Run,
}
