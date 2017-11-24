package services

type Service interface {
	Run(parameters map[string]interface{}) (int, error)
	Name() string
}

var Services = map[string]Service{
	"postgresql":    PostgresqlService{},
	"redis":         RedisService{},
	"elasticsearch": ElasticSearchService{},
}

func GetService(name string) Service {
	s := Services[name]
	if s == nil {
		return nil
	}
	return s
}
