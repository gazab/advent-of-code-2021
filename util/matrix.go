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

func PrintMatrix(matrix [][]int, title string) {

	size := len(matrix)

	print_matrix_header(size, title)
	for row := 0; row < len(matrix); row++ {
		fmt.Print("║ ")
		for col := 0; col < len(matrix[row]); col++ {
			cell := matrix[row][col]
			if cell > 0 {
				ct.Foreground(ct.Red, cell > 1)
			}
			fmt.Printf("%2d ", cell)
			ct.ResetColor()
		}
		fmt.Println("║")
	}
	print_matrix_footer(size)
}
