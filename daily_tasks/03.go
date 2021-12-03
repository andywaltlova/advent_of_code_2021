package main

// Every solution is run along with utils.go
// e.g `go run 03.go utils.go`

import (
	"fmt"
	"strconv"
)

func part1(input_strings []string) int {
	var gamma string = ""
	var epsilon string = ""

	for col := 0; col < len(input_strings[0]); col++ {
		var count_0 int = 0
		var count_1 int = 0

		for row := 0; row < len(input_strings); row++ {
			if input_strings[row][col] == '0' {
				count_0 += 1
			} else {
				count_1 += 1
			}
		}

		if count_0 > count_1 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gamma_decimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_decimal, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gamma_decimal) * int(epsilon_decimal)
}

func getCO2_or_oxygen(input_strings []string, col_index int, find_oxygen bool) int {
	// Recursion should stop
	if len(input_strings) == 1 {
		num, _ := strconv.ParseInt(input_strings[0], 2, 64)
		return int(num)
	}

	// Count number of 0's and 1's
	var count_0 int = 0
	var count_1 int = 0
	for row := 0; row < len(input_strings); row++ {
		if input_strings[row][col_index] == '0' {
			count_0 += 1
		} else {
			count_1 += 1
		}
	}

	// Correct counts based on lookup value
	if count_0 == count_1 {
		// oxygen keeps most common and 1 in case of equality
		// co2 keeps least common and 0 in case of equality
		count_1 += 1
	}

	// Filter input
	var new_input []string
	var to_keep byte = '0'
	if (find_oxygen && count_1 > count_0) || (!find_oxygen && count_1 < count_0) {
		to_keep = '1'
	}
	for _, str_num := range input_strings {
		if str_num[col_index] == to_keep {
			new_input = append(new_input, str_num)
		}
	}
	return getCO2_or_oxygen(new_input, col_index+1, find_oxygen)
}

func part2(input_strings []string) int {
	oxygen := getCO2_or_oxygen(input_strings, 0, true)
	co2 := getCO2_or_oxygen(input_strings, 0, false)
	return co2 * oxygen
}

func main() {
	lines := getInputLines("data/03.txt")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
