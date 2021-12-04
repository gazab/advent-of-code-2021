package util

import (
	"os"
	"strings"
)

func ReadFileDynamicSplit(filepath string, splitstring string) []string {
	data, err := os.ReadFile(filepath)
	Check(err)
	dataString := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(dataString, splitstring)
	return CleanStringSlice(lines)
}

func ReadFile(filepath string) []string {
	return ReadFileDynamicSplit(filepath, "\n")
}

func ReadFileDoubleNewlineSplit(filepath string) []string {
	return ReadFileDynamicSplit(filepath, "\n\n")
}
