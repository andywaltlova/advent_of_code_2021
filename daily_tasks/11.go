package main

// Every solution is run along with utils.go
// e.g `go run 10.go utils.go`

import (
	"fmt"
)

func part1(lines []int) int {
	return 1
}

func part2(lines []int) int {
	return 2
}

func main() {
	lines := getNumberInput("data/11.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
