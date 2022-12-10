package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2022/helpers"
)

func main() {
	content := helpers.GetInput("./day1/input.txt")

	calorieGroups := strings.Split(content, "\n\n")

	eachGroupTotal := []int{}

	for _, group := range calorieGroups {
		group := strings.Split(group, "\n")
		sum := 0
		for _, calorie := range group {
			i, _ := strconv.Atoi(calorie)
			sum += i
		}

		eachGroupTotal = append(eachGroupTotal, sum)
	}

	sort.Slice(eachGroupTotal, func(i, j int) bool {
		return eachGroupTotal[i] > eachGroupTotal[j]
	})

	part1(eachGroupTotal)
	part2(eachGroupTotal, 3)
}

func part1(groupTotals []int) {
	fmt.Println("PART 1 ANSWER:", groupTotals[0])
}

func part2(groupTotals []int, top int) {
	result := 0

	for i := 0; i < top; i++ {
		result += groupTotals[i]
	}

	fmt.Println("PART 2 ANSWER:", result)
}
