package main

// Every solution is run along with utils.go
// e.g `go run 05.go utils.go`

import (
	"fmt"
	"strconv"
	"strings"
)

type HydroThermalVent struct {
	from Coordinates
	to   Coordinates
}

type Coordinates struct {
	x int
	y int
}

func (h HydroThermalVent) isVerticalOrHorizontal() bool {
	return h.isHorizontal() || h.isVertical()
}

func (h HydroThermalVent) isVertical() bool {
	return h.from.y == h.to.y
}

func (h HydroThermalVent) isHorizontal() bool {
	return h.from.x == h.to.x
}

func (h HydroThermalVent) getAllCoordinates() []Coordinates {
	var coordinates []Coordinates
	if h.isHorizontal() {
		// x constant
		start := min(h.from.y, h.to.y)
		end := max(h.from.y, h.to.y)

		for y := start; y <= end; y++ {
			coordinates = append(coordinates, Coordinates{x: h.from.x, y: y})
		}
		return coordinates
	}
	if h.isVertical() {
		// y constant
		start := min(h.from.x, h.to.x)
		end := max(h.from.x, h.to.x)
		for x := start; x <= end; x++ {
			coordinates = append(coordinates, Coordinates{x: x, y: h.from.y})
		}
		return coordinates
	}

	// Diagonal
	var x_inc int = 1
	var y_inc int = 1
	// Determine increment for diagonal coordinates based on direction
	if h.from.x > h.to.x {
		x_inc = -1
	}
	if h.from.y > h.to.y {
		y_inc = -1
	}

	for c := h.from; c != h.to; {
		coordinates = append(coordinates, c)
		c = Coordinates{x: c.x + x_inc, y: c.y + y_inc}
	}
	coordinates = append(coordinates, h.to)
	return coordinates
}

func parseVents(lines []string, includeDiagonal bool) []HydroThermalVent {
	var result []HydroThermalVent
	for _, line := range lines {
		start_to_end := strings.Split(line, " -> ")
		start := strings.Split(start_to_end[0], ",")
		end := strings.Split(start_to_end[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		var from Coordinates = Coordinates{x: x1, y: y1}
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		var to Coordinates = Coordinates{x: x2, y: y2}
		result = append(result, HydroThermalVent{from: from, to: to})
	}
	return result
}

func getOverlaps(input_strings []string, with_diagonal bool) int {
	vents := parseVents(input_strings, false)
	all_coordinates := make(map[Coordinates]int)

	for _, vent := range vents {
		if !with_diagonal && vent.isVerticalOrHorizontal() {
			for _, c := range vent.getAllCoordinates() {
				all_coordinates[c] = all_coordinates[c] + 1
			}
		}
		if with_diagonal {
			for _, c := range vent.getAllCoordinates() {
				all_coordinates[c] = all_coordinates[c] + 1
			}
		}
	}
	var count int = 0
	for _, element := range all_coordinates {
		if element > 1 {
			count += 1
		}
	}
	return count
}

func part1(input_strings []string) int {
	return getOverlaps(input_strings, false)
}

func part2(input_strings []string) int {
	return getOverlaps(input_strings, true)
}

func main() {
	lines := getInputLines("data/05.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
