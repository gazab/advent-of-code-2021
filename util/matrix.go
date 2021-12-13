package util

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
)

func print_matrix_header(size int, title string) {
	fmt.Printf("╔╣ ")
	ct.Foreground(ct.Green, false)
	fmt.Printf("%s", title)
	ct.ResetColor()
	fmt.Printf(" ╠")
	indexSize := len(title)
	for i := 0; i <= size*3-indexSize-4; i++ {
		fmt.Printf("═")
	}

	fmt.Printf("╗\n")
}

func print_matrix_footer(size int) {
	fmt.Printf("╚")
	for i := 0; i <= size*3; i++ {
		fmt.Printf("═")
	}

	fmt.Printf("╝\n")
}

type highlighter func(int) bool

func standardHighlighter(x int) bool {
	return x > 0
}

func standardBrighter(x int) bool {
	return x > 1
}

func PrintMatrix(matrix [][]int, title string) {
	PrintMatrixWithHightlighter(matrix, title, standardHighlighter, standardBrighter)
}

func PrintMatrixWithHightlighter(matrix [][]int, title string, fn highlighter, fn2 highlighter) {

	size := len(matrix)

	print_matrix_header(size, title)
	for row := 0; row < len(matrix); row++ {
		fmt.Print("║ ")
		for col := 0; col < len(matrix[row]); col++ {
			cell := matrix[row][col]
			if fn(cell) {
				ct.Foreground(ct.Red, fn2(cell))
			}
			if fn2(cell) {
				ct.Foreground(ct.Red, true)
			}
			fmt.Printf("%2d ", cell)
			ct.ResetColor()
		}
		fmt.Println("║")
	}
	print_matrix_footer(size)
}
