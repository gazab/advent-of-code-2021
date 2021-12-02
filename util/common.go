package util

import "strconv"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ToInt(input string) int {
	value, err := strconv.Atoi(input)
	Check(err)

	return value
}
