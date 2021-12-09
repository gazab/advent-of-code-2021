package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gazab/advent-of-code-2021/util"
)

func Solve08A(input []string) {
	sum := 0
	for _, s := range input {
		_, output := parseLine(s)
		sum += countKnownDigits(output)
	}
	fmt.Println(sum)
}

func countKnownDigits(output []string) int {
	sum := 0
	for _, s := range output {
		if len(s) != 5 && len(s) != 6 {
			sum++
		}
	}

	return sum
}

func parseLine(input string) ([]string, []string) {
	parts := strings.Split(input, "|")

	return util.CleanStringSlice(strings.Split(parts[0], " ")), util.CleanStringSlice(strings.Split(parts[1], " "))
}

func Solve08B(input []string) {

	sum := 0
	for i := 0; i < len(input); i++ {
		signal, output := parseLine(input[i])
		fmt.Println(signal)
		fmt.Println(output)

		sets := []map[string]bool{}

		signalSet := createSets(signal)
		outputSet := createSets(output)
		sets = append(sets, signalSet...)
		sets = append(sets, outputSet...)

		decoded := decode(sets)
		if debug {
			fmt.Println(decoded)
		}

		fmt.Println(crackCode(signalSet, decoded))
		code, err := strconv.Atoi(crackCode(outputSet, decoded))
		util.Check(err)
		sum += code
		fmt.Println(code)
	}
	fmt.Println(sum)
}

func crackCode(outputSet []map[string]bool, decoded map[int]map[string]bool) string {

	result := ""

	for _, n := range outputSet {
		for key, d := range decoded {
			if reflect.DeepEqual(n, d) {
				result += strconv.Itoa(key)
			}

		}
	}

	return result
}

func decode(sets []map[string]bool) map[int]map[string]bool {
	decoded := map[int]map[string]bool{}

	// Beautiful monster
	for len(decoded) < 10 {
		for _, s := range sets {
			if (len(s)) == 2 {
				decoded[1] = s
			} else if len(s) == 3 {
				decoded[7] = s
			} else if len(s) == 4 {
				decoded[4] = s
			} else if len(s) == 7 {
				decoded[8] = s
			} else if len(s) == 6 {
				if len(decoded[4]) > 0 && len(util.Difference(s, decoded[4])) == 0 {
					decoded[9] = s
				} else if len(decoded[1]) > 0 && len(util.Difference(s, decoded[1])) == 0 {
					decoded[0] = s
				} else {
					decoded[6] = s
				}
			} else if len(s) == 5 {
				if len(decoded[1]) > 0 && len(util.Difference(s, decoded[1])) == 0 {
					decoded[3] = s
				} else if len(decoded[4]) > 0 && len(util.Difference(s, decoded[4])) == 1 {
					decoded[5] = s
				} else if len(decoded[3]) > 0 && len(decoded[5]) > 0 {
					decoded[2] = s
				}

			}
		}
	}

	return decoded
}

func createSets(input []string) []map[string]bool {
	sets := []map[string]bool{}

	for _, n := range input {
		sets = append(sets, createSet(n))
	}

	return sets
}

func createSet(number string) map[string]bool {
	s := map[string]bool{}

	for _, r := range number {
		s[string(r)] = true
	}

	return s
}
