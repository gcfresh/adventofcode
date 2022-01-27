package bingo

import (
	"strconv"
)

const sideLength = 5

type Board struct {
	Nums         [][]int
	RowStatus    []int
	ColumnStatus []int
	Sum          int
}

func NewBoard(nums [][]string) Board {
	b := Board{}
	b.Nums = make([][]int, sideLength)
	for i := 0; i < sideLength; i++ {
		b.Nums[i] = make([]int, sideLength)
		for j := 0; j < sideLength; j++ {
			b.Nums[i][j], _ = strconv.Atoi(nums[i][j])
			b.Sum = b.Sum + b.Nums[i][j]
		}
	}
	b.RowStatus = make([]int, sideLength)
	b.ColumnStatus = make([]int, sideLength)
	return b
}

func (b *Board) DidWin(num int) bool {
	for i := 0; i < sideLength; i++ {
		for j := 0; j < sideLength; j++ {
			if b.Nums[i][j] == num {
				b.RowStatus[i]++
				b.ColumnStatus[j]++
				b.Sum = b.Sum - num
				return b.RowStatus[i] == sideLength || b.ColumnStatus[j] == sideLength
			}
		}
	}
	return false
}
