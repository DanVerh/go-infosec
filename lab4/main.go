package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var message = "Hello, world"

func main() {
	fmt.Println("Message: " + message)
	message = convertToBinary()
	fmt.Println("Binary message: " + message + "\n")

	img := generateImage(generateImageSize())
	encodeImageToFile("image.png", img)
	fmt.Println("Image genearated with 16-color palette")

	embedBits(img, message)
	encodeImageToFile("imageEncrypted.png", img)
	fmt.Println("Text encrypted in first row of pixels" + "\n")

	decryptedMessage := decryptMessage(img)
	asciiDecryptedMessage := binaryToASCII(decryptedMessage)
	fmt.Println("Decrypted binary message: " + decryptedMessage)
	fmt.Println("Decrypted messag: " + asciiDecryptedMessage)
}

// Convert message to binary string
func convertToBinary() string {
	var result strings.Builder
	for _, r := range message {
		result.WriteString(fmt.Sprintf("%08b", r))
	}

	return result.String()
}

// Generate image size: width = binary message length and random height < binary message length
func generateImageSize() (width int, height int) {
	width = len(message)
	height = rand.Intn(len(message))

	return width, height
}

// Generate image with the following possible 16-color palette
func generateImage(width, height int) *image.Paletted {

	// This will create image with the first mentioned color
	img := image.NewPaletted(image.Rect(0, 0, width, height), color.Palette{
		color.Black,
		color.White,
		color.RGBA{255, 0, 0, 255},     // Red
		color.RGBA{0, 255, 0, 255},     // Green
		color.RGBA{0, 0, 255, 255},     // Blue
		color.RGBA{255, 255, 0, 255},   // Yellow
		color.RGBA{255, 0, 255, 255},   // Magenta
		color.RGBA{0, 255, 255, 255},   // Cyan
		color.RGBA{128, 128, 128, 255}, // Gray
		color.RGBA{255, 165, 0, 255},   // Orange
		color.RGBA{128, 0, 128, 255},   // Purple
		color.RGBA{0, 128, 128, 255},   // Teal
		color.RGBA{128, 0, 0, 255},     // Maroon
		color.RGBA{0, 0, 128, 255},     // Navy
		color.RGBA{192, 192, 192, 255}, // Silver
		color.RGBA{255, 192, 203, 255}, // Pink
	})

	// Randomize color of each pixel
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			// Generate a random color index
			colorIndex := uint8(rand.Intn(len(img.Palette)))
			// Set the color index for the current pixel
			img.SetColorIndex(x, y, colorIndex)
		}
	}

	return img
}

// Create image file and encode the data in it
func encodeImageToFile(filename string, img *image.Paletted){
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

// Add one bit of message to each line of pixels in the first row
func embedBits(img *image.Paletted, bits string) {
	index := 0
	for x := 0; x < img.Bounds().Dx(); x++ {
		colorIndex := img.ColorIndexAt(x, 0)
		// Change every last bit of the pixel to the bit from the message
		img.SetColorIndex(x, 0, colorIndex&^1|(bits[index]&1))
		// Move to the next index
		index = (index + 1) % len(bits)
	}
}

// Function to decrypt one-bit message from the image
func decryptMessage(img *image.Paletted) string {
	var decryptedMessage strings.Builder

	// Iterate over each pixel in the first row
	for x := 0; x < img.Bounds().Dx(); x++ {
		// Get the color index for the current pixel
		colorIndex := img.ColorIndexAt(x, 0)
		// Extract the least significant bit and append it to the decrypted message
		decryptedMessage.WriteByte((colorIndex & 1) + '0')
	}

	return decryptedMessage.String()
}

// Convert binary string to ASCII
func binaryToASCII(binaryString string) (asciiResult string) {
	for i := 0; i < len(binaryString); i += 8 {
		// Divide on 8-bit slices
		segment := binaryString[i : i+8]
		// Convert binary string to decimal
		decimal, _ := strconv.ParseInt(segment, 2, 64)
		// Convert decimal to ASCII
		asciiChar := string(decimal)
		asciiResult += asciiChar
	}

	return asciiResult
}
