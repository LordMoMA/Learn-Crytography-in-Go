/*
Key Derivation Functions
A key derivation function (KDF) is a special kind of hash function that specializes in deriving one or more secret keys from a secret value, such as a password.

All KDFs are hash functions, but not all hash functions are KDFs.

hash kdf venn

There are 4 primary reasons to use a KDF, and I've ordered them from most the common use cases to the least common:

Key Stretching
Key Separation
Key Whitening
Key Strengthening
1. Key Stretching
Key stretching takes a short key and stretches it into a longer key. This is useful when you need a longer key than what you have available. For example, it's common to take a human-friendly password and stretch it into a 256-bit key.

2. Key Separation
KDFs can create multiple child keys from a master key. You may remember that we talked about this process when we worked with the Feistel Cipher. This is helpful because you only need to remember or store the master, and can deterministically generate the child keys when you need them.

3. Key Whitening
Key whitening increases the security of block ciphers by combining portions of the data with the key.

4. Key Strengthening
Less common than stretching, strengthing extends a key with a random salt, but then deletes the salt so it can’t be used again. This strengthens the resulting key, but comes at a strong convenience because even legitimate users won't be able to derive the original easily.

Bcrypt - Storing Passwords
KDFs are the best way to store passwords in web applications! As a back-end developer, this is critical to understand.

Can I store passwords in plain text?
Storing passwords in a database in plain text is a huge security risk. If someone gets access to your database, they can see all of your users' passwords.

Can I hash passwords with SHA-256?
No. SHA-256 is a hash function, but it's not a KDF. SHA-256 is a very fast hash function. Good KDFs like Bcrypt are designed to be slow. We'll talk more about why that's important soon.

Assignment
At Passly, we store passwords securely (it would be sad if we didn't). Each user has a master password that they use to log into their cloud account. That password is hashed with Bcrypt before being stored.

Use the golang.org/x/crypto/bcrypt package to complete the hashPassword() and checkPasswordHash() functions. You do not need to modify the function signatures, just implement the Bcrypt library's API and do the []byte <-> string conversions.

Use a cost factor of 13
Docs for bcrypt.GenerateFromPassword
Docs for bcrypt.CompareHashAndPassword
*/

package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	passwd := []byte(password)
	cost := 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwd, cost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	passwd := []byte(password)
	hashed := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashed, passwd)
	if err != nil {
		return false
	} else {
		return true
	}
}

// don't touch below this line

func test(password1, password2 string) {
	defer fmt.Println("========")
	fmt.Printf("Hashing '%s'...\n", password1)
	hashed, err := hashPassword(password1)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return
	}
	fmt.Printf("Bcrypt output generated with len: %v\n", len(hashed))
	match := checkPasswordHash(password2, hashed)
	fmt.Printf("%v has a matching hash: %v\n", password2, match)
}

func main() {
	test("thisIsAPassword", "thisIsAPassword")
	test("thisIsAPassword", "thisIsAnotherPassword")
	test("corr3ct h0rse", "corr3ct h0rse")
}

/*

Hashing 'thisIsAPassword'...

Bcrypt output generated with len: 60

thisIsAPassword has a matching hash: true

========

Hashing 'thisIsAPassword'...

Bcrypt output generated with len: 60

thisIsAnotherPassword has a matching hash: false

========

Hashing 'corr3ct h0rse'...

Bcrypt output generated with len: 60

corr3ct h0rse has a matching hash: true

========
*/
