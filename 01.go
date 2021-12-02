package main

import (
	"fmt"
	"strconv"
)

func Solve01A(input []int) {

	lines := input

	lastValue := -1
	increased := 0
	for i := 0; i < len(lines); i++ {
		value := lines[i]
		if lastValue != -1 && value > lastValue {
			increased++
			fmt.Println(strconv.Itoa(value) + ": Increased!")
		} else {
			fmt.Println(strconv.Itoa(value) + ": Didn't increase!")
		}
		lastValue = value
	}

	fmt.Println(increased)
}

func Solve01B(input []int) {
	increased := 0
	sumOne := 0
	sumTwo := 0
	for i := 3; i < len(input); i++ {
		sumOne = input[i-3] + input[i-2] + input[i-1]
		sumTwo = input[i-2] + input[i-1] + input[i]

		fmt.Println(sumOne)
		fmt.Println(sumTwo)

		if sumOne < sumTwo {
			increased++
		}
	}

	fmt.Println(increased)
}
