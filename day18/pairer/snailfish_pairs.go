package pairer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	s := ""
	var p *Node
	//wasLoaded := false
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(s) <= 0 {
			continue
		}

		//fmt.Println(s)
		if p == nil {
			p = loadPair(s)
		} else {
			newPair := loadPair(s)
			p = &Node{Left: p, Right: newPair}
		}
		p.Reduce()
	}
	fmt.Println("Problem1", "count", p.GetMagnitude())
}
func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	s := ""
	var p []*Node
	//wasLoaded := false
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(s) <= 0 {
			continue
		}

		p = append(p, loadPair(s))
	}
	fmt.Println(p[0].GetString())
	num := p[0].Copy()
	fmt.Println(num.Left.Left.Left.Value)
	fmt.Println(num.GetString())

	maxMag := 0
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			if i != j {
				//fmt.Println("testing", i, j)
				pTest := &Node{
					Left:  p[i].Copy(),
					Right: p[j].Copy(),
				}
				pTest.Reduce()
				m := pTest.GetMagnitude()
				if m > maxMag {
					maxMag = m
				}

			}
		}
	}

	fmt.Println("Problem2", "maxMag", maxMag)
}

func loadPair(s string) *Node {
	//fmt.Println("loadPair", s)
	p := &Node{}
	//fmt.Println(s)
	s = s[1 : len(s)-1] // unwrap
	if s[0] == '[' {
		// deeper to go
		bCount := 1
		closeIndex := 0
		for i := 1; i < len(s) && bCount > 0; i++ {
			closeIndex = i
			if s[i] == '[' {
				bCount++
			} else if s[i] == ']' {
				bCount--
			}
		}
		p.Left = loadPair(s[:closeIndex+1]) // +1 is comma index

		if closeIndex < len(s) { // not sure if we need this check
			// check if after comma is value or needs more processing
			p.Right = loadRight(s[closeIndex+2:]) // +2 is after comma index
		}

	} else {
		num := ""
		for i := 0; i < len(s); i++ {
			_, err := strconv.Atoi(string(s[i]))
			if err != nil {
				break
			}
			num = num + string(s[i])
		}
		v, _ := strconv.Atoi(num)
		p.Left = &Node{Value: v, Parent: p}
		p.Right = loadRight(s[1+len(num):])
		p.Right.Parent = p
	}
	return p
}

func loadRight(s string) *Node {
	_, err := strconv.Atoi(string(s[0]))
	if err != nil {
		return loadPair(s)
	} else {
		num := ""
		for i := 0; i < len(s); i++ {
			_, err = strconv.Atoi(string(s[i]))
			if err != nil {
				break
			}
			num = num + string(s[i])
		}
		v, _ := strconv.Atoi(num)

		return &Node{Value: v}
	}
}
