package main

// Every solution is run along with utils.go
// e.g `go run 10.go utils.go`

import (
	"fmt"
)

func part1(lines []string) int {
	return 1
}

func part2(lines []string) int {
	return 2
}

func main() {
	lines := getInputLines("data/10.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
