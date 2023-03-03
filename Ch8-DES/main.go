/*
Feistel Network
Many block ciphers, including DES which we'll be talking about soon, use a Feistel network (or "Feistel Cipher") as a key component of their encryption algorithms.

Feistel structures have the useful property that encryption and decryption operations are very similar and sometimes identical. Decryption only requires a reversal of the key schedule. This keeps the implementation of the algorithms simple and concise.

Feistel Network

n = The number of ciphering rounds
L0 = Left half of plaintext
R0 = Right half of plaintext
K0 - Kn = The round keys (from a key schedule)
F = The round function (specific to the ciphering algorithm, like DES for example)
One of the most interesting things about Feistel networks is that the round function does NOT have to be reversible.

A Feistel cipher is not a fully-fledged encryption algorithm but is rather a framework that more complete cipher implementations (like DES) utilize.

Assignment
For marketing purposes, Passly has decided to create its own Feistel network. It will use the Go standard library's SHA-256 hash function as the round function.

Here's some psuedocode:

feistel(msg, key []byte, rounds int) []byte
Split the message into equal left and right-hand sides of equal length.
For each round:
nextRHS = xor(lhs, hash(rhs+key))
nextLHS = oldRHS
When you're done with all the rounds, return the concatenation of the right and left-hand sides (right first, then left)
The hash() function is provided for you.

*/

/*
The roundKeys parameter is a slice of slices of bytes ([][]byte), not a slice of bytes ([]byte).

In the feistel function, roundKeys is a two-dimensional slice of bytes where each inner slice is a round key. The outer slice represents all the round keys needed for all the rounds of the Feistel cipher.

For example, if you have a Feistel cipher with 10 rounds, you will need 10 round keys, so roundKeys will be a slice of 10 slices of bytes.

examples:
roundKeys := [][]byte{
    {0x01, 0x02, 0x03, 0x04},
    {0x05, 0x06, 0x07, 0x08},
    {0x09, 0x0A, 0x0B, 0x0C},
}

*/

package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/bits"
)

func feistel(msg []byte, roundKeys [][]byte) []byte {
	lhs := msg[:len(msg)/2]
	rhs := msg[len(msg)/2:]
	for _, key := range roundKeys {
		nextRHS := xor(lhs, hash(rhs, key, len(rhs)))
		nextLHS := rhs
		rhs = nextRHS
		lhs = nextLHS
	}
	return append(rhs, lhs...)
}

// don't touch below this line

func test(msg []byte, key []byte, rounds int) {
	roundKeys := [][]byte{}
	for i := 0; i < rounds; i++ {
		ui := binary.BigEndian.Uint32(key)
		rotated := bits.RotateLeft(uint(ui), i)
		finalRound := make([]byte, len(key))
		binary.LittleEndian.PutUint32(finalRound, uint32(rotated))
		roundKeys = append(roundKeys, finalRound)
	}

	fmt.Printf("Encrypting '%v' with %v round keys...\n", string(msg), rounds)
	encrypted := feistel(msg, roundKeys)
	decrypted := feistel(encrypted, reverse(roundKeys))
	fmt.Printf("Decrypted: '%v'\n", string(decrypted))
	fmt.Println("========")
}
func main() {
	test(
		[]byte("General Kenobi!!!!"),
		[]byte("thesecret"),
		8,
	)
	test(
		[]byte("Hello there!"),
		[]byte("@n@kiN"),
		16,
	)
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

// outputLength should be equal to the key length
// when used in feistel so that the XOR operates on
// inputs of the same size
func hash(lhs, rhs []byte, outputLength int) []byte {
	h := sha256.New()
	h.Write(append(lhs, rhs...))
	return h.Sum(nil)[:outputLength]
}

/*

Encrypting 'General Kenobi!!!!' with 8 round keys...

Decrypted: 'General Kenobi!!!!'

========

Encrypting 'Hello there!' with 16 round keys...

Decrypted: 'Hello there!'

========
*/
