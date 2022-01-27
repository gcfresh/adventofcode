package sonar

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currDepth := math.MaxInt32
	increases := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		newDepth, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if currDepth < newDepth {
			increases++
		}
		currDepth = newDepth
	}

	fmt.Println("Problem1", increases)
}

func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	increases := 0
	count := 0
	window := []int{0, 0, 0, 0}
	for scanner.Scan() { //&& count < len(tdata){
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		newDepth, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		//newDepth = tdata[count]

		if count >= 3 {
			oldSum := 0
			newSum := 0
			// shift new value in
			for i := len(window) - 1; i > 0; i-- {
				window[i] = window[i-1]
			}
			window[0] = newDepth
			oldSum = window[1] + window[2] + window[3]
			newSum = window[1] + window[2] + window[0]
			if newSum > oldSum {
				increases++
				//fmt.Println("count", count,"oldSum", oldSum,"newSum", newSum, "window", window)
			}
		} else {
			window[len(window)-2-count] = newDepth
			//fmt.Println("count", count, "window", window)
		}

		count++
	}

	fmt.Println("Problem2", increases)
}

var tdata = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}
