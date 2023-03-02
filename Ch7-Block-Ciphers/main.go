/*
Block Ciphers
A block cipher is a deterministic algorithm that operates on fixed-length groups of data, called blocks. Like stream ciphers, block ciphers are a kind of symmetric encryption.

Block ciphers are widely used in real-world applications to encrypt large amounts of data.

Assignment
The Go standard library has built-in support for the AES and DES block ciphers, which we will talk about in more detail later.

We've been asked by leadership to check on the block sizes of each algorithm and report back. Complete the getBlockSize function.

getBlockSize(keyLen, cipherType int) (int, error)
This function accepts a keyLen and cipherType and returns the block size of the cipher along with any errors encountered.

The cipherType is an enum of typeAES or typeDES. Depending on the cipher type, create a new cipher using:

aes.NewCipher or
des.NewCipher
The value of the key passed in doesn't matter here, all that matters is its length.

Return the .BlockSize() of the new cipher, and any error values without changing them if they are encountered. Return an "invalid cipher type" error if the cipherType isn't one of the valid values.

Notes
Notice the relationship (or lack thereof) between the key length and the block size.
It's expected that some of the test cases will result in error messages
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
)

const (
	typeAES = iota
	typeDES
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	var block cipher.Block

	switch cipherType {
	case typeAES:
		var err error
		block, err = aes.NewCipher(make([]byte, keyLen))
		if err != nil {
			return 0, err
		}
	case typeDES:
		var err error
		block, err = des.NewCipher(make([]byte, keyLen))
		if err != nil {
			return 0, err
		}
	default:
		return 0, errors.New("invalid cipher type")
	}

	return block.BlockSize(), nil
}

// don't touch below this line

func test(keyLen, cipherType int) {
	fmt.Printf(
		"Getting block size of %v cipher with key length %v...\n",
		getCipherTypeName(cipherType),
		keyLen,
	)
	blockSize, err := getBlockSize(keyLen, cipherType)
	if err != nil {
		fmt.Println(err)
		fmt.Println("========")
		return
	}
	fmt.Println("Block size:", blockSize)
	fmt.Println("========")
}

func getCipherTypeName(cipherType int) string {
	switch cipherType {
	case typeAES:
		return "AES"
	case typeDES:
		return "DES"
	}
	return "unknown"
}

func main() {
	test(16, typeAES)
	test(24, typeAES)
	test(32, typeAES)
	test(64, typeAES)

	test(8, typeDES)
	test(16, typeDES)
	test(24, typeDES)
	test(1, 124)
}

/*

Getting block size of AES cipher with key length 16...

Block size: 16

========

Getting block size of AES cipher with key length 24...

Block size: 16

========

Getting block size of AES cipher with key length 32...

Block size: 16

========

Getting block size of AES cipher with key length 64...

crypto/aes: invalid key size 64

========

Getting block size of DES cipher with key length 8...

Block size: 8

========

Getting block size of DES cipher with key length 16...

crypto/des: invalid key size 16

========

Getting block size of DES cipher with key length 24...

crypto/des: invalid key size 24

========

Getting block size of unknown cipher with key length 1...

invalid cipher type

========
*/
