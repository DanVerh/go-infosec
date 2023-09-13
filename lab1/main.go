package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var symbolsMax = 32

func main() {
	fmt.Println("Input text:")
	text := "Go is expressive, concise, clean and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines"
	fmt.Println(text, "\n")

	fmt.Println("Encrypted text:")
	parts, err := getPartsAmount(text) // write output and error values (for check of input length)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		text = encrypt(stringToSlice(text, parts))
		fmt.Println(text)
	}
}

// get amount of substring of the divided on 32 string
func getPartsAmount(text string) (parts int, err error) { // two values in output to catch the error
	if len(text) < symbolsMax {
		return 0, errors.New("text contains less than 32 characters")
	}

	parts = len(text) / symbolsMax
	if parts != len(text) {
		return parts + 1, nil // for last part < 32 characters
	} else {
		return parts, nil
	}
}

// converting divided on parts string to slice where each part has 32 characters (except last one)
func stringToSlice(text string, parts int) (divSlice []string) {
	var newStr string
	for i, r := range text { // looping through each char of string, i - index, r - rune (unicode value)
		newStr += fmt.Sprintf("%c", r)                // convert symbol unicode to string
		if len(newStr) == 32 && i <= (parts-1)*32-1 { // append to 32 chars part
			divSlice = append(divSlice, string(newStr))
			newStr = "" // empty for next 32 characters
		} else if i > (parts-1)*32-1 && len(newStr) == len(text)-(parts-1)*32 { // for the last part
			divSlice = append(divSlice, string(newStr))
		}
	}
	return divSlice
}

// final ecryption function
func encrypt(textSlice []string) (encryptedText string) {
	var randInt int
	var key []int
	for _, textPart := range textSlice { // looping parts of slice, _ - not used index, textPart - part of slice
		randInt = rand.Intn(32) // random int 0-32
		key = append(key, randInt)
		for _, unicodeElement := range textPart {
			unicodeElement = unicodeElement + int32(randInt) // moving in unicode format
			if unicodeElement > 126 {
				unicodeElement = 31 + unicodeElement - 126 // to use only Latin range
			}
			encryptedText += fmt.Sprintf("%c", unicodeElement) // from unicode to string
		}
	}
	fmt.Println("Key:", key) // print the key
	return encryptedText
}
