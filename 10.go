package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var openings = "([{<"
var closings = ")]}>"

func Solve10(input []string) {

	sum := 0
	var autocompleteScore []int

	for _, row := range input {
		fmt.Println("--- New row:", row)
		var stack []rune
		for _, r := range row {
			if isOpening(r) {
				// Push to stack
				stack = push(r, stack)
			} else if isClosing(r) {
				// Pop stack
				var element rune
				element, stack = pop(stack)

				if matches(element, r) {
					if debug {
						fmt.Println(strconv.QuoteRune(element), "and", strconv.QuoteRune(r), " matched!")
					}
				} else {
					if debug {
						fmt.Println(strconv.QuoteRune(element), "and", strconv.QuoteRune(r), " did not match!", strconv.QuoteRune(r), "is illegal char with value of ", getIllegalCharPoint(r))
					}
					sum += getIllegalCharPoint(r)
					stack = nil
					break
				}

			}
		}
		if len(stack) > 0 {
			fmt.Println("Incomplete row. Stack is:", string(stack))
			score := 0
			for len(stack) > 0 {
				score *= 5
				var element rune
				element, stack = pop(stack)
				score += getAutocompleteCharPoint(element)
			}
			autocompleteScore = append(autocompleteScore, score)
		}
	}

	sort.Ints(autocompleteScore)

	fmt.Println("A: ", sum)
	fmt.Println("B: ", autocompleteScore[len(autocompleteScore)/2])
}

func getIllegalCharPoint(c rune) int {
	switch string(c) {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}

	return 0
}

func getAutocompleteCharPoint(c rune) int {
	switch string(c) {
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	}

	return 0
}

func matches(r1 rune, r2 rune) bool {
	if isOpening(r1) {
		return strings.Index(openings, string(r1)) == strings.Index(closings, string(r2))
	} else {
		return strings.Index(closings, string(r1)) == strings.Index(openings, string(r2))
	}
}

func isOpening(r rune) bool {
	return strings.ContainsRune(openings, r)
}

func isClosing(r rune) bool {
	return strings.ContainsRune(closings, r)
}

func push(element rune, stack []rune) []rune {
	return append(stack, element)
}

func pop(stack []rune) (rune, []rune) {
	element := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return element, stack
}
