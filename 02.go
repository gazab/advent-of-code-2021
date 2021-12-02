package main

import (
	"fmt"
	"strings"

	"github.com/gazab/advent-of-code-2021/util"
)

func Solve02A(input []string, debug bool) {

	pos := 0
	depth := 0
	for i, s := range input {
		if debug {
			fmt.Println(i, s)
		}
		instruction := strings.Split(s, " ")
		command := instruction[0]
		arg := util.ToInt(instruction[1])

		switch command {
		case "forward":
			pos += arg
		case "up":
			depth -= arg
		case "down":
			depth += arg
		}
	}

	fmt.Println("Horizontal position: ", pos)
	fmt.Println("Depth: ", depth)
	fmt.Println("Multiplied: ", pos*depth)
}

func Solve02B(input []string, debug bool) {

	pos := 0
	depth := 0
	aim := 0
	for i, s := range input {
		if debug {
			fmt.Println(i, s)
		}
		instruction := strings.Split(s, " ")
		command := instruction[0]
		arg := util.ToInt(instruction[1])

		switch command {
		case "forward":
			pos += arg
			depth = depth + (aim * arg)
		case "up":
			aim -= arg
		case "down":
			aim += arg
		}
	}

	fmt.Println("Horizontal position: ", pos)
	fmt.Println("Depth: ", depth)
	fmt.Println("Multiplied: ", pos*depth)
}
