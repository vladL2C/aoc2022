package main

import (
	"aoc2022/helpers"
	"fmt"
	"strings"
)

func main() {
	input := strings.TrimSpace(helpers.GetInput("./day3/input.txt"))
	fmt.Println("PART 1: ", part1(input))
	fmt.Println("PART 2: ", part2(input))
}

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += getItemValue(findOne(line))
	}
	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for i := 0; i < len(lines)-2; i += 3 {
		sum += getItemValue(findOne2(lines[i], lines[i+1], lines[i+2]))
	}
	return sum
}

func getItemValue(item string) int {
	value := int(item[0])
	if value >= 65 && value <= 90 { // capital letter
		return value - 38
	} else if value >= 97 && value <= 122 { // lower case letter
		return value - 96
	}

	return value
}

func findOne(line string) string {
	mid := len(line) / 2
	left := line[:mid]
	right := line[mid:]

	found := ""

	for _, char := range strings.Split(left, "") {
		if strings.Contains(right, char) {
			found = char
			break
		}
	}

	return found
}

func findOne2(line1, line2, line3 string) string {
	found := ""

	for _, char := range strings.Split(line1, "") {
		if strings.Contains(line2, char) && strings.Contains(line3, char) {
			found = char
			break
		}
	}

	return found
}
