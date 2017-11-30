package services

// Service is a service checker.
type Service func(parameters map[string]interface{}) (int, error)

// Services are built-in services.
var Services = map[string]Service{
	"postgresql":    Postgres,
	"redis":         Redis,
	"elasticsearch": Elasticsearch,
}
