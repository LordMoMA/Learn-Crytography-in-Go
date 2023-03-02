/*
Block Sizes
Block ciphers use fixed-size blocks and fixed-size keys. However, just because the blocks are fixed-size doesn't mean the message must be a fixed size as well. Block ciphers can operate on messages of any length, and the key doesn't even need to be the same length as the message or the block.

For example:

3DES: blocksize=64 bits, keysize=168 bits
AES-256: blocksize=128 bits, keysize=256 bits
Chunks and Padding
Messages in a block cipher are broken up into blocks, and each block is encrypted separately.

For example, let's say we are using a cipher that requires 256-bit blocks. We have a message of 650 bits that we want to encrypt. The algorithm would break that message up into three blocks:

block1 = first 256 bits
block2 = next 256 bits
block3 = last 138 bits
The last block is then padded with extra garbage bits so that it also has 256 bits of data. That padding is stripped off when the message is decrypted.

Assignment
While robust cryptography libraries take care of block padding for us, it's important to understand how it works. As part of employee onboarding at Passly, we require our cryptography engineers to write a padding function from scratch so they can understand how it works.

Complete the padWithZeros function. It takes a message (block) as a []byte and a desired output size. It should increase the size of block by adding zero-value-bytes until it has a length of desiredSize. Return the new byte slice.
*/

package main

import "fmt"

func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, byte(0))
	}
	return block
}

// don't touch below this line

func test(block []byte, desiredSize int) {
	fmt.Printf("Padding %v for a total of %v bytes...\n",
		block,
		desiredSize,
	)
	padded := padWithZeros(block, desiredSize)
	fmt.Printf("Result: %v\n", padded)
	fmt.Println("========")
}

func main() {
	test([]byte{0xFF}, 4)
	test([]byte{0xFA, 0xBC}, 8)
	test([]byte{0x12, 0x34, 0x56}, 12)
	test([]byte{0xFA}, 16)
}

/*

Padding [255] for a total of 4 bytes...

Result: [255 0 0 0]

========

Padding [250 188] for a total of 8 bytes...

Result: [250 188 0 0 0 0 0 0]

========

Padding [18 52 86] for a total of 12 bytes...

Result: [18 52 86 0 0 0 0 0 0 0 0 0]

========

Padding [250] for a total of 16 bytes...

Result: [250 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

========
*/
