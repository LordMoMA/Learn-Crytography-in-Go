/*
State Ciphers
Stream ciphers are also referred to as "state ciphers" because the encryption and decryption of each character depends on the current state of the cipher.

"Current state" just means the "current progress", that is, how many bytes we've already encrypted or decrypted so far.

For example, let's take the following data for a one-time pad cipher:

plaintext = "010011"
key       = "101001"
When the "state" of the cipher is at index 0, we will only have encrypted the first digit:

plaintext = "010011"
key       = "101001"

state     = index 0
result    = "1"

state     = index 1
result    = "11"

state     = index 2
result    = "111"

state     = index 3
result    = "1110"

state     = index 4
result    = "11101"

state     = index 5
result    = "111010"
Assignment
We've been asked to update our code to log the progress of the cipher as it encrypts and decrypts each character. There have been some very large encryption tasks, and people are getting confused when the cipher doesn't seem to be doing anything for a while.

Update the crypt function. It should log this message after each character is sent to the result channel:

Crypted byte: NUM
Where NUM is the number of bytes that have been encrypted so far, starting at 1.
*/

import (
	"errors"
	"fmt"
)

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	count := 0
	for {
		textChar, textOk := <-textCh
		if !textOk {
			return
		}
		keyChar, keyOk := <-keyCh
		if !keyOk {
			return
		}
		result <- textChar ^ keyChar
		count++
		fmt.Printf("Crypted byte: %d\n", count)
	}
}

// don't touch below this line

func encrypt(plaintext, key []byte) ([]byte, error) {
	if len(plaintext) != len(key) {
		return nil, errors.New("plaintext and key must be the same length")
	}

	plaintextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(plaintextCh)
		for _, v := range plaintext {
			plaintextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(plaintextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}
	return res, nil
}

func decrypt(ciphertext, key []byte) ([]byte, error) {
	if len(ciphertext) != len(key) {
		return nil, errors.New("ciphertext and key must be the same length")
	}

	ciphertextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(ciphertextCh)
		for _, v := range ciphertext {
			ciphertextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(ciphertextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}
	return res, nil
}

func test(plaintext, key []byte) {
	fmt.Printf("Encrypting '%s' using key '%s'\n", string(plaintext), string(key))
	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypted ciphertext bytes: %v\n", ciphertext)
	decrypted, err := decrypt(ciphertext, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted message: %v\n", string(decrypted))
	fmt.Println("========")
}

func main() {
	test([]byte("Shazam"), []byte("Sk7p13"))
	test([]byte("I'm lovin it"), []byte("mysecurepass"))
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

Encrypting 'Shazam' using key 'Sk7p13'

Crypted byte: 1

Crypted byte: 2

Crypted byte: 3

Crypted byte: 4

Crypted byte: 5

Crypted byte: 6

Encrypted ciphertext bytes: [0 3 86 10 80 94]

Crypted byte: 1

Crypted byte: 2

Crypted byte: 3

Crypted byte: 4

Crypted byte: 5

Crypted byte: 6

Decrypted message: Shazam

========

Encrypting 'I'm lovin it' using key 'mysecurepass'

Crypted byte: 1

Crypted byte: 2

Crypted byte: 3

Crypted byte: 4

Crypted byte: 5

Crypted byte: 6

Crypted byte: 7

Crypted byte: 8

Crypted byte: 9

Crypted byte: 10

Crypted byte: 11

Crypted byte: 12

Encrypted ciphertext bytes: [36 94 30 69 15 26 4 12 30 65 26 7]

Crypted byte: 1

Crypted byte: 2

Crypted byte: 3

Crypted byte: 4

Crypted byte: 5

Crypted byte: 6

Crypted byte: 7

Crypted byte: 8

Crypted byte: 9

Crypted byte: 10

Crypted byte: 11

Crypted byte: 12

Decrypted message: I'm lovin it

========
*/
