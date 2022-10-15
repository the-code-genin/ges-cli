package cli

import (
	"fmt"

	"github.com/the-code-genin/ges-cli/core"
	"github.com/urfave/cli/v2"
)

var (
	keygenFlags  = []cli.Flag{
		&cli.Uint64Flag{
			Name: "size",
			Usage: "The key size in bits. Must be a multiple of 32 bits.",
			Value: 32,
			Required: true,
			HasBeenSet: true,
		},
		&cli.StringFlag{
			Name: "format",
			Usage: "The output format. Available options are \"hex\" and \"base64\".",
			Value: "hex",
			Required: true,
			HasBeenSet: true,
		},
	}

	keygenCommand = &cli.Command{
		Name: "keygen",
		Usage: "Generate a new encryption key.",
		Flags: keygenFlags,
		Action: keygenAction,
	}
)

func keygenAction(ctx *cli.Context) error {
	keySize := ctx.Uint64("size")
	if keySize % 32 != 0 {
		return fmt.Errorf("key sizes must be a multiple of 32 bits")
	}

	key, err := core.RandomBytes(keySize / 8)
	if err != nil {
		return err
	}

	encodingFormat := ctx.String("format")
	encodedKey, err := core.EncodeBytes(key, encodingFormat)
	if err != nil {
		return err
	}

	fmt.Printf(encodedKey)
	return nil
}
