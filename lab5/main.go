package main

import "fmt"

var x1 = []uint32{0x7A8FBCD3}
var x2 = []uint32{0xE19D2F4A}
var x3 = []uint32{0x45C67082}
var x4 = []uint32{0xF632A1E7}
var k = [2]uint32{0x1B9E84F0, 0x8D376A59}
var n = 12

func main() {
	fmt.Printf("Input message: %x %x %x %x\n", x1[0], x2[0], x3[0], x4[0])
	encryptFeistel()
	fmt.Printf("Encrypted message: %x %x %x %x\n", x1[11], x2[11], x3[11], x4[11])
	decryptFeistel()
	fmt.Printf("Decrypted message: %x %x %x %x", x1[0], x2[0], x3[0], x4[0])
}

func rol(value uint32, shift uint) uint32 {
	return (value << shift) | (value >> (32 - shift))
}

func ror(value uint32, shift uint) uint32 {
	return (value >> shift) | (value << (32 - shift))
}

func getH(i uint) uint32 {
	return rol(k[0], i) ^ ror(k[1], i)
}

func getV(i uint) uint32 {
	return (x1[i-1] * getH(i)) % (1<<n + 1)
}

func encryptFeistel() {
	for i := 1; i < n; i++ {
		x1 = append(x1, x2[i-1]^getV(uint(i)))
		x2 = append(x2, x3[i-1])
		x3 = append(x3, x4[i-1])
		x4 = append(x4, x1[i-1])
	}
}

func decryptFeistel() {
	for i := n - 1; i >= 1; i-- {
		x4 = append(x4, x1[i-1])
		x3 = append(x3, x2[i-1])
		x2 = append(x2, x3[i-1])
		x1 = append(x1, x4[i-1]^getV(uint(i)))
	}
}
