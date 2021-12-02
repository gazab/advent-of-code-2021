package util

import (
	"os"
	"strings"
)

func ReadFile(filepath string) []string {
	dat, err := os.ReadFile(filepath)
	Check(err)
	lines := strings.Split(string(dat), "\n")
	return CleanStringArray(lines)
}
