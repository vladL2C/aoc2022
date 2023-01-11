package main

import (
	"aoc2022/helpers"
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func main() {
	raw := helpers.GetInput("./day9/sample.txt")

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
			tail = MoveTail(tail, head)
		}

	}

	fmt.Printf("%#v", head)
	fmt.Println(len(visibility))
}

func MoveTail(tail, head coord) coord {
	diffX := tail.x - head.x
	diffY := tail.y - head.y

	fmt.Println(diffX, diffY)
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
