package cmd

import (
	"os"

	"github.com/urfave/cli"

	"github.com/ulule/dekiteru/cmd/check"
)

// Run runs the CLI application.
func Run() error {
	a := cli.NewApp()
	a.Name = "Dekiteru"
	a.Version = "0.1.0"
	a.Usage = "Check if a service is ready to use"
	a.Flags = []cli.Flag{}
	a.Commands = []cli.Command{check.Command}
	return a.Run(os.Args)
}
