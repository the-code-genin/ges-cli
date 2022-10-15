package cli

import "github.com/urfave/cli/v2"

var (
	keySizeFlag = &cli.Uint64Flag{
		Name:       "size",
		Aliases:    []string{"s"},
		Usage:      "The key size in bits. Must be a multiple of 32 bits.",
		Value:      32,
		Required:   true,
		HasBeenSet: true,
	}

	formatFlag = &cli.StringFlag{
		Name:       "format",
		Aliases:    []string{"f"},
		Usage:      "The output format. Available options are \"hex\" and \"base64\".",
		Value:      "hex",
		Required:   true,
		HasBeenSet: true,
	}

	outputFlag = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
	}
)
