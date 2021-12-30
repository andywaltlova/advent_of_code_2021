package main

// Every solution is run along with utils.go
// e.g `go run 17.go utils.go`

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(yMin int) int {
	// Took me quite a long time to discover it's simple equation
	// n(n+1)/2 is the sum of numbers from 1 to n
	// Reasoning behind this is nicely explained in one of the other solutions I found on Reddit
	// -> https://github.com/prendradjaja/advent-of-code-2021/blob/main/17--trick-shot/a.py

	return (yMin + 1) * yMin / 2
}

// Brute force...
func part2(xMin int, xMax int, yMin int, yMax int) int {
	noVelocities, n := 0, int(math.Pow(float64(xMin*2), 0.5)-1) // n-th member of arithmetic progression

	for yInit := yMin; yInit < -yMin; yInit++ {
		for xInit := n; xInit < xMax+1; xInit++ {

			// x,y are coordinates, _x, _y are velocities
			x, y, _x, _y := 0, 0, xInit, yInit
			for x <= xMax && y >= yMin && (_x == 0 && xMin <= x || _x != 0) {
				x += _x // The probe's x position increases by its x velocity.
				y += _y // The probe's y position increases by its y velocity.

				// Due to drag, the probe's x velocity changes by 1 toward the value 0;
				// that is, it decreases by 1 if it is greater than 0, increases by 1
				// if it is less than 0, or does not change if it is already 0.
				if _x > 0 { // shouldn't be negative...
					_x -= 1
				}
				_y -= 1 // Due to gravity, the probe's y velocity decreases by 1.

				if xMin <= x && x <= xMax && yMin <= y && y <= yMax {
					noVelocities += 1
					break
				}
			}
		}
	}
	return noVelocities
}

func getTarget(input string) (xMin int, xMax int, yMin int, yMax int) {
	data := strings.TrimLeft(input, "target area: ")
	areaData := strings.Split(data, ", ")
	xData := strings.TrimLeft(areaData[0], "x=")
	yData := strings.TrimLeft(areaData[1], "y=")
	xRange := strings.Split(xData, "..")
	yRange := strings.Split(yData, "..")

	xMin, _ = strconv.Atoi(xRange[0])
	xMax, _ = strconv.Atoi(xRange[1])
	yMin, _ = strconv.Atoi(yRange[0])
	yMax, _ = strconv.Atoi(yRange[1])
	return xMin, xMax, yMin, yMax
}

func main() {
	line := getInputLines("data/17.txt")[0]
	xMin, xMax, yMin, yMax := getTarget(line)

	fmt.Println(part1(yMin))
	fmt.Println(part2(xMin, xMax, yMin, yMax))
}
