package api

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

type CmdApi struct {
}

func (this *CmdApi) RunCmd() error {
	app := cli.NewApp()
	app.Version = "18.01.v1"
	app.Commands = this.cmdConfig()
	return app.Run(os.Args)
}

func (this *CmdApi) cmdConfig() []cli.Command {
	return []cli.Command{
		{
			Name:    "assign",
			Aliases: []string{"a"},
			Usage:   "exec a assignment",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
	}
}
