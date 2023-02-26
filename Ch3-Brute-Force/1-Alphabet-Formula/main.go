Alphabet Formula
The number of total characters in an alphabet depends on how many bits each character in an encoded string can represent.

1 bit -> 2 chars (0 and 1, binary)
2 bits -> 4 chars
4 bits -> 16 chars (i.e. hexadecimal)
7 bits -> 128 chars (i.e. ASCII)
8 bits -> 256 chars
Notice how the number of characters grows exponentially based on the number of bits. The number of characters is equal to the number of combinations that can be created by the individual bits.

2-Bit Example
00 = char 0
01 = char 1
10 = char 2
11 = char 3
So with 2 bits, we can represent 4 characters in total.

Assignment
Complete the alphabetSize() function. It accepts a number of bits per character as input, and returns the size of the alphabet that can be represented by that number of bits.

package main

import (
	"fmt"
	"math"
)

func alphabetSize(numBits int) float64 {
	return math.Pow(2, float64(numBits))
}

// don't touch below this line

func test(num int) {
	fmt.Printf("Alphabet size for %v bits: %v\n", num, alphabetSize(num))
}

func main() {
	for i := 1; i < 17; i++ {
		test(i)
	}
}
