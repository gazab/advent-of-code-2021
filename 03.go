package main

import (
	"fmt"
	"strconv"

	"github.com/gazab/advent-of-code-2021/util"
)

func Solve03A(input []string, d bool) {

	oneCount := make([]int, len(input[0]))

	for i, row := range input {
		fmt.Println(i, row)
		for j, rune := range row {
			if string(rune) == "1" {
				oneCount[j] += 1
			}
		}
	}

	fmt.Println(oneCount)

	result := 0
	complement := 0
	for _, count := range oneCount {
		result = result << 1
		complement = complement << 1
		if count > len(input)-count {
			result++
		} else {
			complement++
		}
	}

	fmt.Println("Decimal: ", result)
	fmt.Printf("Binary: %b\n", result)
	fmt.Println("Complement decimal: ", complement)
	fmt.Printf("Complement binary: %b\n", complement)
	fmt.Println("Multiplied: ", result*complement)
}

func Solve03B(input []string, d bool) {

	oxygenRating := find_value(input, 0, true, d)
	co2ScrubberRating := find_value(input, 0, false, d)

	fmt.Printf("Oxygen rating binary: %b\n", oxygenRating)
	fmt.Println("Oxygen rating decimal: ", oxygenRating)
	fmt.Println("CO2 scrubber decimal: ", co2ScrubberRating)
	fmt.Printf("CO2 scrubber binary: %b\n", co2ScrubberRating)

	fmt.Println("Multiplied: ", oxygenRating*co2ScrubberRating)
}

func find_value(input []string, index int, mostCommon bool, d bool) int {

	// End recursion and return value
	if len(input) == 1 {
		val, err := strconv.ParseInt(input[0], 2, 0)
		util.Check(err)
		return int(val)
	}

	var oneValues, zeroValues []string

	for i, row := range input {
		fmt.Println(i, row)
		if string(row[index]) == "1" {
			oneValues = append(oneValues, row)
		} else {
			zeroValues = append(zeroValues, row)
		}
	}

	fmt.Println("--------")
	if len(oneValues) >= len(zeroValues) == mostCommon {
		return find_value(oneValues, index+1, mostCommon, d)
	}
	return find_value(zeroValues, index+1, mostCommon, d)
}
