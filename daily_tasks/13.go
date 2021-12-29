package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Every solution is run along with utils.go
// e.g `go run 13.go utils.go`

func part1(thermalMap thermalImaging, folds []fold) int {
	newThermalMap := thermalImaging{}
	if folds[0].isHorizontal() {
		newThermalMap = thermalMap.horizontalFold(folds[0].value)
	} else {
		newThermalMap = thermalMap.verticalFold(folds[0].value)
	}
	return len(newThermalMap.coordinates)
}

func part2(thermalMap thermalImaging, folds []fold) {
	for _, f := range folds {
		if f.isHorizontal() {
			thermalMap = thermalMap.horizontalFold(f.value)
		} else {
			thermalMap = thermalMap.verticalFold(f.value)
		}
	}

	for y := 0; y < 6; y++ {
		for x := 0; x < 39; x++ {
			c := coordinate{x: x, y: y}
			if _, ok := thermalMap.coordinates[c]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

type thermalImaging struct {
	coordinates map[coordinate]bool
}

func (t thermalImaging) horizontalFold(line int) thermalImaging {
	// y coordinate is folded up
	// e.g folding by line 7 will fold 8 line on 6 line, 9 line on 5 line ...
	newCoordinates := make(map[coordinate]bool)
	for coord := range t.coordinates {
		if coord.y <= line {
			newCoordinates[coordinate{x: coord.x, y: coord.y}] = true
			continue
		}
		newY := 2*line - coord.y
		newCoordinates[coordinate{x: coord.x, y: newY}] = true
	}
	newImaging := thermalImaging{coordinates: newCoordinates}
	return newImaging
}

func (t thermalImaging) verticalFold(col int) thermalImaging {
	// x coordinate is folded to the left
	// e.g folding by col 5 will fold 6 col on 4 col, 9 col on 1 col ...
	newCoordinates := make(map[coordinate]bool)
	for coord := range t.coordinates {
		if coord.x <= col {
			newCoordinates[coordinate{x: coord.x, y: coord.y}] = true
			continue
		}
		newX := 2*col - coord.x
		newCoordinates[coordinate{x: newX, y: coord.y}] = true
	}
	newImaging := thermalImaging{coordinates: newCoordinates}
	return newImaging
}

type coordinate struct {
	x int
	y int
}

type fold struct {
	axis  string
	value int
}

func (f fold) isHorizontal() bool {
	return f.axis == "y"
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
	part2(imaging, folds)
}
