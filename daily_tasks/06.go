package main

// Every solution is run along with utils.go
// e.g `go run 06.go utils.go`

import (
	"fmt"
)

func naiveSimulate(lanternfishes []int, days int) int {
	for i := 0; i < days; i++ {
		var new_fishes []int
		for _, fish := range lanternfishes {
			if fish == 0 {
				new_fishes = append(new_fishes, 6, 8)
			} else {
				new_fishes = append(new_fishes, fish-1)
			}
		}
		lanternfishes = new_fishes
	}
	return len(lanternfishes)
}

func betterSimulate(lanternfishes []int, days int) int {
	fishes := make(map[int]int)
	for _, num := range lanternfishes {
		fishes[num] = fishes[num] + 1
	}

	for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		_, ok := fishes[num]
		if !ok {
			fishes[num] = 0
		}
	}

	for i := 0; i < days; i++ {
		new_fishes := make(map[int]int)
		for i := 8; i > 0; i-- {
			new_fishes[i-1] = fishes[i]
		}
		new_fishes[6] = new_fishes[6] + fishes[0]
		new_fishes[8] = new_fishes[8] + fishes[0]
		fishes = new_fishes
	}

	var count int
	for _, elem := range fishes {
		count += elem
	}
	return count
}

func part1(lanternfishes []int) int {
	return naiveSimulate(lanternfishes, 80)
}

func part2(lanternfishes []int) int {
	return betterSimulate(lanternfishes, 256)
}

func main() {
	lanternfishes := getNumbersOnLine("data/06.txt")

	fmt.Println(part1(lanternfishes))
	fmt.Println(part2(lanternfishes))
}
