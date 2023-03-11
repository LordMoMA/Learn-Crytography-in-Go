/*
Multiplicative Inverse
The multiplicative inverse of a number, say x, is written as x-1 or 1/x.

It's a number that when multiplied by x equals 1.

Modular multiplicative inverse
In the modular world, a multiplicative inverse of a is x as defined by this formula:

a * x ≡ 1 (mod m)
For example, an inverse of 3 in "mod 11" is 4 because:

3 * 4 ≡ 1 (mod 11)
As another example, an inverse of 5 in "mod 11" is 9 because:

5 * 9 ≡ 1 (mod 11)
Some numbers don't have a multiplicative inverse in a given mod. For example, 2 has no multiplicative inverse in "mod 4" because, in the following equation, there is no x that will result in a congruence:

2 * x ≡ 1 (mod 4)
Who cares?
This will be important later when we talk about decryption, but it's important to understand that this formula:

y^ax (mod m)
can be reduced to:

y^1 (mod m)
which reduces to:

y (mod m)
Of course, this all assumes that a is the modular multiplicative inverse of x in "mod m".

Private Key
While (n, e) is the public key, d is the private key. That said, d can be derived from p and q, so you can also consider the pair of p and q the private key, but it's just different ways to look at it.

At the end of the day, p, q, tot and d should all be kept secret, there's no reason to share anything other than (n, e).

Generating d
d is simply the modular multiplicative inverse of e (mod tot).

Assignment
Complete the getD function.

getD(e, tot *big.Int) *big.Int
Use the ModInverse method to calculate d and return it.
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

// Get the private exponent
func getD(e, tot *big.Int) *big.Int {
	return new(big.Int).ModInverse(e, tot)
}

// don't touch below this line

func encrypt(m, e, n *big.Int) *big.Int {
	c := new(big.Int)
	c.Exp(m, e, n)
	return c
}

func gettot(p, q *big.Int) *big.Int {
	tot := new(big.Int)
	tot.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return tot
}

// Choose a public exponent
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

	d := getD(e, tot)
	fmt.Printf("Generated d: %v it has %v digits\n", firstNDigits(*d, 10), len(d.String()))

	fmt.Println("========")
}

func main() {
	test("I hid the cash under the sink", 512)
	test("Don't you think they will look there??", 512)
	test("They'll look at everything but the kitchen sink", 1024)
}

/*

Selection deleted
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
package main

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
)

// Get the private exponent
func getD(e, tot *big.Int) *big.Int {
	return new(big.Int).ModInverse(e, tot)
}

// don't touch below this line

func encrypt(m, e, n *big.Int) *big.Int {
	c := new(big.Int)
	c.Exp(m, e, n)
	return c
}

func gettot(p, q *big.Int) *big.Int {
	tot := new(big.Int)
	tot.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return tot
}

// Choose a public exponent
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
}

Submit



 Run






Encrypting 'I hid the cash under the sink' with key size 512

Generated p: 1244436454... it has 155 digits

Generated q: 1198026774... it has 155 digits

Generated n: 1490868191... it has 309 digits

Generated tot: 1490868191... it has 309 digits

Generated e: 6905949906... it has 308 digits

Encrypted message as a number: 1093960760...

Generated d: 1594991729... it has 308 digits

========

Encrypting 'Don't you think they will look there??' with key size 512

Generated p: 1131593906... it has 155 digits

Generated q: 1117516671... it has 155 digits

Generated n: 1264575056... it has 309 digits

Generated tot: 1264575056... it has 309 digits

Generated e: 6138527131... it has 308 digits

Encrypted message as a number: 5612332892...

Generated d: 1857109338... it has 308 digits

========

Encrypting 'They'll look at everything but the kitchen sink' with key size 1024

Generated p: 1684475698... it has 309 digits

Generated q: 1502288675... it has 309 digits

Generated n: 2530568766... it has 617 digits

Generated tot: 2530568766... it has 617 digits

Generated e: 1554307925... it has 616 digits

Encrypted message as a number: 1551668194...

Generated d: 1585633466... it has 617 digits

========
*/
