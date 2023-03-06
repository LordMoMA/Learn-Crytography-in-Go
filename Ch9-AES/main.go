/*
Advanced Encryption Standard
AES is a block cipher just like DES, but there are some key differences:

AES uses 128, 192, or 256 bit keys. DES uses 56 bit keys
AES block size is 128 bits, DES block size is 64 bits
DES is insecure by modern standards, AES is still considered secure
AES uses a subs-perm network, DES uses Feistel
GCM Mode
When working with DES, we used CBC, or block mode. When working with AES, we'll be using GCM, or Galois Counter Mode. While AES also supports CBC, GCM has a couple of properties that make it more suitable for production use:

There is no need to fuss with block padding, the implementation of GCM handles it for us
GCM supports authenticated encryption. GCM decryption will fail if it wasn't us that encrypted the message in the first place
Assignment
At Passly, the production cipher we use to encrypt password values is AES in GCM mode! Complete the decrypt function.

Create a new cipher block using the key.
Use the cipher block to create a new GCM
Use the GCM (which implements the AEAD interface) to decrypt the ciphertext using aesgcm.Open
Return the plaintext as a []byte
Return any errors that occur without modifying them.
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/rand"
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil

}

// don't touch below this line

func encrypt(key, plaintext, nonce []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func test(key, plaintext, nonce []byte) {
	ciphertext, err := encrypt(key, plaintext, nonce)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Encrypting '%v' with key '%v'...\n", string(plaintext), string(key))
	decryptedText, err := decrypt(key, ciphertext, nonce)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted: '%v'\n", string(decryptedText))
	fmt.Println("========")
}

func generateNonce(length int) []byte {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	return randomBytes
}

func main() {
	rand.Seed(0)
	test(
		[]byte("d00c5215-60f6-4ac4-9648-532b5dad"),
		[]byte("I wonder what he's thinking about me??"),
		generateNonce(12),
	)
	test(
		[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"),
		[]byte("I knew it, Becky has been cheating this whole time!"),
		generateNonce(12),
	)
}

/*




 Run






Encrypting 'I wonder what he's thinking about me??' with key 'd00c5215-60f6-4ac4-9648-532b5dad'...

Decrypted: 'I wonder what he's thinking about me??'

========

Encrypting 'I knew it, Becky has been cheating this whole time!' with key 'db50ecaaa-23ed-43eb-9f8b-6fc5931'...

Decrypted: 'I knew it, Becky has been cheating this whole time!'

========
*/
