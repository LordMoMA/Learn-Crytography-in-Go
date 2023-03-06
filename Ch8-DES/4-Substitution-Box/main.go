/*
Substitution Box
Substitution boxes, also known as s-boxes are a key component of many symmetric key algorithms, particularly block ciphers. The primary goal of an s-box, at least in this context, is to obscure the relationship between the key and the ciphertext. An attacker should not be able to derive information about the key from the ciphertext.

Example
Here's an S-box that maps a 6-bit input into a 4-bit output:

sub box

011011 -> 1001
011111 -> 0110
At the end of the day, an s-box is just a lookup table.

Assignment
Complete the sBox() function. It maps the last 4 bits of a byte down to the last 2 bits of a byte. Use the following lookup table:

00	01	10	11
00	00	10	01	11
01	10	00	11	01
10	01	11	00	10
11	11	01	10	00
sBox(0b0000) -> 0b00
sBox(0b0001) -> 0b10
sBox(0b1111) -> 0b00
sBox(0b0110) -> 0b11 (e.g. 0b00000110 -> 0b00000011)
Each column represents the first two bits of input, the rows represent the second two bits. This is not the same as the table above!

If an input value is outside the range of the table, return an error with the text "invalid input".

Note on the byte values
While we're working with entire byte values, just understand that we only care about the last (smallest in value) 4 and 2 bits for the input and output respectively. For example:

0b0110 -> 0b11 is the same as 0b00000110 -> 0b00000011)

In the binary literal 0b00000011, the b prefix indicates that this is a binary literal, and the digits following the prefix represent the binary value.
In this case, the binary value is 00000011, which is equivalent to the decimal value 3 and the hexadecimal value 0x03. Therefore, b is not a variable or a symbol, but rather part of the syntax used to represent a binary value in Go.
*/

package main

import (
	"fmt"
)

func sBox(b byte) (byte, error) {
	// Define the s-box lookup table
	table := [][]byte{
		{0x00, 0x02, 0x01, 0x03},
		{0x02, 0x00, 0x03, 0x01},
		{0x01, 0x03, 0x00, 0x02},
		{0x03, 0x01, 0x02, 0x00},
	}

	// Check if input value is within the valid range
	// (b & 0xF0) checks if the first 4 bits of the byte b are zero.
	// If they are not zero, then the input value is outside the range of the table and an error should be returned.
	// This is because the lookup table only operates on the last 4 bits of the input byte, and any other value outside this range is not defined in the table.
	if (b & 0xF0) != 0 {
		return 0x00, fmt.Errorf("invalid input")
	}

	// Map the input byte to the output byte using the lookup table
	// the >> operator is used for right shift operation, which shifts the bits of the number to the right by the specified number of positions.
	// In the expression b >> 2, the variable b is being right shifted by 2 positions.
	// This means that the bits of b will be shifted to the right by 2 positions and the bits that fall off the right end will be discarded.
	// For example, if b is 0b11010110, then b >> 2 will be 0b00110101.
	row := (b >> 2) & 0x03
	col := b & 0x03
	return table[row][col], nil
}

// don't touch below this line

func main() {
	for i := 0; i <= 16; i++ {
		b := byte(i)
		subbed, err := sBox(b)
		if err != nil {
			fmt.Printf("Error with input %04b: %v\n", i, err)
			continue
		}
		fmt.Printf("%04b -> %02b\n", i, subbed)
	}
}

// solution 2:

package main

import (
	"fmt"
)

func sBox(b byte) (byte, error) {
    // Define the s-box as a two-dimensional slice
    sbox := [][]byte{
        {0b00, 0b10, 0b01, 0b11},
        {0b10, 0b00, 0b11, 0b01},
        {0b01, 0b11, 0b00, 0b10},
        {0b11, 0b01, 0b10, 0b00},
    }

    // Extract the row and column indices
    row := (b >> 2) & 0b11
    col := b & 0b11

	// did not work here
    // if int(row) >= len(sbox) || int(col) >= len(sbox[0]) {
    //     return 0, fmt.Errorf("invalid input")
    // }

	// Check if the upper 4 bits of the input are nonzero
	if (b & 0b11110000) != 0 {
		return 0x00, fmt.Errorf("invalid input")
	}

    // Look up the output in the s-box
    output := sbox[row][col]

    // Return the output and no error
    return output, nil
}


// don't touch below this line

func main() {
	for i := 0; i <= 16; i++ {
		b := byte(i)
		subbed, err := sBox(b)
		if err != nil {
			fmt.Printf("Error with input %04b: %v\n", i, err)
			continue
		}
		fmt.Printf("%04b -> %02b\n", i, subbed)
	}
}




/*

0000 -> 00

0001 -> 10

0010 -> 01

0011 -> 11

0100 -> 10

0101 -> 00

0110 -> 11

0111 -> 01

1000 -> 01

1001 -> 11

1010 -> 00

1011 -> 10

1100 -> 11

1101 -> 01

1110 -> 10

1111 -> 00

Error with input 10000: invalid input
*/
