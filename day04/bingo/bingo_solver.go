package bingo

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Forward = "forward"
	Up      = "up"
	Down    = "down"
)

func Problem1(input string) {
	nums, boards := loadData(input)

	done := false
	winner := Board{}
	winningNum := -1
	for i := 0; i < len(nums) && !done; i++ {
		for j := 0; j < len(boards); j++ {
			num, _ := strconv.Atoi(nums[i])
			if boards[j].DidWin(num) {
				done = true
				winner = boards[j]
				winningNum = num
				break
			}
		}
	}

	fmt.Println("Problem1", "winner", winner.Nums, "sum", winner.Sum, "winningNum", winningNum, "answer", winner.Sum*winningNum)
}

func Problem2(input string) {
	nums, boards := loadData(input)

	done := false
	winner := Board{}
	winningNum := -1
	boardsToRemove := make(map[int]interface{})
	for i := 0; i < len(nums) && !done; i++ {
		for j := 0; j < len(boards); j++ {
			num, _ := strconv.Atoi(nums[i])
			if boards[j].DidWin(num) && boardsToRemove[j] != true {
				if len(boards)-len(boardsToRemove) == 1 {
					done = true
					winner = boards[j]
					winningNum = num
					break
				} else {
					boardsToRemove[j] = true
					//fmt.Println("removing", j)
				}
			}
		}
	}

	fmt.Println("Problem2", "winner", winner.Nums, "sum", winner.Sum, "winningNum", winningNum, "answer", winner.Sum*winningNum)
}

func loadData(input string) ([]string, []Board) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	nums, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(nums)
	r.Comma = ' '
	r.TrimLeadingSpace = true
	r.FieldsPerRecord = 5

	boardRecords, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(boardRecords)

	boards := []Board{}
	for i := 0; i < len(boardRecords); i = i + 5 {
		boards = append(boards, NewBoard(boardRecords[i:i+5]))
	}
	return nums, boards
}
