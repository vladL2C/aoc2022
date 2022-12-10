package main

import (
	"aoc2022/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var totalDiskSpace = 70000000
var unusedAtLeast = 30000000

func makeAboslutePath(stack []string, currentPath string) string {
	prefix := stack[len(stack)-1]
	path := ""

	if prefix == "/" {
		path = prefix + currentPath
	} else {
		path = prefix + "/" + currentPath
	}

	return path
}

func Parse(text string) map[string]int {
	store := map[string]int{
		"/": 0,
	}

	lines := strings.Split(text, "\n")

	stack := []string{"/"}

	for _, line := range lines {
		commands := strings.Split(line, " ")

		prefix := commands[0]
		command := commands[1]
		path := commands[len(commands)-1]

		if prefix == "$" && command == "cd" && path == ".." {
			stack = stack[:len(stack)-1]
		}

		if prefix == "$" && command == "cd" && path != ".." && path != "/" {
			absPath := makeAboslutePath(stack, path)
			stack = append(stack, absPath)
		}

		fileSize, _ := strconv.Atoi(commands[0])
		currentPath := stack[len(stack)-1]
		store[currentPath] += fileSize
	}

	return store
}

func main() {
	text := helpers.GetInput("./day7/input.txt")

	store := Parse(text)
	directories := SumTotalDirSize(store)
	//part 1
	fmt.Println("PART 1: ", Sum(Below100000(directories)))
	// part 2
	unusedSpace := totalDiskSpace - directories["/"]
	required := unusedAtLeast - unusedSpace
	fmt.Println("PART 2", MinRequiredForUpdate(directories, required))

}

func MinRequiredForUpdate(directories map[string]int, required int) int {
	min := []int{}

	for _, t := range directories {
		if t >= required {
			min = append(min, t)
		}
	}

	sort.Slice(min, func(i, j int) bool {
		return min[i] < min[j]
	})

	return min[0]
}

func SumTotalDirSize(directories map[string]int) (totals map[string]int) {
	totals = map[string]int{}

	for path := range directories {
		sum := 0
		for p, z := range directories {

			if strings.HasPrefix(p, path) {
				sum += z
			}
		}
		totals[path] = sum
	}

	return totals
}

func Below100000(directories map[string]int) (result []int) {
	result = []int{}

	for _, v := range directories {
		if v <= 100000 {
			result = append(result, v)
		}
	}

	return result
}

func Sum(d []int) int {
	total := 0
	for _, n := range d {
		total += n
	}
	return total
}
