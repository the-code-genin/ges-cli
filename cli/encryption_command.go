package cli

import "github.com/urfave/cli/v2"

var (
	encryptionCommand = &cli.Command{
		Name: "encrypt",
		Usage: "Run encryption algorithm.",
		Action: encryptionAction,
	}
)

func encryptionAction(ctx *cli.Context) error {
	return nil
}