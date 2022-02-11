package main

import (
	"os"
	"runtime"

	"github.com/heart-dance/seed/src/app"
	"github.com/urfave/cli/v2"
)

var RootCmd = cli.NewApp()

var (
	version  = "0.0.1"
	hostFlag = &cli.StringFlag{
		Name:     "host",
		Required: true,
		Usage:    "listen host",
	}
	webFlag = &cli.StringFlag{
		Name:  "web",
		Usage: "web assets path",
	}
	profileFlag = &cli.StringFlag{
		Name:     "profile",
		Required: true,
		Usage:    "profile path",
	}
)

func init() {
	RootCmd.Copyright = "Copyright © 2020 gsxhnd"
	RootCmd.Usage = "A Download Tool"
	RootCmd.Commands = nil
	RootCmd.HideHelpCommand = true
	RootCmd.Version = version + " " + runtime.GOARCH + "/" + runtime.GOOS
	RootCmd.Flags = []cli.Flag{
		hostFlag,
		profileFlag,
		webFlag,
	}
	RootCmd.Action = func(c *cli.Context) error {
		var a, err = app.NewApplication(c.String("host"), c.String("profile"), c.String("web"))
		if err != nil {
			return err
		}
		return a.Run()
	}
}

func main() {
	err := RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
