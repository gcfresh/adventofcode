package display

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = '|'
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for i := 0; i < len(data); i++ {
		output := strings.Split(data[i][1], " ")
		for j := 0; j < len(output); j++ {

			if len(output[j]) == 2 { // if 1
				count++
			} else if len(output[j]) == 4 { // 4
				count++
			} else if len(output[j]) == 3 { //7
				count++
			} else if len(output[j]) == 7 { // 8
				count++
			}
		}
	}

	fmt.Println("Problem1", "count", count)
}

func Problem2(input string) {
	t := time.Now()
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = '|'
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for i := 0; i < len(data); i++ {
		signals := strings.Fields(data[i][0])
		//signals = signals
		//fmt.Println(signals, len(signals))

		sigOrder := make([]string, len(signals))
		for j := 0; j < len(signals); j++ {
			if len(signals[j]) == 2 { // if 1
				sigOrder[1] = signals[j]
			} else if len(signals[j]) == 4 { // 4
				sigOrder[4] = signals[j]
			} else if len(signals[j]) == 3 { // 7
				sigOrder[7] = signals[j]
			} else if len(signals[j]) == 7 { // 8
				sigOrder[8] = signals[j]
			}
		}

		// find those with 6
		for j := 0; j < len(signals); j++ {
			if len(signals[j]) == 6 { // 0,6,9
				// check if it has 1
				if strings.ContainsRune(signals[j], rune(sigOrder[1][0])) &&
					strings.ContainsRune(signals[j], rune(sigOrder[1][1])) {
					// could be 0 or 9
					// check if it has all of 4
					if strings.ContainsRune(signals[j], rune(sigOrder[4][0])) &&
						strings.ContainsRune(signals[j], rune(sigOrder[4][1])) &&
						strings.ContainsRune(signals[j], rune(sigOrder[4][2])) &&
						strings.ContainsRune(signals[j], rune(sigOrder[4][3])) {
						sigOrder[9] = signals[j]

					} else {
						sigOrder[0] = signals[j]
					}
				} else {
					sigOrder[6] = signals[j]
				}
			}
			if len(signals[j]) == 5 { // 2,3,5
				// check if it has 1
				if strings.ContainsRune(signals[j], rune(sigOrder[1][0])) &&
					strings.ContainsRune(signals[j], rune(sigOrder[1][1])) {
					sigOrder[3] = signals[j]
				} else {
					// 2 and 5 are a bit more complex
					count := 0
					for _, fourRune := range sigOrder[4] {
						if strings.ContainsRune(signals[j], fourRune) {
							count++
						}
					}
					if count == 2 {
						sigOrder[2] = signals[j]
					} else {
						sigOrder[5] = signals[j]
					}
				}
			}
		}

		//fmt.Println(sigOrder)

		output := strings.Fields(data[i][1])
		//fmt.Println(output)
		num := 0
		for j := 0; j < len(output); j++ {
			for k := 0; k < len(sigOrder); k++ {
				reg := regexp.MustCompile("^[" + sigOrder[k] + "]{" + strconv.Itoa(len(sigOrder[k])) + "}$")
				if reg.MatchString(output[j]) {
					//fmt.Println(k)

					num = num + k*int(math.Pow10(len(output)-1-j))
					break
				}
			}
		}
		//fmt.Println(num)
		total = total + num
		//break

	}

	t2 := time.Now()
	duration := t2.Sub(t)
	fmt.Println("Problem2", "total", total, "duration", duration)
}
