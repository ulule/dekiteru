package services

import (
	"github.com/ulule/dekiteru/services/elasticsearch"
	"github.com/ulule/dekiteru/services/postgres"
	"github.com/ulule/dekiteru/services/redis"
)

// ServiceFunc is the service check function.
type ServiceFunc func(parameters map[string]interface{}) (int, error)

// Services are built-in service checkers.
var Services = map[string]ServiceFunc{
	"postgres":      postgres.Run,
	"redis":         redis.Run,
	"elasticsearch": elasticsearch.Run,
}
