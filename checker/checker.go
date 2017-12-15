package checker

import (
	"errors"
	"time"

	"github.com/ulule/dekiteru/services"
)

// Run runs the given service
func Run(service string, interval int, retries int, parameters map[string]interface{}) error {
	var (
		delta time.Duration
		start time.Time
		err   error
		code  int
	)
	s, ok := services.Services[service]
	if !ok {
		return errors.New("this service does not exist")
	}
	for t := 1; t <= retries; t++ {
		start = time.Now()

		code, err = s.Run(parameters)
		switch code {
		// Everything is ok
		case 0:
			return err
		// Soft error
		case 1:
			return err
		// Hard error
		case 2:
		}

		if t+1 > retries {
			return err
		}

		delta = time.Now().Sub(start)
		if delta < time.Duration(interval)*time.Second {
			time.Sleep(time.Duration(interval)*time.Second - delta)
		}
	}
	return err
}