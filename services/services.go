package services

// Service represents a service.
type Service interface {
	Run(parameters map[string]interface{}) error
	Name() string
	Parameters() []string
}

// Services are built-in services.
var Services = map[string]Service{
	"postgresql":    Postgresql{},
	"redis":         Redis{},
	"elasticsearch": ElasticSearch{},
}
