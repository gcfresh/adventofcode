package dice

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Problem1(input string) {
	players := loadData(input)
	d := Dice{}
	done := false
	for !done {
		for i := 0; i < len(players) && !done; i++ {
			players[i].Move(d.Roll() + d.Roll() + d.Roll())
			if playerWon(players, 1000) {
				done = true
				break
			}
		}
	}
	loserScore := players[0].Score
	if players[0].Score > players[1].Score {
		loserScore = players[1].Score
	}

	fmt.Println("Problem1", "players", players, "loserScore", loserScore, "rollCount", d.RollCount, "loserScore * d.RollCount", loserScore*d.RollCount)
}

func Problem2(input string) {
	players := loadData(input)

	rollWeights := getRollWeights()
	fmt.Println("rollWeights", rollWeights)

	w1, w2 := playTurn(players[0], players[1], 0, rollWeights, false)
	fmt.Println("Problem2", "players", players, "w1", w1, "w2", w2, "max", int64(math.Max(float64(w1), float64(w2))))
}

func playTurn(p1, p2 Player, move int, rollWeights map[int]int64, moveP1 bool) (int64, int64) {
	if move != 0 { // if 0 then we shouldn't move since this is the start of the game
		if moveP1 {
			p1.Move(move)
		} else {
			p2.Move(move)
		}
	}

	// any winners?
	if p1.Score >= 21 {
		return 1, 0
	}
	if p2.Score >= 21 {
		return 0, 1
	}

	winSum1 := int64(0)
	winSum2 := int64(0)

	for i := 3; i <= 9; i++ { // faster than making an iterator on map
		w1, w2 := playTurn(p1, p2, i, rollWeights, !moveP1)
		winSum1 = winSum1 + w1*rollWeights[i]
		winSum2 = winSum2 + w2*rollWeights[i]
	}
	return winSum1, winSum2
}

func getRollWeights() map[int]int64 {
	rollWeights := make(map[int]int64)
	const max = 3
	for i := 1; i <= max; i++ {
		for j := 1; j <= max; j++ {
			for k := 1; k <= max; k++ {
				rollWeights[i+j+k] = rollWeights[i+j+k] + 1
			}
		}
	}

	// sums
	// 3 -> 1
	// 4 -> 3
	// 5 -> 6
	// 6 -> 7
	// 7 -> 6
	// 8 -> 3
	// 9 -> 1
	return rollWeights
}

func playerWon(players []Player, winScore int) bool {
	for _, p := range players {
		if p.Score >= winScore {
			return true
		}
	}
	return false
}

func loadData(input string) []Player {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var s string
	//wasLoaded := false
	var players []Player
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(s) <= 0 {
			continue
		}

		fields := strings.Fields(s)
		id, _ := strconv.Atoi(fields[1])
		start, _ := strconv.Atoi(fields[4])
		players = append(players, Player{
			ID:      id,
			Start:   start,
			Current: start,
		})
	}

	return players
}
