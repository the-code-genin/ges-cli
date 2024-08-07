package main

import "github.com/urfave/cli/v2"

var (
	keyFileFlag = &cli.StringFlag{
		Name:     "key.file",
		Usage:    "The key file path",
		Required: true,
	}

	keyFormatFlag = &cli.StringFlag{
		Name:       "key.format",
		Usage:      "The key file format. Available options are \"binary\", \"hex\" and \"base64\"",
		Value:      "binary",
		Required:   true,
		HasBeenSet: true,
	}

	inputFileFlag = &cli.StringFlag{
		Name:  "input.file",
		Usage: "The input file path",
	}

	outputFormatFlag = &cli.StringFlag{
		Name:       "output.format",
		Usage:      "The output format. Available options are \"binary\", \"hex\" and \"base64\"",
		Value:      "binary",
		Required:   true,
		HasBeenSet: true,
	}

	outputFileFlag = &cli.StringFlag{
		Name:  "output.file",
		Usage: "The output file path. Required for \"binary\" output format",
	}
)
