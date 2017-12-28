package services

import (
	"errors"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

// RabbitMQ service
type RabbitMQ struct{}

// Run implements Service interface.
func (RabbitMQ) Run(parameters map[string]interface{}) error {
	var (
		ok            bool
		uri           string
		timeoutString string
		timeout       int
		err           error
	)

	uri, ok = parameters["uri"].(string)
	if !ok || uri == "" {
		uri = "amqp://guest:guest@127.0.0.1/"
	}

	timeoutString, ok = parameters["timeout"].(string)
	if ok {
		timeout, err = strconv.Atoi(timeoutString)
		if err != nil {
			return &HardError{errors.New("invalid `timeout` parameter")}

		}
		if timeout < 1 {
			return &HardError{errors.New("invalid `timeout` parameter (cannot be below 1 second)")}
		}

	} else {
		timeout = 2
	}

	log.Printf(`uri    : "%s"`, uri)
	log.Printf(`timeout: "%d"`, timeout)

	_, err = amqp.DialConfig(uri, amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Duration(timeout)*time.Second)
		},
	})
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
		"timeout",
	}
}
