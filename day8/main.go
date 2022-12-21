package main

import (
	"aoc2022/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	raw := helpers.GetInput("./day8/sample.txt")

	parse := strings.Split(raw, "\n")

	grid := []string{}
	grid = append(grid, parse...)
	edges := []string{grid[0], grid[len(grid)-1]}
	edges = append(edges, GetRemainingEdges(grid[1:len(grid)-1])...)

	totalEdges := GetEdgeCount(edges)

	visible := 0

	for row := 1; row < len(grid)-1; row++ {

		for col := 1; col < len(grid)-1; col++ {
			tree, _ := strconv.Atoi(string(grid[row][col]))

			if CheckRight(tree, col+1, grid[row]) {
				visible += 1
				continue
			}

			if CheckLeft(tree, col-1, grid[row]) {
				visible += 1
				continue
			}

			if CheckUp(tree, row, col, grid) {
				visible += 1
				continue
			}

			if CheckDown(tree, row, col, grid) {
				visible += 1
				continue
			}

		}

	}

	fmt.Println(totalEdges + visible)

}

func CheckDown(currentTree, row, col int, grid []string) bool {
	isVisible := false

	for i := row + 1; i < len(grid); i++ {
		nextTree, _ := strconv.Atoi(string(grid[i][col]))
		if currentTree > nextTree {
			isVisible = true
		} else {
			return false
		}
	}

	return isVisible

}

func CheckUp(currentTree, row, col int, grid []string) bool {
	isVisible := false

	for i := row - 1; i >= 0; i-- {
		nextTree, _ := strconv.Atoi(string(grid[i][col]))
		if currentTree > nextTree {
			isVisible = true
		} else {
			return false
		}
	}

	return isVisible
}

func CheckLeft(currentTree, row int, grid string) bool {
	isVisible := false

	for i := row; i >= 0; i-- {
		nextTree, _ := strconv.Atoi(string(grid[i]))
		if currentTree > nextTree {
			isVisible = true
		} else {
			return false
		}
	}

	return isVisible
}

func CheckRight(currentTree, row int, grid string) bool {
	isVisible := false

	for i := row; i < len(grid); i++ {
		nextTree, _ := strconv.Atoi(string(grid[i]))
		if currentTree > nextTree {
			isVisible = true
		} else {
			return false
		}
	}

	return isVisible
}

func GetEdgeCount(edges []string) int {
	count := 0

	for _, e := range edges {
		count += len(e)
	}

	return count
}

func GetRemainingEdges(trees []string) []string {
	edges := []string{}

	for _, item := range trees {
		left := string(item[0])
		right := string(item[len(item)-1])

		edges = append(edges, left, right)
	}

	return edges
}
