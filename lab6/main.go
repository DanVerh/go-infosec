package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var t = 10

func main() {
	startTime := time.Now()
	p, i := generatePrimeNumber()
	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime).Seconds()

	fmt.Printf("Prime number generation time: %.6f seconds\n", elapsedTime)
	fmt.Printf("Amount of iterations: %d \n", i)
	fmt.Printf("Prime number: %d\n", p)
}

func getOddP() int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(10000) + 1) | 1
}

func getB(n int) int {
	b := 0
	for n%2 == 0 {
		n /= 2
		b++
	}
	return b
}

func getM(p, b int) (m int) {
	m = (p - 1) / int(math.Pow(2, float64(b)))
	if m == 0 {
		m = 1
	}
	return m
}

func getA(p int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(p-1) + 1
}

func getZ(a, m, p int) int {
	am := int(math.Pow(float64(a), float64(m)))
	return am - p*(am/p)
}

func checkDivisionByPrimes(p int) bool {
	if p%3 == 0 || p%5 == 0 || p%7 == 0 || p%11 == 0 {
		return true
	}
	return false
}

func testRabinMiller(z, j, p, b int) (result bool) {
	// point 3
	if z == 1 || z == p-1 {
		result = true
		return
	}

	// point 5
	j += 1
	if j < b && z < p-1 {
		z2 := int(math.Pow(float64(z), 2))
		z = z2 - p*(z2/p)
		if j > 0 && z == 1 {
			result = false
			return
		} else if z == p-1 {
			result = true
			return
		}
	}

	// point 6
	if j == b && z != p-1 {
		result = false
		return
	}

	return result
}

func generatePrimeNumber() (p int, i int) {
	primeCheck := false
	i = 0
	for primeCheck == false {
		i++
		p = getOddP()
		fmt.Println(p, checkDivisionByPrimes(p))
		if checkDivisionByPrimes(p) == false && p != 1 {
			for i := 1; i <= t; i++ {
				b := getB(p - 1)
				m := getM(p, b)
				a := getA(p)
				j := 0
				z := getZ(a, m, p)
				//fmt.Printf("p = %d, b = %d, m = %d, a = %d, j = %d, z = %d \n", p, b, m, a, j, z)
				primeCheck = testRabinMiller(z, j, p, b)
				//fmt.Println(primeCheck)
			}
		} else {
			primeCheck = false
		}
	}
	return p, i
}
