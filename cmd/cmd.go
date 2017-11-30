package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ulule/dekiteru/checker"
	"github.com/urfave/cli"
)

// Cmd is the CLI application.
type Cmd struct {
	App *cli.App
}

// New returns a new Cmd instance.
func New() *Cmd {
	a := cli.NewApp()
	a.Name = "Dekiteru"
	a.Version = "0.1.0"
	a.Usage = "Check if a service is ready to use"
	a.Flags = []cli.Flag{}
	a.Commands = []cli.Command{
		{
			Name:  "check",
			Usage: "add a task to the list",
			Action: func(ctx *cli.Context) error {
				if ctx.String("service") == "" {
					fmt.Println("Error: --service parameter is missing")
					cli.ShowAppHelp(ctx)
					return cli.NewExitError("", 1)
				}

				parameters := map[string]interface{}{}
				for _, value := range ctx.StringSlice("parameters") {
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
					Name:  "parameters, p",
					Usage: "Parameter to send to service checker",
				},
			},
		},
	}

	return &Cmd{App: a}
}

func (c Cmd) Run() {
	c.App.Run(os.Args)
}
