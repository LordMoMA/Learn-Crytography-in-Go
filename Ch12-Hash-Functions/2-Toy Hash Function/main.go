/*
Toy Hash Function
Let's build a toy hash function!

Our goal will be to build a function that:

Is hard to reverse
Has a fixed output size
Is deterministic
Our function will not be suitable for production use, but will be useful for the Passly marketing team to explain how our security systems work from a high level.

Shifting bits in Go
To shift a byte left by numBits:

numBits := 3
// original = 11110000
shifted := original << numBits
// shifted = 10000000
XOR in Go
result := a ^ b
Assignment
Complete the hash() function. It takes an arbitrarily sized []byte and returns a fixed size [8]byte.

It should do the following:

Rotate the bits in each byte left 3 bits, do this one byte at a time
Shift the bits in each byte left 2 bits, do this one byte at a time
Create a final empty array of exactly 4 bytes
The value at each index in final array is the XOR of all the values of the indexes in the rotated and shifted input that equal that index in modulo 4. For example, if the (rotated and shifted) input is:
[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]
Example
00001111 01010101 00110011 11110000 11111111 00000000
01111000 10101010 10011001 10000111 11111111 00000000
11100000 10101000 01100100 00011100 11111100 00000000
(11100000 ^ 11111100) (10101000 ^ 00000000) 01100100 00011100
00011100 10101000 01100100 00011100
*/

package main

import (
	"fmt"
	"math/bits"
)

/*
This expression performs a bit rotation and bit shifting operation on a single byte of the input.

Let's break it down step by step:

input[i] << 2 - This shifts the bits in the byte input[i] to the left by 2 positions. For example, if the byte was originally 0b11010110, this would shift it to 0b01011000.

input[i] >> 6 - This shifts the bits in the byte input[i] to the right by 6 positions. For example, if the byte was originally 0b11010110, this would shift it to 0b00000011.

(input[i] << 2) | (input[i] >> 6) - This performs a bitwise OR operation between the results of the previous two steps. Essentially, it takes the two binary numbers produced in steps 1 and 2 and "merges" them together by setting all the bits that are 1 in either number to 1 in the result. In our example, this would give us 0b01011011.

The end result of this expression is a rotated and shifted version of the original byte, where the first 2 bits have been moved to the end of the byte. This can be useful in creating "avalanche effect" in a hash function, where small changes in the input result in large changes in the output, making it difficult for an attacker to predict or manipulate the output.
*/

func hash(input []byte) [4]byte {
	// Rotate bits left 3 bits and shift left 2 bits
	for i := 0; i < len(input); i++ {
		input[i] = (input[i] << 2) | (input[i] >> 6)
		input[i] = (input[i] << 3) | (input[i] >> 5)
	}

	// Create final array
	final := [4]byte{}

	// XOR values from rotated and shifted input
	for i := 0; i < len(input); i++ {
		final[i%4] ^= input[i]
	}

	return final
}

// solution2

func hash(input []byte) [4]byte {
	for i, b := range input {
		rotated := bits.RotateLeft(uint(b), 3)
		input[i] = byte(rotated)
	}

	for i, b := range input {
		shifted := b << 2
		input[i] = byte(shifted)
	}

	final := [4]byte{}
	for i, b := range input {
		final[i%len(final)] ^= b
	}

	return final
}

// don't touch below this line

func test(input string) {
	fmt.Printf("Hashing '%s'...\n", input)
	fmt.Printf("Output: %X\n", hash([]byte(input)))
	fmt.Println("========")
}

func main() {
	test("Example message")
	test("This is a slightly longer example to hash")
	test("This is much a longer example of some text to hash, maybe it's the opening paragraph of a blog post")
}

/*

Hashing 'Example message'...

Output: 27C242C7

========

Hashing 'This is a slightly longer example to hash'...

Output: 288AC849

========

Hashing 'This is much a longer example of some text to hash, maybe it's the opening paragraph of a blog post'...

Output: 418C844B

========
*/
