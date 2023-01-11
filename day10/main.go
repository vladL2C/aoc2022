package main

import (
	"aoc2022/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		raw                = helpers.GetInput("./day10/input.txt")
		lines              = strings.Split(raw, "\n")
		x                  = 1
		cycle              = 0
		sum                = 0
		cycles map[int]int = map[int]int{
			20:  20,
			60:  60,
			100: 100,
			140: 140,
			180: 180,
			220: 220,
		}
		seen map[int]bool = map[int]bool{
			20:  false,
			60:  false,
			100: false,
			140: false,
			180: false,
			220: false,
		}
	)

	for _, line := range lines {
		instruction, val := GetInstruction(line)
		cycle += 1

		if CheckCycle(cycles, cycle) && !(seen[cycle]) {
			sum += cycle * x
			seen[cycle] = true
		}

		if instruction == "addx" {
			cycle += 1
		}

		if CheckCycle(cycles, cycle) && !(seen[cycle]) {
			sum += cycle * x
			seen[cycle] = true
		}

		x += val
	}

	fmt.Println(sum)

	Part2()
}

func CheckCycle(cycles map[int]int, cycle int) bool {
	_, ok := cycles[cycle]
	return ok
}

func GetInstruction(line string) (string, int) {
	parsed := strings.Split(line, " ")

	instruction := parsed[0]
	val := 0

	if len(parsed) > 1 {
		val, _ = strconv.Atoi(parsed[1])
	}

	return instruction, val
}

type cpu struct {
	x           int
	cycleNumber int

	currentRow   string
	renderedRows []string
}

func Part2() {
	raw := helpers.GetInput("./day10/input.txt")

	cpu := cpu{
		x:            1,
		cycleNumber:  -1,
		currentRow:   "",
		renderedRows: make([]string, 0),
	}

	for _, line := range strings.Split(raw, "\n") {
		instruction, val := GetInstruction(line)

		if instruction == "noop" {
			cpu.noop()
		} else {
			cpu.addx(val)
		}

	}

	for _, row := range cpu.renderedRows {
		fmt.Println(row)
	}
}

func (c *cpu) noop() {
	c.incrementCycleNumber()
}

func (c *cpu) addx(x int) {
	c.incrementCycleNumber()
	c.incrementCycleNumber()

	c.x += x
}

func (c *cpu) incrementCycleNumber() {
	c.cycleNumber++

	currentPixel := c.currentPixel()
	if c.x-1 <= currentPixel && currentPixel <= c.x+1 {
		c.currentRow += "#"
	} else {
		c.currentRow += "."
	}

	if c.currentPixel() == 39 {
		c.renderedRows = append(c.renderedRows, c.currentRow)
		c.currentRow = ""
	}
}

func (c *cpu) currentPixel() int {
	return c.cycleNumber % 40
}
