package cli

import "github.com/urfave/cli/v2"

var (
	decryptionCommand = &cli.Command{
		Name: "decrypt",
		Usage: "Run decryption algorithm.",
		Action: decryptionAction,
	}
)

func decryptionAction(ctx *cli.Context) error {
	return nil
}