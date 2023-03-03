/*
Initialization Vectors
As we briefly mentioned before, an iv, or initialization vector, is a random value that is used to initialize a block cipher. It is used to ensure that the same plaintext always encrypts to a different ciphertext.

Nonce
If the iv doesn't need to be truly random, and the only requirement is that it's only used once, then it's often called a "nonce", which is just short for "a number used only once".

Whether the iv for an encryption algorithm needs to be cryptographically random, pseudorandom, or just unique will be specified in the documentation for that algorithm. It's important to read the documentation for the algorithm you're using to ensure that you're providing the correct level of security for your iv.

Assignment
Update the generateIV function. It's currently returning unique values due to a global count variable, but it's not random, which is what Passly's cipher requires.

Use the math/rand package's rand.Read function to generate a random iv of the specified byte length.

A note on math/rand vs crypto/rand
In production, we would swap out the math/rand package for the crypto/rand package. The math/rand package is not cryptographically secure and should only be used for testing when we need to generate secure random values.
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
)

var count = 0

func generateIV(length int) ([]byte, error) {
	iv := make([]byte, length)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}

// don't touch below this line

func main() {
	rand.Seed(0)
	for i := 8; i < 17; i++ {
		iv, err := generateIV(i)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("%v-byte iv: %0*X\n", i, 2*i, iv)
	}
}

/*

8-byte iv: 0194FDC2FA2FFCC0

9-byte iv: 41D3FF12045B73C86E

10-byte iv: 4FF95FF662A5EEE82ABD

11-byte iv: F44A2D0B75FB180DAF48A7

12-byte iv: 9EE0B10D394651850FD4A178

13-byte iv: 892EE285ECE1511455780875D6

14-byte iv: 4EE2D3D0D0DE6BF8F9B44CE85FF0

15-byte iv: 44C6B1F83B8E883BBF857AAB99C5B2

16-byte iv: 52C7429C32F3A8AEB79EF856F659C18F
*/
