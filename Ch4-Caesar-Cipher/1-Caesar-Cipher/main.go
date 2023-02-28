/*
Caesar Cipher
Example of Caesar cipher with right shift 3:

Alphabet:    ABCDEFGHIJKLMNOPQRSTUVWXYZ
Shifted:     DEFGHIJKLMNOPQRSTUVWXYZABC
Each letter in the plaintext alphabet maps directly to its corresponding match in the cipher alphabet.

For example:

"A" -> "D"
"B" -> "E"
Assignment
To demonstrate the superiority of our modern ciphers at Passly, we've been asked to implement a Caesar cipher. It's so the marketing team can show off how much better our algorithms are than the ones used by the Romans. Go figure.

Complete the encrypt, decrypt, crypt, and getOffsetChar functions. The encrypt and decrypt functions do what you would expect, crypt and getOffsetChar are helper functions.

encrypt
Should call and return the result of crypt().

decrypt
Should call and return the result of crypt(), but should pass in the negative value of key so that the offset is reversed.

crypt
This is where the meat of the encryption/decryption happens. Iterate over each character in the text, and use getOffsetChar to get the proper substitution character. Return a string of all of the substitution characters.

getOffsetChar
This function takes a single character (as a rune) and an offset. It should return the character (as a string) that is offset positions away from the given character. The alphabet has been provided for you.

For example, when offset = 5:

a -> f
b -> g
...
x -> c
z -> e
Or when offset = 1:

a -> b
b -> c
...
x -> y
z -> a
Remember that offset can be negative, and remember that the alphabet wraps around. For example, when offset = -5:

a -> v
b -> w
c -> x
d -> y
e -> z
If you can't find the character in the alphabet, return an empty string.

/////////////////////////////

package main

import "fmt"

func encrypt(plaintext string, key int) string {
	// ?
}

func decrypt(ciphertext string, key int) string {
	// ?
}

func crypt(text string, key int) string {
	// ?
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	// ?
}

// don't touch below this line

func test(plaintext string, key int) {
	fmt.Printf("Encrypting %v with key %v\n", plaintext, key)
	encrypted := encrypt(plaintext, 5)
	fmt.Printf("Encrypted text: %v\n", encrypted)
	decrypted := decrypt(encrypted, 5)
	fmt.Printf("Decrypted text: %v\n", decrypted)
	fmt.Println("========")
}

func main() {
	test("abcdefghi", 1)
	test("hello", 5)
	test("correcthorsebatterystaple", 16)
	test("onetwothreefourfivesixseveneightnineten", 25)
}

*/

package main

import "fmt"

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	result := ""
	for _, c := range text {
		offset := getOffsetChar(c, key)
		result += string(offset)
	}
	return result

}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	for i, curr := range alphabet {
		if curr == c {
			modI := (i + offset) % len(alphabet)
			if modI < 0 {
				modI += len(alphabet)
			}
			return string(alphabet[modI])
		}
	}
	return ""
}

// don't touch below this line

func test(plaintext string, key int) {
	fmt.Printf("Encrypting %v with key %v\n", plaintext, key)
	encrypted := encrypt(plaintext, 5)
	fmt.Printf("Encrypted text: %v\n", encrypted)
	decrypted := decrypt(encrypted, 5)
	fmt.Printf("Decrypted text: %v\n", decrypted)
	fmt.Println("========")
}

func main() {
	test("abcdefghi", 1)
	test("hello", 5)
	test("correcthorsebatterystaple", 16)
	test("onetwothreefourfivesixseveneightnineten", 25)
}
