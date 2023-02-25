## Better Keys

Keys are typically encoded in software as raw binary data. In Go, this means we usually use a slice of bytes as the data type. For example:

[]byte{0xAA, 0xBB}

0xAA and 0xBB are hexadecimal values, and hex is one of the more common ways to write raw binary in code. They represent the raw binary data 10101010 and 10111011 respectively.

## Assignment

Update the code so that compiles and creates the cipher as intended. Here are the docs for aes.NewCipher()

link: https://pkg.go.dev/crypto/aes#NewCipher

## Cryptology, Cryptography and Cryptanalysis

Refer to the following and answer the questions

Cryptology
The scientific study of cryptography and cryptanalysis. Wikipedia

Cryptology is the science of secret messages. Therefore, anything that has to do with making or breaking codes falls into cryptology’s domain.

Cryptography
Cryptography is the practice and study of techniques for secure communication in the presence of third parties called adversaries.

People often lazily use Cryptography in place of the word cryptology, but in reality, cryptography focuses only on creating secure cryptosystems. For example, the design of RSA would have been cryptography work.

Cryptanalysis
You use Cryptanalysis to breach cryptographic security systems and gain access to the contents of encrypted messages, even if the cryptographic key is unknown.

Cryptanalysis is the inverse of cryptography. It's the study of how to break secret codes, not make them. Having a solid understanding of cryptanalysis is fundamental in cryptography, however, as one must know their enemy.

Honorable Mention – Cryptocurrency
A cryptocurrency is a digital asset designed to work as a medium of exchange that uses strong cryptography to secure financial transactions, control the creation of additional units, and also verify the transfer of assets. Wikipedia

In the past few years, the slang term “crypto” has been hijacked. It used to mean cryptography or cryptanalysis but has recently come to be understood as cryptocurrency.
