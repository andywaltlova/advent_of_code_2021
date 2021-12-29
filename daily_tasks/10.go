package main

// Every solution is run along with utils.go
// e.g `go run 10.go utils.go`

import (
	"fmt"
)

var CLOSING_BRACKETS_MAP map[string]string = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
	">": "<",
}

var SCORE map[string]int = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func validate_chunks(chunks string) int {
	stack := []string{}
	var score int
	for _, runeChar := range chunks {
		closing := string(runeChar)
		if opening, ok := CLOSING_BRACKETS_MAP[closing]; ok {
			stackLastIndex := len(stack) - 1
			if stack[stackLastIndex] != opening {
				score = SCORE[closing]
				break
			}
			stack = stack[:stackLastIndex]
		} else {
			stack = append(stack, closing)
		}
	}
	return score
}

func part1(lines []string) int {
	var score int = 0
	for _, chunks := range lines {
		score += validate_chunks(chunks)
	}
	return score
}

func part2(lines []string) int {
	return 2
}

func main() {
	lines := getInputLines("data/10.txt")

	fmt.Println(part1(lines))
	// fmt.Println(part2(lines))
}
