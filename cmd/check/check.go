package check

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/ulule/dekiteru/checker"
)

const (
	defaultInterval = 1
	defaultRetry    = 10
)

// Command is the check command.
var Command = cli.Command{
	Name:   "check",
	Usage:  "add a task to the list",
	Action: run,
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

// run is the check command action.
func run(ctx *cli.Context) error {
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

	err := checker.Run(service, interval, retry, params)
	if err != nil {
		return cli.NewExitError(err, 10)
	}

	return err
}
