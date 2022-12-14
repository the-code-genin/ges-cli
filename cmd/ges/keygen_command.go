package main

import (
	"fmt"

	"github.com/the-code-genin/ges-cli/core"
	"github.com/urfave/cli/v2"
)

var (
	keygenFlags = []cli.Flag{
		outputFormatFlag,
		outputFileFlag,
	}

	keygenCommand = &cli.Command{
		Name:   "keygen",
		Usage:  "Generate a new encryption key",
		Flags:  keygenFlags,
		Action: keygenAction,
	}
)

func keygenAction(ctx *cli.Context) error {
	key, err := core.RandomBytes(64 / 8)
	if err != nil {
		return err
	}

	outputFilePath := ctx.String("output.file")
	encodingFormat := ctx.String("output.format")
	if encodingFormat == "binary" {
		if outputFilePath == "" {
			return fmt.Errorf("output file path is required for binary encoding")
		}

		file, err := core.OpenFile(outputFilePath)
		if err != nil {
			return err
		}

		err = core.WriteToFile(file, 0, key)
		if err != nil {
			return err
		}

		err = file.Sync()
		if err != nil {
			return err
		}
	} else {
		encodedKey, err := core.EncodeBytes(key, encodingFormat)
		if err != nil {
			return err
		}

		if outputFilePath == "" {
			fmt.Print(encodedKey)
		} else {
			file, err := core.OpenFile(outputFilePath)
			if err != nil {
				return err
			}
	
			err = core.WriteToFile(file, 0, []byte(encodedKey))
			if err != nil {
				return err
			}
	
			err = file.Sync()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
