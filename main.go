package main

import (
	"os"

	cli "github.com/the-code-genin/ges-cli/cli"
)

func main() {
	app := cli.CreateNewApp()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
