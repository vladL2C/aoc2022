package main

import (
	"aoc2022/helpers"
	"fmt"
)

type Stream struct {
	Data string
}

func NewStreamer(data string) Stream {
	return Stream{
		Data: data,
	}
}

func main() {
	stream := NewStreamer(helpers.GetInput("./day6/input.txt"))
	stream.FindUniqueBy(3)
	stream.FindUniqueBy(13)
}

func (s Stream) FindUniqueBy(n int) {
	for i := n; i < len(s.Data); i++ {
		section := s.Data[i-n : i+1]

		if isUnique(section) {
			fmt.Println(i + 1)
			break
		}
	}
}

func isUnique(substr string) bool {
	store := make(map[rune]bool)

	for _, c := range substr {
		_, ok := store[c]
		if ok {
			return false
		}
		store[c] = true
	}

	return true
}
