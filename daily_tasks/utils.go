package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetInput(filename string) []int {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		result = append(result, num)
	}

	file.Close()
	return result
}
