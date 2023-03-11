/*
Modular
Modular arithmetic and modular exponentiation are widely used in cryptography, and RSA is no exception.

What is "mod"?
In modular arithmetic, we're working with a set of numbers that are all less than a certain number, and we call that the "mod". It's similar to the modulus operator, %, but instead of returning the remainder, we're working within the world of the possible remainders.

The world of "mod 2"
0 in the world of "mod 2" is "congruent" with 0. So we would write:

0 ≡ 0 (mod 2)
1 ≡ 1 (mod 2)
2 ≡ 0 (mod 2)
3 ≡ 1 (mod 2)
4 ≡ 0 (mod 2)
5 ≡ 1 (mod 2)
...
≡ is the congruence symbol.

The world of "mod 3"
As you saw above, in the world of "mod 2", there are only 2 possible values: {0, 1}. However, in the world of "mod 3", there are 3 possible values: {0, 1, 2}.

1 ≡ 1 (mod 3)
2 ≡ 2 (mod 3)
3 ≡ 0 (mod 3)
4 ≡ 1 (mod 3)
5 ≡ 2 (mod 3)
...

Modular Arithmetic
It will be important to understand a simple transformation that's possible with modular arithmetic. There 2 formulas are effectively equivalent:

n ≡ r (mod q)
n = qk + r
Where n is the number we're working with, r is the remainder, q is the quotient, and k is the number of times q goes into n. This allows us to convert an equation in the world of modular arithmetic into a more familiar equation the normal world.

For example,

9 ≡ 1 (mod 2)
Can be converted to:

9 = 2k + 1
This will be important later.

Encryption
The public key
Aside from the message itself, all we need to perform encryption are the numbers e and n. Together, these are the public key.

The security of RSA relies on the fact that given n, it's really hard to guess what p and q (the private keys) were. Finding prime factors of very large numbers is computationally expensive. Trying to brute-force p and q would take more than trillions of years on modern hardware.

Remember, n = p * q.

The encryption formula
The math for encrypting a message with RSA follows this formula:

ciphertext = m^e (mod n)

m = message
e = public key exponent
n = public key modulus
We'll talk more about this math later, but for now, let's just implement it.

Assignment
Complete the encrypt function.

encrypt(m, e, n *big.Int) *big.Int
Return the result of m^e (mod n). Use the .Exp method.
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

func encrypt(m, e, n *big.Int) *big.Int {
	return new(big.Int).Exp(m, e, n)
}

// don't touch below this line

func gettot(p, q *big.Int) *big.Int {
	tot := new(big.Int)
	tot.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return tot
}

func getE(tot *big.Int) *big.Int {
	totMinusTwo := new(big.Int)
	totMinusTwo.Sub(tot, big.NewInt(1))

	e, _ := crand.Int(randReader, totMinusTwo)
	e.Add(e, big.NewInt(2))
	for gcd(e, tot).Cmp(big.NewInt(1)) != 0 {
		e, _ = crand.Int(randReader, totMinusTwo)
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

	tot := gettot(p, q)
	fmt.Printf("Generated tot: %v it has %v digits\n", firstNDigits(*tot, 10), len(tot.String()))

	e := getE(tot)
	fmt.Printf("Generated e: %v it has %v digits\n", firstNDigits(*e, 10), len(e.String()))

	plaintext := big.NewInt(0)
	plaintext.SetBytes(msgBytes)
	ciphertext := encrypt(plaintext, e, n)
	fmt.Printf("Encrypted message as a number: %v\n", firstNDigits(*ciphertext, 10))

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

Generated tot: 1490868191... it has 309 digits

Generated e: 6905949906... it has 308 digits

Encrypted message as a number: 1093960760...

========

Encrypting 'Don't you think they will look there??' with key size 512

Generated p: 1131593906... it has 155 digits

Generated q: 1117516671... it has 155 digits

Generated n: 1264575056... it has 309 digits

Generated tot: 1264575056... it has 309 digits

Generated e: 6138527131... it has 308 digits

Encrypted message as a number: 5612332892...

========

Encrypting 'They'll look at everything but the kitchen sink' with key size 1024

Generated p: 1684475698... it has 309 digits

Generated q: 1502288675... it has 309 digits

Generated n: 2530568766... it has 617 digits

Generated tot: 2530568766... it has 617 digits

Generated e: 1554307925... it has 616 digits

Encrypted message as a number: 1551668194...

========
*/
