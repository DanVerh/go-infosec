package main

import (
	"fmt"
	"strconv"
	"strings"
)

var w []string

// converting binary string to slice where each part has 32 bits
func stringToSlice(text string) {
	var newStr string
	for i, r := range text { // looping through each char of string, i - index, r - rune (unicode value)
		newStr += fmt.Sprintf("%c", r)     // convert symbol unicode to string
		if len(newStr) == 32 && i <= 512 { // append to 32 chars part until the index is 512
			w = append(w, newStr)
			newStr = "" // empty for next 32 characters
		}
	}
}

// add 48 32-bit words consist of 0s
func append48bytes() {
	for i := 0; i < 48; i++ {
		w = append(w, strings.Repeat("0", 32))
	}
}

// modify the zero-ed indexes at the end of the array from w[16-63]
func createSchedule() {
	//for i := 16; i < 64; i++ {}
	wuint32, _ := strconv.ParseUint(w[15], 2, 32)
	fmt.Printf("%032b", wuint32)
}
