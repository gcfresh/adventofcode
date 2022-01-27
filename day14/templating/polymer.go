package templating

import (
	"fmt"
	"math"
)

type Polymer struct {
	Template       string
	InsertionRules map[string]string
}

func (p *Polymer) RunStep(s string) string {
	newString := ""
	for i := 0; i < len(s)-1; i++ {
		newString = newString + string(s[i]) + p.InsertionRules[s[i:i+2]]
	}
	newString = newString + string(s[len(s)-1])
	return newString
}

func (p *Polymer) RunStep2(m map[string]int64) map[string]int64 {
	if len(m) == 0 {
		m = make(map[string]int64)
		for i := 0; i < len(p.Template)-1; i++ {
			m[p.Template[i:i+2]]++
		}
	}
	newMap := make(map[string]int64)

	for i, v := range m {
		newChar := p.InsertionRules[i]
		newCode1 := string(i[0]) + newChar
		newCode2 := newChar + string(i[1])
		newMap[newCode1] = newMap[newCode1] + v
		newMap[newCode2] = newMap[newCode2] + v
	}

	return newMap
}

func (p *Polymer) Count(s string) int {
	counts := make(map[string]int)
	for i := 0; i < len(s); i++ {
		counts[string(s[i])]++
	}

	max := 0
	maxS := ""
	min := math.MaxInt32
	minS := ""
	for i, c := range counts {
		if c > max {
			max = c
			maxS = i
		}
		if c < min {
			min = c
			minS = i
		}
	}

	fmt.Println("minS", minS, min, "maxS", maxS, max)
	return max - min
}

func (p *Polymer) Count2(m map[string]int64) int64 {
	counts := make(map[string]int64)
	for i, v := range m {
		counts[string(i[0])] = counts[string(i[0])] + v
		counts[string(i[1])] = counts[string(i[1])] + v
	}

	max := int64(0)
	maxS := ""
	min := int64(math.MaxInt64)
	minS := ""
	for i, c := range counts {
		if c > max {
			max = c
			maxS = i
		}
		if c < min {
			min = c
			minS = i
		}
	}

	if minS == string(p.Template[0]) {
		fmt.Println("dec min1")
		min++
	}

	if minS == string(p.Template[len(p.Template)-1]) {
		fmt.Println("dec min2")
		min++
	}

	if maxS == string(p.Template[0]) {
		fmt.Println("dec max1")
		max++
	}

	if maxS == string(p.Template[len(p.Template)-1]) {
		fmt.Println("dec max2")
		max++
	}

	fmt.Println("minS", minS, min, "maxS", maxS, max)
	return (max - min) / 2
}

type Point struct {
	X int
	Y int
}
