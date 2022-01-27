package vents

import (
	"fmt"
	"math"
)

const (
	Horizontal = "horizontal"
	Vertical   = "vertical"
	Diagonal   = "diagonal"
	Other      = "other"
)

type Line struct {
	Point1   Point
	Point2   Point
	LineType string
}

type Point struct {
	X int
	Y int
}

const sideLength = 1000

type Map struct {
	data [][]int
}

func NewMap() Map {
	m := Map{}
	m.data = make([][]int, sideLength)
	for i := 0; i < sideLength; i++ {
		m.data[i] = make([]int, sideLength)
	}
	return m
}

func (m *Map) TotalOverlaps() int {
	overlaps := 0
	for i := 0; i < sideLength; i++ {
		for j := 0; j < sideLength; j++ {
			if m.data[i][j] > 1 {
				overlaps++
			}
		}
	}
	return overlaps
}

func (m *Map) AddLine(line Line) {
	if line.LineType == Horizontal {
		minX := int(math.Min(float64(line.Point1.X), float64(line.Point2.X)))
		maxX := int(math.Max(float64(line.Point1.X), float64(line.Point2.X)))
		for i := minX; i <= maxX; i++ {
			m.data[line.Point1.Y][i]++
		}
	} else if line.LineType == Vertical {
		minY := int(math.Min(float64(line.Point1.Y), float64(line.Point2.Y)))
		maxY := int(math.Max(float64(line.Point1.Y), float64(line.Point2.Y)))
		for i := minY; i <= maxY; i++ {
			m.data[i][line.Point1.X]++
		}
	} else if line.LineType == Diagonal {
		left := line.Point1
		right := line.Point2
		if line.Point1.X > line.Point2.X {
			left = line.Point2
			right = line.Point1
		}

		//if left.X == 2 && left.Y == 0 {
		//	fmt.Println()
		//	m.Print()
		//	fmt.Println()
		//}

		for i := left.X; i <= right.X; i++ {
			if left.Y < right.Y {
				m.data[left.Y+(i-left.X)][i]++
			} else {
				m.data[left.Y-(i-left.X)][i]++
			}
		}
		//if left.X == 2 && left.Y == 0 {
		//	fmt.Println()
		//	m.Print()
		//	fmt.Println()
		//}

		//minY := int(math.Min(float64(line.Point1.Y), float64(line.Point2.Y)))
		//maxY := int(math.Max(float64(line.Point1.Y), float64(line.Point2.Y)))
		//for i := minY; i <= maxY; i++ {
		//	m.data[i][line.Point1.X]++
		//}
	} else {
		// todo
	}
}

func (m *Map) Print() {
	for i := 0; i < sideLength; i++ {
		fmt.Println(m.data[i])

	}
}
