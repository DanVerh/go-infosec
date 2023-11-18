package main

import (
	"fmt"
	"strings"
)

var divSlice []string

// converting binary string to slice where each part has 32 bits
func stringToSlice(text string) {
	var newStr string
	for i, r := range text { // looping through each char of string, i - index, r - rune (unicode value)
		newStr += fmt.Sprintf("%c", r)     // convert symbol unicode to string
		if len(newStr) == 32 && i <= 512 { // append to 32 chars part until the index is 512
			divSlice = append(divSlice, newStr)
			newStr = "" // empty for next 32 characters
		}
	}
}

// add 48 32-bit words consist of 0s
func append48bytes() {
	for i := 0; i < 48; i++ {
		divSlice = append(divSlice, strings.Repeat("0", 32))
	}
}
