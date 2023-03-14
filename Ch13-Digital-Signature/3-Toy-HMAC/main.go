/*
Toy HMAC
At Passly, we use HMACs to authenticate messages between our internal servers when they need to make requests to each other over the public internet.

To demonstrate to our Luddite manager why we should use an open-source crypto library instead of writing our own HMAC implementation, we decided to write our own and then prove it's inferiority. Hopefully we don't get fired instead of getting our way.

Assignment
Complete the hmac function. It should:

Split the key into two halves. The second half should be the larger half if key's length is odd
Return the result of sha256(keyFirstHalf + sha256(keySecondHalf + message)) as a string in lowercase hex
Cast strings directly to slices of bytes and don't use any delimiters when concatenating the data.
*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hmac(message, key string) string {
	half1 := key[:len(key)/2]
	half2 := key[len(key)/2:]
	if len(key)%2 != 0 {
		half1 = key[:len(key)/2-1]
		half2 = key[len(key)/2-1:]
	}
	h := sha256.New()
	h.Write([]byte(message + half2))
	sum1 := h.Sum(nil)
	hexStr1 := hex.EncodeToString(sum1)

	h.Write([]byte(half1 + hexStr1))
	sum2 := h.Sum(nil)
	hexStr2 := hex.EncodeToString(sum2)
	return hexStr2
}

// solution2
func hmac(message, key string) string {
	keyBytes := []byte(key)
	if len(keyBytes)%2 == 1 {
		keyBytes = append(keyBytes, byte(0))
	}

	halfLen := len(keyBytes) / 2
	keyFirstHalf := keyBytes[:halfLen-1]
	keySecondHalf := keyBytes[halfLen-1:]

	h := sha256.New()
	h.Write(keySecondHalf)
	h.Write([]byte(message))
	hashedSecondHalfMessage := h.Sum(nil)

	h.Reset()
	h.Write(keyFirstHalf)
	h.Write(hashedSecondHalfMessage)
	finalHash := h.Sum(nil)

	return hex.EncodeToString(finalHash)
}

// The reason why h.Write(keySecondHalf + []byte(message)) does not work is that keySecondHalf is a slice of bytes, and the + operator is not defined for slices.
// To concatenate two slices of bytes, we need to use the append function. So to concatenate keySecondHalf and message, we would use append(keySecondHalf, []byte(message)...). However, we cannot pass the result of append directly to h.Write because h.Write expects a slice of bytes, not a slice of bytes combined with another slice of bytes. Instead, we need to write each slice of bytes to the hash function separately using two calls to h.Write.
// Thus, we use h.Write(keySecondHalf) followed by h.Write([]byte(message)) to write the key second half and the message to the hash function separately.
// Regarding the h.Reset() step, it is necessary to reset the hash function state before writing new data to it. Otherwise, the new data would be appended to the previous data, and the resulting hash would be incorrect.

// don't touch below this line

func test(message, key string) {
	fmt.Printf("Calculating HMAC of '%v' with key '%v'...\n", message, key)
	checksum := hmac(message, key)
	fmt.Printf("HMAC: %v\n", checksum)
	fmt.Println("========")
}

func main() {
	test("I hope no one finds the Bitcoin keys I keep under my mailbox", "super_secret_password")
	test("No really, they're just written on a piece of paper", "correct horse battery staple")
	test("It's like a gazillion satoshis worth of BTC", "aFiveDoll@rWr3nch")
}

/*

Calculating HMAC of 'I hope no one finds the Bitcoin keys I keep under my mailbox' with key 'super_secret_password'...

HMAC: 09603a12641b87f3ef7ab139f4f39d911e6f02c821caa51cf33cb93d037b3c1b

========

Calculating HMAC of 'No really, they're just written on a piece of paper' with key 'correct horse battery staple'...

HMAC: ff2da418333b6f48e6937bab496fd0161170e022dd695283f17eb92c4592c0c9

========

Calculating HMAC of 'It's like a gazillion satoshis worth of BTC' with key 'aFiveDoll@rWr3nch'...

HMAC: 13c800bdf9a561984cb160f91ad03e10b31903ae7770232cd23bd8e053dc8984

========
*/
