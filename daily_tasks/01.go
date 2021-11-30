package main

// Every solution is run along with utils.go
// e.g `go run 01.go utils.go`

import (
	"fmt"
)

func main() {
	nums := GetInput("data/example_input.txt")
	fmt.Println(nums)
}
