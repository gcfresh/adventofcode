package templating

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Fold struct {
	Axis  string
	Index int
}

func Problem1(input string) {
	p := loadData(input)

	steps := 10
	s := p.Template
	for i := 0; i < steps; i++ {
		s = p.RunStep(s)
		//fmt.Println(s)
	}

	diff := p.Count(s)
	fmt.Println("Problem1", "diff", diff)
}

func Problem2(input string) {
	t1 := time.Now()
	p := loadData(input)

	steps := 40
	m := make(map[string]int64)
	for i := 0; i < steps; i++ {
		m = p.RunStep2(m)
		//fmt.Println(s)
	}

	diff := p.Count2(m)
	t2 := time.Now()
	//fmt.Println("len(s)", len(s))
	d := t2.Sub(t1)
	fmt.Println("Problem2", "diff", diff, "time", d.String())

}

func loadData(input string) Polymer {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	s := scanner.Text()
	p := Polymer{Template: s}
	p.InsertionRules = make(map[string]string)

	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if len(s) == 0 {
			continue
		}

		ss := strings.Split(s, " -> ")
		if len(ss) == 2 {
			p.InsertionRules[ss[0]] = ss[1]

		}
	}

	return p
}
