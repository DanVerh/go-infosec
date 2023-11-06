package main

import (
	"fmt"
	"math"
	"math/big"
)

var p = 17
var q = 29
var message = "СБР23П"
var cipher = []int{1, 172, 225, 32, 335, 443, 469, 379}

func main() {
	n := getN(p, q)
	f := eulerFunc(p, q)
	e := getE(p)
	publicKey := getKeyPair(e, n)
	d := getD(float64(f), float64(e))
	privateKey := getKeyPair(d, n)

	fmt.Println("Public Key:", publicKey)
	fmt.Println("Message: " + message)
	fmt.Println("Message encrypted with RSA public key:", encryptRSA(publicKey, message), "\n")

	fmt.Println("Private Key:", privateKey)
	fmt.Println("Encrypted Message:", cipher)
	fmt.Println("Message decrypted with RSA private key:", decryptRSA(privateKey, cipher))
}

func getN(p int, q int) int {
	return p * q
}

func eulerFunc(p int, q int) int {
	return (p - 1) * (q - 1)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func getE(p int) int {
	for n := p - 1; n > 0; n-- {
		if isPrime(n) {
			return n
		}
	}
	return -1
}

func getKeyPair(n1 int, n2 int) (publicKey [2]int) {
	publicKey[0] = n1
	publicKey[1] = n2

	return publicKey
}

func getD(f float64, e float64) int {
	notFloor := true
	k := 1.0
	for notFloor {
		d := ((k * f) + 1.0) / e
		if d == math.Floor(d) {
			notFloor = false
			return int(d)
		} else {
			k++
		}
	}
	return -1
}

func alphabetMap() (m map[string]int) {
	m = make(map[string]int)
	var unicode int32
	count := 1
	for unicode = 1040; unicode <= 1071; unicode++ {
		if count == 29 {
			m[fmt.Sprintf("%c", 1168)] = count
			unicode--

		} else if count == 28 {
			m[fmt.Sprintf("%c", 1028)] = count
			unicode--

		} else if count == 31 {
			m[fmt.Sprintf("%c", 1030)] = count
			unicode--

		} else if count == 7 {
			m[fmt.Sprintf("%c", 1111)] = count
			unicode--

		} else if unicode == 1066 || unicode == 1067 || unicode == 1069 {
			count--
		} else {
			m[fmt.Sprintf("%c", unicode)] = count

		}
		count++
	}
	m[fmt.Sprintf("%c", 32)] = count
	count++
	for unicode = 48; unicode <= 57; unicode++ {
		m[fmt.Sprintf("%c", unicode)] = count
		count++
	}
	return m
}

func encryptRSA(publicKey [2]int, message string) []int {
	alphabet := alphabetMap()
	var encryptedMessage []int
	var c int
	for _, r := range message {
		t := alphabet[fmt.Sprintf("%c", r)]
		c = int(math.Pow(float64(t), float64(publicKey[0]))) - (publicKey[1] * (int(math.Pow(float64(t), float64(publicKey[0]))) / publicKey[1]))
		encryptedMessage = append(encryptedMessage, c)
	}
	return encryptedMessage
}

func getKeyByValue(m map[string]int, value int) string {
	for key, v := range m {
		if v == value {
			return key
		}
	}
	return ""
}

func decryptRSA(privateKey [2]int, cipher []int) string {
	var decryptedMessage string

	for _, r := range cipher {
		// Declare and set big int values
		base := big.NewInt(int64(r))
		exponent := big.NewInt(int64(privateKey[0]))
		privateKey1 := big.NewInt(int64(privateKey[1]))

		// Operations to decrypt step-by-step
		expResult := new(big.Int).Exp(base, exponent, nil)
		modResult := new(big.Int).Div(expResult, privateKey1)
		product := new(big.Int).Mul(privateKey1, modResult)
		result := new(big.Int).Sub(expResult, product).Int64()

		// Decrypt by each symbol and add to string
		decryptedSymbol := getKeyByValue(alphabetMap(), int(result))
		decryptedMessage += decryptedSymbol
	}
	return decryptedMessage
}
