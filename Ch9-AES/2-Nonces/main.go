/*
Nonces
We talked a bit about nonces when we talked about IVs, but let's take a closer look at them.

A nonce is an arbitrary number that must be used only once in a cryptographic communication. Nonces are not secret and can be transported publicly.

Assignment
Because we use AES-256 at Passly, we're very concerned about using truly random nonces so we don't accidentally reuse once. Write the nonceStrength function. It should return the strength of a nonce as an integer.

We refer to the strength of a nonce as the total number of possible combinations of bits that could exist in the nonce.

Example
0b10010100
0b00010001
0b11000000
Each nonce above has a strength of 256 because there are 256 possible combinations of bits that could exist in an 8-bit nonce.
*/