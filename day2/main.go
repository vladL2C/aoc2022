package main

import (
	"aoc2022/helpers"
	"fmt"
	"strings"
)

var Scores = map[string]int{
	"B X": 1,
	"C Y": 2,
	"A Z": 3,
	"A X": 4,
	"B Y": 5,
	"C Z": 6,
	"C X": 7,
	"A Y": 8,
	"B Z": 9,
}

var ScoresRigged = map[string]int{
	"B X": 1,
	"C X": 2,
	"A X": 3,
	"A Y": 4,
	"B Y": 5,
	"C Y": 6,
	"C Z": 7,
	"A Z": 8,
	"B Z": 9,
}

type RockPaperScissors struct {
	Rounds      []string
	Board       map[string]int
	RiggedBoard map[string]int
}

func (game *RockPaperScissors) Score() int {
	points := 0

	for _, g := range game.Rounds {
		score := game.Board[g]
		points += score
	}

	return points
}

func (game *RockPaperScissors) RiggedScore() int {
	points := 0

	for _, g := range game.Rounds {
		score := game.RiggedBoard[g]
		points += score
	}

	return points
}

func main() {
	input := helpers.GetInput("./day2/input.txt")

	games := strings.Split(input, "\n")

	gamer := RockPaperScissors{
		Rounds:      games,
		Board:       Scores,
		RiggedBoard: ScoresRigged,
	}

	fmt.Println("Part 1:", gamer.Score())
	fmt.Println("Part 2:", gamer.RiggedScore())
}
