package main

// Every solution is run along with utils.go
// e.g `go run 11.go utils.go`

import (
	"fmt"
	"strconv"
	"strings"
)

type octopusGrid map[string]struct {
	octopusValue int
	neighbours   []string
}
type queueItem struct {
	coords      string
	instruction string
}

func part1(grid octopusGrid) int {
	flashes_count := 0
	for i := 1; i < 101; i++ {
		flashes_count += doOneStep(grid)
	}
	return flashes_count
}

func part2(grid octopusGrid) int {
	var synchronized_step int
	for step := 1; synchronized_step == 0; step++ {
		if doOneStep(grid) == len(grid) {
			synchronized_step = step
		}
	}
	return synchronized_step
}

func doOneStep(grid octopusGrid) int {
	// returns number of flashed octopuses

	queue := make([]queueItem, len(grid))
	flashedCoords := make(map[string]bool)

	i := 0
	for key := range grid {
		queue[i] = queueItem{coords: key, instruction: "+"}
		i++
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		// already flashed in this step
		if _, ok := flashedCoords[item.coords]; ok {
			continue
		}

		entry := grid[item.coords]

		switch item.instruction {
		case "flash":
			entry.octopusValue = 0
			for _, neighbour := range entry.neighbours {
				queue = append(queue, queueItem{coords: neighbour, instruction: "+"})
			}
			flashedCoords[item.coords] = true
		case "+":
			entry.octopusValue++
			if entry.octopusValue == 10 {
				queue = append(queue, queueItem{coords: item.coords, instruction: "flash"})
			}
		}

		grid[item.coords] = entry
	}
	return len(flashedCoords)
}

func getCoords(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func createOctopusGrid(lines []string) octopusGrid {
	grid := octopusGrid{}

	for row, line := range lines {
		octopuses := strings.Split(line, "")
		for col, octopus_val := range octopuses {
			coords := getCoords(col, row)
			entry := grid[coords]
			entry.octopusValue, _ = strconv.Atoi(octopus_val)

			notFirstRow, notLastRow := row > 0, row < (len(lines)-1)
			notFirstCol, notLastCol := col > 0, col < (len(octopuses)-1)

			if notFirstRow {
				entry.neighbours = append(entry.neighbours, getCoords(col, row-1))
			}
			if notLastRow {
				entry.neighbours = append(entry.neighbours, getCoords(col, row+1))
			}

			if notFirstCol {
				entry.neighbours = append(entry.neighbours, getCoords(col-1, row))
				if notFirstRow {
					entry.neighbours = append(entry.neighbours, getCoords(col-1, row-1))
				}
				if notLastRow {
					entry.neighbours = append(entry.neighbours, getCoords(col-1, row+1))
				}
			}
			if notLastCol {
				entry.neighbours = append(entry.neighbours, getCoords(col+1, row))
				if notFirstRow {
					entry.neighbours = append(entry.neighbours, getCoords(col+1, row-1))
				}
				if notLastRow {
					entry.neighbours = append(entry.neighbours, getCoords(col+1, row+1))
				}
			}

			grid[coords] = entry
		}
	}

	return grid
}

func main() {
	lines := getInputLines("data/11.txt")
	grid := createOctopusGrid(lines)

	fmt.Println(part1(grid))
	fmt.Println(part2(grid) + 100) // because 100 steps was already made in part1
}
