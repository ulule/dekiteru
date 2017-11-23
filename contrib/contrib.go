package contrib

type Service interface {
	Run() error
}

func GetServices() map[string]func() error {
	return map[string]func() error{
		"postgresql":    CheckPostgresql,
		"redis":         CheckRedis,
		"elasticsearch": CheckElasticSearch,
	}
}
