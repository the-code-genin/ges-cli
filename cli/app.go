package cli

import (
	cli "github.com/urfave/cli/v2"
)

func CreateNewApp() *cli.App {
	return &cli.App{
		Name: "cryptware",
		Usage: "A suite of cryptography tools to securly communicate over the internet.",
	}
}