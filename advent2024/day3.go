package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// i need to remove all extraneous characters that are not
// mul() or 1-3 digit numbers, and commas.
// then I need to multiply the mul expressions and add the results together.
func day3_1(input string) {
	output := 0
	//easiest way is to regex out what I need.
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matching := re.FindAllStringSubmatch(input, -1)

	for _, match := range matching {
		numerouno, _ := strconv.Atoi(match[1])
		numerodos, _ := strconv.Atoi(match[2])
		output += numerouno * numerodos
	}

	fmt.Println("understand my task but need youtube to figure out how to do the regex required. output for part 1 is: ", output)
}

func day3_2(input string) {
	output := 0
	//easiest way is to regex out what I need.
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matching := re.FindAllStringSubmatch(input, -1)
	do := true
	for _, match := range matching {
		if match[0] == "do()" {
			do = true
			continue
		} else if match[0] == "don't()" {
			do = false
			continue
		} else if !do {
			continue
		}
		numerouno, _ := strconv.Atoi(match[1])
		numerodos, _ := strconv.Atoi(match[2])
		output += numerouno * numerodos
	}

	fmt.Println("learning how to do these if statements with array objects is quite interesting. output for part 2 is: ", output)
}
