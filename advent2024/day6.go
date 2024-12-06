package main

import (
	"fmt"
	"strings"
)

func day6_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	output := 0

	guardRow := -1
	guardCol := -1
	guardDirection := 0

	for _, row := range grid {
		if guardRow >= 0 {
			break
		}
		for col, _:= range grid[row] {
			if grid[row][col] == '^' {
				guardRow = row
				guardCol = col
				posFoune = true
				break

			}
		}
	}
	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{-1, 0}
	}
	fmt.Println("Day 6 part 1 is: ", output)
}

// func day6_2(input string) {
//	fmt.Println("Day 6 part 2 is: ", output)
// }
