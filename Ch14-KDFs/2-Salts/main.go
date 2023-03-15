/*
Salts
Rainbow Tables
Rainbow tables are a way for attackers to get around slow KDFs. Attackers can pre-hash all of the common passwords once, then compare this list (called a "rainbow table") to the list of hashes in a compromised database and see if any match.

If the hashes match, then the attacker will know the original password, which they might then be able to use to log in to the user's bank account (or any other place they use the same password).

Salts to the rescue
A salt is a random chunk of data added to a password before it is hashed so that its output hash will differ from the hash of the same password with a different salt.

For example:

digest := hash(password+salt)
// save the digest AND salt
// checking the hash now involves the salt, hash, and the password to check
A rainbow table is ineffective against a salted database, as long as the salt generated for each password is unique.

Assignment
We have an existing database of users and their passwords. We're working with a third-party vendor, and unfortunately, they didn't salt their passwords and they're using SHA-256 instead of a slow KDF. We can only tackle one problem at a time, so for now, we've just been asked to salt and re-hash the passwords. We can migrate from SHA-256 to Bcrypt later.

To do that, we need all of our users to reset their passwords, because we don't store the plaintext.

Complete the generateSalt and hashPassword functions.

generateSalt
Use crypto/rand to generate a random salt of the specified length. Use rand.Read().

hashPassword
Append the salt directly to the end of the password, then hash it with SHA-256. Use crypto/sha256. Return the result of the hash.

Why are KDFs slow?
SHA-256 and MD5 are examples of fast hashes and are bad for storing password hashes. KDFs like Scrypt and Bcrypt are both extremely slow and are good for passwords.

Always use slow KDFs for storing user passwords, never fast hashes.

If an attacker gains access to a database of password hashes made with fast algorithms it is easy for them to "reverse" the hashes by guessing different inputs and seeing if the outputs match. For example, let's say the attacker finds the following entry in a database:

Email	Password Hash
bob@example.com	ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f
They can try hashing common passwords like:

password -> 0b14d501a594442a01c6859541bcb3e8164d183d32937b851835442f69d5c94e
01011994 -> 6cf615d5bcaac778352a8f1f3360d23f02f34ec182e259897fd6ce485d7870d4
password123 -> ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f
By trying common passwords, they can quickly find a matching hash! In the example above, password123 matched.

Now the attacker knows that bob@example.com is likely to use the password password123 on other sites and can go hack their accounts. This is only possible because the attacker can quickly compute many hashes per second and guess millions of potential passwords. If they were going through a login form, the server would probably block them after a few failed attempts.

KDFs fix this
In contrast, a slow hash function (KDF) will run so painfully slow that the attacker is unable to make enough guesses to be able to brute force their way to the plaintext of any passwords.
*/

package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func generateSalt(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func hashPassword(password, salt []byte) []byte {
	password = append(password, salt...)
	h := sha256.New()
	h.Write(password)
	sum := h.Sum(nil)
	return sum
}

//////////
// When appending a byte slice to another byte slice, the second argument must be a slice of bytes ([]byte), not a single byte. To append a byte to a byte slice, you can wrap it in a slice of bytes first.

// Here's an example of how to do it:
salt := []byte("somesalt")
password := []byte("password")
hashedPassword := append(password, salt...)
//In this example, the append function is used to concatenate the password byte slice and the salt byte slice. The ... syntax is used to expand the salt byte slice into individual bytes, which are then appended to the password byte slice. 
//The resulting byte slice is stored in the hashedPassword variable.

// solution 2

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %v", err)
	}
	return salt, nil
}

// hashPassword appends the salt directly to the end of the password, then hashes it with SHA-256
func hashPassword(password, salt []byte) []byte {
	passwordWithSalt := append(password, salt...)
	hash := sha256.Sum256(passwordWithSalt)
	return hash[:]
}
// don't touch below this line

func test(password1, password2 string, saltLen int) {
	defer fmt.Println("========")
	fmt.Printf("Hashing '%s' with salt length %v...\n", password1, saltLen)
	salt, err := generateSalt(saltLen)
	if err != nil {
		fmt.Printf("Error generating salt: %v", err)
		return
	}
	hashed := hashPassword([]byte(password1), salt)
	fmt.Println("Hash generated")

	fmt.Printf("Checking first hash against hash of '%v'...\n", password2)
	hashed2 := hashPassword([]byte(password2), salt)

	if string(hashed) == string(hashed2) {
		fmt.Println("Hashes match!")
	} else {
		fmt.Println("Hashes don't match!")
	}
}

func main() {
	test("samepass", "samepass", 16)
	test("passone", "passtwo", 24)
	test("correct horse battery staple", "correct horse battery staple", 32)
}

/*

Hashing 'samepass' with salt length 16...

Hash generated

Checking first hash against hash of 'samepass'...

Hashes match!

========

Hashing 'passone' with salt length 24...

Hash generated

Checking first hash against hash of 'passtwo'...

Hashes don't match!

========

Hashing 'correct horse battery staple' with salt length 32...

Hash generated

Checking first hash against hash of 'correct horse battery staple'...

Hashes match!

========
*/
