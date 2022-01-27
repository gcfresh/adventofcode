package thermals

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fold struct {
	Axis  string
	Index int
}

func Problem1(input string) {
	folds, thermals := loadData(input)
	numToRun := 1
	for i := 0; i < numToRun; i++ {
		if folds[i].Axis == "x" {
			thermals.FoldOnX(folds[i].Index)
		} else {
			thermals.FoldOnY(folds[i].Index)
		}

	}
	//thermals.PrintMap()
	fmt.Println("Problem1", "points", thermals.Count())
}

func Problem2(input string) {
	folds, thermals := loadData(input)
	numToRun := len(folds)
	for i := 0; i < numToRun; i++ {
		if folds[i].Axis == "x" {
			thermals.FoldOnX(folds[i].Index)
		} else {
			thermals.FoldOnY(folds[i].Index)
		}

	}
	thermals.PrintMap()
	fmt.Println("Problem2", "points", thermals.Count())
}

func loadData(input string) ([]Fold, Thermal) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	coords := []Point{}
	maxX := 0
	maxY := 0
	folds := []Fold{}
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		ss := strings.Split(s, ",")
		if len(ss) == 2 {
			x, _ := strconv.Atoi(ss[0])
			y, _ := strconv.Atoi(ss[1])
			coords = append(coords, Point{X: x, Y: y})

			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		} else if len(s) == 0 {
			//do nothing
		} else {
			f := strings.Split(s, "=")
			//fmt.Println("here", f[0])
			c := f[0][len(f[0])-1]
			if c == 'x' {
				x, _ := strconv.Atoi(f[1])
				folds = append(folds, Fold{Axis: "x", Index: x})
			} else {
				y, _ := strconv.Atoi(f[1])
				folds = append(folds, Fold{Axis: "y", Index: y})

			}
		}
	}

	thermals := NewThermal(coords, maxX+1, maxY+1)
	return folds, thermals
}
