package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func keyToCipher(key string) (cipher.Block, error) {
	return aes.NewCipher(key)
}

// don't touch below this line

func main() {
	const symmetricKey = "thisIsMySecretKeyIHopeNoOneFinds"
	cipher, err := keyToCipher(symmetricKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Cipher created successfully with block size: %v\n", cipher.BlockSize())
}
