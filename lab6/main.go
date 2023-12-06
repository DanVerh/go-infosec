package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var n = 128
var t = 10

func main() {
	p := generateRandomBigInt(n)
	fmt.Println(p)
	b := findMaxPowerOfTwo(new(big.Int).Sub(p, big.NewInt(64)))
	fmt.Println(b)
	m := getM(p, b)
	fmt.Println(m)
	a := generateRandomBigIntLessP(n, p)
	fmt.Println(a)
	j := big.NewInt(0)
	fmt.Println(j)
	z := getZ(a, m, p)
	fmt.Println(z)
	fmt.Printf("p = %d, b = %d, m = %d, a = %d, j = %d, z = %d \n", p, b, m, a, j, z)
	testResult := testRabinMiller(z, j, p, b)
	fmt.Println(testResult)
}

func generateRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 2
}

func generateRandomBigInt(bitLength int) *big.Int {
	// Generate random bytes
	randomBytes := make([]byte, (bitLength+7)/8)

	// Create a big.Int from the random bytes
	randomBigInt := new(big.Int).SetBytes(randomBytes)

	// Set the top bit to ensure the number is positive and larger than 2^64
	minValue := new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)

	randomBigInt.Add(randomBigInt, minValue)
	randomBigInt.SetBit(randomBigInt, bitLength-1, 1)

	randomBigInt.SetBit(randomBigInt, 0, 1)

	return randomBigInt
}

func findMaxPowerOfTwo(n *big.Int) *big.Int {
	b := new(big.Int)
	two := big.NewInt(2)

	// Iterate while n is even
	for n.Bit(0) == 0 { // Check if the least significant bit is 0
		n.Div(n, two)           // Divide n by 2
		b.Add(b, big.NewInt(1)) // Increment the count of powers of 2
	}
	return b
}

func getM(p, b *big.Int) *big.Int {
	subp1 := new(big.Int).Sub(p, big.NewInt(1))
	exp := new(big.Int).Exp(big.NewInt(2), b, nil)
	result := new(big.Int).Div(subp1, exp)

	return result
}

func generateRandomBigIntLessP(bitLength int, upperLimit *big.Int) *big.Int {
	for {
		// Generate random bytes
		randomBytes := make([]byte, (bitLength+7)/8)
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err) // Handle the error appropriately in your actual code
		}

		// Create a big.Int from the random bytes
		randomBigInt := new(big.Int).SetBytes(randomBytes)

		// Check if the generated value is less than the upper limit
		if randomBigInt.Cmp(upperLimit) < 0 {
			return randomBigInt
		}
	}
}

func getZ(a, m, p *big.Int) *big.Int {
	// Calculate a^m
	am := new(big.Int).Exp(a, m, nil)

	// Calculate result: am - p*(am/p)
	amDivP := new(big.Int).Div(am, p)
	result := new(big.Int).Sub(am, new(big.Int).Mul(p, amDivP))

	return result
}

func testRabinMiller(z, j, p, b *big.Int) bool {
	// point 3
	one := big.NewInt(1)
	pMinusOne := new(big.Int).Sub(p, one)
	if z.Cmp(one) == 0 || z.Cmp(pMinusOne) == 0 {
		return true
	}

	// point 5
	j.Add(j, one)
	bigB := new(big.Int).Set(b)
	if j.Cmp(bigB) < 0 && z.Cmp(pMinusOne) < 0 {
		z2 := new(big.Int).Exp(z, big.NewInt(2), nil)
		z.Mod(z2.Sub(z2, p.Mul(z2.Div(z2, p), p)), p)
		if j.Cmp(one) > 0 && z.Cmp(one) == 0 {
			return false
		} else if z.Cmp(pMinusOne) == 0 {
			return true
		}
	}

	// point 6
	if j.Cmp(bigB) == 0 && z.Cmp(pMinusOne) != 0 {
		return false
	}

	return false
}
