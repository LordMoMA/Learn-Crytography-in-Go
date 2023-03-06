/*
Nonces
We talked a bit about nonces when we talked about IVs, but let's take a closer look at them.

A nonce is an arbitrary number that must be used only once in a cryptographic communication. Nonces are not secret and can be transported publicly.

Assignment
Because we use AES-256 at Passly, we're very concerned about using truly random nonces so we don't accidentally reuse once. Write the nonceStrength function. It should return the strength of a nonce as an integer.

We refer to the strength of a nonce as the total number of possible combinations of bits that could exist in the nonce.

Example
0b10010100
0b00010001
0b11000000
Each nonce above has a strength of 256 because there are 256 possible combinations of bits that could exist in an 8-bit nonce.

Nonce Review
Nonces in AES
In some modes of AES, a nonce is used to prevent attacks that exploit the reuse of keys. It can be catastrophic if the same key is used to encrypt two different messages because some information will be leaked that an attacker can utilize.

To be able to use the same key while staying secure, we combine the key with a random nonce for each new encryption. This provides uniqueness to each encryption and prevents information leakage.

Nonces in Bitcoin
In Bitcoin, a nonce is used to mine new coins. Miners increment (or randomly generate) new nonces and check if the result has a specific value. If it does, the miner earns some Bitcoin as a reward.

We'll talk more about hashes later, but just know that at the end of the day, Bitcoin miners are just spending electricity to continuously loop through nonces trying to find the "golden nonce".
*/

package main

import (
	"fmt"
	"math/rand"
)

// nonceStrength returns the number of bits of entropy in the nonce.
func nonceStrength(nonce []byte) int {
	return 2 ^ len(nonce)
}

// don't touch below this line

func generateIV(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func main() {
	rand.Seed(0)
	for i := 1; i < 6; i++ {
		nonce, _ := generateIV(i)
		nonceStr := nonceStrength(nonce)
		fmt.Printf("A random nonce of %v bytes has strength of %v\n", i, nonceStr)
	}
}

/*

A random nonce of 1 bytes has strength of 3

A random nonce of 2 bytes has strength of 0

A random nonce of 3 bytes has strength of 1

A random nonce of 4 bytes has strength of 6

A random nonce of 5 bytes has strength of 7
*/
