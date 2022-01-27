package density

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Fold struct {
	Axis  string
	Index int
}

func Problem1(input string) {
	densityMap := loadData(input)
	costMap := make([][]int, len(densityMap))
	for i := 0; i < len(densityMap); i++ {
		costMap[i] = make([]int, len(densityMap[i]))
		for j := 0; j < len(densityMap[i]); j++ {
			if i == 0 {
				if j == 0 {
					costMap[i][j] = 0
				} else {
					costMap[i][j] = costMap[i][j-1] + densityMap[i][j]
				}
			} else {
				if j == 0 {
					costMap[i][j] = costMap[i-1][j] + densityMap[i][j]
				} else {
					costMap[i][j] = int(math.Min(float64(costMap[i-1][j]), float64(costMap[i][j-1]))) + densityMap[i][j]
				}
			}
		}
	}

	fmt.Println("Problem1", "min", costMap[len(costMap)-1][len(costMap[0])-1])
}

func Problem2(input string) {
	densityMap := loadData(input)
	densityMap = scaleMap(densityMap, 5)
	//PrintMap(densityMap)
	costMap := make([][]int, len(densityMap))
	for i := 0; i < len(densityMap); i++ {
		costMap[i] = make([]int, len(densityMap[i]))
		for j := 0; j < len(densityMap[i]); j++ {
			if i == 0 {
				if j == 0 {
					costMap[i][j] = 0
				} else {
					costMap[i][j] = costMap[i][j-1] + densityMap[i][j]
				}
			} else {
				if j == 0 {
					costMap[i][j] = costMap[i-1][j] + densityMap[i][j]
				} else {
					costMap[i][j] = int(math.Min(float64(costMap[i-1][j]), float64(costMap[i][j-1]))) + densityMap[i][j]
				}
			}
		}
	}
	done := false
	count := 0
	adjustments := 0
	lastValue := costMap[len(costMap)-1][len(costMap)-1]
	for !done {
		done = true
		count++
		if count%1000 == 0 {
			fmt.Println(count)
		}
		for i := 0; i < len(costMap); i++ {
			for j := 0; j < len(costMap[i]); j++ {
				if i > 0 {
					if costMap[i][j] < costMap[i-1][j]-densityMap[i-1][j] {
						done = false
						costMap[i-1][j] = costMap[i][j] + densityMap[i-1][j]
						adjustments++
						for k := j; k > 0; k-- {
							if costMap[i-1][k] < costMap[i-1][k-1]-densityMap[i-1][k-1] {
								costMap[i-1][k-1] = costMap[i-1][k] + densityMap[i-1][k-1]
								adjustments++
							} else {
								break
							}
						}
						for k := j; k < len(costMap[i])-1; k++ {
							if costMap[i-1][k] < costMap[i-1][k+1]-densityMap[i-1][k+1] {
								costMap[i-1][k+1] = costMap[i-1][k] + densityMap[i-1][k+1]
								adjustments++
							} else {
								break
							}
						}
					}
				}

				if i < len(costMap)-1 {
					if costMap[i][j] < costMap[i+1][j]-densityMap[i+1][j] {
						done = false
						costMap[i+1][j] = costMap[i][j] + densityMap[i+1][j]
						adjustments++
						for k := j; k > 0; k-- {
							if costMap[i+1][k] < costMap[i+1][k-1]-densityMap[i+1][k-1] {
								costMap[i+1][k-1] = costMap[i+1][k] + densityMap[i+1][k-1]
								adjustments++
							} else {
								break
							}
						}
						for k := j; k < len(costMap[i])-1; k++ {
							if costMap[i+1][k] < costMap[i+1][k+1]-densityMap[i+1][k+1] {
								costMap[i+1][k+1] = costMap[i+1][k] + densityMap[i+1][k+1]
								adjustments++
							} else {
								break
							}
						}
					}
				}

				if j > 0 {
					if costMap[i][j] < costMap[i][j-1]-densityMap[i][j-1] {
						done = false
						costMap[i][j-1] = costMap[i][j] + densityMap[i][j-1]
						adjustments++

						for k := i; k > 0; k-- {
							if costMap[k][j-1] < costMap[k-1][j-1]-densityMap[k-1][j-1] {
								costMap[k-1][j-1] = costMap[k][j-1] + densityMap[k-1][j-1]
								adjustments++
							} else {
								break
							}
						}
						for k := i; k < len(costMap)-1; k++ {
							if costMap[k][j-1] < costMap[k+1][j-1]-densityMap[k+1][j-1] {
								costMap[k+1][j-1] = costMap[k][j-1] + densityMap[k+1][j-1]
								adjustments++
							} else {
								break
							}
						}
					}
				}

				if j < len(costMap[i])-1 {
					if costMap[i][j] < costMap[i][j+1]-densityMap[i][j+1] {
						done = false
						costMap[i][j+1] = costMap[i][j] + densityMap[i][j+1]
						adjustments++

						for k := i; k > 0; k-- {
							if costMap[k][j+1] < costMap[k-1][j+1]-densityMap[k-1][j+1] {
								costMap[k-1][j+1] = costMap[k][j+1] + densityMap[k-1][j+1]
								adjustments++
							} else {
								break
							}
						}
						for k := i; k < len(costMap)-1; k++ {
							if costMap[k][j+1] < costMap[k+1][j+1]-densityMap[k+1][j+1] {
								costMap[k+1][j+1] = costMap[k][j+1] + densityMap[k+1][j+1]
								adjustments++
							} else {
								break
							}
						}
					}
				}
			}
		}
		if lastValue != costMap[len(costMap)-1][len(costMap)-1] {
			lastValue = costMap[len(costMap)-1][len(costMap)-1]
			fmt.Println("new lastvalue", lastValue)
		}

	}

	//PrintMap(costMap)
	fmt.Println(adjustments, count)
	// old guess 3019
	fmt.Println("Problem2", "min", costMap[len(costMap)-1][len(costMap[0])-1])
}

func scaleMap(densityMap [][]int, scale int) [][]int {
	newMap := make([][]int, len(densityMap)*scale)
	//fmt.Println(len(newMap))
	for i := 0; i < len(newMap); i++ {
		newMap[i] = make([]int, len(densityMap)*scale)
		for j := 0; j < len(newMap[i]); j++ {
			newMap[i][j] = densityMap[i%len(densityMap)][j%len(densityMap[0])] + i/len(densityMap) + j/(len(densityMap[0]))
			for newMap[i][j] > 9 {
				newMap[i][j] = newMap[i][j] - 9
			}
		}
	}
	//fmt.Println(len(newMap[0]))

	return newMap
}

func loadData(input string) [][]int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var densityMap [][]int
	line := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(s) == 0 {
			continue
		}

		densityMap = append(densityMap, make([]int, len(s)))
		for i := 0; i < len(s); i++ {
			densityMap[line][i], _ = strconv.Atoi(string(s[i]))
		}
		line++
	}

	return densityMap
}
func PrintMap(m [][]int) {
	fmt.Println()
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] < 10 {
				fmt.Print("0")
			}
			if m[i][j] < 100 {
				fmt.Print("0")
			}
			fmt.Print(m[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
