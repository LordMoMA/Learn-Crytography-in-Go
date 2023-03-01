/*
One Time Pad
A cipher is said to have perfect security if an attacker who has access to only the ciphertext can infer absolutely nothing of interest about the plaintext. Such perfect ciphers do exist, one such example is the "one-time pad".

The XOR operator in Go
The ^ operator will XOR two bytes. This cheat-sheet may be helpful.

byte1 := 0b01110000
byte2 := 0b10101000
xorRes := byte1 ^ byte2
fmt.Printf("%b\n", xorRes)
// 11011000
Assignment
While the One Time Pad is a very powerful cipher, it's also very difficult to use in practice. That said, our product managers have found a use case in Passly that isn't overly inconvenient, so they've asked us to implement it.

Complete the crypt function. It accepts some data and a key as slices of bytes, and returns the result of an XOR operation on all the bits.

For example:

data       = "0110100001100101011011000110110001101111"
key        = "0111001101101010011001100111010101100100"

output     = "0001101100001111000010100001100100001011"
Note on software design
Because XOR encryption is the perfect inverse of XOR decryption, the encrypt() and decrypt() will be identical in functionality. That's why we're writing a single crypt function that's used directly by the encrypt and decrypt functions. It makes the code easier to understand and use.
*/

package main

import "fmt"

func crypt(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i := range plaintext {
		keyByte := key[i%len(key)]
		ciphertext[i] = plaintext[i] ^ keyByte
	}
	return ciphertext
}

// don't touch below this line

func encrypt(plaintext, key []byte) []byte {
	return crypt(plaintext, key)
}

func decrypt(ciphertext, key []byte) []byte {
	return crypt(ciphertext, key)
}

func test(plaintext, key []byte) {
	ciphertext := encrypt(plaintext, key)
	fmt.Printf("Encrypting '%s' using key '%s'\n", string(plaintext), string(key))
	fmt.Printf("Encrypted ciphertext bytes: %v\n", ciphertext)
	decrypted := decrypt(ciphertext, key)
	fmt.Printf("Decrypted message: %v\n", string(decrypted))
	fmt.Println("========")
}

func main() {
	test([]byte("Shazam"), []byte("Sk7p13"))
	test([]byte("I'm lovin it"), []byte("mysecurepass"))
	test([]byte("Don't tell him I'm in love"), []byte("c5f149783abf22a96e9a7bb999"))
}

/*
Encrypting 'Shazam' using key 'Sk7p13'

Encrypted ciphertext bytes: [0 3 86 10 80 94]

Decrypted message: Shazam

========

Encrypting 'I'm lovin it' using key 'mysecurepass'

Encrypted ciphertext bytes: [36 94 30 69 15 26 4 12 30 65 26 7]

Decrypted message: I'm lovin it

========

Encrypting 'Don't tell him I'm in love' using key 'c5f149783abf22a96e9a7bb999'

Encrypted ciphertext bytes: [39 90 8 22 64 25 67 93 95 13 66 14 91 95 65 112 17 8 25 8 89 66 14 86 79 92]

Decrypted message: Don't tell him I'm in love

========

Perfect Security
The One Time Pad has perfect security because an attacker who has access to only the ciphertext can infer absolutely nothing of interest about the plaintext.

Ciphers used in production are usually not perfectly secure
Most production ciphers are not perfectly secure, but are "close enough". In short, trade-offs are made that add to the practical security of the system while sacrificing the perfect theoretical security of the cipher itself.

The big problem with the One Time Pad is that the key needs to be the same length as the message. That means to encrypt a 128 GB hard drive, I'd need a 128 GB key!! That's just not practical.

Some ciphers are just crap
The Caesar cipher is a great example of a cipher that is NOT perfectly secure OR practically secure. As we demonstrated earlier, when given access to the ciphertext of a Caesar cipher, an attacker can see the positions and patterns of characters in the plaintext.
*/
