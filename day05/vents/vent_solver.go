package vents

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Problem1(input string) {

	overlaps := countOverlaps(input, false)
	fmt.Println("Problem1", "overlaps", overlaps)

}
func Problem2(input string) {

	overlaps := countOverlaps(input, true)
	fmt.Println("Problem2", "overlaps", overlaps)

}

func countOverlaps(input string, useDiagonals bool) int {
	ventMap := NewMap()

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		ss := strings.Split(s, " -> ")
		if len(ss) != 2 {
			log.Fatal("invalid split size")
		}
		p1 := StringToPoint(strings.Split(ss[0], ","))
		p2 := StringToPoint(strings.Split(ss[1], ","))
		lType := Other
		if p1.X == p2.X {
			lType = Vertical
		} else if p1.Y == p2.Y {
			lType = Horizontal
		} else if math.Abs(float64(p1.Y-p2.Y)) == math.Abs(float64(p1.X-p2.X)) {
			lType = Diagonal
		}

		line := Line{
			Point1:   p1,
			Point2:   p2,
			LineType: lType,
		}

		if !useDiagonals && line.LineType == Diagonal {
			continue
		}
		ventMap.AddLine(line)
	}
	//ventMap.Print()
	overlaps := ventMap.TotalOverlaps()
	return overlaps
}

func StringToPoint(p []string) Point {
	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])
	return Point{X: x, Y: y}
}
