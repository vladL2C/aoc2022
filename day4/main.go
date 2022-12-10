package main

import (
	"aoc2022/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	rawData := helpers.GetInput("./day4/input.txt")

	rawSlice := strings.Split(rawData, "\n")

	Part1(rawSlice)
	Part2(rawSlice)
}

func Part1(rawData []string) {
	result := 0

	for _, line := range rawData {
		dataPoints := strings.Split(line, ",")

		first, second := GetRange(dataPoints[0])
		third, fourth := GetRange(dataPoints[1])

		if first <= third && second >= fourth || first >= third && second <= fourth {
			result++
		}

	}

	fmt.Println(result)
}

func Part2(rawData []string) {
	result := 0

	for _, line := range rawData {
		dataPoints := strings.Split(line, ",")

		first, second := GetRange(dataPoints[0])
		third, fourth := GetRange(dataPoints[1])

		if !(second < third || first > fourth) {
			result += 1
		}

	}

	fmt.Println(result)
}

func GetRange(point string) (int, int) {
	data := strings.Split(point, "-")

	a, _ := strconv.Atoi(data[0])
	b, _ := strconv.Atoi(data[1])

	return a, b
}
