package beacon_scanner

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ProbeScanner struct {
	ID           int
	Beacons      []Point
	Orientations [24][]Point
	Location     Point
}

func (ps *ProbeScanner) FillOrientations() {
	for i := 0; i < len(ps.Orientations); i++ {
		ps.Orientations[i] = make([]Point, len(ps.Beacons))
	}
	for i := 0; i < len(ps.Orientations); i++ {
		for j := 0; j < len(ps.Beacons); j++ {
			ps.Orientations[i][j] = ps.Beacons[j]
		}
	}
	for i := 0; i < len(ps.Beacons); i++ {
		// 0:	x,	y,	z
		// 1:	x,	z,	-y
		// 2:	x,	-y,	-z
		// 3:	x,	-z,	y
		ps.Orientations[1][i].Y = ps.Beacons[i].Z
		ps.Orientations[1][i].Z = -ps.Beacons[i].Y

		ps.Orientations[2][i].Y = -ps.Beacons[i].Y
		ps.Orientations[2][i].Z = -ps.Beacons[i].Z

		ps.Orientations[3][i].Y = -ps.Beacons[i].Z
		ps.Orientations[3][i].Z = ps.Beacons[i].Y

		// 4:	-x, y,	-z
		// 5:	-x,	-z,	-y
		// 6:	-x,	-y,	z
		// 7:	-x,	z,	y
		ps.Orientations[4][i].X = -ps.Beacons[i].X
		ps.Orientations[4][i].Z = -ps.Beacons[i].Z

		ps.Orientations[5][i].X = -ps.Beacons[i].X
		ps.Orientations[5][i].Y = -ps.Beacons[i].Z
		ps.Orientations[5][i].Z = -ps.Beacons[i].Y

		ps.Orientations[6][i].X = -ps.Beacons[i].X
		ps.Orientations[6][i].Y = -ps.Beacons[i].Y

		ps.Orientations[7][i].X = -ps.Beacons[i].X
		ps.Orientations[7][i].Y = ps.Beacons[i].Z
		ps.Orientations[7][i].Z = ps.Beacons[i].Y

		// 8:	y,	-x,	z
		// 9:	y,	z,	x
		// 10:	y,	x,	-z
		// 11:	y,	-z,	-x

		ps.Orientations[8][i].X = ps.Beacons[i].Y
		ps.Orientations[8][i].Y = -ps.Beacons[i].X

		ps.Orientations[9][i].X = ps.Beacons[i].Y
		ps.Orientations[9][i].Y = ps.Beacons[i].Z
		ps.Orientations[9][i].Z = ps.Beacons[i].X

		ps.Orientations[10][i].X = ps.Beacons[i].Y
		ps.Orientations[10][i].Y = ps.Beacons[i].X
		ps.Orientations[10][i].Z = -ps.Beacons[i].Z

		ps.Orientations[11][i].X = ps.Beacons[i].Y
		ps.Orientations[11][i].Y = -ps.Beacons[i].Z
		ps.Orientations[11][i].Z = -ps.Beacons[i].X

		// 12:	-y,	x,	z
		// 13:	-y,	z,	-x
		// 14:	-y,	-x,	-z
		// 15:	-y,	-z,	x

		ps.Orientations[12][i].X = -ps.Beacons[i].Y
		ps.Orientations[12][i].Y = ps.Beacons[i].X

		ps.Orientations[13][i].X = -ps.Beacons[i].Y
		ps.Orientations[13][i].Y = ps.Beacons[i].Z
		ps.Orientations[13][i].Z = -ps.Beacons[i].X

		ps.Orientations[14][i].X = -ps.Beacons[i].Y
		ps.Orientations[14][i].Y = -ps.Beacons[i].X
		ps.Orientations[14][i].Z = -ps.Beacons[i].Z

		ps.Orientations[15][i].X = -ps.Beacons[i].Y
		ps.Orientations[15][i].Y = -ps.Beacons[i].Z
		ps.Orientations[15][i].Z = ps.Beacons[i].X

		// 16:	z,	y,	-x
		// 17:	z,	-x,	-y
		// 18:	z,	-y,	x
		// 19:	z,	x,	y
		ps.Orientations[16][i].X = ps.Beacons[i].Z
		ps.Orientations[16][i].Z = -ps.Beacons[i].X

		ps.Orientations[17][i].X = ps.Beacons[i].Z
		ps.Orientations[17][i].Y = -ps.Beacons[i].X
		ps.Orientations[17][i].Z = -ps.Beacons[i].Y

		ps.Orientations[18][i].X = ps.Beacons[i].Z
		ps.Orientations[18][i].Y = -ps.Beacons[i].Y
		ps.Orientations[18][i].Z = ps.Beacons[i].X

		ps.Orientations[19][i].X = ps.Beacons[i].Z
		ps.Orientations[19][i].Y = ps.Beacons[i].X
		ps.Orientations[19][i].Z = ps.Beacons[i].Y

		// 20:	-z,	y,	x
		// 21:	-z,	x,	-y
		// 22:	-z,	-y,	-x
		// 23:	-z,	-x,	y
		ps.Orientations[20][i].X = -ps.Beacons[i].Z
		ps.Orientations[20][i].Z = ps.Beacons[i].X

		ps.Orientations[21][i].X = -ps.Beacons[i].Z
		ps.Orientations[21][i].Y = ps.Beacons[i].X
		ps.Orientations[21][i].Z = -ps.Beacons[i].Y

		ps.Orientations[22][i].X = -ps.Beacons[i].Z
		ps.Orientations[22][i].Y = -ps.Beacons[i].Y
		ps.Orientations[22][i].Z = -ps.Beacons[i].X

		ps.Orientations[23][i].X = -ps.Beacons[i].Z
		ps.Orientations[23][i].Y = -ps.Beacons[i].X
		ps.Orientations[23][i].Z = ps.Beacons[i].Y
	}
}

type Point struct {
	X int
	Y int
	Z int
}

func SolveBoth(input string) {
	probeScanners := loadData(input)
	ref := probeScanners[0]
	var sharedBeacons []Point
	queue := make([]int, len(probeScanners)-1)
	for i := 0; i < len(queue); i++ {
		queue[i] = i + 1
	}

	foundScanners := []int{0}

	sharedBeacons = ref.Beacons
	for len(queue) > 0 {
		scannerToLocate := queue[0]
		queue = queue[1:]
		found := false
		for i := 0; i < len(foundScanners); i++ {
			correctedBeacons, shift := getSharedBeacons(probeScanners[foundScanners[i]], probeScanners[scannerToLocate])
			if len(correctedBeacons) >= 12 {
				probeScanners[scannerToLocate].Beacons = correctedBeacons
				sharedBeacons = append(sharedBeacons, correctedBeacons...)
				foundScanners = append(foundScanners, scannerToLocate)
				probeScanners[scannerToLocate].Location = shift
				found = true
				break
			}
		}
		if !found {
			queue = append(queue, scannerToLocate)
		}
		//fmt.Println(queue, scannerToLocate)
	}
	sharedBeacons = removeDuplicates(sharedBeacons)

	maxDist := 0
	for i := 0; i < len(probeScanners); i++ {
		for j := 0; j < len(probeScanners); j++ {
			if i != j {
				dist := int(math.Abs(float64(probeScanners[i].Location.X-probeScanners[j].Location.X)) +
					math.Abs(float64(probeScanners[i].Location.Y-probeScanners[j].Location.Y)) +
					math.Abs(float64(probeScanners[i].Location.Z-probeScanners[j].Location.Z)))
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
	}

	fmt.Println("Problem1", "count", len(sharedBeacons))
	fmt.Println("Problem2", "maxDist", maxDist)
}

func getSharedBeacons(ref, s ProbeScanner) ([]Point, Point) {
	// go through and check how many overlap if beacon j is beacon i
	for i := 0; i < len(ref.Beacons); i++ {
		for j := 0; j < len(s.Beacons); j++ {
			beacons, shift := getSharedBeaconsGivenRef(i, j, ref, s)
			if len(beacons) >= 12 {
				return beacons, shift
			}
		}
	}

	return nil, Point{}
}

func getSharedBeaconsGivenRef(refIndex, sIndex int, ref, s ProbeScanner) ([]Point, Point) {
	// 24 orientations because every face (6 faces) can be rotated 4 different way
	// adjust values relative

	// trying default orientation

	// var for orientation to test against
	// var for orientation with max beacons
	translatedBeacons := make([]Point, len(s.Beacons))

	maxMatches := 0

	shift := Point{}
	for i := 0; i < len(s.Orientations); i++ {
		matchCount := 0
		shift.X = -s.Orientations[i][sIndex].X + ref.Beacons[refIndex].X
		shift.Y = -s.Orientations[i][sIndex].Y + ref.Beacons[refIndex].Y
		shift.Z = -s.Orientations[i][sIndex].Z + ref.Beacons[refIndex].Z
		for j := 0; j < len(s.Beacons); j++ {
			translatedBeacons[j].X = s.Orientations[i][j].X + shift.X
			translatedBeacons[j].Y = s.Orientations[i][j].Y + shift.Y
			translatedBeacons[j].Z = s.Orientations[i][j].Z + shift.Z
		}
		for j := 0; j < len(translatedBeacons); j++ {
			for k := 0; k < len(ref.Beacons); k++ {
				if translatedBeacons[j] == ref.Beacons[k] {
					matchCount++
					break
				}
			}
		}
		if matchCount > maxMatches {
			maxMatches = matchCount
		}
		if maxMatches >= 12 {
			maxMatches = matchCount
			//fmt.Println("maxMatches", maxMatches, "translatedBeacons", translatedBeacons, "shift", shift)
			return translatedBeacons, shift
		}
	}
	return nil, Point{}
}

func loadData(input string) []ProbeScanner {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	s := ""
	var probeScanners []ProbeScanner
	count := 0
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(s) <= 0 {
			continue
		}
		if s == "break" {
			break
		}

		re := regexp.MustCompile(`--- scanner \d+ ---`)
		if re.MatchString(s) {
			probeScanners = append(probeScanners, ProbeScanner{ID: count})
			count++
		} else {
			coords := strings.Split(s, ",")
			b := Point{}
			b.X, _ = strconv.Atoi(coords[0])
			b.Y, _ = strconv.Atoi(coords[1])
			b.Z, _ = strconv.Atoi(coords[2])
			probeScanners[len(probeScanners)-1].Beacons = append(probeScanners[len(probeScanners)-1].Beacons, b)
		}
	}

	for i := 0; i < len(probeScanners); i++ {
		probeScanners[i].FillOrientations()
	}

	return probeScanners
}

func removeDuplicates(intSlice []Point) []Point {
	keys := make(map[Point]bool)
	list := []Point{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
