/*
Key Schedules
The simplest versions of block ciphers operate on single blocks of data, however, those blocks of data typically go through many rounds of encryption. Each round of encryption needs its own key, but it's nearly impossible for users to keep track of more than a single key.

A key schedule is an algorithm that a block cipher employs to split an original key into multiple "round keys" or "subkeys". These round keys are deterministically derived from the original key, meaning that the same original key will always produce the same round keys.

For example the following "master key":

key = 1101011100010010
might produce the following 8 "round keys" in an 8-round block cipher:

roundKey0 = 1101011100010010
roundKey1 = 1101011100010010
roundKey2 = 1101011100010010
roundKey3 = 1101011100010010
roundKey4 = 1101011100010010
roundKey5 = 1101011100010010
roundKey6 = 1101011100010010
roundKey7 = 1101011100010010
Assignment
In the real world, there are many different production algorithms for key schedules. At Passly, we've been asked to implement a simple key schedule that can be used in our test suite. It doesn't need to be "production-level-secure", it just needs to get the high-level idea across.

Complete the deriveRoundKey() function. It accepts a master key, which is 4 bytes long and represented as a slice of bytes. It will also accept a "round number", which is just an int representing which round key is being derived. The key schedule produces a round key where each byte in the round key is the original byte from the master key XORed with the binary representation of the round number.

For example:

masterKey     = 01101100 01110000...
roundNumber   = 00000001
roundKey      = 01101101 01110001...
or when roundNumber = 5

masterKey     = 01101100 01110000...
roundNumber   = 00000101
roundKey      = 01101001 01110101...
*/

package main

import "fmt"

/*

this is wrong, and byte(roundNumber) is still wrong

	func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	roundKey := []byte{}
	for b := range masterKey {
		xor := b ^ roundNumber
		roundKey = append(roundKey, xor)
	}
	return roundKey
}
*/

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	var roundKey [4]byte
	for i := 0; i < 4; i++ {
		roundKey[i] = byte(roundNumber) ^ masterKey[i]
	}
	return roundKey
}

// don't touch below this line

func test(masterKey [4]byte) {
	fmt.Printf("Deriving round keys from master: %X...\n", masterKey)
	for i := 1; i < 9; i++ {
		roundKey := deriveRoundKey(masterKey, i)
		fmt.Printf(" - Round key %v: %X\n", i, roundKey)
	}
	fmt.Println("========")
}

func main() {
	test([4]byte{0xAA, 0xFF, 0x11, 0xBC})
	test([4]byte{0xEB, 0xCD, 0x13, 0xFC})
}

/*

Deriving round keys from master: AAFF11BC...

 - Round key 1: ABFE10BD

 - Round key 2: A8FD13BE

 - Round key 3: A9FC12BF

 - Round key 4: AEFB15B8

 - Round key 5: AFFA14B9

 - Round key 6: ACF917BA

 - Round key 7: ADF816BB

 - Round key 8: A2F719B4

========

Deriving round keys from master: EBCD13FC...

 - Round key 1: EACC12FD

 - Round key 2: E9CF11FE

 - Round key 3: E8CE10FF

 - Round key 4: EFC917F8

 - Round key 5: EEC816F9

 - Round key 6: EDCB15FA

 - Round key 7: ECCA14FB

 - Round key 8: E3C51BF4

========
*/
