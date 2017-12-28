package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/ulule/dekiteru/checker"
)

const (
	name    = "Dekiteru"
	version = "0.1.0"
	usage   = "Check if a service is ready to use"
)

// Run runs the CLI application.
func Run() error {
	a := cli.NewApp()
	a.Name = name
	a.Version = version
	a.Usage = usage
	a.Flags = []cli.Flag{}
	a.Commands = []cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "Check if a service is ready to use",
			Action: func(ctx *cli.Context) error {
				if ctx.String("service") == "" {
					fmt.Println("Error: --service parameter is missing")
					err := cli.ShowSubcommandHelp(ctx)
					if err != nil {
						return err
					}
					return cli.NewExitError("", 1)
				}
				parameters := map[string]interface{}{}
				for _, value := range ctx.StringSlice("parameter") {
					splits := strings.Split(value, "=")
					key := splits[0]
					val := strings.Join(splits[1:], "=")
					parameters[key] = val
				}

				err := checker.Run(ctx.String("service"), ctx.Int("interval"), ctx.Int("retry"), parameters)
				if err != nil {
					return cli.NewExitError(err, 10)
				}
				return err
			},
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
					Name:  "parameter, p",
					Usage: "Parameter to send to service checker (repeatable)",
				},
			},
		},
	}
	return a.Run(os.Args)
}
