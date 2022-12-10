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

func main() {
	input := helpers.GetInput("./day2/input.txt")

	games := strings.Split(input, "\n")

	Part1(games)
	Part2(games)
}

func Part1(games []string) {
	points := 0

	for _, g := range games {
		score := Scores[g]
		points += score
	}

	fmt.Println(points)
}

func Part2(games []string) {
	points := 0

	for _, g := range games {
		score := ScoresRigged[g]
		points += score
	}

	fmt.Println(points)
}
