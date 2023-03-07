/*
Asymmetric cryptography, also known as "public-key cryptography", has the same goal as symmetric cryptography: to encrypt and decrypt data.
However, asymmetric cryptography uses two different keys to do it.

Each pair of keys has a public key and a corresponding private key. The security of public-key cryptography depends on keeping the private key secret. The public key can be given freely to anyone without compromising security.

Assignment
At Passly, we use Elliptic Curve Cryptography (ECC) for our public-key cryptography needs in production.

Complete the genKeys() function. Use the ecdsa.GenerateKey function from the standard library, along with the elliptic.P256() curve.

We won't be checking the values of the keys, so feel free to use a secure random source from the crypto/rand package.

https://blog.boot.dev/cryptography/elliptic-curve-cryptography/

Private Key: Is kept secret and is used to decrypt messages
Public Key: Is shared with the world and is used to encrypt messages

Why two keys?
With symmetric encryption, if I want to send a message to my friend, I need to give them a copy of my key. Not only is it hard to securely transmit a key to them, but I also might not want to trust them with my encryption keys.

Public key cryptography makes it easy to communicate securely over an untrusted medium (like the internet) with other parties. Symmetric encryption is generally only used to encrypt data for oneself.
*/

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {
	// Generate a private key using the P256 curve
	privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Get the public key from the private key
	pubKey = &privKey.PublicKey

	// Return the keys and no error
	return pubKey, privKey, nil
}

// don't touch below this line

func keysArePaired(pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey) bool {
	msg := "a test message"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		return false
	}

	return ecdsa.VerifyASN1(pubKey, hash[:], sig)
}

func test(i int) {
	fmt.Printf("Generating key pair %v...\n", i)
	pub, priv, err := genKeys()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Verifying key pair...")
	arePaired := keysArePaired(pub, priv)
	fmt.Printf("Keys are securely paired: %v\n", arePaired)
	fmt.Println("========")
}

func main() {
	for i := 1; i < 4; i++ {
		test(i)
	}
}

/*

Generating key pair 1...

Verifying key pair...

Keys are securely paired: true

========

Generating key pair 2...

Verifying key pair...

Keys are securely paired: true

========

Generating key pair 3...

Verifying key pair...

Keys are securely paired: true

========
*/
