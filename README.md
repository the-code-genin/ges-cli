# GES CLI

The software utilizes a simple block cipher encryption algorithm to encrypt and decrypt data. It uses an **weak** home-grown encryption algorithm called **GES** based off of the feistel cipher structure.

Developed and maintained by [Mohammed Adekunle](https://mohammedadekunle.com.ng)

## GES Algorithm

Based off of the feistel cipher structure, the block size of the algorithm is 128 bits with a 64 bit key size. It utilizes an initial NXOR operation between the block and the key to scramble the data rather than utilizing a P-box. A XOR operation is utilized for the round function between the right-half block and the round key.

## Installation

You can compile the binary manually using the `go build` command or you can compile via the build system using the `make` tool.

### Via Go Build

To compile via `go build`, you will require a recent version of the `go` binary installed. Then follow the steps below:

- Run `cd cmd/ges` to change your terminal path
- Run `go build .` to generate a build of the `ges cli` for your operating system architecture
- Run `./ges` to get started

### Via make

To compile via `make`, you will require a recent version of the `make` utility installed. Then follow the steps below:

- Run `make all` to trigger an automated build of the project
- Run `cd build/bin`to change your terminal path to the build directory
- Run `./ges` to get started

## Usage

Before you can get started with encryption or decryption you need to generate an cipher key. Cipher keys for the `GES` algorithm need to be 64 bits large. An utiliy is provided for this in the `ges` binary called `keygen`. To generate a cipher key, simply run `ges keygen --output.file <file>` with `<file>` being the path of the output file for the key relative to the terminal path.

Next, we can try a simple encryption process by following this step:

- Run `./ges encrypt --key.file <key_file> --output.file <output_file> <plaintext_file>` with `plaintext_file` referring to the plain text file you wish to encrypt, `<key_file>` being your cipher key, and `<output_file>` being your cipher text output file. *Note:* The software works the files contents, it does not utilize the file name.

Decryption is similar with the following process:

- Run `./ges decrypt --key.file <key_file> --output.file <output_file> <ciphertext_file>` with `ciphertext_file` referring to the cipher text file you wish to decrypt, `<key_file>` being your cipher key, and `<output_file>` being your plain text output file. *Note:* The software works the files contents, it does not utilize the file name.

***Note:***

- Like an other cipher, attempting to decrypt a cipher text with the wrong cipher key will not give a proper plain text. But neverthelesss, a plain text will be generated.
- The GES algorithm **should not** be used in a production environment as it does not meet any international standards for encryption(It's just a fun project). 

## License

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
