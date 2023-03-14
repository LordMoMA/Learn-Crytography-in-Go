/*
Toy HMAC
At Passly, we use HMACs to authenticate messages between our internal servers when they need to make requests to each other over the public internet.

To demonstrate to our Luddite manager why we should use an open-source crypto library instead of writing our own HMAC implementation, we decided to write our own and then prove it's inferiority. Hopefully we don't get fired instead of getting our way.

Assignment
Complete the hmac function. It should:

Split the key into two halves. The second half should be the larger half if key's length is odd
Return the result of sha256(keyFirstHalf + sha256(keySecondHalf + message)) as a string in lowercase hex
Cast strings directly to slices of bytes and don't use any delimiters when concatenating the data.
*/