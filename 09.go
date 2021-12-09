package main

import (
	"fmt"
	"sort"
	"strconv"

	ct "github.com/daviddengcn/go-colortext"
)

func Solve09A(input []string) {

	var lowPoints []int

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			num, _ := strconv.Atoi(string(input[row][col]))
			adjacent := getAdjacent(row, col, input)
			lowPoint := true
			for _, a := range adjacent {
				if a <= num {
					lowPoint = false
				}
			}
			if lowPoint {
				lowPoints = append(lowPoints, num)
				ct.Foreground(ct.Red, false)
				fmt.Printf("%d", num)
				ct.ResetColor()
			} else {
				fmt.Printf("%d", num)
			}
		}
		fmt.Printf("\n")
	}

	sum := 0
	for _, n := range lowPoints {
		sum += n + 1
	}

	fmt.Println("Answer:", sum)
}

func getAdjacent(row int, col int, input []string) []int {
	var adjacent []int
	if col > 0 {
		num, _ := strconv.Atoi(string(input[row][col-1])) // y, x-1
		adjacent = append(adjacent, num)
	}

	if col < len(input[row])-1 {
		num, _ := strconv.Atoi(string(input[row][col+1])) // y, x+1
		adjacent = append(adjacent, num)
	}

	if row > 0 {
		num, _ := strconv.Atoi(string(input[row-1][col])) // y-1, x
		adjacent = append(adjacent, num)
	}

	if row < len(input)-1 {
		num, _ := strconv.Atoi(string(input[row+1][col])) // y+1, x
		adjacent = append(adjacent, num)
	}

	return adjacent
}

type coord struct {
	row int
	col int
}

func Solve09B(input []string) {

	lowPoints := make(map[coord]bool)

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			num, _ := strconv.Atoi(string(input[row][col]))
			adjacent := getAdjacent(row, col, input)
			lowPoint := true
			for _, a := range adjacent {
				if a <= num {
					lowPoint = false
				}
			}
			if lowPoint {
				lowPoints[coord{row: row, col: col}] = true
			}
		}
	}

	fmt.Println("Low points:", lowPoints)

	basinPoints := make(map[coord]bool)
	var basinSizes []int
	for point := range lowPoints {
		basinPointsForLowPoint := getBasinPoints(point, input)
		basinSizes = append(basinSizes, len(basinPointsForLowPoint))
		fmt.Println("Found", len(basinPointsForLowPoint), " basin points for", point)
		for key := range basinPointsForLowPoint {
			basinPoints[key] = true
		}
	}

	printCave(lowPoints, basinPoints, input)

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println("Result:", basinSizes[0]*basinSizes[1]*basinSizes[2])
}

func printCave(lowPoints map[coord]bool, basinPoints map[coord]bool, input []string) {
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			coord := coord{row: row, col: col}
			num := input[coord.row][coord.col]
			if lowPoints[coord] {
				ct.Foreground(ct.Blue, false)
			} else if basinPoints[coord] {
				ct.Foreground(ct.Blue, true)
			}
			fmt.Print(string(num))
			ct.ResetColor()
		}
		fmt.Printf("\n")
	}
}

func getAdjacentBasinCoords(c coord, input []string) map[coord]bool {
	adjacent := make(map[coord]bool)
	if c.col > 0 {
		num, _ := strconv.Atoi(string(input[c.row][c.col-1])) // y, x-1
		if num != 9 {
			adjacent[coord{row: c.row, col: c.col - 1}] = true
		}
	}

	if c.col < len(input[c.row])-1 {
		num, _ := strconv.Atoi(string(input[c.row][c.col+1])) // y, x+1
		if num != 9 {
			adjacent[coord{row: c.row, col: c.col + 1}] = true
		}
	}

	if c.row > 0 {
		num, _ := strconv.Atoi(string(input[c.row-1][c.col])) // y-1, x
		if num != 9 {
			adjacent[coord{row: c.row - 1, col: c.col}] = true
		}
	}

	if c.row < len(input)-1 {
		num, _ := strconv.Atoi(string(input[c.row+1][c.col])) // y+1, x
		if num != 9 {
			adjacent[coord{row: c.row + 1, col: c.col}] = true
		}
	}

	return adjacent
}

func getBasinPoints(c coord, input []string) map[coord]bool {
	points := make(map[coord]bool)

	pointsToCheck := make(map[coord]bool)
	pointsToCheck[c] = true
	for len(pointsToCheck) > 0 {
		newPointsToCheck := make(map[coord]bool)
		for p := range pointsToCheck {
			adjacentPoints := getAdjacentBasinCoords(p, input)
			for key := range adjacentPoints {
				if !points[key] {
					points[key] = true
					newPointsToCheck[key] = true
				}
			}
		}
		pointsToCheck = newPointsToCheck
	}

	return points
}
