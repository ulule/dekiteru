package cmd

import (
	"os"

	"github.com/urfave/cli"

	"github.com/ulule/dekiteru/cmd/check"
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
		check.Command,
	}

	return &Cmd{App: a}
}

// Run runs the command.
func (c Cmd) Run() error {
	return c.App.Run(os.Args)
}
