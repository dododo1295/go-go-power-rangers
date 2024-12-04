package main

import (
	"fmt"
	"strings"
)

// I think the best way is to remove extraneous letters and then count the combinations left over and add it to a count?
// let's see. This is for sure wrong though...
func day4_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	output := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{-1, -1},
		{1, -1},
		{1, 1},
	}
	word := "XMAS"

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {

			for _, dir := range directions {
				rowIndex := dir[0]
				colIndex := dir[1]
				isXmas := true

				for charIndex := 0; charIndex < len(word); charIndex++ {
					offsetRow := row + (rowIndex * charIndex)
					offsetCol := col + (colIndex * charIndex)

					if offsetRow < 0 || offsetRow >= len(grid) || offsetCol < 0 || offsetCol >= len(grid[row]) {
						isXmas = false
						break
					}

					if grid[offsetRow][offsetCol] != rune(word[charIndex]) {
						isXmas = false
					}
				}
				if isXmas {
					output++
				}
			}
		}
	}

	fmt.Println(" output day 4 part 1: ", output)
}

func day4_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	output := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2][2]int{
		{
			{-1, 1},
			{1, -1},
		},
		{
			{-1, -1},
			{1, 1},
		},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {

			if grid[row][col] != 'A' {
				continue
			}

			isXmas := true

			for _, dir := range directions {
				row1Index := row + dir[0][0]
				col1Index := col + dir[0][1]

				row2Index := row + dir[1][0]
				col2Index := col + dir[1][1]

				if row1Index < 0 || row1Index >= len(grid) || col1Index < 0 || col1Index >= len(grid[row]) || row2Index < 0 || row2Index >= len(grid) || col2Index < 0 || col2Index >= len(grid[row]) {
					isXmas = false
					break
				}

				if (grid[row1Index][col1Index] == 'M' && grid[row2Index][col2Index] == 'S') || (grid[row1Index][col1Index] == 'S' && grid[row2Index][col2Index] == 'M') {
					continue
				}

				isXmas = false
				break
			}

			if isXmas {
				output++
			}
		}
	}

	fmt.Println(" Excruciating day of full tutorial on how to complete... I understand the code but oof, this was confusing. Output Day 4 Part 2", output)
}
