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
*/