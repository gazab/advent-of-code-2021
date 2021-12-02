package util

import (
	"strconv"
	"strings"
)

func StringArrayToIntArray(input []string) []int {
	var result []int

	for _, x := range input {
		value, err := strconv.Atoi(x)
		Check(err)
		result = append(result, value)
	}

	return result
}

func CleanStringArray(input []string) []string {
	var result []string

	for _, x := range input {
		value := strings.Trim(x, "\r\n ")
		if len(value) > 0 {
			result = append(result, value)
		}
	}

	return result
}
