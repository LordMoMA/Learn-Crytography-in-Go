## Better Keys

Keys are typically encoded in software as raw binary data. In Go, this means we usually use a slice of bytes as the data type. For example:

[]byte{0xAA, 0xBB}

0xAA and 0xBB are hexadecimal values, and hex is one of the more common ways to write raw binary in code. They represent the raw binary data 10101010 and 10111011 respectively.

## Assignment

Update the code so that compiles and creates the cipher as intended. Here are the docs for aes.NewCipher()

link: https://pkg.go.dev/crypto/aes#NewCipher
