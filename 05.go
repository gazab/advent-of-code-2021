package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gazab/advent-of-code-2021/util"
)

var coordRegex = regexp.MustCompile(`(?P<x1>\d*),(?P<y1>\d*) -> (?P<x2>\d*),(?P<y2>\d*)`)

func Solve05A(input []string) {
	solve(input, true)
}

func Solve05B(input []string) {
	solve(input, false)
}

func solve(input []string, onlyStraight bool) {
	// Create matrix
	matrixSize := 10
	if len(input) > 15 {
		matrixSize = 1000
	}
	board := make([][]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		board[i] = make([]int, matrixSize)
	}

	if debug {
		println(input)
	}

	// Process input and mark matrix
	for _, s := range input {
		x1, y1, x2, y2 := getCoords(s)

		if !onlyStraight || (x1 == x2 || y1 == y2) {
			if debug {
				fmt.Printf("Drawing: (%d,%d) -> (%d,%d)\n", x1, y1, x2, y2)
			}

			xDir := 0
			yDir := 0
			if x1 < x2 {
				xDir = 1
			} else if x1 > x2 {
				xDir = -1
			}

			if y1 < y2 {
				yDir = 1
			} else if y1 > y2 {
				yDir = -1
			}

			if debug {
				fmt.Println("Dir: ", xDir, yDir)
			}

			xPos := x1
			yPos := y1

			for {
				if debug {
					fmt.Printf("Marking : (%d,%d)\n", xPos, yPos)
				}
				board[yPos][xPos]++

				// Check if we should stop
				if xPos == x2 && yPos == y2 {
					break
				}

				// If not, one more step
				xPos += xDir
				yPos += yDir
			}
		}

	}

	// Print output
	if debug {
		util.PrintMatrix(board, "Lines")
	}

	fmt.Println("Number of overlaps: ", countOverlaps(board))
}

func countOverlaps(board [][]int) int {

	overlaps := 0
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] > 1 {
				overlaps++
			}
		}
	}

	return overlaps
}

func getCoords(input string) (int, int, int, int) {

	matches := coordRegex.FindStringSubmatch(input)
	x1 := getCoord(matches, "x1")
	y1 := getCoord(matches, "y1")
	x2 := getCoord(matches, "x2")
	y2 := getCoord(matches, "y2")

	if debug {
		fmt.Println(x1, y1, x2, y2)
	}

	return x1, y1, x2, y2
}

func getCoord(matches []string, coord string) int {
	val := matches[coordRegex.SubexpIndex(coord)]
	coordVal, err := strconv.Atoi(val)
	util.Check(err)
	return coordVal
}
