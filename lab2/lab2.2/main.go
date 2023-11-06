package main

import (
	"fmt"
	"math"
)

var message = "0Ф26ОЯ7ЇВ 8ЖС6ММЬ"
var gamma = "ДАЛІБОР"

func main() {
	fmt.Println("Message: " + message)
	fmt.Println("Gamma: " + gamma + "\n")

	gamma = gammaLengthToMessage(message, gamma)
	fmt.Println("Decrypted message: " + decrypt(message, gamma))
}

func getLength(text string) (length int) {
	count := 0
	for range text {
		count++
	}
	return count
}

func getPartsAmount(message string, gamma string) (parts int) {
	if getLength(message)%getLength(gamma) == 0 {
		return getLength(message) / getLength(gamma)
	} else {
		return getLength(message)/getLength(gamma) + 1
	}
}

func alphabetMap() (m map[string]int) {
	m = make(map[string]int)
	var unicode int32
	count := 1
	for unicode = 1040; unicode <= 1071; unicode++ {
		if count == 5 {
			m[fmt.Sprintf("%c", 1168)] = count
			unicode--

		} else if count == 8 {
			m[fmt.Sprintf("%c", 1028)] = count
			unicode--

		} else if count == 12 {
			m[fmt.Sprintf("%c", 1030)] = count
			unicode--

		} else if count == 13 {
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

func gammaLengthToMessage(message string, gamma string) (editedGamma string) {
	parts := getPartsAmount(message, gamma)
	if parts > 1 {
		for p := 0; p <= parts; p++ {
			editedGamma = gamma + gamma
		}
	}
	lettersLeft := getLength(message) - getLength(editedGamma)
	count := 0
	for _, r := range gamma {
		if count <= lettersLeft-1 {
			editedGamma += fmt.Sprintf("%c", r)
		}
		count++
	}

	return editedGamma
}

func getKeyByValue(m map[string]int, value int) string {
	for key, v := range m {
		if v == value {
			return key // Found the key for the value
		}
	}
	return "" // Value not found in the map
}

func decrypt(message string, gamma string) string {
	key := alphabetMap()
	var ecryptedMessage []int
	var ecryptedGamma []int
	var encryptedDiff []int
	var decryptedText string

	for _, r := range message {
		ecryptedMessage = append(ecryptedMessage, key[fmt.Sprintf("%c", r)])
	}

	for _, r := range gamma {
		ecryptedGamma = append(ecryptedGamma, key[fmt.Sprintf("%c", r)])
	}

	for i := range ecryptedMessage {
		encryptedDiff = append(encryptedDiff, ecryptedMessage[i]-ecryptedGamma[i])
		if encryptedDiff[i] < 0 {
			encryptedDiff[i] = int(math.Abs(float64(encryptedDiff[i]))) + 44 - ecryptedGamma[i]
		}
	}

	for i := range encryptedDiff {
		decryptedText += getKeyByValue(key, encryptedDiff[i])
	}

	return decryptedText
}
