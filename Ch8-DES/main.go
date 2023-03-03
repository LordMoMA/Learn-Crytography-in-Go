/*
Feistel Network
Many block ciphers, including DES which we'll be talking about soon, use a Feistel network (or "Feistel Cipher") as a key component of their encryption algorithms.

Feistel structures have the useful property that encryption and decryption operations are very similar and sometimes identical. Decryption only requires a reversal of the key schedule. This keeps the implementation of the algorithms simple and concise.

Feistel Network

n = The number of ciphering rounds
L0 = Left half of plaintext
R0 = Right half of plaintext
K0 - Kn = The round keys (from a key schedule)
F = The round function (specific to the ciphering algorithm, like DES for example)
One of the most interesting things about Feistel networks is that the round function does NOT have to be reversible.

A Feistel cipher is not a fully-fledged encryption algorithm but is rather a framework that more complete cipher implementations (like DES) utilize.

Assignment
For marketing purposes, Passly has decided to create its own Feistel network. It will use the Go standard library's SHA-256 hash function as the round function.

Here's some psuedocode:

feistel(msg, key []byte, rounds int) []byte
Split the message into equal left and right-hand sides of equal length.
For each round:
nextRHS = xor(lhs, hash(rhs+key))
nextLHS = oldRHS
When you're done with all the rounds, return the concatenation of the right and left-hand sides (right first, then left)
The hash() function is provided for you.

*/

/*
The roundKeys parameter is a slice of slices of bytes ([][]byte), not a slice of bytes ([]byte).

In the feistel function, roundKeys is a two-dimensional slice of bytes where each inner slice is a round key. The outer slice represents all the round keys needed for all the rounds of the Feistel cipher.

For example, if you have a Feistel cipher with 10 rounds, you will need 10 round keys, so roundKeys will be a slice of 10 slices of bytes.

examples:
roundKeys := [][]byte{
    {0x01, 0x02, 0x03, 0x04},
    {0x05, 0x06, 0x07, 0x08},
    {0x09, 0x0A, 0x0B, 0x0C},
}

*/