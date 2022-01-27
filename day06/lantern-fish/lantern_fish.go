package lantern_fish

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = ','
	data, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	fish := make([]int, len(data))
	for i, d := range data {
		fish[i], _ = strconv.Atoi(d)
	}

	const numDays = 18
	for i := 0; i < numDays; i++ {
		fishToAdd := 0
		for j := 0; j < len(fish); j++ {
			if fish[j] == 0 {
				fish[j] = 6
				fishToAdd++
			} else {
				fish[j]--
			}
		}

		if fishToAdd > 0 {
			newFish := make([]int, fishToAdd)
			for j := 0; j < len(newFish); j++ {
				newFish[j] = 8
			}
			fish = append(fish, newFish...)
		}
	}

	fmt.Println("Problem1", "len(fish)", len(fish))
}

func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = ','
	data, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	fish := make([]int64, 9)
	for _, d := range data {
		num, _ := strconv.Atoi(d)
		fish[num]++
	}

	const numDays = 256

	for i := 0; i < numDays; i++ {
		fishToAdd := fish[0]
		for j := 0; j < len(fish)-1; j++ {
			fish[j] = fish[j+1]
			// fish 0 go to 6
			// other fish decrement
			if j == 6 {
				fish[j] = fish[j] + fishToAdd
			}
		}
		fish[8] = fishToAdd
		//fmt.Println(fish)
	}

	numFish := int64(0)
	for j := 0; j < len(fish); j++ {
		numFish = numFish + fish[j]
	}

	fmt.Println("Problem2", "numFish", numFish)
}
