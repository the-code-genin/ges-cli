package cli

import (
	cli "github.com/urfave/cli/v2"
)

func CreateNewApp() *cli.App {
	app := &cli.App{
		Name:  "ges-cli",
		Usage: "A simple encryption algorithm to securly communicate over the internet and obfuscate your data.",
		Commands: []*cli.Command{
			keygenCommand,
			encryptionCommand,
			decryptionCommand,
		},
	}

	return app
}
