package check

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli"

	"github.com/ulule/dekiteru/services"
)

const (
	defaultInterval = 1
	defaultRetry    = 10
)

// Command is the check command.
var Command = cli.Command{
	Name:   "check",
	Usage:  "add a task to the list",
	Action: action,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "service, s",
			Usage: "Service name to check (required)",
		},
		cli.IntFlag{
			Name:  "interval, i",
			Value: defaultInterval,
			Usage: "Interval between retries in second",
		},
		cli.IntFlag{
			Name:  "retry, r",
			Value: defaultRetry,
			Usage: "Number of retry",
		},
		cli.StringSliceFlag{
			Name:  "parameters, p",
			Usage: "Parameter to send to service checker",
		},
	},
}

// action is the check command action.
func action(ctx *cli.Context) error {
	var (
		service    = ctx.String("service")
		interval   = ctx.Int("interval")
		retry      = ctx.Int("retry")
		parameters = ctx.StringSlice("parameters")
	)

	if service == "" {
		fmt.Println("Error: --service parameter is missing")

		err := cli.ShowAppHelp(ctx)
		if err != nil {
			return err
		}

		return cli.NewExitError("", 1)
	}

	params := map[string]interface{}{}

	for _, value := range parameters {
		splits := strings.Split(value, "=")
		params[splits[0]] = strings.Join(splits[1:], "=")
	}

	err := check(service, interval, retry, params)
	if err != nil {
		return cli.NewExitError(err, 10)
	}

	return err
}

// check runs the given service.
func check(service string, interval int, retries int, parameters map[string]interface{}) error {
	var (
		delta time.Duration
		start time.Time
		err   error
		code  int
	)

	checkr, ok := services.Services[service]
	if !ok {
		return fmt.Errorf("%s service does not exist", service)
	}

	for t := 1; t <= retries; t++ {
		start = time.Now()

		code, err = checkr(parameters)
		if code > 1 && err == nil {
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
