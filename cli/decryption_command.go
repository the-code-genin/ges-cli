package cli

import "github.com/urfave/cli/v2"

var (
	decryptionCommand = &cli.Command{
		Name: "decrypt",
		Usage: "Decrypt data.",
		Action: decryptionAction,
	}
)

func decryptionAction(ctx *cli.Context) error {
	return nil
}