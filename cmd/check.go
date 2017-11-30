package cmd

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/ulule/dekiteru/checker"
)

var checkCommand = cli.Command{
	Name:   "check",
	Usage:  "add a task to the list",
	Action: checkAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "service, s",
			Usage: "Service name to check (required)",
		},
		cli.IntFlag{
			Name:  "interval, i",
			Value: 1,
			Usage: "Interval between retries in second",
		},
		cli.IntFlag{
			Name:  "retry, r",
			Value: 10,
			Usage: "Number of retry",
		},
		cli.StringSliceFlag{
			Name:  "parameters, p",
			Usage: "Parameter to send to service checker",
		},
	},
}

// check is the check action.
func checkAction(ctx *cli.Context) error {
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
