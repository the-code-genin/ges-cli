package cli

import (
	cli "github.com/urfave/cli/v2"
)

var (
	generalFlags = []cli.Flag{
		&cli.StringFlag{
			Name: "file",
			Aliases: []string{"f"},
		},
	}
)

func CreateNewApp() *cli.App {
	app := &cli.App{
		Name: "ges-cli",
		Usage: "A simple encryption algorithm to securly communicate over the internet and obfuscate your data to prying eyes.",
		Commands: []*cli.Command{
			encryptionCommand,
			decryptionCommand,
		},
	}

	app.Flags = append(app.Flags, generalFlags...)

	return app
}