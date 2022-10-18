# GES CLI

The software utilizes a simple block cipher encryption algorithm to encrypt and decrypt data. It uses an **weak** home-grown encryption algorithm called **GES** based off of the feistel cipher structure.

Developed and maintained by [Mohammed Adekunle](https://mohammedadekunle.com.ng)

## GES Algorithm

Based off of the feistel cipher structure, the block size of the algorithm is 128 bits with a 64 bit key size. It utilizes an initial NXOR operation between the block and the key to scramble the data rather than utilizing a P-box. A XOR operation is utilized for the round function between the right-half block and the round key.

## Installation

You can compile the binary manually using the `go` command or you can compile via the build system using the `make` utility.

### Via Go

### Via make

## Usage

## License

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
