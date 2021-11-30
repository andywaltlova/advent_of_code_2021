package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetNumberInput(filename string) []int {
	var result []int
	for _, s := range GetInputLines(filename) {
		num, _ := strconv.Atoi(s)
		result = append(result, num)
	}
	return result
}

func GetInputLines(filename string) []string {
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
