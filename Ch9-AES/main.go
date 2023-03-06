/*
Advanced Encryption Standard
AES is a block cipher just like DES, but there are some key differences:

AES uses 128, 192, or 256 bit keys. DES uses 56 bit keys
AES block size is 128 bits, DES block size is 64 bits
DES is insecure by modern standards, AES is still considered secure
AES uses a subs-perm network, DES uses Feistel
GCM Mode
When working with DES, we used CBC, or block mode. When working with AES, we'll be using GCM, or Galois Counter Mode. While AES also supports CBC, GCM has a couple of properties that make it more suitable for production use:

There is no need to fuss with block padding, the implementation of GCM handles it for us
GCM supports authenticated encryption. GCM decryption will fail if it wasn't us that encrypted the message in the first place
Assignment
At Passly, the production cipher we use to encrypt password values is AES in GCM mode! Complete the decrypt function.

Create a new cipher block using the key.
Use the cipher block to create a new GCM
Use the GCM (which implements the AEAD interface) to decrypt the ciphertext using aesgcm.Open
Return the plaintext as a []byte
Return any errors that occur without modifying them.
*/