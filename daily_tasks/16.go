package main

import (
	"fmt"
	"strings"
)

// Every solution is run along with utils.go
// e.g `go run 16.go utils.go`

func part1(hex []string) int {
	return 1
}

func part2(hex []string) int {
	return 2
}

func main() {
	lines := getInputLines("data/16.txt")
	hexRepr := strings.Split(lines[0], "")

	fmt.Println(hexRepr)
	fmt.Println(part1(hexRepr))
	// fmt.Println(part2(hexRepr))
}
