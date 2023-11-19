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

/*
	modify the zero-ed indexes at the end of the array from w[16-63] with the next algorithm:

s0 = (w[i-15] rightrotate 7) xor (w[i-15] rightrotate 18) xor (w[i-15] rightshift 3)
s1 = (w[i- 2] rightrotate 17) xor (w[i- 2] rightrotate 19) xor (w[i- 2] rightshift 10)
w[i] = w[i-16] + s0 + w[i-7] + s1
*/
func createSchedule() {
	for i := 16; i < 64; i++ {
		wi15uint32, _ := strconv.ParseUint(w[i-15], 2, 32)
		wi2uint32, _ := strconv.ParseUint(w[i-2], 2, 32)
		wi16uint32, _ := strconv.ParseUint(w[i-16], 2, 32)
		wi7uint32, _ := strconv.ParseUint(w[i-7], 2, 32)
		s0 := rightRotate(uint32(wi15uint32), 7) ^ rightRotate(uint32(wi15uint32), 18) ^ rightShift(uint32(wi15uint32), 3)
		s1 := rightRotate(uint32(wi2uint32), 17) ^ rightRotate(uint32(wi2uint32), 19) ^ rightShift(uint32(wi2uint32), 10)
		w[i] = fmt.Sprintf("%032b", uint32(wi16uint32)+s0+uint32(wi7uint32)+s1)
	}
}

func rightRotate(value uint32, shift uint) uint32 {
	return (value >> shift) | (value << (32 - shift))
}

func rightShift(value uint32, shift uint) uint32 {
	return value >> shift
}
