package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day1_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)

		// NOTE: leftNum & rightNum converts the first number in the array into an integer
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		// NOTE: len() is quick way to find the length of an array or string or integer.
		if distance < 0 {
			distance *= -1
		}

		totalDistance += distance
	}

	fmt.Println("Output of day 1 part 1 is: ", totalDistance)

}

// NOTE: need to find the frequency of a number in both arrays.
// i need to multiply the frequencies of the array numbers together if their in both arrays.
// Then i need to add the frequncy scores together.

func day1_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	rightMap := make(map[int]int)

	for _, num := range right {
		rightMap[num]++
	}

	similarityScore := 0

	for _, num := range left {
		similarityScore += num * rightMap[num]
	}
	fmt.Println("Look at this, needed a bit of help to organize BUT we did it, part 2 is: ", similarityScore)

}

// NOTE: I figured out how to do the array mapping myself but needed some internet help
// to organize my loops, this was super fun and really informative. on to day 2!
