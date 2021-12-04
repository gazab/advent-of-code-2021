package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/gazab/advent-of-code-2021/util"
)

type bingosquare struct {
	number int
	marked bool
}

func Solve04A(input []string, debug bool) {
	numbers, cards := setup(input, debug)

	fmt.Println("Marking cards")
	winner, index, number := mark_and_check_for_bingo(numbers, cards, debug)
	fmt.Println("Last number: ", number)
	print_card(winner, index)
	cardSum := get_card_sum(winner)
	fmt.Println("Card sum: ", cardSum)
	fmt.Println("Multiplied: ", cardSum*number)
}

func setup(input []string, debug bool) ([]int, [][][]bingosquare) {
	numbers := util.StringSliceToIntSlice(strings.Split(input[0], ","))
	fmt.Println("Numbers: ", numbers)

	var cards [][][]bingosquare

	for i := 1; i < len(input); i++ {
		cards = append(cards, create_bingo_card(input[i]))
	}

	fmt.Println("Created", len(cards), "cards")
	return numbers, cards
}

func Solve04B(input []string, debug bool) {
	numbers, cards := setup(input, debug)

	for {

		fmt.Println("Cards left: ", len(cards))
		winner, index, number := mark_and_check_for_bingo(numbers, cards, debug)
		if len(cards) == 1 {
			fmt.Println("Last number: ", number)
			print_card(winner, index)
			cardSum := get_card_sum(winner)
			fmt.Println("Card sum: ", cardSum)
			fmt.Println("Multiplied: ", cardSum*number)
			break
		} else {
			// Remove winner card and go again
			copy(cards[index:], cards[index+1:])
			cards[len(cards)-1] = nil
			cards = cards[:len(cards)-1]
		}
	}
}

func get_card_sum(card [][]bingosquare) int {
	sum := 0
	for row := 0; row < len(card); row++ {
		for col := 0; col < len(card[row]); col++ {
			if !card[row][col].marked {
				sum += card[row][col].number
			}
		}
	}

	return sum
}

func mark_and_check_for_bingo(numbers []int, cards [][][]bingosquare, debug bool) ([][]bingosquare, int, int) {
	for _, number := range numbers {
		if debug {
			fmt.Println("New number: ", number)
		}
		for index, card := range cards {
			find_number_on_card(number, card, true)
			if has_bingo(card) {
				print_bingo()
				return card, index, number
			}

		}
		if debug {
			print_cards(cards)
		}
	}

	return nil, 0, 0
}

func has_bingo(card [][]bingosquare) bool {

	// Horizontal bingo
	for row := 0; row < len(card); row++ {
		for col := 0; col < len(card[row]); col++ {
			if !card[row][col].marked {
				break
			} else if col == 4 {
				return true
			}
		}
	}

	// Vertical bingo
	for col := 0; col < len(card); col++ {
		for row := 0; row < len(card[col]); row++ {
			if !card[row][col].marked {
				break
			} else if row == 4 {
				return true
			}
		}
	}

	return false
}

func print_bingo() {
	ct.Foreground(ct.Red, true)
	fmt.Print("B")
	ct.Foreground(ct.Cyan, true)
	fmt.Print("I")
	ct.Foreground(ct.Green, true)
	fmt.Print("N")
	ct.Foreground(ct.Yellow, true)
	fmt.Print("G")
	ct.Foreground(ct.Magenta, true)
	fmt.Print("O")
	ct.Foreground(ct.Blue, true)
	fmt.Print("!\n")
	ct.ResetColor()

}

func print_cards(cards [][][]bingosquare) {
	for i, card := range cards {
		print_card(card, i)
	}
}

func print_card_header(index int) {
	fmt.Printf("╔╣ ")
	ct.Foreground(ct.Green, false)
	fmt.Printf("%d", index)
	ct.ResetColor()
	fmt.Printf(" ╠")
	indexSize := len(strconv.Itoa(index))
	for i := 0; i < 12-indexSize; i++ {
		fmt.Printf("═")
	}

	fmt.Printf("╗\n")
}

func print_card(card [][]bingosquare, index int) {
	print_card_header(index)
	for row := 0; row < len(card); row++ {
		fmt.Print("║ ")
		for col := 0; col < len(card[row]); col++ {
			square := card[row][col]
			if square.marked {
				ct.Foreground(ct.Red, false)
				fmt.Printf("%2d ", square.number)
				ct.ResetColor()
			} else {
				fmt.Printf("%2d ", square.number)
			}
		}
		fmt.Println("║")
	}
	fmt.Printf("╚════════════════╝\n\n")
}

func create_bingo_card(input string) [][]bingosquare {
	var card [][]bingosquare

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		rowNumberSlice := util.StringSliceToIntSlice(util.WhiteSpaceSplit(row))
		var bingoSquareRow []bingosquare
		for _, number := range rowNumberSlice {
			bingoSquareRow = append(bingoSquareRow, bingosquare{number: number})
		}

		card = append(card, bingoSquareRow)
	}

	return card
}

func find_number_on_card(number int, card [][]bingosquare, mark bool) (int, int, error) {

	for row := 0; row < len(card); row++ {
		for col := 0; col < len(card[row]); col++ {
			if card[row][col].number == number {
				if mark {
					card[row][col].marked = true
				}
				return row, col, nil
			}
		}
	}

	return 0, 0, errors.New("Number not found")
}
