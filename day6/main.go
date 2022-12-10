package main

import (
	"aoc2022/helpers"
	"fmt"
)

func main() {
	stream := helpers.GetInput("./day6/input.txt")

	Part1(stream)
	Part2(stream)
}

func Part2(stream string) {
	for i := 13; i < len(stream); i++ {
		section := stream[i-13 : i+1]

		if isUnique(section) {
			fmt.Println(i + 1)
			break
		}
	}
}

func Part1(stream string) {
	for i := 3; i < len(stream); i++ {
		section := stream[i-3 : i+1]

		if isUnique(section) {
			fmt.Println(i + 1)
			break
		}
	}
}

func isUnique(substr string) bool {
	store := make(map[rune]bool)

	for _, c := range substr {
		_, ok := store[c]

		if ok {
			return false
		}

		store[c] = true
	}

	return true

}
