/*
What is a Message Integrity?
Message integrity ensures that a message has not been changed.

Let's say you need to transmit a private key from your computer to your phone. Because your network connection isn't great, there's a chance that information is lost in transit. You want to be sure that the private key makes it to your phone in pristine condition. If there is data loss, you'll want to know about it so you can try again.

To accomplish this, we can take a hash of the private key, and send it alongside the book. When the private key is transmitted, you'll be able to check on your phone by hashing the private key and checking if the hashes match.

Checksums
This "hash and check" method is called a checksum.

checksum = hash(data)
send data
send checksum
if (hash(data) == checksum) then the message has not been altered
Assignment
Complete the checksumMatches function.

It should hash the message after converting it directly to a slice of bytes using sha256. Check if the lowercase hexadecimal encoding of the hash matches the checksum argument. If it does, return true. Otherwise, return false.
*/
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func checksumMatches(message string, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message))
	sum := h.Sum(nil)
	hexStr := hex.EncodeToString(sum)
	if hexStr == checksum {
		return true
	} else {
		return false
	}
}

// don't touch below this line

func test(message, checksum string) {
	fmt.Printf("Checking message '%v'...\n", message)
	fmt.Printf("Expected checksum: %v\n", checksum)
	if checksumMatches(message, checksum) {
		fmt.Println("Checksum matches!")
	} else {
		fmt.Println("Checksum does not match!")
	}
	fmt.Println("========")
}

func main() {
	test("pa$$w0rd", "4b358ed84b7940619235a22328c584c7bc4508d4524e75231d6f450521d16a17")
	test("buil4WithB1ologee", "1c489a153271aaf3b234aa154b1a2eef5248eb9ab402e4d3c8b7bc3d81fed1a8")
	test("br3ak1ngB@d1sB3st", "5d178e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546")
	test("b3ttterC@llS@ulI$B3tter", "8d42f2dc81476123974619969a42b27b8d8a4fa507be99c9623f614ad2d859f7")
}

/*

Checking message 'pa$$w0rd'...

Expected checksum: 4b358ed84b7940619235a22328c584c7bc4508d4524e75231d6f450521d16a17

Checksum matches!

========

Checking message 'buil4WithB1ologee'...

Expected checksum: 1c489a153271aaf3b234aa154b1a2eef5248eb9ab402e4d3c8b7bc3d81fed1a8

Checksum does not match!

========

Checking message 'br3ak1ngB@d1sB3st'...

Expected checksum: 5d178e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546

Checksum does not match!

========

Checking message 'b3ttterC@llS@ulI$B3tter'...

Expected checksum: 8d42f2dc81476123974619969a42b27b8d8a4fa507be99c9623f614ad2d859f7

Checksum matches!

========
*/
