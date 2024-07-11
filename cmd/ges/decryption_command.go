package main

import (
	"io"
	"os"

	"github.com/the-code-genin/ges-cli/core"
	"github.com/the-code-genin/ges-cli/internal"
	"github.com/urfave/cli/v2"
)

var (
	decryptionCommand = &cli.Command{
		Name:   "decrypt",
		Usage:  "Decrypt data",
		Action: decryptionAction,
		Flags: []cli.Flag{
			keyFileFlag,
			keyFormatFlag,
			inputFileFlag,
			outputFileFlag,
		},
	}
)

func decryptionAction(ctx *cli.Context) error {
	// Read the key file
	key, err := os.ReadFile(ctx.String(keyFileFlag.Name))
	if err != nil {
		return err
	}

	switch keyFormat := ctx.String(keyFormatFlag.Name); keyFormat {
	case internal.EncodingFormatHex.String(), internal.EncodingFormatBase64.String():
		key, err = internal.DecodeBytes(internal.EncodingFormat(keyFormat), string(key))
		if err != nil {
			return err
		}
	}

	if len(key) != 16 {
		return internal.ErrInvalidKeyLength
	}

	// Get input stream for decryption
	var inputStream *os.File
	inputFilePath := ctx.String(inputFileFlag.Name)

	// Reading from standard input
	if inputFilePath == "" {
		inputStream = os.Stdin
	}

	// Reading from specified file
	if inputFilePath != "" {
		var err error
		inputStream, err = os.Open(inputFilePath)
		if err != nil {
			return err
		}
	}

	defer inputStream.Close()

	// Get the output stream for decryption
	var outputStream *os.File

	outputFilePath := ctx.String(outputFileFlag.Name)

	// Writing to standard output
	if outputFilePath == "" {
		outputStream = os.Stdout
	}

	// Writing to specified file
	if outputFilePath != "" {
		var err error
		outputStream, err = os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
	}

	defer outputStream.Close()

	// Use ECB method to incrementally read and decrypt data blocks
	cipherBlock := make([]byte, 16)
	for {
		_, readErr := inputStream.Read(cipherBlock)
		if readErr != nil && readErr != io.EOF {
			return readErr
		}

		plainBlock, err := core.Decrypt(cipherBlock, key)
		if err != nil {
			return err
		}

		_, err = outputStream.Write(plainBlock)
		if err != nil {
			return err
		}

		if readErr == io.EOF {
			break
		}
	}

	return nil
}
