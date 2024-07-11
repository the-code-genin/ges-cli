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
		outputStream, err = os.OpenFile(outputFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
	}

	defer outputStream.Close()

	// Use ECB method to incrementally read and decrypt data blocks
	var cipherBlock []byte
	for {
		cipherBlock = make([]byte, 16)
		readBytes, readErr := inputStream.Read(cipherBlock)
		if readErr != nil && readErr != io.EOF {
			return readErr
		}

		if readBytes == 0 || readErr == io.EOF {
			break
		}

		plainBlock, err := core.Decrypt(cipherBlock, key)
		if err != nil {
			return err
		}

		// If last block contains null byte
		// strip the plain block till the padded bit
		if plainBlock[readBytes-1] == byte(0) {
			for i := len(plainBlock) - 2; i >= 0; i-- {
				if plainBlock[i] == byte(1)<<7 {
					plainBlock = plainBlock[:i]
					break
				}
			}
		}

		if _, err = outputStream.Write(plainBlock); err != nil {
			return err
		}
	}

	if err := outputStream.Sync(); err != nil {
		return err
	}

	return nil
}
