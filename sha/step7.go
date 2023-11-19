package main

import "fmt"

func modifyValues() {
	h0 = h0 + a
	h1 = h1 + b
	h2 = h2 + c
	h3 = h3 + d
	h4 = h4 + e
	h5 = h5 + f
	h6 = h6 + g
	h7 = h7 + h
}

func formatHashResults() string {
	hashValue := fmt.Sprintf("%x", h0)
	hashValue += fmt.Sprintf("%x", h1)
	hashValue += fmt.Sprintf("%x", h2)
	hashValue += fmt.Sprintf("%x", h3)
	hashValue += fmt.Sprintf("%x", h4)
	hashValue += fmt.Sprintf("%x", h5)
	hashValue += fmt.Sprintf("%x", h6)
	hashValue += fmt.Sprintf("%x", h7)

	return hashValue
}
