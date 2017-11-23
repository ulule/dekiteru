package checker

import (
	"errors"
	"time"

	"github.com/ulule/dekiteru/contrib"
)

type Checker struct {
	Services map[string]func() error
}

func New() *Checker {
	return &Checker{Services: contrib.GetServices()}
}

func (a Checker) Run(service string, interval int, retries int) error {
	var (
		delta time.Duration
		start time.Time
		err   error
	)
	s := contrib.GetServices()[service]
	if s == nil {
		return errors.New("This service does not exist.")
	}
	for t := 1; t <= retries; t++ {
		start = time.Now()

		err = s()
		if err == nil {
			return err
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
