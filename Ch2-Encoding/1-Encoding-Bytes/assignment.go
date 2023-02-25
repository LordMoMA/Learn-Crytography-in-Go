/*
Encoding Bytes
Encoding and encryption are not the same. When we talk about encoding, we are talking about converting raw data (binary) into a format that can be stored or transmitted as text.

The are many popular ways to encode data as text. Some include:

Hexadecimal
Base64
ASCII
Typically, it is a bad idea to create your own encoding scheme. However, for the sake of learning, we are going to create one for fun.

Alphabets
Encoding schemes have an alphabet. An alphabet is just the set of available characters in the scheme.

Arabic numerals alphabet
0123456789

Hexadecimal alphabet
0123456789abcdef

Base64 alphabet
ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/

The number of characters in the alphabet is the "base". Base64 has 64 characters.

Binary Literals in Go
Binary integer literals are written in Go as follows:

// zeroWrittenInBinary is just an int with the value 0
zeroWrittenInBinary := 0b0000
fmt.Println(zeroWrittenInBinary)
// prints: 0

// oneWrittenInBinary is just an int with the value 1
oneWrittenInBinary := 0b0001
fmt.Println(oneWrittenInBinary)
// prints: 1
Assignment
Complete the base8Char function. It accepts a single byte of data and returns the associated character from our alphabet: ABCDEFGH.

For example, given these binary numbers, it will return the following characters:

0000 -> A
0001 -> B
0010 -> C
0011 -> D
etc
Note that because we only have 8 characters in our alphabet, we will ignore any numbers higher than 8. Remember, an entire byte can contain the numbers 0 through 255.

Steps
Convert the byte to an integer using int().
If the number is out of the range of our alphabet, return an empty string.
Return a string containing the character associated with the index represented by the integer.
*/

package main

import "fmt"

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	// ?
}

// don't touch below this line

func test(rawMessage []byte) {
	decodedMsg := ""
	for _, b := range rawMessage {
		decodedMsg += base8Char(b)
	}
	fmt.Printf("Encoding %04b in custom base 8...\n", rawMessage)
	fmt.Printf("Encoded result: %v\n", decodedMsg)
	fmt.Println("========")
}

func main() {
	test([]byte{0b010, 0b000, 0b001})
	test([]byte{0b011, 0b000, 0b011})
	test([]byte{0b110, 0b000, 0b001})
}

/*
Encoding != Encryption
Encoding and encryption are not the same.

Encoding
When we talk about encoding in the context of cryptography, we are talking about converting raw data (binary) into a format that can be stored or transmitted as text.

It is not safe to simply encode a secret message. Anyone who has access to the encoded message can decode it, there are no encryption keys involved. Anyone can trivially figure out the alphabet used to encode the message and decode it.

Encryption
Encryption, on the other hand, is a secure way to transmit secret information. The critical difference is that encryption requires a key to decrypt the message. Without the key, the message is unreadable.
*/
