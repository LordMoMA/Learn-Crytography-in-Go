/*
Data Encryption Standard
The Data Encryption Standard is an example of a symmetric-key block cipher that utilizes a Feistel network.

DES was developed in the early 1970s at IBM and in 1977 a slightly modified version was published as an official Federal Information Processing Standard for the United States.

In 1997 the DESCHALL Project was the first group to publicly break DES encryption and won $10,000 for their efforts! DES was broken by a simple brute force attack, which is possible due to its small key size of just 56 bits.

What's an iv?
An IV, or initialization vector, is a random value that is used to initialize a block cipher. It is used to ensure that the same plaintext always encrypts to a different ciphertext. Without an IV, the same plaintext would always encrypt to the same ciphertext which is a big security vulnerability.

Assignment
For Cryptanalysis purposes, Passly keeps a DES implementation around so that our engineers can practice breaking it. Complete the encrypt function. It should use the standard library's crypto/des package to encrypt the plaintext with the given key.

Complete the encrypt function and its helper padMsg function. The decrypt function is already written for you.

padMsg(plaintext []byte, blockSize int) []byte
The padWithZeros function is provided for you, but it only pads a single block. You'll need to find the last block in the message and pad that one. Essentially you need to ensure that the entire message length is a multiple of the block size.

encrypt(key, plaintext []byte) ([]byte, error)
We'll be using DES in CBC mode. Here's an example from the Go documentation that shows how to encrypt a message.

Create a new cipher block
Pad the plaintext with zeros using padMsg
Generate a random iv and append it to the beginning of the ciphertext. It should be the same length as the block size.
Create a new encrypter
Encrypt the blocks and return the entire ciphertext
Return any errors that occur.

Tip
Don't be afraid to debug your code by printing out the values of different variables. Just be sure to remove any debug code before submitting your answer.
*/

/*
slice1 := []int{1, 2, 3, 4, 5} 
slice2 := []int{1, 2 } 

copy(slice2, slice1)
fmt.Println(slice2)

the copy() function copies the elements from the source slice (slice1) into the destination slice (slice2) until either the destination slice is full or the source slice is exhausted. In this case, slice2 is only two elements long, so only the first two elements of slice1 are copied into slice2.
*/

// Solution 1
package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	paddedPlaintext := padMsg(plaintext, blockSize)

	ciphertext := make([]byte, blockSize+len(paddedPlaintext))

	// Generate IV
	iv := make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		iv[i] = byte(i)
	}
	copy(ciphertext[:blockSize], iv)

	// Encrypt using CBC mode
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], paddedPlaintext)

	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	padSize := blockSize - len(plaintext)%blockSize
	padding := make([]byte, padSize)
	return append(plaintext, padding...)
}

// Solution 2

import (
    "crypto/cipher"
    "crypto/des"
    "crypto/rand"
)

func padMsg(plaintext []byte, blockSize int) []byte {
    padding := blockSize - (len(plaintext) % blockSize)
    if padding == 0 {
        padding = blockSize
    }
    padtext := make([]byte, len(plaintext)+padding)
    copy(padtext, plaintext)
    return padtext
}

func encrypt(key, plaintext []byte) ([]byte, error) {
    // Create a new cipher block
    block, err := des.NewCipher(key)
    if err != nil {
        return nil, err
    }

    // Pad the plaintext with zeros
    blockSize := block.BlockSize()
    paddedText := padMsg(plaintext, blockSize)

    // Generate a random iv and append it to the beginning of the ciphertext
    iv := make([]byte, blockSize)
    if _, err := rand.Read(iv); err != nil {
        return nil, err
    }
    ciphertext := make([]byte, len(paddedText)+blockSize)
    copy(ciphertext[:blockSize], iv)

    // Create a new encrypter
    mode := cipher.NewCBCEncrypter(block, iv)

    // Encrypt the blocks
    mode.CryptBlocks(ciphertext[blockSize:], paddedText)

    // Return the entire ciphertext
    return ciphertext, nil
}


// don't touch below this line

func decrypt(key, ciphertext []byte) (plaintext []byte, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, 0)
	}
	return block
}

func test(key, plaintext []byte) {
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypting '%v' with key '%v'...\n", string(plaintext), string(key))
	decryptedText, err := decrypt(key, ciphertext)
	if err != nil {
		fmt.Println(err)
		return
	}
	decryptedText = bytes.Trim(decryptedText, "\x00")
	fmt.Printf("Decrypted: '%v'\n", string(decryptedText))
	fmt.Println("========")
}

func main() {
	test(
		[]byte("12344321"),
		[]byte("Today I met my crush, what a hunk"),
	)

	test(
		[]byte("p@$$w0rd"),
		[]byte("I hope my boyfriend never finds out about this"),
	)
}
