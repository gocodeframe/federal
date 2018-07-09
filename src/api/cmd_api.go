package api

import (
	"domain"
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
			Aliases: []string{},
			Usage:   "exec a assignment",
			Flags:   this.asignFlags(),
			Action: func(c *cli.Context) error {
				assignmentList := &domain.AssigmentList{}
				assignmentList.Load(c.String("file"))
				assignmentList.Run()
				return nil
			},
		},
	}
}

func (this *CmdApi) asignFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "assigments(yaml file)",
			Value: "default.yaml",
		}}
}
