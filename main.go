package main

import (
	"os"

	"github.com/heart-dance/seed/src/cmd"
)

func main() {
	err := cmd.RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
