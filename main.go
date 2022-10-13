package main

import (
	"os"

	cli "github.com/the-code-genin/cryptware/cli"
)

func main() {
	app := cli.CreateNewApp()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
