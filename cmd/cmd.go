package cmd

import (
	"fmt"
	"os"

	"github.com/ulule/dekiteru/checker"
	"github.com/urfave/cli"
)

type Cmd struct {
	App     *cli.App
	Checker *checker.Checker
}

func New() *Cmd {
	c := checker.New()

	a := cli.NewApp()
	a.Name = "Dekiteru"
	a.Version = "0.1.0"
	a.Usage = "Check if a service is ready to use"
	a.Flags = []cli.Flag{
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
	}
	a.Action = func(ctx *cli.Context) error {
		if ctx.String("service") == "" {
			fmt.Println("Error: --service parameter is missing")
			cli.ShowAppHelp(ctx)
			return cli.NewExitError("", 1)
		}
		err := c.Run(ctx.String("service"), ctx.Int("interval"), ctx.Int("retry"))
		if err != nil {
			return cli.NewExitError(err, 10)
		}
		return err
	}
	return &Cmd{App: a, Checker: c}
}

func (c Cmd) Run() {
	c.App.Run(os.Args)
}
