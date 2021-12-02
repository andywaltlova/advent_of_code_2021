package main

// Every solution is run along with utils.go
// e.g `go run 02.go utils.go`

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(input_strings []string) int {
	depth := 0
	horizontal := 0
	for _, s := range input_strings {
		command_units := strings.Fields(s)
		command := command_units[0]
		units, _ := strconv.Atoi(command_units[1])

		if command == "forward" {
			horizontal += units
		}
		if command == "down" {
			depth += units
		}
		if command == "up" {
			depth -= units
		}
	}
	return depth * horizontal
}

func part2(input_strings []string) int {
	aim := 0
	depth := 0
	horizontal := 0
	for _, s := range input_strings {
		command_units := strings.Fields(s)
		command := command_units[0]
		units, _ := strconv.Atoi(command_units[1])

		if command == "forward" {
			horizontal += units
			depth += aim * units
		}
		if command == "down" {
			aim += units
		}
		if command == "up" {
			aim -= units
		}
	}
	return depth * horizontal
}

func main() {
	lines := getInputLines("data/02.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
