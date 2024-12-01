package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// this imports my inputs/day#.txt files so remember to use this to make life much easier.
func getInput(day int) string {
	inputDir := "inputs"
	inputFile := filepath.Join(inputDir, fmt.Sprintf("day%d.txt", day))

	if _, err := os.Stat(inputFile); err == nil {
		content, _ := os.ReadFile(inputFile)
		return string(content)
	}
	return string(inputFile)
}

func main() {
	day := 1
	input := getInput(day)
	// actual code
	day1_1(input)
	day1_2(input)
}