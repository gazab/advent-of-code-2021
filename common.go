package main

import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ToInt(input string) int {
	value, err := strconv.Atoi(input)
	check(err)

	return value
}
