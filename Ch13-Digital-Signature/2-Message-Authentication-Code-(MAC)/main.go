/*
Message Authentication Code (MAC)
A MAC takes advantage of the irreversibility property of hash functions. A MAC is the hash of a message concatenated with a key.

mac = hash(message + key)
As with a checksum, a mac can be sent along with the message in a communication. A man in the middle won't be able to alter the message secretly, because they can't produce a new valid mac without access to the secret key.

mac = hash(message + key)
send mac
send message
if mac == hash(message + key) then the message is valid
The disadvantage of a MAC is that the receiver will also need a copy of the secret key. This is why MACs are generally only used when the sender and the receiver are the same entity.

Assignment
Fix the macMatches function. It currently uses a plain checksum. Concatenate the key to the end of the message before hashing it.

Hash-Based Message Authentication Code
An HMAC is a kind of MAC. All HMACs are MACs but not all MACs are HMACs. The main difference is that an HMAC uses two rounds of hashing instead of one (or none). Each round of hashing uses a child key that's derived from the secretKey.

Here's a naive implementation:

secretKey = 'thisIsASecretKey1234'
childKey1 = 'thisIsASe'
childKey2 = 'cretKey1234'
hash(childKey1 + hash(childKey2 + 'the message we want to send'))
This is a simplified version of the function given in RFC-2104.

Why use HMAC? Why do we need to hash twice?
With some MACs, depending on the hash function, it is possible to change the message (without knowing the key) and obtain another valid MAC. This is called a length extension attack. There are no known extension attacks against the current HMAC specification, so you should prefer HMACs over MACs.

Note about the child keys
In the naive example above, we created 2 child keys by splitting the original key. That's not the way a secure HMAC would be implemented, instead, we would derive child keys using a slightly more complex process. That said, the principle is the same: using a single key we can derive two separate child keys.

If you're curious about how that might work in production, you can check out the implementation here.
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func macMatches(message, key, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message + key))
	return checksum == fmt.Sprintf("%x", h.Sum(nil))
}

// don't touch below this line

func test(message, key, checksum string) {
	fmt.Printf("Checking message '%v' with key '%v'...\n", message, key)
	fmt.Printf("Expected checksum: %v\n", checksum)
	if macMatches(message, key, checksum) {
		fmt.Println("Checksum matches!")
	} else {
		fmt.Println("Checksum does not match!")
	}
	fmt.Println("========")
}

func main() {
	test("pa$$w0rd", "abdf6b86cb", "7b1dede90198e94c7432358e9bd937b91015cc45c84af5dfbedcd1a3c764ff71")
	test("buil4WithB1ologee", "6ddf6b86cb", "1cddede90198e94c7432358e9bd937b91015cc45c84af5dfbedcd1a3c764ff71")
	test("br3ak1ngB@d1sB3st", "7adf6b86cb", "2c678e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546")
	test("b3ttterC@llS@ulI$B3tter", "12df6b86cb", "eb4d4516bd4141322c3ab160bc2b3010eaf7bd19f821d0c48f480791d32af359")
}

/*

Checking message 'pa$$w0rd' with key 'abdf6b86cb'...

Expected checksum: 7b1dede90198e94c7432358e9bd937b91015cc45c84af5dfbedcd1a3c764ff71

Checksum matches!

========

Checking message 'buil4WithB1ologee' with key '6ddf6b86cb'...

Expected checksum: 1cddede90198e94c7432358e9bd937b91015cc45c84af5dfbedcd1a3c764ff71

Checksum does not match!

========

Checking message 'br3ak1ngB@d1sB3st' with key '7adf6b86cb'...

Expected checksum: 2c678e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546

Checksum does not match!

========

Checking message 'b3ttterC@llS@ulI$B3tter' with key '12df6b86cb'...

Expected checksum: eb4d4516bd4141322c3ab160bc2b3010eaf7bd19f821d0c48f480791d32af359

Checksum matches!

========
*/
