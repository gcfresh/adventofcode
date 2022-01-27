package syntax_solver

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var brackets = map[uint8]uint8{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var points = map[uint8]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		queue := []uint8{}
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(s); i++ {
			if _, isOpen := brackets[s[i]]; isOpen {
				queue = append(queue, (s[i]))
			} else {
				if len(queue) > 0 &&
					brackets[queue[len(queue)-1]] == s[i] {
					queue = queue[0 : len(queue)-1]
				} else {
					//invalid
					//fmt.Println(queue, s[i])
					//fmt.Println(string(s[i]), points[s[i]])
					score = score + points[s[i]]
					break
				}
			}
		}

	}
	//fmt.Println(heightMap)

	fmt.Println("Problem1", "score", score)
}

var closeBrackets = map[uint8]uint8{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}
var closePoints = map[uint8]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scores := []int{}
	for scanner.Scan() {
		queue := []uint8{}
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		corrupted := false
		for i := 0; i < len(s); i++ {
			if _, isOpen := brackets[s[i]]; isOpen {
				queue = append(queue, (s[i]))
			} else {
				if len(queue) > 0 &&
					brackets[queue[len(queue)-1]] == s[i] {
					queue = queue[0 : len(queue)-1]
				} else {
					corrupted = true
					break
				}
			}
		}

		if !corrupted {
			// now go back through the queue
			//printQueue(queue)
			score := 0
			for i := len(queue) - 1; i >= 0; i-- {
				score = score*5 + closePoints[queue[i]]
			}
			scores = append(scores, score)
		}

	}

	sort.Ints(scores)
	fmt.Println(len(scores))
	//fmt.Println(scores)
	fmt.Println("Problem2", "score", scores[len(scores)/2])
}

func printQueue(queue []uint8) {
	s := ""
	for i := 0; i < len(queue); i++ {
		s = s + string(queue[i]) + ","
	}
	s = strings.TrimSuffix(s, ",")
	fmt.Println(s)
}
