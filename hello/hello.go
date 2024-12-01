package main

import (
	"fmt"
	"math/rand"
)

func Hello() string {

	return "Hello World"
}

func main() {

	fmt.Println(Hello(), ". A random number is:", rand.Intn(10))

	sum := 0
	for i := 0; i <= 10; i++ {
		sum = i * 2

	}

	fmt.Println(sum)

	sumTwo := 1
	for sumTwo < 1000 {
		sumTwo += sumTwo
	}
	fmt.Println(sumTwo)
}
