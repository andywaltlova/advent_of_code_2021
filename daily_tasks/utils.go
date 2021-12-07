package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func arrayMax(s []int) int {
	m := s[0]
	for _, v := range s {
		if v > m {
			m = v
		}
	}
	return m
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func arrayMin(s []int) int {
	m := s[0]
	for _, v := range s {
		if v < m {
			m = v
		}
	}
	return m
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func getNumberInput(filename string) []int {
	var result []int
	for _, s := range getInputLines(filename) {
		num, _ := strconv.Atoi(s)
		result = append(result, num)
	}
	return result
}

func getInputLines(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	file.Close()
	return result
}

func getNumbersOnLine(filename string) []int {
	string_nums := strings.Split(getInputLines(filename)[0], ",")

	var result []int
	for _, str_num := range string_nums {
		num, _ := strconv.Atoi(str_num)
		result = append(result, num)
	}
	return result
}
