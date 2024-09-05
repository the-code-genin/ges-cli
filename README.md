# GES CLI

The software utilizes a simple block cipher encryption algorithm to encrypt and decrypt data. It uses a home-grown encryption algorithm called **GES** based off of the feistel cipher structure.

Developed and maintained by [Mohammed Adekunle](https://mohammedadekunle.com.ng)

## GES Algorithm

Based off of the feistel cipher structure, the block size of the algorithm is 128 bits with a 128 bit key. It utilizes a parity drop for every 8th-bit to reduce the key size to 112 bits, followed by a shift function which condenses the key to 96 bits. Then DES S-Box substitution is done to generate unique 64-bit round keys.

## Installation

You can compile the binary manually using the `go build` command or you can compile via the build system using the `make` tool.

### Via Go Install
To install locally via `go install`. Run the command below:

```shell
go install github.com/the-code-genin/ges-cli/cmd/ges@latest
ges --version
```

### Via Go Build

To compile via `go build`, you will require at least `go 1.22` installed. Then follow the steps below:

- Run `cd cmd/ges` to change your terminal path
- Run `go build .` to generate a build of the `ges cli` for your operating system architecture
- Run `./ges` to get started

### Via make

To compile via `make`, you will require a recent version of the `make` utility installed. Then follow the steps below:

- Run `make build` to trigger an automated build of the project
- Run `cd build/bin`to change your terminal path to the build directory
- Run `./ges` to get started

## Usage

Before you can get started with encryption or decryption you need to generate an cipher key. Cipher keys for the `GES` algorithm need to be 128-bits large. An utiliy is provided for this in the `ges` binary called `keygen`.

To encrypt data, an `encrypt` utility is provided in the `ges` binary which can either encrypt a provided input file or encrypt data directly from the standard input.

Decryption is similar with with the `decrypt` utility in the `ges` binary which can also decrypt a provided input file or decrypt directly from the standard input.

***Special Notes:***

- Like any other cipher, attempting to decrypt a cipher text with the wrong cipher key will not give the expected plain text. But neverthelesss, a plain text will be generated.
- The GES algorithm **should not** be used in a production environment as it does not meet any international standards for encryption (It is a proof of concept).

## License

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
