package main

import (
	"os"

	cli "github.com/urfave/cli/v2"
)

var (
	app *cli.App
)

func init() {
	app = &cli.App{
		Name:  "ges",
		Usage: "A simple encryption algorithm to securly communicate over the internet and obfuscate your data.",
		Commands: []*cli.Command{
			keygenCommand,
			encryptionCommand,
			decryptionCommand,
		},
		Version: "1.0.2",
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
