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
	message = convertToBinary()
	fmt.Println(message)

	img := generateImage(generateImageSize())

	// Запис зображення в файл
	file, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	// Вбудовування однобітового повідомлення в зображення
	embedBits(img, message)

	// Запис зображення в файл
	file, err = os.Create("imageEncrypted.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	// Split binary string into 8-bit segments
	var asciiResult string
	for i := 0; i < len(decryptMessage(img)); i += 8 {
		segment := decryptMessage(img)[i : i+8]
		asciiChar, err := binaryToASCII(segment)
		if err != nil {
			fmt.Println("Error converting binary to ASCII:", err)
			return
		}
		asciiResult += asciiChar
	}

	fmt.Println(asciiResult)
}

// Convert message to binary representation
func convertToBinary() string {
	var result strings.Builder
	for _, r := range message {
		result.WriteString(fmt.Sprintf("%08b", r))
	}

	return result.String()
}

// Generate image size: width = message length and random height
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

func binaryToASCII(binaryString string) (string, error) {
	// Convert binary string to decimal
	decimal, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		return "", err
	}

	// Convert decimal to ASCII
	asciiChar := string(decimal)

	return asciiChar, nil
}
