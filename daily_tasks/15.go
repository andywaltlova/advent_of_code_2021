package main

import (
	"container/heap" // for implementing priority queue
	"fmt"
	"strconv"
)

// Every solution is run along with utils.go
// e.g `go run 15.go utils.go`

type Point struct {
	x int
	y int
}

type Node struct {
	Point     // node is basically point, to get easier access to x,y
	risk      int
	pathScore int // for priority queue implementation
	neighbors []*Node
}

type Path []*Node

func (pq Path) Len() int {
	return len(pq)
}

func (pq Path) Less(i, j int) bool {
	return pq[i].pathScore < pq[j].pathScore
}

func (pq Path) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Path) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *Path) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return node
}

func loadGrid(lines []string, part2 bool) [][]*Node {
	riskGrid := [][]*Node{}

	for y, row := range lines {
		riskRow := []*Node{}
		for x, r := range row {
			num, _ := strconv.Atoi(string([]rune{r}))
			riskRow = append(riskRow, &Node{risk: num, Point: Point{x, y}})
		}
		riskGrid = append(riskGrid, riskRow)
	}

	if part2 {
		riskGrid = expandGrid(riskGrid)
	}

	for y, row := range riskGrid {
		for x, n := range row {
			if y != 0 {
				n.neighbors = append(n.neighbors, riskGrid[y-1][x])
			}
			if x != 0 {
				n.neighbors = append(n.neighbors, riskGrid[y][x-1])
			}
			if x < len(row)-1 {
				n.neighbors = append(n.neighbors, riskGrid[y][x+1])
			}
			if y < len(riskGrid)-1 {
				n.neighbors = append(n.neighbors, riskGrid[y+1][x])
			}
		}
	}
	return riskGrid
}

func expandGrid(baseGrid [][]*Node) [][]*Node {
	baseY, baseX := len(baseGrid), len(baseGrid[0])
	fullGrid := [][]*Node{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for y, row := range baseGrid {
				if j == 0 {
					fullGrid = append(fullGrid, []*Node{})
				}
				for x, n := range row {
					incRisk := (n.risk + j + i) % 9
					if incRisk == 0 {
						incRisk = 9
					}
					pn := &Node{risk: incRisk, Point: Point{baseX*j + x, baseY*i + y}}
					fullGrid[baseY*i+y] = append(fullGrid[baseY*i+y], pn)
				}
			}
		}
	}
	return fullGrid
}

func findLowestRisk(start, end *Node) []Point {
	pq := Path{}
	heap.Init(&pq)
	heap.Push(&pq, start)
	cameFrom := map[Point]Point{}
	graphScore := map[Point]int{start.Point: 0}
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Node)

		// end search
		if cur.Point == end.Point {
			return constructPathFromEnd(end, cameFrom)
		}

		// based on Djikstra algorithm for shortest path using priority queue
		for _, n := range cur.neighbors {
			gs := graphScore[cur.Point] + n.risk
			if graphScore[n.Point] == 0 || gs < graphScore[n.Point] {
				cameFrom[n.Point] = cur.Point
				graphScore[n.Point] = gs
				newNode := &Node{
					Point:     n.Point,
					risk:      n.risk,
					pathScore: gs + (end.x - n.x) + (end.y - n.y),
					neighbors: n.neighbors,
				}
				heap.Push(&pq, newNode)
			}
		}
	}
	return nil
}

func constructPathFromEnd(end *Node, cameFrom map[Point]Point) []Point {
	path := []Point{end.Point}
	nextNode := cameFrom[end.Point]
	start := Point{0, 0}
	for nextNode != start {
		path = append(path, nextNode)
		nextNode = cameFrom[nextNode]
	}
	return path
}

func getPathRisk(path []Point, riskGrid [][]*Node) int {
	totalRisk := 0
	for _, p := range path {
		totalRisk += riskGrid[p.y][p.x].risk
	}
	return totalRisk
}

func part1(lines []string) int {
	riskGrid := loadGrid(lines, false)
	start := riskGrid[0][0]
	maxX, maxY := len(riskGrid[0])-1, len(riskGrid)-1
	end := riskGrid[maxY][maxX]
	path := findLowestRisk(start, end)
	return getPathRisk(path, riskGrid)
}

func part2(lines []string) int {
	riskGrid := loadGrid(lines, true)
	start := riskGrid[0][0]
	maxX, maxY := len(riskGrid[0])-1, len(riskGrid)-1
	end := riskGrid[maxY][maxX]
	path := findLowestRisk(start, end)
	return getPathRisk(path, riskGrid)
}

func main() {
	lines := getInputLines("data/15.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
