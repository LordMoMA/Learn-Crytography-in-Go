/*
Crack an Insecure Key
At Passly we practice Cryptanalysis, meaning that we try to break our own encryption algorithms. We do this to make sure that our encryption is secure!

Assignment
Complete the findKey function. It accepts the encrypted and decrypted versions of messages as arguments. It should return the key that was used to encrypt the password.

If a key is found, return it and a nil error. Otherwise, return an empty byte slice and an error that reads key not found.

Steps
Iterate over all the numbers from 0 to 2^24.
Convert each number to a byte slice (the next key to try) using the intToBytes function.
Decrypt the encrypted message using the current key and the crypt function.
If the decrypted value matches the expected decrypted value (which was passed in as an argument), you found the key! (You can cast the byte slice directly to a string to compare it to the expected value)
Keep in mind, this code might take a while to run, after all, you're brute forcing 24 bits of data!

*/

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	for i := 0; i < int(math.Pow(2, 24)); i++ {
		key := intToBytes(i)
		decryptedMsg := crypt([]byte(encrypted), key)

		if bytes.Equal([]byte(decrypted), decryptedMsg) {
			return key, nil
		}
	}
	return nil, fmt.Errorf("key not found")
}

// don't touch below this line

func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		fmt.Println("Error in intToBytes:", err)
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}

func test(encrypted []byte, decrypted string) {
	fmt.Printf("Encrypted: %x, decrypted: %s\n", []byte(encrypted), decrypted)
	fmt.Println("Starting brute force search...")
	key, err := findKey(encrypted, decrypted)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Key found: %x\n", key)
	fmt.Println("========")
}

func main() {
	test([]byte{0x1b, 0x2c, 0x3d}, "yes")
	test([]byte{0x2a, 0xff, 0xea}, "car")
	test([]byte{0x7d, 0x31, 0x32}, "she")
}
