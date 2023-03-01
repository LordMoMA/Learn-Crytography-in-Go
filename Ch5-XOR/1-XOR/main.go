/*
Assignment
Write a function xor() that takes two boolean values as input and
returns the result of the "exclusive or" operation.
*/

package main

import "fmt"

func xor(lhs, rhs bool) bool {
	if lhs == rhs {
		return false
	} else {
		return true
	}
}

// don't touch below this line

func test(lhs, rhs bool) {
	res := xor(lhs, rhs)
	fmt.Printf("%v XOR %v = %v\n", lhs, rhs, res)
}

func main() {
	test(true, true)
	test(true, false)
	test(false, true)
	test(false, false)
}
