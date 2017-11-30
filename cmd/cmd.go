package cmd

import (
	"os"

	"github.com/urfave/cli"

	"github.com/ulule/dekiteru/cmd/check"
)

const (
	name        = "Dekiteru"
	version     = "0.1.0"
	description = "Check if a service is ready to use"
)

// Run runs the CLI application.
func Run() error {
	a := cli.NewApp()
	a.Name = name
	a.Version = version
	a.Usage = description
	a.Flags = []cli.Flag{}
	a.Commands = []cli.Command{check.Command}
	return a.Run(os.Args)
}
