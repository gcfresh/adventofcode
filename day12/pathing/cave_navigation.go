package pathing

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Problem1(input string) {
	caves := loadCaves(input)
	numPaths := caves["start"].NavigateToEnd(1)

	fmt.Println("Problem1", "numPaths", numPaths)
}

func Problem2(input string) {
	caves := loadCaves(input)
	numPaths := caves["start"].NavigateToEnd(2)

	fmt.Println("Problem2", "numPaths", numPaths)
}

func loadCaves(input string) map[string]*Cave {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	caves := make(map[string]*Cave)
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		ss := strings.Split(s, "-")

		// add cave if it doesn't exist
		if _, ok := caves[ss[0]]; !ok {
			caves[ss[0]] = NewCave(ss[0])
		}
		if _, ok := caves[ss[1]]; !ok {
			caves[ss[1]] = NewCave(ss[1])
		}

		// connect caves
		caves[ss[0]].Connect(caves[ss[1]])
		//caves[ss[0]].Print()
	}

	//for _, c := range caves {
	//	c.Print()
	//}
	////fmt.Println()

	return caves
}
