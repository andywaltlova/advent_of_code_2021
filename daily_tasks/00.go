package main

// Every solution is run along with utils.go
// e.g `go run 00.go utils.go`

import (
	"fmt"
)

func main() {
	nums := GetNumberInput("data/example_input.txt")
	fmt.Println(nums)
}
