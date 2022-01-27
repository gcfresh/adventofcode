package thermals

import "fmt"

type Thermal struct {
	//Coordinates []Point
	Map [][]int
}
type Point struct {
	X int
	Y int
}

func NewThermal(coords []Point, lenx, leny int) Thermal {
	t := Thermal{}
	t.Map = make([][]int, leny)
	for i := 0; i < leny; i++ {
		t.Map[i] = make([]int, lenx)
	}
	for _, c := range coords {
		t.Map[c.Y][c.X] = 1
	}
	return t
}

func (t *Thermal) PrintMap() {
	fmt.Println()
	for i := 0; i < len(t.Map); i++ {
		for j := 0; j < len(t.Map[i]); j++ {
			if t.Map[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func (t *Thermal) FoldOnY(y int) {
	for i := y + 1; i < len(t.Map); i++ {
		for j := 0; j < len(t.Map[i]); j++ {
			if t.Map[i][j] == 1 && y-(i-y) >= 0 {
				t.Map[y-(i-y)][j] = 1
			}
		}
	}
	t.Map = t.Map[:y][:]
}
func (t *Thermal) FoldOnX(x int) {
	for i := 0; i < len(t.Map); i++ {
		for j := x + 1; j < len(t.Map[i]); j++ {
			if t.Map[i][j] == 1 && x-(j-x) >= 0 {
				t.Map[i][x-(j-x)] = 1
			}
		}
	}
	for i := 0; i < len(t.Map); i++ {
		t.Map[i] = t.Map[i][:x]
	}

}

func (t *Thermal) Count() interface{} {
	sum := 0
	for i := 0; i < len(t.Map); i++ {
		for j := 0; j < len(t.Map[i]); j++ {
			if t.Map[i][j] == 1 {
				sum++
			}
		}
	}
	return sum
}
