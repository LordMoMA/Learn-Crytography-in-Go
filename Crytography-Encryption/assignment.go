package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
)

func debugEncryptDecrypt(masterKey, iv, password string) {
	// ?
}

// don't touch below this line

func main() {
	const masterKey = "kjhgfdsaqwertyuioplkjhgfdsaqwert"
	const iv = "1234567812345678"

	test(masterKey, iv, "k33pThisPasswordSafe")
	test(masterKey, iv, "12345")
	test(masterKey, iv, "thePasswordOnMyLuggage")
	test(masterKey, iv, "pizza_the_HUt")
}

func test(masterKey, iv, password string) {
	debugEncryptDecrypt(masterKey, iv, password)
	fmt.Println("========")
}

func encrypt(plainText, key, iv string) string {
	bytes := []byte(plainText)
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println(err)
		return ""
	}
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	stream.XORKeyStream(bytes, bytes)
	return fmt.Sprintf("%x", bytes)
}

func decrypt(cipherText, key, iv string) string {
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println(err)
		return ""
	}
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	bytes, err := hex.DecodeString(cipherText)
	if err != nil {
		log.Println(err)
		return ""
	}
	stream.XORKeyStream(bytes, bytes)
	return string(bytes)
}
