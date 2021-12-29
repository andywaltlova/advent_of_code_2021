package main

// Every solution is run along with utils.go
// e.g `go run 10.go utils.go`

import (
	"fmt"
	"sort"
)

var CLOSING_BRACKETS_MAP map[string]string = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
	">": "<",
}

var CLOSING_SCORE map[string]int = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var OPENING_SCORE map[string]int = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func validate_chunks(chunks string) (int, []string) {
	stack := []string{}
	var score int
	for _, runeChar := range chunks {
		closing := string(runeChar)
		if opening, ok := CLOSING_BRACKETS_MAP[closing]; ok {
			stackLastIndex := len(stack) - 1
			if stack[stackLastIndex] != opening {
				score = CLOSING_SCORE[closing]
				break
			}
			stack = stack[:stackLastIndex]
		} else {
			stack = append(stack, closing)
		}
	}
	return score, stack
}

func part1(lines []string) int {
	var total_score int = 0
	for _, chunks := range lines {
		score, _ := validate_chunks(chunks)
		total_score += score
	}
	return total_score
}

func part2(lines []string) int {
	scores := []int{}
	for _, chunks := range lines {
		score, stack := validate_chunks(chunks)
		if score != 0 {
			continue
		}

		for i := len(stack) - 1; i >= 0; i-- {
			symbolScore := OPENING_SCORE[stack[i]]
			score *= 5
			score += symbolScore
		}
		scores = append(scores, score)
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	return scores[len(scores)/2]
}

func main() {
	lines := getInputLines("data/10.txt")

	fmt.Println(part1(lines)) //316851
	fmt.Println(part2(lines)) //2182912364
}
