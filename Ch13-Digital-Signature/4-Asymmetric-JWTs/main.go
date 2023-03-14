/*
Asymmetric JWTs
JWTs that use HMACs are symmetric signatures. We talked about how they use the same key to sign and validate the tokens. This works well in some scenarios, but let's take the example of third-party authentication, like "Sign in with Google".

ECDSA JWTs
Asymmetric JWTs use a public key to sign the JSON payload, and a private key to validate it, it's similar to asymmetric encryption, but instead of encrypting data, we're signing it.

ECDSA, or Elliptic Curve Digital Signature Algorithm, is a type of asymmetric signature algorithm. It works using the same elliptic curve math that we talked about in the chapter on asymmetric encryption.

Lifecycle of an ECDSA (asymmetric) JWT
Let's take a look at how the Boot.dev web app uses asymmetric JWTs for our "Sign in with Google" button.

User clicks "Sign in with Google" and enters their credentials
If the credentials are valid, Google creates a JWT by signing it with their private key
The signed JWT is given to the client
The client sends the JWT in an HTTP header along with every request that requires authentication
The Boot.dev server uses Google's public key to validate the JWT in each subsequent request
Assignment
"Sign in with Google", "Sign in with GitHub", "Sign in with Twitter", etc. are all examples of third-party authentication. They reduce the friction of creating an account on a new website, so many websites use them.

Passly wants to be an issuer of JWTs. Our massive egos demand that we push clients into using "Sign in with Passly".

Complete the createECDSAMessage function.

createECDSAMessage()
Hash the message using SHA-256
Create a signature of the hashed message using the private key and the SignASN1 function.
Return a new token in the following format:
MESSAGE.signature
Where MESSAGE is the original message, and signature is the signature of the hashed message in lowercase hex.

Keep in mind, this isn't a full JWT, it's an arbitrary message and a signature.

Asymmetric JWT Review
Asymmetric JWTs vs Symmetric JWTs
Asymmetric JWTs use a private key to sign the JSON payload, and a public key to validate it, it's similar to asymmetric encryption. but instead of encrypting data, we're signing it.

Symmetric JWTs use the same key to sign and validate the tokens.

Lifecycle of an Asymmetric JWT\
An example of "Sign in with Google":

User clicks "Sign in with Google" and enters their credentials
If the credentials are valid, Google creates a JWT by signing it with their private key
The signed JWT is given to the client
The client sends the JWT in an HTTP header along with every request that requires authentication
The Boot.dev server uses Google's public key to validate the JWT in each subsequent request
ECDSA is just one signing algorithm that can be used with asymmetric JWTs. There are many others, including RSA.

Lifecycle of a Symmetric JWT
An example of signing into a website with a username and password:

User clicks "Sign in" and enters their credentials
If the credentials are valid, the server creates a JWT by signing it with their private key
The signed JWT is sent to the client
The Boot.dev server uses their private key again to validate the JWT in each subsequent request
*/

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"strings"
)

// my solution:

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write([]byte(message))
	sum := h.Sum(nil)
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, sum)
	if err != nil {
		return "", err
	}
	final := hex.EncodeToString(signature)
	return final, nil
}

// solution 2:

// don't touch below this line

func verifyECDSAMessage(token string, publicKey *ecdsa.PublicKey) error {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return errors.New("invalid token sections")
	}
	sig, err := hex.DecodeString(parts[1])
	if err != nil {
		return err
	}
	hash := sha256.Sum256([]byte(parts[0]))

	valid := ecdsa.VerifyASN1(publicKey, hash[:], sig)
	if !valid {
		return errors.New("invalid signature")
	}
	return nil
}

func test(message string) {
	defer fmt.Println("========")
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Creating token for: '%v'...\n", message)
	token, err := createECDSAMessage(message, privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Verifying token...")
	err = verifyECDSAMessage(token, &privateKey.PublicKey)
	if err != nil {
		fmt.Printf("Token is invalid: %v\n", err)
		return
	}
	fmt.Println("Token is valid!")
}

func main() {
	test("userid:2f9c584e-5d25-4516-a0ed-ddfa6e152006")
	test("userid:0e803af6-292f-4432-a285-84a7591e25a8")
	test("userid:f77e36d6-0edc-44ef-964e-af4a5b1ebd5f")
}

/*

Creating token for: 'userid:2f9c584e-5d25-4516-a0ed-ddfa6e152006'...

Verifying token...

Token is invalid: invalid token sections

========

Creating token for: 'userid:0e803af6-292f-4432-a285-84a7591e25a8'...

Verifying token...

Token is invalid: invalid token sections

========

Creating token for: 'userid:f77e36d6-0edc-44ef-964e-af4a5b1ebd5f'...

Verifying token...

Token is invalid: invalid token sections

========
*/
