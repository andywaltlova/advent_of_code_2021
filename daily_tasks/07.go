package main

// Every solution is run along with utils.go
// e.g `go run 07.go utils.go`

import (
	"fmt"
)

func makePositionCrabMap(numbers []int) (map[int]int, []int) {
	// create map position:numberOfCrabs
	positionsMap := make(map[int]int)
	for _, num := range numbers {
		positionsMap[num] += 1
	}

	positions := makeRange(arrayMin(numbers), arrayMax(numbers))
	return positionsMap, positions
}

func getRequiredFuel(positions map[int]int, target int, part2 bool) int {
	var fuel int
	for position, numOfCrabs := range positions {
		distance := abs(position - target)
		if part2 {
			// Gauss sum of consecutive numbers
			// n * (n + 1) / 2
			fuel += ((distance * (distance + 1)) / 2) * numOfCrabs
		} else {
			fuel += distance * numOfCrabs
		}
	}
	return fuel
}

func getMinimalRequiredFuel(numbers []int, part2 bool) int {
	positions, keys := makePositionCrabMap(numbers)

	var minFuel int = getRequiredFuel(positions, keys[len(keys)-1], part2)
	for _, v := range keys {
		requiredFuel := getRequiredFuel(positions, v, part2)
		if requiredFuel < minFuel {
			minFuel = requiredFuel
		}
	}
	return minFuel
}

func part1(numbers []int) int {
	return getMinimalRequiredFuel(numbers, false)
}

func part2(numbers []int) int {
	return getMinimalRequiredFuel(numbers, true)
}

func main() {
	numbers := getNumbersOnLine("data/07.txt")

	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}
