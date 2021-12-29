package main

// Every solution is run along with utils.go
// e.g `go run 11.go utils.go`

import (
	"fmt"
	"strings"
)

func part1(caves map[string]*cave) int {
	return 1
}

func part2(caves map[string]*cave) int {
	return 2
}

type cave struct {
	name  string
	edges []string
}

func (c cave) isSmall() bool {
	return c.name == strings.ToLower(c.name)
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

func main() {
	lines := getInputLines("data/12.txt")
	caves := loadCaves(lines)
	fmt.Println(caves)

	fmt.Println(part1(caves))
	fmt.Println(part2(caves))
}
