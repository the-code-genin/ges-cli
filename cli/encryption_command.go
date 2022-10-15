package cli

import "github.com/urfave/cli/v2"

var (
	encryptionFlags = []cli.Flag{
		
	}

	encryptionCommand = &cli.Command{
		Name: "encrypt",
		Usage: "Encrypt data.",
		Action: encryptionAction,
	}
)

func encryptionAction(ctx *cli.Context) error {
	return nil
}