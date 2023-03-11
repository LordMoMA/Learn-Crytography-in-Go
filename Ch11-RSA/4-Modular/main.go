/*
Modular
Modular arithmetic and modular exponentiation are widely used in cryptography, and RSA is no exception.

What is "mod"?
In modular arithmetic, we're working with a set of numbers that are all less than a certain number, and we call that the "mod". It's similar to the modulus operator, %, but instead of returning the remainder, we're working within the world of the possible remainders.

The world of "mod 2"
0 in the world of "mod 2" is "congruent" with 0. So we would write:

0 ≡ 0 (mod 2)
1 ≡ 1 (mod 2)
2 ≡ 0 (mod 2)
3 ≡ 1 (mod 2)
4 ≡ 0 (mod 2)
5 ≡ 1 (mod 2)
...
≡ is the congruence symbol.

The world of "mod 3"
As you saw above, in the world of "mod 2", there are only 2 possible values: {0, 1}. However, in the world of "mod 3", there are 3 possible values: {0, 1, 2}.

1 ≡ 1 (mod 3)
2 ≡ 2 (mod 3)
3 ≡ 0 (mod 3)
4 ≡ 1 (mod 3)
5 ≡ 2 (mod 3)
...

Modular Arithmetic
It will be important to understand a simple transformation that's possible with modular arithmetic. There 2 formulas are effectively equivalent:

n ≡ r (mod q)
n = qk + r
Where n is the number we're working with, r is the remainder, q is the quotient, and k is the number of times q goes into n. This allows us to convert an equation in the world of modular arithmetic into a more familiar equation the normal world.

For example,

9 ≡ 1 (mod 2)
Can be converted to:

9 = 2k + 1
This will be important later.
*/