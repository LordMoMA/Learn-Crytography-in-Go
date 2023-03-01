/*
Block Ciphers
A block cipher is a deterministic algorithm that operates on fixed-length groups of data, called blocks. Like stream ciphers, block ciphers are a kind of symmetric encryption.

Block ciphers are widely used in real-world applications to encrypt large amounts of data.

Assignment
The Go standard library has built-in support for the AES and DES block ciphers, which we will talk about in more detail later.

We've been asked by leadership to check on the block sizes of each algorithm and report back. Complete the getBlockSize function.

getBlockSize(keyLen, cipherType int) (int, error)
This function accepts a keyLen and cipherType and returns the block size of the cipher along with any errors encountered.

The cipherType is an enum of typeAES or typeDES. Depending on the cipher type, create a new cipher using:

aes.NewCipher or
des.NewCipher
The value of the key passed in doesn't matter here, all that matters is its length.

Return the .BlockSize() of the new cipher, and any error values without changing them if they are encountered. Return an "invalid cipher type" error if the cipherType isn't one of the valid values.

Notes
Notice the relationship (or lack thereof) between the key length and the block size.
It's expected that some of the test cases will result in error messages
*/