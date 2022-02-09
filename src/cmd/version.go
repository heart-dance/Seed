package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name: "version",
	Action: func(c *cli.Context) error {
		fmt.Println("Seed Version: " + runtime.GOARCH + "/" + runtime.GOOS)
		return nil
	},
}
