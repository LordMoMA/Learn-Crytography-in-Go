/*
Decryption
Decryption in RSA is almost unbelievably simple from a mathematics perspective but remains extremely secure.

The only known way to efficiently crack a secure RSA system is with a $5 wrench.

Assignment
Complete the decrypt function.

decrypt(c, d, n *big.Int) *big.Int
Use the big.Int.Exp function to raise the encrypted message (c) to the power of the private key (d) within (mod n).
*/

package main

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
)

func decrypt(c, d, n *big.Int) *big.Int {
	return new(big.Int).Exp(c, d, n)
}

// don't touch below this line

func getD(e, phi *big.Int) *big.Int {
	d := new(big.Int)
	d.ModInverse(e, phi)
	return d
}

func encrypt(m, e, n *big.Int) *big.Int {
	c := new(big.Int)
	c.Exp(m, e, n)
	return c
}

func getPhi(p, q *big.Int) *big.Int {
	phi := new(big.Int)
	phi.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return phi
}

func getE(phi *big.Int) *big.Int {
	phiMinusTwo := new(big.Int)
	phiMinusTwo.Sub(phi, big.NewInt(1))

	e, _ := crand.Int(randReader, phiMinusTwo)
	e.Add(e, big.NewInt(2))
	for gcd(e, phi).Cmp(big.NewInt(1)) != 0 {
		e, _ = crand.Int(randReader, phiMinusTwo)
		e.Add(e, big.NewInt(2))
	}
	return e
}

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, _ := getBigPrime(keysize)
	q, _ := getBigPrime(keysize)
	return p, q
}

func getN(p, q *big.Int) *big.Int {
	n := new(big.Int)
	n.Mul(p, q)
	return n
}

func gcd(x, y *big.Int) *big.Int {
	xCopy := new(big.Int).Set(x)
	yCopy := new(big.Int).Set(y)
	for yCopy.Cmp(big.NewInt(0)) != 0 {
		xCopy, yCopy = yCopy, xCopy.Mod(xCopy, yCopy)
	}
	return xCopy
}

func firstNDigits(n big.Int, numDigits int) string {
	if len(n.String()) < numDigits {
		return fmt.Sprintf("%v", n.String())
	}
	return fmt.Sprintf("%v...", n.String()[:numDigits])
}

var randReader = mrand.New(mrand.NewSource(0))

func getBigPrime(bits int) (*big.Int, error) {
	if bits < 2 {
		return nil, errors.New("prime size must be at least 2-bit")
	}
	b := uint(bits % 8)
	if b == 0 {
		b = 8
	}
	bytes := make([]byte, (bits+7)/8)
	p := new(big.Int)
	for {
		if _, err := io.ReadFull(randReader, bytes); err != nil {
			return nil, err
		}
		bytes[0] &= uint8(int(1<<b) - 1)
		if b >= 2 {
			bytes[0] |= 3 << (b - 2)
		} else {
			bytes[0] |= 1
			if len(bytes) > 1 {
				bytes[1] |= 0x80
			}
		}
		bytes[len(bytes)-1] |= 1
		p.SetBytes(bytes)
		if p.ProbablyPrime(20) {
			return p, nil
		}
	}
}

func test(msg string, keySize int) {
	msgBytes := []byte(msg)

	fmt.Printf("Encrypting '%v' with key size %v\n", msg, keySize)

	p, q := generatePrivateNums(keySize)
	fmt.Printf("Generated p: %v it has %v digits\n", firstNDigits(*p, 10), len(p.String()))
	fmt.Printf("Generated q: %v it has %v digits\n", firstNDigits(*q, 10), len(q.String()))

	n := getN(p, q)
	fmt.Printf("Generated n: %v it has %v digits\n", firstNDigits(*n, 10), len(n.String()))

	phi := getPhi(p, q)
	fmt.Printf("Generated phi: %v it has %v digits\n", firstNDigits(*phi, 10), len(phi.String()))

	e := getE(phi)
	fmt.Printf("Generated e: %v it has %v digits\n", firstNDigits(*e, 10), len(e.String()))

	plaintext := big.NewInt(0)
	plaintext.SetBytes(msgBytes)
	ciphertext := encrypt(plaintext, e, n)
	fmt.Printf("Encrypted message as a number: %v\n", firstNDigits(*ciphertext, 10))

	d := getD(e, phi)
	fmt.Printf("Generated d: %v it has %v digits\n", firstNDigits(*d, 10), len(d.String()))

	decrypted := decrypt(ciphertext, d, n)
	fmt.Printf("Decrypted message: '%s'\n", string(decrypted.Bytes()))

	fmt.Println("========")
}

func main() {
	test("I hid the cash under the sink", 512)
	test("Don't you think they will look there??", 512)
	test("They'll look at everything but the kitchen sink", 1024)
}

/*

Encrypting 'I hid the cash under the sink' with key size 512

Generated p: 1244436454... it has 155 digits

Generated q: 1198026774... it has 155 digits

Generated n: 1490868191... it has 309 digits

Generated phi: 1490868191... it has 309 digits

Generated e: 6905949906... it has 308 digits

Encrypted message as a number: 1093960760...

Generated d: 1594991729... it has 308 digits

Decrypted message: 'I hid the cash under the sink'

========

Encrypting 'Don't you think they will look there??' with key size 512

Generated p: 1131593906... it has 155 digits

Generated q: 1117516671... it has 155 digits

Generated n: 1264575056... it has 309 digits

Generated phi: 1264575056... it has 309 digits

Generated e: 6138527131... it has 308 digits

Encrypted message as a number: 5612332892...

Generated d: 1857109338... it has 308 digits

Decrypted message: 'Don't you think they will look there??'

========

Encrypting 'They'll look at everything but the kitchen sink' with key size 1024

Generated p: 1684475698... it has 309 digits

Generated q: 1502288675... it has 309 digits

Generated n: 2530568766... it has 617 digits

Generated phi: 2530568766... it has 617 digits

Generated e: 1554307925... it has 616 digits

Encrypted message as a number: 1551668194...

Generated d: 1585633466... it has 617 digits

Decrypted message: 'They'll look at everything but the kitchen sink'

========
*/
