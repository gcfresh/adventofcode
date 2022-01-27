package submarines

import (
	"encoding/csv"
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
	r := csv.NewReader(file)
	r.Comma = ','
	data, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	crabs := make([]int, len(data))
	fuelCosts := make(map[int]int)
	positions := make(map[int]int)
	for i, d := range data {
		crabs[i], _ = strconv.Atoi(d)
		positions[crabs[i]]++
		fuelCosts[i] = 0
	}

	minCost := math.MaxInt32
	for i, _ := range fuelCosts {
		for j, p := range positions {
			fuelCosts[i] = fuelCosts[i] + int(math.Abs(float64(i-j)))*p
		}
		if minCost > fuelCosts[i] {
			minCost = fuelCosts[i]
		}
	}

	fmt.Println("Problem1", "minCost", minCost)
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

	crabs := make([]int, len(data))
	fuelCosts := make(map[int]int)
	positions := make(map[int]int)
	for i, d := range data {
		crabs[i], _ = strconv.Atoi(d)
		positions[crabs[i]]++
		fuelCosts[i] = 0
	}

	minCost := math.MaxInt32
	for i, _ := range fuelCosts {
		for j, p := range positions {
			dist := int(math.Abs(float64(i - j)))
			for cost := 1; cost <= dist; cost++ {
				fuelCosts[i] = fuelCosts[i] + cost*p
			}
		}
		if minCost > fuelCosts[i] {
			minCost = fuelCosts[i]
		}
	}

	fmt.Println("Problem2", "minCost", minCost)
}
