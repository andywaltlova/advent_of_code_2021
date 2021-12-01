package main

// Every solution is run along with utils.go
// e.g `go run 01.go utils.go`

import (
	"fmt"
)

func getIncreased(nums []int) int {
	var increased_count int = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increased_count += 1
		}
	}
	return increased_count
}

func groupByThreeSums(nums []int) []int {
	var grouped_nums []int
	for i := 0; i < len(nums); i++ {
		if (i+2 < len(nums)) && (i+1 < len(nums)) {
			grouped_nums = append(grouped_nums, nums[i]+nums[i+1]+nums[i+2])
		}
	}
	return grouped_nums
}

func main() {
	nums := getNumberInput("data/01.txt")
	part1 := getIncreased(nums)
	part2 := getIncreased(groupByThreeSums(nums))
	fmt.Println(part1)
	fmt.Println(part2)
}
