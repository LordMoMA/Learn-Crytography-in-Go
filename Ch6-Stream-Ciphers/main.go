/*
Stream Ciphers
The Ceasar cipher, other substitution ciphers, and the "one-time pad" are all examples of stream ciphers.

A stream cipher is a symmetric key cipher where plaintext digits are combined with a key stream. In a stream cipher, each plaintext digit is encrypted one at a time with the corresponding digit of the keystream, to give a digit of the ciphertext stream.

Assignments
We've been asked to update our "One Time Pad". For performance reasons, we'll now be reading and writing the data using Go channels. This will allow us to encrypt the data as it's read in from an external database. Rather than storing the entire message in memory (which could be many Gigabytes) and then decrypting it all at once, this new crypt function can do it one character at a time.

crypt(textCh, keyCh <-chan byte, result chan<- byte)
Read one byte at a time from the textCh and keyCh channels. Perform an XOR operation on the two bytes and write the result to the result channel. Stop once either channel is closed.

Be sure to close the result channel when you're done writing to it.
*/

package main

import (
	"errors"
	"fmt"
)

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result) // close result channel when we're done

	for {
		textByte, ok1 := <-textCh // read from textCh channel
		keyByte, ok2 := <-keyCh   // read from keyCh channel

		if !ok1 || !ok2 { // if either channel is closed, we're done
			return
		}

		resultByte := textByte ^ keyByte // XOR operation
		result <- resultByte             // write result to result channel
	}
}

// don't touch below this line

func encrypt(plaintext, key []byte) ([]byte, error) {
	if len(plaintext) != len(key) {
		return nil, errors.New("plaintext and key must be the same length")
	}

	plaintextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(plaintextCh)
		for _, v := range plaintext {
			plaintextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(plaintextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		fmt.Printf("Got encrypted byte: %08b\n", v)
		res = append(res, v)
	}
	return res, nil
}

func decrypt(ciphertext, key []byte) ([]byte, error) {
	if len(ciphertext) != len(key) {
		return nil, errors.New("ciphertext and key must be the same length")
	}

	ciphertextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(ciphertextCh)
		for _, v := range ciphertext {
			ciphertextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(ciphertextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		fmt.Printf("Got decrypted byte: %08b\n", v)
		res = append(res, v)
	}
	return res, nil
}

func test(plaintext, key []byte) {
	fmt.Printf("Encrypting '%s' using key '%s'\n", string(plaintext), string(key))
	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypted ciphertext bytes: %v\n", ciphertext)
	decrypted, err := decrypt(ciphertext, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted message: %v\n", string(decrypted))
	fmt.Println("========")
}

func main() {
	test([]byte("Shazam"), []byte("Sk7p13"))
	test([]byte("I'm lovin it"), []byte("mysecurepass"))
}

/*

Encrypting 'Shazam' using key 'Sk7p13'

Got encrypted byte: 00000000

Got encrypted byte: 00000011

Got encrypted byte: 01010110

Got encrypted byte: 00001010

Got encrypted byte: 01010000

Got encrypted byte: 01011110

Encrypted ciphertext bytes: [0 3 86 10 80 94]

Got decrypted byte: 01010011

Got decrypted byte: 01101000

Got decrypted byte: 01100001

Got decrypted byte: 01111010

Got decrypted byte: 01100001

Got decrypted byte: 01101101

Decrypted message: Shazam

========

Encrypting 'I'm lovin it' using key 'mysecurepass'

Got encrypted byte: 00100100

Got encrypted byte: 01011110

Got encrypted byte: 00011110

Got encrypted byte: 01000101

Got encrypted byte: 00001111

Got encrypted byte: 00011010

Got encrypted byte: 00000100

Got encrypted byte: 00001100

Got encrypted byte: 00011110

Got encrypted byte: 01000001

Got encrypted byte: 00011010

Got encrypted byte: 00000111

Encrypted ciphertext bytes: [36 94 30 69 15 26 4 12 30 65 26 7]

Got decrypted byte: 01001001

Got decrypted byte: 00100111

Got decrypted byte: 01101101

Got decrypted byte: 00100000

Got decrypted byte: 01101100

Got decrypted byte: 01101111

Got decrypted byte: 01110110

Got decrypted byte: 01101001

Got decrypted byte: 01101110

Got decrypted byte: 00100000

Got decrypted byte: 01101001

Got decrypted byte: 01110100

Decrypted message: I'm lovin it

========
*/
