package enhancer

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Problem1(input string) {
	algorithm, image := loadData(input)

	numLit := iterate(image, algorithm, 2)

	fmt.Println("Problem1", "numLit", numLit)
}

func Problem2(input string) {
	algorithm, image := loadData(input)

	numLit := iterate(image, algorithm, 50)

	fmt.Println("Problem2", "numLit", numLit)
}

func loadData(input string) (string, []string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		s         string
		algorithm string
		image     []string
	)
	//wasLoaded := false
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(s) <= 0 {
			continue
		}

		if len(s) == 512 {
			algorithm = s
		} else {
			image = append(image, s)
		}
	}
	//fmt.Println(image)
	//fmt.Println(algorithm)
	return algorithm, image
}

func iterate(image []string, algorithm string, iterations int) int {
	boundary := "."
	var newImage []string
	newImage = image
	for i := 0; i < iterations; i++ {
		newImage = enhance(newImage, algorithm, boundary)
		if boundary == "." {
			boundary = string(algorithm[0])
		} else {
			boundary = string(algorithm[len(algorithm)-1])
		}
	}

	numLit := 0
	for i := 0; i < len(newImage); i++ {
		numLit = numLit + strings.Count(newImage[i], "#")
	}
	return numLit
}

func enhance(image []string, algorithm string, boundary string) []string {
	// determine new image size
	// process boundary

	// grow
	imageCopy := make([]string, len(image))
	copy(imageCopy, image)
	rowEdge := boundary + boundary
	for i := 0; i < len(imageCopy); i++ {
		imageCopy[i] = rowEdge + imageCopy[i] + rowEdge
	}
	newRow := strings.Repeat(boundary, len(imageCopy[0]))
	imageCopy = append([]string{newRow, newRow}, imageCopy...)
	imageCopy = append(imageCopy, newRow, newRow)
	//fmt.Println("after resizing")
	//printImage(imageCopy)
	//fmt.Println()

	// run enhance
	var newImage []string
	// add rows to top and bottom
	newImage = make([]string, len(imageCopy)-2)
	for i := 1; i < len(imageCopy)-1; i++ {
		for j := 1; j < len(imageCopy[i])-1; j++ {
			bin := imageCopy[i-1][j-1:j+2] + imageCopy[i][j-1:j+2] + imageCopy[i+1][j-1:j+2]
			newPixel := enhancePixel(bin, algorithm)
			newImage[i-1] = newImage[i-1] + newPixel
		}
	}
	//fmt.Println("after enhance")
	//printImage(newImage)
	//fmt.Println()

	return newImage
}

func enhancePixel(s, algorithm string) string {
	return numToCode(toNum(s), algorithm)
}

func toNum(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			num = num + int(math.Pow(2, float64(len(s)-i-1)))
		}
	}
	return num
}
func numToCode(num int, algorithm string) string {
	return string(algorithm[num])
}

func printImage(s []string) {
	fmt.Println(strings.Join(s, "\n"))
}
