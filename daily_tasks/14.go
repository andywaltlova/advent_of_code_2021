package main

import (
	"fmt"
	"math"
	"strings"
)

// Every solution is run along with utils.go
// e.g `go run 14.go utils.go`

var POLYMER string
var INSERTIONS map[string]string = make(map[string]string)

// brute force ... not working for part 2
func naiveReinforceSubmarine(steps int) map[string]int {
	var charCounts map[string]int = make(map[string]int)

	var newPolymer string = POLYMER[:]
	for step := 0; step < steps; step++ {
		var tmpPolymer string = ""
		for i := 0; i < len(newPolymer)-1; i++ {
			pair := string(newPolymer[i]) + string(newPolymer[i+1])
			insertion := INSERTIONS[pair]
			charCounts[insertion]++

			tmpPolymer += string(pair[0]) + insertion
		}
		newPolymer = tmpPolymer[:] + newPolymer[len(newPolymer)-1:]
	}
	return charCounts
}

func recursiveReinforceSubmarine(polymer string, steps int, memory map[string]map[string]int) map[string]int {
	charCounts := map[string]int{}

	if steps == 0 {
		return charCounts
	}

	key := fmt.Sprint(polymer, steps)
	if res, ok := memory[key]; ok {
		return res
	}

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		insertion := INSERTIONS[pair]
		charCounts[insertion]++

		newPolymer := pair[:1] + insertion + pair[1:]
		recurseCounts := recursiveReinforceSubmarine(newPolymer, steps-1, memory)

		for k, v := range recurseCounts {
			charCounts[k] += v
		}
	}

	memory[key] = charCounts
	return charCounts
}

func part1(initialCounts map[string]int) int {
	charCounts := naiveReinforceSubmarine(10)

	//add initial count
	for k, v := range initialCounts {
		charCounts[k] += v
	}

	most, least := getMostAndLeast(charCounts)
	return most - least
}

func part2(initialCounts map[string]int) int {
	empty_memory := map[string]map[string]int{}
	charCounts := recursiveReinforceSubmarine(POLYMER[:], 40, empty_memory)

	//add initial count
	for k, v := range initialCounts {
		charCounts[k] += v
	}

	most, least := getMostAndLeast(charCounts)
	return most - least
}

func getMostAndLeast(charCounts map[string]int) (most int, least int) {
	most, least = 0, math.MaxInt64
	for k := range charCounts {
		val := charCounts[k]
		if val > most {
			most = val
		}
		if val < least {
			least = val
		}
	}
	return most, least
}

func loadData(lines []string) (initialCount map[string]int) {
	var charCounts map[string]int = make(map[string]int)

	POLYMER = lines[0]
	for _, c := range POLYMER {
		charCounts[string(c)]++
	}
	for _, v := range lines[2:] {
		data := strings.Split(v, " -> ")
		INSERTIONS[data[0]] = data[1]
	}
	return charCounts
}

func main() {
	lines := getInputLines("data/14.txt")
	initialCount := loadData(lines)

	fmt.Println(part1(initialCount))
	fmt.Println(part2(initialCount))
}
