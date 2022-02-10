package main

import (
	"fmt"
	"os"
	"runtime"

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
	RootCmd.Before = func(c *cli.Context) error {
		return nil
	}
	RootCmd.Action = func(c *cli.Context) error {
		fmt.Println(c.String("host"))
		fmt.Println(c.String("profile"))
		fmt.Println(c.String("web"))
		// var app = src.NewApplication()
		// app.Run()
		return nil
	}
}

func main() {
	err := RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
