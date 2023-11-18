package main

import (
	"fmt"
	"strings"
)

// this string will be modified during all operations
var result strings.Builder

// change each char of input string to binary representation
func inputBinary(input string) {
	for _, char := range input {
		result.WriteString(fmt.Sprintf("%08b", char))
	}
}

func append1() {
	result.WriteString("1")
}

// append 0s until we have len = 504 (last 8 symbols for input len)
func append0s() {
	zeroes := strings.Repeat("0", 512-len(result.String())-8)

	result.WriteString(zeroes)
}

// add binary len of input at the end
func inputLenBinary() {
	inputLenBinary := fmt.Sprintf("%08b", len(input)*8)

	result.WriteString(inputLenBinary)
}
