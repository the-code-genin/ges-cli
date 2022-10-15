package cli

import (
	"fmt"

	"github.com/the-code-genin/ges-cli/core"
	"github.com/urfave/cli/v2"
)

var (
	decryptionFlags = []cli.Flag{
		keyFileFlag,
		keyFormatFlag,
		inputFormatFlag,
		outputFormatFlag,
		outputFileFlag,
	}

	decryptionCommand = &cli.Command{
		Name: "decrypt",
		Usage: "Decrypt data. The block size is always double the key size.",
		Action: decryptionAction,
		Flags: decryptionFlags,
	}
)

func decryptionAction(ctx *cli.Context) error {
	// Open the cipher text file
	args := ctx.Args()
	if args.Len() != 1 {
		return fmt.Errorf("expected the path of the file to be encrypted as the argument")
	}

	cipherTextFilePath := args.First()
	cipherTextFile, err := core.OpenFile(cipherTextFilePath)
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

	// Decrypt the cipher text
	keySize := uint64(len(key) * 8)
	if keySize != 64 {
		return fmt.Errorf("key sizes must be 64 bits")
	}

	cipherTextFileLen, err := core.LengthOfFile(cipherTextFile)
	if err != nil {
		return err
	}

	cipherText, err := core.ReadFile(cipherTextFile, 0, int(cipherTextFileLen))
	if err != nil {
		return err
	}

	inputFormat := ctx.String("input.format")
	if inputFormat != "binary" {
		cipherText, err = core.DecodeBytes(string(cipherText), inputFormat)
		if err != nil {
			return err
		}
	}

	cipher, err := core.NewGESCipher()
	if err != nil {
		return err
	}

	plainText, err := cipher.Decrypt(cipherText, key)
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

		err = core.WriteToFile(file, 0, plainText)
		if err != nil {
			return err
		}

		err = file.Sync()
		if err != nil {
			return err
		}
	} else {
		encodedPlainText, err := core.EncodeBytes(plainText, encodingFormat)
		if err != nil {
			return err
		}

		if outputFilePath == "" {
			fmt.Print(encodedPlainText)
		} else {
			file, err := core.OpenFile(outputFilePath)
			if err != nil {
				return err
			}
	
			err = core.WriteToFile(file, 0, []byte(encodedPlainText))
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