package octopus

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	energies := [][]int{}
	line := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		energies = append(energies, make([]int, len(s)))

		for i := 0; i < len(s); i++ {
			e, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal("bad energy")
			}
			energies[line][i] = e
		}
		line++
	}
	flashes := 0
	for step := 0; step < 100; step++ {
		flashes = flashes + runStep(energies)
	}

	fmt.Println("Problem1", "flashes", flashes)
}

func runStep(energies [][]int) int {
	flashCount := 0
	for i := 0; i < len(energies); i++ {
		for j := 0; j < len(energies[i]); j++ {
			increaseEnergyRec(i, j, energies)
		}
	}

	for i := 0; i < len(energies); i++ {
		for j := 0; j < len(energies[i]); j++ {
			if energies[i][j] > 9 {
				energies[i][j] = 0
				flashCount++
			}
		}
	}
	return flashCount
}

func increaseEnergyRec(i, j int, energies [][]int) {
	energies[i][j]++
	if energies[i][j] == 10 {
		// increment left side
		if i > 0 {
			increaseEnergyRec(i-1, j, energies)
			if j > 0 {
				increaseEnergyRec(i-1, j-1, energies)
			}
			if j < len(energies[i])-1 {
				increaseEnergyRec(i-1, j+1, energies)
			}
		}

		// increment right side
		if i < len(energies)-1 {
			increaseEnergyRec(i+1, j, energies)
			if j > 0 {
				increaseEnergyRec(i+1, j-1, energies)
			}
			if j < len(energies[i])-1 {
				increaseEnergyRec(i+1, j+1, energies)
			}
		}

		// increment middle row
		if j > 0 {
			increaseEnergyRec(i, j-1, energies)
		}
		if j < len(energies[i])-1 {
			increaseEnergyRec(i, j+1, energies)
		}
	}
}
func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	energies := [][]int{}
	line := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		energies = append(energies, make([]int, len(s)))

		for i := 0; i < len(s); i++ {
			e, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal("bad energy")
			}
			energies[line][i] = e
		}
		line++
	}
	size := len(energies) * len(energies[0])
	flashes := 0
	step := 0
	for flashes < size {
		flashes = runStep(energies)
		step++
	}

	fmt.Println("Problem2", "flashes", flashes, "size", size, "step", step)
}

func printQueue(queue []uint8) {
	s := ""
	for i := 0; i < len(queue); i++ {
		s = s + string(queue[i]) + ","
	}
	s = strings.TrimSuffix(s, ",")
	fmt.Println(s)
}
