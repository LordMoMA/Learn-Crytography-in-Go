/*
Hash Functions
Hash functions have 3 basic goals:

Hash functions scramble data deterministically
No matter the input, the output of a hash function always has the same size
The original data can not be retrieved from the scrambled data (one-way function)

Checking integrity
At Passly, we use SHA-256 for many things, but one of the most important is to ensure the integrity of a password vault. Each time we save a vault, we hash the vault's contents and store the hash in our database. Later, if we need to verify the integrity of the vault, we can hash the vault's contents again and compare the two hashes. If they match, we know the vault has not been tampered with.

Assignment
Add the following functions and methods to the program:

newHasher
h.Write
h.GetHex
newHasher
Returns a pointer to a new hasher. Uses sha256.New() to create a new hash.Hash.

h.Write
A method on a pointer to a hasher. Uses h.Write() to write data to the hasher. It should accept a string and cast the string to a []byte. It should pass along the return values, that is, it returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early.

h.GetHex
A method on a pointer to a hasher. Uses h.Sum() to get the hash value of the data written to the hasher. It should encode the hash value as a lowercase hex string and return it.
*/

