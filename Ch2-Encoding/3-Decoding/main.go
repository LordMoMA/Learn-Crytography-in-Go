package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	final := []byte{}
	parts := strings.Split(s, ":")
	for _, v := range parts {
		dec, err := hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
		final = append(final, dec...)
	}
	return final, nil
}

// don't touch below this line

func testHex(s string) {
	myBytes, err := getHexBytes(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Hex: '%s', String: %v\n", s, string(myBytes))
	fmt.Println("========")
}

func main() {
	testHex("48:65:6c:6c:6f")
	testHex("57:6f:72:6c:64")
	testHex("50:61:73:73:77:6f:72:64")
}
