/*
Totient and E
Now that we have p, q and n, we need to calculate:

The totient of n, which we'll call tot.
A random number that is relatively prime to tot, which we'll call e
The public key for RSA encryption is the pair of numbers: (n, e)

The totient function (phi)
Euler's totient function counts the positive integers up to a given integer, n in our case, that are relatively prime to it.

In other words, the totient is the number of integers between 2 and n whose greatest common divisor is 1.

tot = ϕ(n) = (p - 1) * (q - 1)

Remember p and q are prime, which means their totient's are just p-1 and q-1. Because we know n = p * q, we know the totient of n is (p - 1) * (q - 1).

e
e is a random number between 1 and tot that is relatively prime to tot. This means that the greatest common divisor of e and tot is 1.

Assignment
Complete the getTot and getE functions.

getTot(p, q *big.Int) *big.Int
Use the math/big package to calculate (p-1)(q-1) and return it as a pointer to a big.Int. This is the "totient" of n, which we can also call "phi of n", or ϕ(n).

getE(tot *big.Int) *big.Int
Use the math/big package to generate a random number e that adheres to the following constraints:

e is greater than 1
e is less than tot
e and tot have a greatest common divisor of 1
The gcd function is provided for you. It calculates the greatest common divisor of two big ints.

Generate random e values in the range of [2, tot) until you find one that satisfies the constraints. Use crand.Int to generate random big ints. You'll need to do some manual arithmetic to get the range you want because crand.Int only generates random numbers in the range of [0, max)
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

// Calculate ϕ(n) = (p-1)(q-1)
func getTot(p, q *big.Int) *big.Int {
	// Compute (p-1)(q-1)
	pMinusOne := new(big.Int).Sub(p, big.NewInt(1))
	qMinusOne := new(big.Int).Sub(q, big.NewInt(1))
	tot := new(big.Int).Mul(pMinusOne, qMinusOne)

	return tot
}

// Choose a public exponent
func getE(tot *big.Int) *big.Int {
	// Generate a random integer in the range [2, tot-1]
	two := big.NewInt(2)
	max := new(big.Int).Sub(tot, two)
	for {
		// Generate a random integer in the range [2, tot-1]
		e, _ := crand.Int(randReader, max)
		e.Add(e, two)
		if gcd(e, tot).Cmp(big.NewInt(1)) == 0 {
			return e
		}
	}
}

//// solution2

// totMinusTwo := new(big.Int)
// totMinusTwo.Sub(tot, big.NewInt(1))

// e, _ := crand.Int(randReader, totMinusTwo)
// e.Add(e, big.NewInt(2))
// for gcd(e, tot).Cmp(big.NewInt(1)) != 0 {
// 	e, _ = crand.Int(randReader, totMinusTwo)
// 	e.Add(e, big.NewInt(2))
// }
// return e

// don't touch below this line

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

func test(keySize int) {
	p, q := generatePrivateNums(keySize)
	fmt.Printf("Generated p: %v it has %v digits\n", firstNDigits(*p, 10), len(p.String()))
	fmt.Printf("Generated q: %v it has %v digits\n", firstNDigits(*q, 10), len(q.String()))

	n := getN(p, q)
	fmt.Printf("Generated n: %v it has %v digits\n", firstNDigits(*n, 10), len(n.String()))

	tot := getTot(p, q)
	fmt.Printf("Generated tot: %v it has %v digits\n", firstNDigits(*tot, 10), len(tot.String()))

	e := getE(tot)
	fmt.Printf("Generated e: %v it has %v digits\n", firstNDigits(*e, 10), len(e.String()))
	fmt.Println("========")
}

func main() {
	test(512)
	test(1024)
}

/*

Generated p: 1244436454... it has 155 digits

Generated q: 1198026774... it has 155 digits

Generated n: 1490868191... it has 309 digits

Generated tot: 1490868191... it has 309 digits

Generated e: 6905949906... it has 308 digits

========

Generated p: 1684475698... it has 309 digits

Generated q: 1502288675... it has 309 digits

Generated n: 2530568766... it has 617 digits

Generated tot: 2530568766... it has 617 digits

Generated e: 1554307925... it has 616 digits

========
*/
