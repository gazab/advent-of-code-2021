package main

import (
	"fmt"
	"math"
	"sort"
)

func Solve07A(input []int) {

	sort.Ints(input)

	lowestFuel := math.MaxInt
	lowestFuelPoint := 0
	for i := input[0]; i < input[len(input)-1]; i++ {
		fuelForPoint := getTotalFuelToPointPartA(i, input)
		if fuelForPoint < lowestFuel {
			lowestFuel = fuelForPoint
			lowestFuelPoint = i
		}
		if debug {
			fmt.Println("Total fuel for", i, ":", fuelForPoint)
		}
	}
	fmt.Println("Total fuel for", lowestFuelPoint, ":", lowestFuel)
}

func getTotalFuelToPointPartA(point int, numbers []int) int {
	fuelSum := 0
	for _, s := range numbers {
		fuelSum += int(math.Abs(float64(s - point)))
	}
	return fuelSum
}

func Solve07B(input []int) {

	sort.Ints(input)

	lowestFuel := math.MaxInt
	lowestFuelPoint := 0
	for i := input[0]; i < input[len(input)-1]; i++ {
		fuelForPoint := getTotalFuelToPointPartB(i, input)
		if fuelForPoint < lowestFuel {
			lowestFuel = fuelForPoint
			lowestFuelPoint = i
		}
		if debug {
			fmt.Println("Total fuel for", i, ":", fuelForPoint)
		}
	}
	fmt.Println("Total fuel for", lowestFuelPoint, ":", lowestFuel)
}

func getTotalFuelToPointPartB(point int, numbers []int) int {
	fuelSum := 0
	for _, s := range numbers {
		fuelSum += getTriangleNumber(int(math.Abs(float64(s - point))))
	}
	return fuelSum
}

func getTriangleNumber(number int) int {
	return (number * (number + 1)) / 2
}
