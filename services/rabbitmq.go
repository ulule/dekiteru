package services

import (
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQ service
type RabbitMQ struct{}

// Run implements Service interface.
func (RabbitMQ) Run(parameters map[string]interface{}) error {
	var (
		ok  bool
		uri string
		err error
	)

	uri, ok = parameters["uri"].(string)
	if !ok || uri == "" {
		uri = "amqp://guest:guest@127.0.0.1/"
	}

	log.Printf(`uri: "%s"`, uri)

	_, err = amqp.Dial(uri)
	if err != nil {
		return &SoftError{err}
	}

	return nil
}

// Name implements Service interface.
func (RabbitMQ) Name() string {
	return "rabbitmq"
}

// Parameters implements Service interface.
func (RabbitMQ) Parameters() []string {
	return []string{
		"uri",
	}
}
