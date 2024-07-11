package main

import (
	"io"
	"os"

	"github.com/the-code-genin/ges-cli/core"
	"github.com/the-code-genin/ges-cli/internal"
	"github.com/urfave/cli/v2"
)

var (
	encryptionCommand = &cli.Command{
		Name:   "encrypt",
		Usage:  "Encrypt data",
		Action: encryptionAction,
		Flags: []cli.Flag{
			keyFileFlag,
			keyFormatFlag,
			inputFileFlag,
			outputFileFlag,
		},
	}
)

func encryptionAction(ctx *cli.Context) error {
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

	// Get input stream for encryption
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

	// Get the output stream for encryption
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

	// Use ECB method to incrementally read and encrypt data blocks
	plainBlock := make([]byte, 16)
	for {
		_, readErr := inputStream.Read(plainBlock)
		if readErr != nil && readErr != io.EOF {
			return readErr
		}

		cipherBlock, err := core.Encrypt(plainBlock, key)
		if err != nil {
			return err
		}

		_, err = outputStream.Write(cipherBlock)
		if err != nil {
			return err
		}

		if readErr == io.EOF {
			break
		}
	}

	return nil
}
