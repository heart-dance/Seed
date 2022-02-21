package main

import (
	"os"
	"runtime"

	"github.com/heart-dance/seed/app"
	"github.com/urfave/cli/v2"
)

var RootCmd = cli.NewApp()

var (
	version  = "0.0.1"
	hostFlag = &cli.StringFlag{
		Name:  "host",
		Usage: "listen host",
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
	releaseFlag = &cli.BoolFlag{
		Name:  "release",
		Value: true,
	}
	debugFlag = &cli.BoolFlag{
		Name:  "debug",
		Value: false,
	}
)

func init() {
	RootCmd.Copyright = "Copyright © 2020 HeartDance"
	RootCmd.Usage = "A Remote Download Tool"
	RootCmd.Commands = nil
	RootCmd.HideHelpCommand = true
	RootCmd.Version = version + " " + runtime.GOARCH + "/" + runtime.GOOS
	RootCmd.Flags = []cli.Flag{
		hostFlag,
		profileFlag,
		webFlag,
		debugFlag,
		releaseFlag,
	}
	RootCmd.Action = func(c *cli.Context) error {
		var runMode = "dev"
		if c.Bool("release") {
			if c.Bool("debug") {
				runMode = "debug"
			} else {
				runMode = "prod"
			}
		}

		var a, err = app.NewApplication(version, c.String("host"), c.String("profile"), c.String("web"), runMode)
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
