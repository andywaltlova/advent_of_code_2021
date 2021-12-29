package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Every solution is run along with utils.go
// e.g `go run 13.go utils.go`

func part1(thermalMap thermalImaging, folds []fold) int {
	return 1
}

func part2(thermalMap thermalImaging, folds []fold) int {
	return 2
}

type thermalImaging struct {
	coordinates map[coordinate]bool
}

type coordinate struct {
	x int
	y int
}

type fold struct {
	axis  string
	value int
}

func loadThermalImagingAndFolds(lines []string) (thermalImaging, []fold) {
	loadingFolds := false
	folds := []fold{}
	coordinates := make(map[coordinate]bool)
	for _, line := range lines {
		if line == "" {
			loadingFolds = true
			continue
		}
		if loadingFolds {
			params := strings.Split(line, "=")
			axis := string(params[0][len(params[0])-1])
			num, _ := strconv.Atoi(params[1])
			folds = append(folds, fold{axis: axis, value: num})
		} else {
			params := strings.Split(line, ",")
			x, _ := strconv.Atoi(params[0])
			y, _ := strconv.Atoi(params[1])
			coordinates[coordinate{x: x, y: y}] = true
		}
	}
	imaging := thermalImaging{coordinates: coordinates}
	return imaging, folds
}

func main() {
	lines := getInputLines("data/13.txt")
	imaging, folds := loadThermalImagingAndFolds(lines)

	fmt.Println(part1(imaging, folds))
	fmt.Println(part2(imaging, folds))
}
