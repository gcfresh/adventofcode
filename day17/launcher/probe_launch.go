package launcher

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Target struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func Solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	s := ""
	target := Target{}
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		re := regexp.MustCompile(`target area: x=(\d+)..(\d+), y=(-\d+)..(-\d+)`)
		ss := re.FindStringSubmatch(s)
		if len(ss) != 5 {
			log.Fatal("not enough data", ss)
		}
		target.MinX, _ = strconv.Atoi(ss[1])
		target.MaxX, _ = strconv.Atoi(ss[2])
		target.MinY, _ = strconv.Atoi(ss[3])
		target.MaxY, _ = strconv.Atoi(ss[4])
	}

	fmt.Println(target)
	maxHeight := 0
	count := 0
	var x, y int
	for vx := 1; vx <= target.MaxX; vx++ {
		for vy := target.MinY; vy <= 0-target.MinY; vy++ {
			//if vx == 6 && vy == 9 {
			h := simulateProbe(vx, vy, 0, 0, 0, target)

			if h >= 0 {
				count++
			}
			//fmt.Println(h, maxHeight)
			if h > maxHeight {
				maxHeight = h
				x = vx
				y = vy
			}
			//}
		}
	}

	fmt.Println("Problem1", "maxHeight", maxHeight, "|", x, y)
	fmt.Println("Problem2", "count", count)
}

func simulateProbe(vx, vy, maxHeight, x int, y int, target Target) int {
	//fmt.Println(x, y, "|", vx, vy, "|", maxHeight)
	x = x + vx
	y = y + vy
	if vx > 0 {
		vx--
	}
	vy--
	maxHeight = int(math.Max(float64(maxHeight), float64(y)))
	// won't hit
	if x > target.MaxX ||
		y < target.MinY {
		//fmt.Println("miss")
		return -1
	}
	// hit?
	if x >= target.MinX &&
		x <= target.MaxX &&
		y >= target.MinY &&
		y <= target.MaxY {
		//fmt.Println("hit")
		return maxHeight
	}

	return simulateProbe(vx, vy, maxHeight, x, y, target)
}
