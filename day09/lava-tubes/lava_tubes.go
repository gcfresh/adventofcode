package lava_tubes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightMap := [][]int{}
	lineIndex := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		heightMap = append(heightMap, make([]int, len(s)))
		for i := 0; i < len(s); i++ {
			num, _ := strconv.Atoi(string(s[i]))
			heightMap[lineIndex][i] = num
		}
		lineIndex++
	}
	//fmt.Println(heightMap)

	lows := []int{}
	sum := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if (i == 0 || heightMap[i][j] < heightMap[i-1][j]) &&
				(i == len(heightMap)-1 || heightMap[i][j] < heightMap[i+1][j]) &&
				(j == 0 || heightMap[i][j] < heightMap[i][j-1]) &&
				(j == len(heightMap[i])-1 || heightMap[i][j] < heightMap[i][j+1]) {

				lows = append(lows, heightMap[i][j])
				sum = sum + heightMap[i][j] + 1
			}
		}
	}

	fmt.Println("Problem1", "len(lows)", len(lows), "sum", sum)
}

func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightMap := [][]int{}
	lineIndex := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		heightMap = append(heightMap, make([]int, len(s)))
		for i := 0; i < len(s); i++ {
			num, _ := strconv.Atoi(string(s[i]))
			heightMap[lineIndex][i] = num
		}
		lineIndex++
	}
	//fmt.Println(heightMap)

	basins := []int{}
	result := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if (i == 0 || heightMap[i][j] < heightMap[i-1][j]) &&
				(i == len(heightMap)-1 || heightMap[i][j] < heightMap[i+1][j]) &&
				(j == 0 || heightMap[i][j] < heightMap[i][j-1]) &&
				(j == len(heightMap[i])-1 || heightMap[i][j] < heightMap[i][j+1]) {

				var size int
				size = findBasinSize(i, j, heightMap, nil)
				basins = append(basins, size)
			}
		}
	}

	sort.Ints(basins)
	fmt.Println(basins)
	basins = basins[len(basins)-3:]
	result = basins[0] * basins[1] * basins[2]

	fmt.Println("Problem2", "basins", basins, "result", result)
}

func findBasinSize(i int, j int, heightMap [][]int, basinCoords map[[2]int]bool) int {
	if i < 0 ||
		i >= len(heightMap) ||
		j < 0 ||
		j >= len(heightMap[i]) {
		return 0
	}

	if heightMap[i][j] == 9 {
		return 0
	}

	if basinCoords == nil {
		basinCoords = make(map[[2]int]bool)
	}

	if basinCoords[[2]int{i, j}] == false {
		basinCoords[[2]int{i, j}] = true
	} else {
		return 0
	}

	return 1 +
		findBasinSize(i-1, j, heightMap, basinCoords) +
		findBasinSize(i+1, j, heightMap, basinCoords) +
		findBasinSize(i, j-1, heightMap, basinCoords) +
		findBasinSize(i, j+1, heightMap, basinCoords)
}
