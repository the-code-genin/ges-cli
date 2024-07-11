package main

import (
	"fmt"
	"os"

	"github.com/the-code-genin/ges-cli/internal"
	"github.com/urfave/cli/v2"
)

var (
	keygenCommand = &cli.Command{
		Name:  "keygen",
		Usage: "Generate a new encryption key",
		Flags: []cli.Flag{
			outputFormatFlag,
			outputFileFlag,
		},
		Action: keygenAction,
	}
)

func keygenAction(ctx *cli.Context) error {
	key, err := internal.RandomBytes(8)
	if err != nil {
		return err
	}

	outputFilePath := ctx.String("output.file")
	encodingFormat := ctx.String("output.format")

	switch encodingFormat {
	case internal.EncodingFormatHex.String(), internal.EncodingFormatBase64.String():
		encodedKey, err := internal.EncodeBytes(internal.EncodingFormat(encodingFormat), key)
		if err != nil {
			return err
		}

		// Print encoded key to standard output if file isn't specified
		if outputFilePath == "" {
			fmt.Println(encodedKey)
			break
		}

		file, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}

		_, err = file.WriteString(encodedKey)
		if err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}

	default:
		if outputFilePath == "" {
			return internal.ErrRequiredOutputFilePath
		}

		file, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}

		_, err = file.Write(key)
		if err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}
	}

	return nil
}
