package main

// Every solution is run along with utils.go
// e.g `go run 12.go utils.go`

import (
	"fmt"
	"strings"
)

type cave struct {
	name  string
	edges []string
}

func (c cave) isSmall() bool {
	return c.name == strings.ToLower(c.name)
}

func findPaths(currentCave string, caves map[string]*cave, path []string, smallVisited bool) [][]string {
	if currentCave == "end" {
		path = append(path, currentCave)
		return [][]string{path}
	}
	if caves[currentCave].isSmall() {
		for _, c := range path {
			if currentCave == c {
				if currentCave == "start" || smallVisited {
					return [][]string{}
				}
				smallVisited = true
			}
		}
	}
	paths := [][]string{}
	path = append(path, currentCave)
	for _, c := range caves[currentCave].edges {
		// make a copy of slice
		newPath := append([]string{}, path...)
		paths = append(paths, findPaths(c, caves, newPath, smallVisited)...)
	}
	return paths
}

func loadCaves(lines []string) map[string]*cave {
	caves := map[string]*cave{}
	for _, line := range lines {
		edges := strings.Split(line, "-")

		c, caveExists := caves[edges[0]]
		if !caveExists {
			c = &cave{edges[0], []string{}}
			caves[c.name] = c
		}
		c.edges = append(c.edges, edges[1])

		c, caveExists = caves[edges[1]]
		if !caveExists {
			c = &cave{edges[1], []string{}}
			caves[c.name] = c
		}
		c.edges = append(c.edges, edges[0])
	}
	return caves
}

func part1(caves map[string]*cave) int {
	return len(findPaths("start", caves, []string{}, true))
}

func part2(caves map[string]*cave) int {
	// IN ONE PATH
	// single small cave can be visited at most twice
	// and the remaining small caves can be visited at most once
	return len(findPaths("start", caves, []string{}, false))
}

func main() {
	lines := getInputLines("data/12.txt")
	caves := loadCaves(lines)

	fmt.Println(part1(caves))
	fmt.Println(part2(caves))
}
