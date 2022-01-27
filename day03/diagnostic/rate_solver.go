package diagnostic

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	Forward = "forward"
	Up      = "up"
	Down    = "down"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counts := make([]int, 12)
	gammaRate := 0
	epsilonRate := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if len(s) != 12 {
			log.Fatal("invalid string length", s)
		}
		for i := 0; i < len(s); i++ {
			shift := 1
			if string(s[i]) == "0" {
				shift = -1
			}
			counts[i] = counts[i] + shift
		}
	}
	for i := 0; i < len(counts); i++ {
		if counts[i] > 0 {
			gammaRate = gammaRate + int(math.Pow(2, float64(len(counts)-i-1)))
		} else if counts[i] < 0 {
			epsilonRate = epsilonRate + int(math.Pow(2, float64(len(counts)-i-1)))
		} else {
			log.Fatal("oh no we got 0")
		}
	}

	fmt.Println("Problem1", "gammaRate", gammaRate, "epsilonRate", epsilonRate, "answer", gammaRate*epsilonRate)
}

func Problem2(input string) {
	data := getData(input)
	sort.Strings(data)

	const numLength = 12
	//fmt.Println("len(data)", len(data))

	found := false
	minOx := 0
	maxOx := len(data) - 1
	minSc := 0
	maxSc := len(data) - 1
	for bitPos := 0; bitPos < numLength && found == false; bitPos++ {
		if minOx < maxOx {
			line := minOx
			for ; line <= maxOx && string(data[line][bitPos]) != "1"; line++ {
			}

			length := maxOx - minOx + 1
			num0s := line - minOx

			// more or even 1s then raise min
			if num0s <= length-num0s {
				minOx = int(math.Min(float64(line), float64(maxOx)))
			} else {
				// less 1s then decrease max
				maxOx = int(math.Max(float64(line-1), float64(minOx)))
			}

		}

		if minSc < maxSc {
			line := minSc
			for ; line <= maxSc && string(data[line][bitPos]) != "1"; line++ {

			}
			length := maxSc - minSc + 1
			num0s := line - minSc

			// less or even 0s then lower max
			if num0s <= length-num0s {
				maxSc = int(math.Max(float64(line-1), float64(minSc)))
			} else {
				minSc = int(math.Min(float64(line), float64(maxSc)))
			}

		}
		//fmt.Println(bitPos, "minOx", minOx, "maxOx", maxOx, "minSc", minSc, "maxSc", maxSc)
	}

	oxRate, _ := strconv.ParseInt(data[minOx], 2, 32)
	scRate, _ := strconv.ParseInt(data[minSc], 2, 32)

	fmt.Println("Problem2", "oxRate", oxRate, "scRate", scRate, "answer", oxRate*scRate)
}

func getData(dataPath string) []string {
	file, err := os.Open(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ss := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		ss = append(ss, s)
	}
	return ss
}
