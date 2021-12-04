package util

import "strings"

func WhiteSpaceSplit(s string) []string {
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.Trim(s, " \r\n")
	return strings.Split(s, " ")
}
