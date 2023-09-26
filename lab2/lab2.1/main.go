package main

import "fmt"

var message = "ЧЕГДОМИН 5000 Т"
var gamma = "БІЛОМИР"

func main() {
	fmt.Println(getLength(gamma))
	fmt.Println(getLength(message))
	fmt.Println(getPartsAmount(message, gamma))
	fmt.Println(alphabetMap())
	fmt.Println(gammaLengthToMeassage(message, gamma))

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

func gammaLengthToMeassage(message string, gamma string) (editedGamma string) {
	parts := getPartsAmount(message, gamma)
	if parts > 1 {
		for p := 0; p <= parts; p++ {
			editedGamma = gamma + gamma
		}
	}
	lettersLeft := getLength(message) - getLength(editedGamma)
	for i, r := range gamma {
		if i <= lettersLeft-1 {
			editedGamma += fmt.Sprintf("%c", r)
		}
		break
	}
	/*for i, u := range message {}*/

	return editedGamma
}

func alphabetMap() (m map[string]int) {
	m = make(map[string]int)
	var unicode int32
	count := 1
	for unicode = 1040; unicode <= 1071; unicode++ {
		if count == 5 {
			m[fmt.Sprintf("%c", 1168)] = count
			unicode--
			//count++
		} else if count == 8 {
			m[fmt.Sprintf("%c", 1028)] = count
			unicode--
			//count++
		} else if count == 12 {
			m[fmt.Sprintf("%c", 1030)] = count
			unicode--
			//count++
		} else if count == 13 {
			m[fmt.Sprintf("%c", 1111)] = count
			unicode--
			//count++
		} else if unicode == 1066 || unicode == 1067 || unicode == 1069 {
			count--
		} else {
			m[fmt.Sprintf("%c", unicode)] = count
			//count++
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

/*
func alphabetMap() (m map[int]string) {
	m = make(map[int]string)
	var unicode int32
	count := 0
	for unicode = 1040; unicode <= 1048; unicode++ {
		m[count] = fmt.Sprintf("%c", unicode)
		count++
	}
	m[count] = fmt.Sprintf("%c", 1030)
	count++
	m[count] = fmt.Sprintf("%c", 1111)
	count++
	for unicode = 1049; unicode <= 1071; unicode++ {
		m[count] = fmt.Sprintf("%c", unicode)
		count++
	}
	m[count] = fmt.Sprintf("%c", 32)
	count++
	for unicode = 48; unicode <= 57; unicode++ {
		m[count] = fmt.Sprintf("%c", unicode)
		count++
	}
	return m
}
*/
/*func alphabetSlice() (alphabet []string) {
	var unicode int32
	var symbol string
	for unicode = 1040; unicode <= 1048; unicode++ {
		symbol = fmt.Sprintf("%c", unicode)
		alphabet = append(alphabet, symbol)
	}
	alphabet = append(alphabet, fmt.Sprintf("%c", 1030))
	alphabet = append(alphabet, fmt.Sprintf("%c", 1111))
	for unicode = 1049; unicode <= 1071; unicode++ {
		symbol = fmt.Sprintf("%c", unicode)
		alphabet = append(alphabet, symbol)
	}
	alphabet = append(alphabet, fmt.Sprintf("%c", 32))
	for unicode = 48; unicode <= 57; unicode++ {
		symbol = fmt.Sprintf("%c", unicode)
		alphabet = append(alphabet, symbol)
	}
	return alphabet
}*/
