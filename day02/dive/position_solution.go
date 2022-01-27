package dive

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Forward = "forward"
	Up      = "up"
	Down    = "down"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	depth := 0
	forward := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		ss := strings.Split(s, " ")
		if len(ss) != 2 {
			log.Fatal("invalid split size")
		}

		shift, err := strconv.Atoi(ss[1])
		if err != nil {
			log.Fatal(err)
		}

		switch ss[0] {
		case Up:
			depth = depth - shift
		case Down:
			depth = depth + shift
		case Forward:
			forward = forward + shift
		default:
			log.Fatal("invalid command")
		}
	}

	fmt.Println("Problem1", "depth", depth, "forward", forward, "answer", forward*depth)
}

func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	depth := 0
	forward := 0
	aim := 0
	for scanner.Scan() {
		s := scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
		ss := strings.Split(s, " ")
		if len(ss) != 2 {
			log.Fatal("invalid split size")
		}

		shift, err := strconv.Atoi(ss[1])
		if err != nil {
			log.Fatal(err)
		}

		switch ss[0] {
		case Up:
			//depth = depth - shift
			aim = aim - shift
		case Down:
			//depth = depth + shift
			aim = aim + shift
		case Forward:
			forward = forward + shift
			depth = depth + aim*shift
		default:
			log.Fatal("invalid command")
		}
	}

	fmt.Println("Problem2", "depth", depth, "forward", forward, "answer", forward*depth)
}
