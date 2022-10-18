package main

import (
	"fmt"

	"github.com/the-code-genin/ges-cli/core"
	"github.com/urfave/cli/v2"
)

var (
	encryptionFlags = []cli.Flag{
		keyFileFlag,
		keyFormatFlag,
		inputFormatFlag,
		outputFormatFlag,
		outputFileFlag,
	}

	encryptionCommand = &cli.Command{
		Name: "encrypt",
		Usage: "Encrypt data",
		Action: encryptionAction,
		Flags: encryptionFlags,
	}
)

func encryptionAction(ctx *cli.Context) error {
	// Open the plaintext file
	args := ctx.Args()
	if args.Len() != 1 {
		return fmt.Errorf("expected the path of the file to be encrypted as the argument")
	}

	plainFilePath := args.First()
	plainFile, err := core.OpenFile(plainFilePath)
	if err != nil {
		return err
	}

	// Read the key
	keyFilePath := ctx.String("key.file")
	keyFile, err := core.OpenFile(keyFilePath)
	if err != nil {
		return err
	}

	keyFileLen, err := core.LengthOfFile(keyFile)
	if err != nil {
		return err
	}

	key, err := core.ReadFile(keyFile, 0, int(keyFileLen))
	if err != nil {
		return err
	}

	keyFormat := ctx.String("key.format")
	if keyFormat != "binary" {
		key, err = core.DecodeBytes(string(key), keyFormat)
		if err != nil {
			return err
		}
	}

	// Encrypt the plain text
	keySize := uint64(len(key) * 8)
	if keySize != 64 {
		return fmt.Errorf("key sizes must be 64 bits")
	}

	plainFileLen, err := core.LengthOfFile(plainFile)
	if err != nil {
		return err
	}

	plainText, err := core.ReadFile(plainFile, 0, int(plainFileLen))
	if err != nil {
		return err
	}

	inputFormat := ctx.String("input.format")
	if inputFormat != "binary" {
		plainText, err = core.DecodeBytes(string(plainText), inputFormat)
		if err != nil {
			return err
		}
	}

	cipher, err := core.NewGESCipher()
	if err != nil {
		return err
	}

	cipherText, err := cipher.Encrypt(plainText, key)
	if err != nil {
		return err
	}

	// Record the output
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

		err = core.WriteToFile(file, 0, cipherText)
		if err != nil {
			return err
		}

		err = file.Sync()
		if err != nil {
			return err
		}
	} else {
		encodedCipherText, err := core.EncodeBytes(cipherText, encodingFormat)
		if err != nil {
			return err
		}

		if outputFilePath == "" {
			fmt.Print(encodedCipherText)
		} else {
			file, err := core.OpenFile(outputFilePath)
			if err != nil {
				return err
			}
	
			err = core.WriteToFile(file, 0, []byte(encodedCipherText))
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