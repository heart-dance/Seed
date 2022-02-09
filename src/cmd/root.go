package cmd

import (
	"github.com/heart-dance/seed/src"
	"github.com/urfave/cli/v2"
)

var RootCmd = cli.NewApp()

var (
	portFlag = &cli.StringFlag{
		Name:  "port",
		Value: "123",
		Usage: "listen port",
	}
	webFlag = &cli.StringFlag{
		Name:  "web",
		Usage: "web assets path",
	}
	profileFlag = &cli.StringFlag{
		Name:  "profile",
		Usage: "profile path",
	}
)

func init() {
	RootCmd.Copyright = "Copyright © 2020 gsxhnd"
	RootCmd.Usage = "A Download Tool"
	RootCmd.Commands = []*cli.Command{
		versionCmd,
	}
	RootCmd.Flags = []cli.Flag{
		portFlag,
		webFlag,
		profileFlag,
	}
	RootCmd.Action = func(c *cli.Context) error {
		var app = src.NewApplication()
		app.Run()
		return nil
	}
}
