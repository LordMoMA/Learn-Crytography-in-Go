/*
Data Encryption Standard
The Data Encryption Standard is an example of a symmetric-key block cipher that utilizes a Feistel network.

DES was developed in the early 1970s at IBM and in 1977 a slightly modified version was published as an official Federal Information Processing Standard for the United States.

In 1997 the DESCHALL Project was the first group to publicly break DES encryption and won $10,000 for their efforts! DES was broken by a simple brute force attack, which is possible due to its small key size of just 56 bits.

What's an iv?
An IV, or initialization vector, is a random value that is used to initialize a block cipher. It is used to ensure that the same plaintext always encrypts to a different ciphertext. Without an IV, the same plaintext would always encrypt to the same ciphertext which is a big security vulnerability.

Assignment
For Cryptanalysis purposes, Passly keeps a DES implementation around so that our engineers can practice breaking it. Complete the encrypt function. It should use the standard library's crypto/des package to encrypt the plaintext with the given key.

Complete the encrypt function and its helper padMsg function. The decrypt function is already written for you.

padMsg(plaintext []byte, blockSize int) []byte
The padWithZeros function is provided for you, but it only pads a single block. You'll need to find the last block in the message and pad that one. Essentially you need to ensure that the entire message length is a multiple of the block size.

encrypt(key, plaintext []byte) ([]byte, error)
We'll be using DES in CBC mode. Here's an example from the Go documentation that shows how to encrypt a message.

Create a new cipher block
Pad the plaintext with zeros using padMsg
Generate a random iv and append it to the beginning of the ciphertext. It should be the same length as the block size.
Create a new encrypter
Encrypt the blocks and return the entire ciphertext
Return any errors that occur.

Tip
Don't be afraid to debug your code by printing out the values of different variables. Just be sure to remove any debug code before submitting your answer.
*/