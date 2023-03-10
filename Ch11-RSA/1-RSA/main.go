/*
RSA
RSA (which stands for Rivest–Shamir–Adleman, the names of the creators), is another widely used public-key encryption algorithm. ECC is becoming the more modern choice, but RSA is still used in production all across the internet, so it's important to understand how it works.

Assignment
Our web systems at Passly still use RSA, even though we've moved our native desktop encryption to ECC. Complete the encrypt function.

Use the rsa.EncryptOAEP function to encrypt the message with the public key. Use nil as the label because we don't need it. Follow the patterns in the documentation as well as in the decrypt function if you're concerned about the other parameters.

Note on the last test
Because of our simple implementation that just wraps rsa.EncryptOAEP, the last test case is expected to fail for having a message that's too long.

RSA vs ECC
To be clear, RSA is still widely used in production and is still considered secure. Where ECC shines is that it is a bit more efficient in terms of both speed and memory usage. Let's take a look at some of the differences between RSA and ECC.

RSA
RSA key sizes are larger, 2048, 3072, or 4096 bits to remain secure
RSA's security is based on the difficulty of factoring large numbers
RSA is relatively slow. Because of this, it is not commonly used to directly encrypt user data. More often, RSA is used to transmit shared keys for symmetric-key cryptography, as it does in TLS/HTTPS.
ECC
ECC key sizes are smaller, typically 256 bits are enough to remain secure
ECC's security is based on the mathematics of elliptic curves
Comparing key sizes
Security Bits	RSA (bits)	ECC (bits)
80	            1024	    160
112	            2048	    224
128	            3072	    256
192	            7680	    384
256	            15360	    512
To get 2^256 combinations, RSA requires 15360 bits in a private key, while ECC requires only 512 bits. That's a huge difference! You can save a lot of memory and processing power if you're using a lot of keys. That's why systems like Bitcoin use ECC.
*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func encrypt(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	label := []byte(nil)
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, pubKey, msg, label)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// don't touch below this line

func decrypt(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privKey,
		ciphertext,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func genKeys() (pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return &privateKey.PublicKey, privateKey, nil
}

func test(pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, msg string) {
	defer fmt.Println("========")
	fmt.Printf("Encrypting message: '%v'\n", msg)
	ciphertext, err := encrypt(pubKey, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Ciphertext created with length %v\n", len(ciphertext))
	plaintext, err := decrypt(privKey, ciphertext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Plaintext: %v\n", string(plaintext))
}

func main() {
	pub, priv, err := genKeys()
	if err != nil {
		fmt.Println(err)
		return
	}
	test(pub, priv, "Hey Darling, don't come over tonight, I'm out with my people")
	test(pub, priv, "Yes, ten million in cash. No, every penny better be accounted for")
	test(pub, priv, "Do you know what would happen if I suddenly decided to stop going into work? A business big enough that it could be listed on the NASDAQ goes belly up. Disappears! It ceases to exist without me. No, you clearly don't know who you're talking to, so let me clue you in. I am not in danger, Skyler. I am the danger. A guy opens his door and gets shot and you think that of me? No. I am the one who knocks!")
}

/*


Encrypting message: 'Hey Darling, don't come over tonight, I'm out with my people'

Ciphertext created with length 256

Plaintext: Hey Darling, don't come over tonight, I'm out with my people

========

Encrypting message: 'Yes, ten million in cash. No, every penny better be accounted for'

Ciphertext created with length 256

Plaintext: Yes, ten million in cash. No, every penny better be accounted for

========

Encrypting message: 'Do you know what would happen if I suddenly decided to stop going into work? A business big enough that it could be listed on the NASDAQ goes belly up. Disappears! It ceases to exist without me. No, you clearly don't know who you're talking to, so let me clue you in. I am not in danger, Skyler. I am the danger. A guy opens his door and gets shot and you think that of me? No. I am the one who knocks!'

crypto/rsa: message too long for RSA key size

========
*/
