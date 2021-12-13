package main

import (
	"fmt"

	"github.com/gazab/advent-of-code-2021/util"
)

func Solve11A(input [][]int) {
	steps := 1000
	fmt.Println("Initial data")
	printMatrix(input)

	flashes := 0
	flashesAtStep100 := 0

	for i := 1; i <= steps; i++ {
		fmt.Println("After step", i)
		performStep(input)
		newFlashes := countFlashes(input)
		flashes += newFlashes

		if i == 100 {
			flashesAtStep100 = flashes
		}

		printMatrix(input)

		if newFlashes == len(input)*len(input[0]) {
			fmt.Println("All flashed at step", i)
			break
		}

	}

	fmt.Println("Number of flashes after 100 steps:", flashesAtStep100)
}

func printMatrix(input [][]int) {
	util.PrintMatrixWithHightlighter(input, "Squidz", hasFlashed, shouldFlash)
}

func hasFlashed(x int) bool {
	return x == 0
}

func shouldFlash(x int) bool {
	return x > 9
}

func performStep(input [][]int) {
	increaseAll(input)
	for flash(input) {

	}

}

func increaseAll(input [][]int) {
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			input[row][col]++
		}
	}
}

func countFlashes(input [][]int) int {
	flashes := 0
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if hasFlashed(input[row][col]) {
				flashes++
			}
		}
	}

	return flashes
}

func flash(input [][]int) bool {
	shouldFlashAgain := false

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			num := input[row][col]
			if shouldFlash(num) {
				input[row][col] = 0
				for y := util.Max(0, row-1); y <= util.Min(len(input)-1, row+1); y++ {
					for x := util.Max(0, col-1); x <= util.Min(len(input[row])-1, col+1); x++ {
						if !(x == col && y == row) {
							shouldFlashAgain = increasePoint(y, x, input) || shouldFlashAgain
						}
					}
				}
			}
		}
	}

	return shouldFlashAgain
}

func increasePoint(row int, col int, input [][]int) bool {
	num := input[row][col]
	if !hasFlashed(num) {
		input[row][col]++
	}
	return shouldFlash(input[row][col])
}
