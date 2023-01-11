package main

import (
	"aoc2022/helpers"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func main() {
	Part1()
	Part2()
}

func Part1() {
	raw := helpers.GetInput("./day9/input.txt")

	data := strings.Split(raw, "\n")

	head := coord{x: 0, y: 0}
	tail := coord{x: 0, y: 0}

	visibility := make(map[coord]struct{})

	visibility[tail] = struct{}{}

	for _, line := range data {
		instruction := strings.Split(line, " ")
		direction, moves := GetDirection(instruction)

		for i := 0; i < moves; i++ {
			head = MoveHead(head, direction)

			if !Touching(head, tail) {
				tail = MoveTailToHead(head, tail)
				visibility[tail] = struct{}{}
			}

		}

	}

	fmt.Println("PART 1:", len(visibility))
}

func Part2() {
	raw := helpers.GetInput("./day9/input.txt")

	data := strings.Split(raw, "\n")

	head := coord{x: 0, y: 0}
	tails := make([]coord, 9)

	visibility := make(map[coord]struct{})

	visibility[tails[0]] = struct{}{}

	for _, line := range data {
		instruction := strings.Split(line, " ")
		direction, moves := GetDirection(instruction)

		for i := 0; i < moves; i++ {
			head = MoveHead(head, direction)
			previousTail := head
			for j := 0; j < len(tails); j++ {
				tails[j] = MoveTailToHead(previousTail, tails[j])
				previousTail = tails[j]
				if j == len(tails)-1 {
					visibility[tails[j]] = struct{}{}
				}

			}

		}

	}
	fmt.Println("PART 2:", len(visibility))
}
func Touching(head, tail coord) bool {
	diffX := math.Abs(float64(tail.x - head.x))
	diffY := math.Abs(float64(tail.y - head.y))

	return diffX <= 1 && diffY <= 1
}

func MoveTailToHead(head, tail coord) coord {
	if !Touching(head, tail) {
		x, y := head.x-tail.x, head.y-tail.y
		if x > 0 {
			x = 1
		} else if x < 0 {
			x = -1
		}
		if y > 0 {
			y = 1
		} else if y < 0 {
			y = -1
		}
		return coord{tail.x + x, tail.y + y}
	}

	return tail
}

func MoveHead(head coord, direction string) coord {
	switch direction {
	case "R":
		head.x += 1
	case "L":
		head.x -= 1
	case "U":
		head.y += 1
	case "D":
		head.y -= 1
	default:
		panic("invalid direction")
	}

	return head
}

func GetDirection(instruction []string) (string, int) {
	direction := instruction[0]
	moves, _ := strconv.Atoi(instruction[1])
	return direction, moves
}
