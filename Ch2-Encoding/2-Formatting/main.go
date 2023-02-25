package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	result := []string{}
	for _, v := range b {
		result = append(result, fmt.Sprintf("%02x", v))
	}
	return strings.Join(result, ":")
}

func getBinaryString(b []byte) string {
	result := []string{}
	for _, v := range b {
		result = append(result, fmt.Sprintf("%08b", v))
	}
	return strings.Join(result, ":")
}

// don't touch below this line

func testHex(s string) {
	myBytes := []byte(s)
	fmt.Printf("String: '%s', Hex: %v\n", s, getHexString(myBytes))
	fmt.Println("========")
}

func testBinary(s string) {
	myBytes := []byte(s)
	fmt.Printf("String: '%s', Bin: %v\n", s, getBinaryString(myBytes))
	fmt.Println("========")
}

func main() {
	testHex("Hello")
	testHex("World")
	testBinary("Hello")
	testBinary("World")
}
