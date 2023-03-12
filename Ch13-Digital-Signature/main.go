/*
What is a Message Integrity?
Message integrity ensures that a message has not been changed.

Let's say you need to transmit a private key from your computer to your phone. Because your network connection isn't great, there's a chance that information is lost in transit. You want to be sure that the private key makes it to your phone in pristine condition. If there is data loss, you'll want to know about it so you can try again.

To accomplish this, we can take a hash of the private key, and send it alongside the book. When the private key is transmitted, you'll be able to check on your phone by hashing the private key and checking if the hashes match.

Checksums
This "hash and check" method is called a checksum.

checksum = hash(data)
send data
send checksum
if (hash(data) == checksum) then the message has not been altered
Assignment
Complete the checksumMatches function.

It should hash the message after converting it directly to a slice of bytes using sha256. Check if the lowercase hexadecimal encoding of the hash matches the checksum argument. If it does, return true. Otherwise, return false.
*/