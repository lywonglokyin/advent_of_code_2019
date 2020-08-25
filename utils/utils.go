package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadFile read a file line by line, and returns a slice.
func ReadFile(filePath string) []string {
	lines := make([]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// SliceAtoi convert string slice to int64 slice.
func SliceAtoi(texts []string) ([]int64, error) {
	numbers := make([]int64, 0, len(texts))
	for _, text := range texts {
		num, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, int64(num))
	}
	return numbers, nil
}
