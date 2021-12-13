package util

import (
	"strconv"
	"strings"
)

func StringSliceToIntSlice(input []string) []int {
	var result []int

	for _, x := range input {
		value, err := strconv.Atoi(x)
		Check(err)
		result = append(result, value)
	}

	return result
}

func StringSliceToIntMatrix(input []string) [][]int {
	var result [][]int

	for _, x := range input {
		result = append(result, StringSliceToIntSlice(strings.Split(x, "")))
	}

	return result
}

func BinaryStringSliceToIntSlice(input []string) []int {
	var result []int

	for _, x := range input {
		value, err := strconv.ParseInt(x, 2, 64)
		Check(err)
		result = append(result, int(value))
	}

	return result
}

func CleanStringSlice(input []string) []string {
	var result []string

	for _, x := range input {
		value := strings.Trim(x, "\r\n ")
		if len(value) > 0 {
			result = append(result, value)
		}
	}

	return result
}
