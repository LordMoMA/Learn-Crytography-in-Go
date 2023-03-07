/*
RSA
RSA (which stands for Rivest–Shamir–Adleman, the names of the creators), is another widely used public-key encryption algorithm. ECC is becoming the more modern choice, but RSA is still used in production all across the internet, so it's important to understand how it works.

Assignment
Our web systems at Passly still use RSA, even though we've moved our native desktop encryption to ECC. Complete the encrypt function.

Use the rsa.EncryptOAEP function to encrypt the message with the public key. Use nil as the label because we don't need it. Follow the patterns in the documentation as well as in the decrypt function if you're concerned about the other parameters.

Note on the last test
Because of our simple implementation that just wraps rsa.EncryptOAEP, the last test case is expected to fail for having a message that's too long.
*/