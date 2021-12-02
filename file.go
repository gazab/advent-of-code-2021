package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filepath string) []string {
	dat, err := os.ReadFile(filepath)
	check(err)

	lines := strings.Split(string(dat), "\n")

	fmt.Println(lines)

	return CleanStringArray(lines)
}
