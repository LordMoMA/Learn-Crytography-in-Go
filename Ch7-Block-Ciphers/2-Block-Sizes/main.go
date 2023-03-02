/*
Block Sizes
Block ciphers use fixed-size blocks and fixed-size keys. However, just because the blocks are fixed-size doesn't mean the message must be a fixed size as well. Block ciphers can operate on messages of any length, and the key doesn't even need to be the same length as the message or the block.

For example:

3DES: blocksize=64 bits, keysize=168 bits
AES-256: blocksize=128 bits, keysize=256 bits
Chunks and Padding
Messages in a block cipher are broken up into blocks, and each block is encrypted separately.

For example, let's say we are using a cipher that requires 256-bit blocks. We have a message of 650 bits that we want to encrypt. The algorithm would break that message up into three blocks:

block1 = first 256 bits
block2 = next 256 bits
block3 = last 138 bits
The last block is then padded with extra garbage bits so that it also has 256 bits of data. That padding is stripped off when the message is decrypted.

Assignment
While robust cryptography libraries take care of block padding for us, it's important to understand how it works. As part of employee onboarding at Passly, we require our cryptography engineers to write a padding function from scratch so they can understand how it works.

Complete the padWithZeros function. It takes a message (block) as a []byte and a desired output size. It should increase the size of block by adding zero-value-bytes until it has a length of desiredSize. Return the new byte slice.
*/