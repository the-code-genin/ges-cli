package cli

import "github.com/urfave/cli/v2"

var (
	keySizeFlag = &cli.Uint64Flag{
		Name:       "key.size",
		Aliases:    []string{"s"},
		Usage:      "The key size in bits. Must be a multiple of 32 bits.",
		Value:      32,
		Required:   true,
		HasBeenSet: true,
	}

	keyFileFlag = &cli.StringFlag{
		Name:     "key.file",
		Usage:    "The key file path.",
		Required: true,
	}

	keyFormatFlag = &cli.StringFlag{
		Name:       "key.format",
		Usage:      "The key file format. Available options are \"binary\", \"hex\" and \"base64\".",
		Value:      "hex",
		Required:   true,
		HasBeenSet: true,
	}

	outputFormatFlag = &cli.StringFlag{
		Name:       "output.format",
		Aliases:    []string{"f"},
		Usage:      "The output format. Available options are \"binary\", \"hex\" and \"base64\".",
		Value:      "hex",
		Required:   true,
		HasBeenSet: true,
	}

	outputFileFlag = &cli.StringFlag{
		Name:    "output.file",
		Aliases: []string{"o"},
		Usage:   "The output file path. Required for \"binary\" output format",
	}
)
