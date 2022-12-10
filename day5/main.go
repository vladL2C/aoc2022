package main

import (
	"aoc2022/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	move, from, to int
}

func GetCrates() [9][]string {
	return [9][]string{
		{" ", "N", "S", "D", "C", "V", "Q", "T"},
		{" ", " ", " ", " ", " ", "M", "F", "V"},
		{"F", "Q", "W", "D", "P", "N", "H", "M"},
		{" ", " ", " ", "D", "Q", "R", "T", "F"},
		{"R", "F", "M", "N", "Q", "H", "V", "B"},
		{" ", "C", "F", "G", "N", "P", "W", "Q"},
		{" ", " ", "W", "F", "R", "L", "C", "T"},
		{" ", " ", " ", " ", "T", "Z", "N", "S"},
		{"M", "S", "D", "J", "R", "Q", "H", "N"},
	}
}

func main() {
	content := helpers.GetInput("./day5/input.txt")

	moves := strings.Split(content, "\n\n")[1]

	rules := []Rule{}

	for _, move := range strings.Split(moves, "\n") {
		re := regexp.MustCompile("[0-9]+")

		rule := re.FindAllString(move, -1)

		move, _ := strconv.Atoi(rule[0])
		from, _ := strconv.Atoi(rule[1])
		to, _ := strconv.Atoi(rule[2])

		rules = append(rules, Rule{
			move: move,
			from: from - 1,
			to:   to - 1,
		})
	}

	Part1(rules, GetCrates())
	Part2(rules, GetCrates())

}

func Part1(rules []Rule, crates [9][]string) {
	for _, rule := range rules {

		fromCrate := &crates[rule.from]
		moveItems := []string{}

		for i := 0; i < rule.move; i++ {
			v := pop(fromCrate)
			moveItems = append(moveItems, v)
		}

		crates[rule.to] = append(crates[rule.to], moveItems...)

	}

	result := ""

	for _, crate := range crates {
		v := pop(&crate)
		result += v
	}

	fmt.Println("PART 1 RESULT: ", result)
}

func Part2(rules []Rule, crates [9][]string) {

	for _, rule := range rules {

		fromCrate := &crates[rule.from]
		moveItems := []string{}

		for i := 0; i < rule.move; i++ {
			v := pop(fromCrate)
			moveItems = prepend(moveItems, v)
		}

		crates[rule.to] = append(crates[rule.to], moveItems...)

	}

	result := ""

	for _, crate := range crates {
		v := pop(&crate)
		result += v
	}

	fmt.Println("PART 2 RESULT:", result)
}

func pop(alist *[]string) string {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func prepend(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y
	return x
}
