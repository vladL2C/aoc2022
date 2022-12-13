package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2022/helpers"
)

type Calories struct {
	Groups []string
}

func (c *Calories) Group() []int {
	totals := []int{}

	for _, group := range c.Groups {
		calories := strings.Split(group, "\n")
		sum := 0
		for _, calorie := range calories {
			cal, _ := strconv.Atoi(calorie)
			sum += cal
		}
		totals = append(totals, sum)
	}

	return totals
}

func Highest(groupTotals []int) {
	fmt.Println("PART 1 ANSWER:", groupTotals[0])
}

func Top(groupTotals []int, top int) {
	result := 0

	for i := 0; i < top; i++ {
		result += groupTotals[i]
	}

	fmt.Println("PART 2 ANSWER:", result)
}

func main() {
	content := helpers.GetInput("./day1/input.txt")

	calorieGroups := strings.Split(content, "\n\n")

	c := Calories{
		Groups: calorieGroups,
	}

	totals := c.Group()

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	Highest(totals)
	Top(totals, 3)
}
