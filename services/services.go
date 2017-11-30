package services

type Service interface {
	Run(parameters map[string]interface{}) (int, error)
	Name() string
}

var Services = map[string]Service{
	"postgresql":    Postgresql{},
	"redis":         Redis{},
	"elasticsearch": Elasticsearch{},
}

func Get(name string) Service {
	s := Services[name]
	if s == nil {
		return nil
	}
	return s
}
