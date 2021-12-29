package main

import (
	"fmt"
	"strings"
)

// Every solution is run along with utils.go
// e.g `go run 14.go utils.go`

func part1() int {
	return 1
}

func part2(template Polymer) int {
	return 2
}

type Polymer struct {
	template string
}

var POLYMER string
var INSERTIONS map[string]string = make(map[string]string)

func loadData(lines []string) {
	POLYMER = lines[0]
	for _, v := range lines[2:] {
		data := strings.Split(v, " -> ")
		INSERTIONS[data[0]] = data[1]
	}
}

func main() {
	lines := getInputLines("data/14.txt")
	loadData(lines)
	fmt.Println(POLYMER, INSERTIONS)

	fmt.Println(part1())
	// fmt.Println(part2(polymer))
}
