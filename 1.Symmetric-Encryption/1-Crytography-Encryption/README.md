## Crytography is about Encryption

Simply put, Cryptography provides a method for secure communication. By using cryptography, we can send secret messages to others across an insecure medium (like the internet), and be sure that only the intended recipients will be able to read the message.

In this course, we'll be writing code in Go that makes up part of "Passly", a password manager. To start, let's write a function that will debug the logic that encrypts and decrypts passwords!

We encrypt passwords so that if an attacker gains access to the computer on which the passwords are stored, they still won't be able to read the passwords.

## Assignment

Complete the debugEncryptDecrypt function.

1. Call the provided encrypt function
2. Print the following message to the console: Encrypted password: ENCRYPTED, where ENCRYPTED is the result of the encrypt function.
3. Call the provided decrypt function
4. Print the following message to the console: Decrypted password: DECRYPTED, where DECRYPTED is the result of the decrypt function.

Make sure you terminate both lines with a newline character (\n) if you're using fmt.Printf.
