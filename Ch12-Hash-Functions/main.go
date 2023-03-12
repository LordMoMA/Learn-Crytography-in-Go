/*
Hash Functions
Hash functions have 3 basic goals:

Hash functions scramble data deterministically
No matter the input, the output of a hash function always has the same size
The original data can not be retrieved from the scrambled data (one-way function)

Checking integrity
At Passly, we use SHA-256 for many things, but one of the most important is to ensure the integrity of a password vault. Each time we save a vault, we hash the vault's contents and store the hash in our database. Later, if we need to verify the integrity of the vault, we can hash the vault's contents again and compare the two hashes. If they match, we know the vault has not been tampered with.

Assignment
Add the following functions and methods to the program:

newHasher
h.Write
h.GetHex
newHasher
Returns a pointer to a new hasher. Uses sha256.New() to create a new hash.Hash.

h.Write
A method on a pointer to a hasher. Uses h.Write() to write data to the hasher. It should accept a string and cast the string to a []byte. It should pass along the return values, that is, it returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early.

h.GetHex
A method on a pointer to a hasher. Uses h.Sum() to get the hash value of the data written to the hasher. It should encode the hash value as a lowercase hex string and return it.
*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
	return &hasher{
		hash: sha256.New(),
	}
}

func (h *hasher) Write(data string) (int, error) {
	bytes := []byte(data)
	return h.hash.Write(bytes)
}

func (h *hasher) GetHex() string {
	sum := h.hash.Sum(nil)
	hexStr := hex.EncodeToString(sum)
	return hexStr
}

// don't touch below this line

func test(passwords []string) {
	fmt.Printf("Hashing vault of length %v...\n", len(passwords))
	h := newHasher()
	for _, password := range passwords {
		h.Write(password)
		fmt.Printf("Adding '%v' to vault hash...\n", password)
	}
	fmt.Printf("Vault hash: %v\n", h.GetHex())
	fmt.Println("========")
}

func main() {
	test([]string{"password1", "password2", "password3"})
	test([]string{"abercromni3", "f1tch", "123456", "abcdefg1234"})
	test([]string{"IHeartNanciedrake", "m7B1rthd@y"})
}
