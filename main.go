package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gazab/advent-of-code-2021/util"
)

var debug = false

func main() {

	currentDay := GetDayString()
	debugPtr := flag.Bool("debug", false, "true/false")
	examplePtr := flag.Bool("example", false, "Use example input")
	dayPtr := flag.String("day", currentDay, "01-24")
	flag.Parse()

	// File path from parameters
	fileName := "input.txt"
	if *examplePtr {
		fileName = "example.txt"
	}
	filePath := "./data/" + *dayPtr + "/" + fileName

	debug = *debugPtr

	solver(*dayPtr, filePath, *debugPtr)
}

func solver(day string, filePath string, debug bool) {
	switch day {
	case "01":
		Solve01A(util.StringSliceToIntSlice(util.ReadFile(filePath)))
		Solve01B(util.StringSliceToIntSlice(util.ReadFile(filePath)))
	case "02":
		Solve02A(util.ReadFile(filePath), debug)
		Solve02B(util.ReadFile(filePath), debug)
	case "03":
		Solve03A(util.ReadFile(filePath), debug)
		Solve03B(util.ReadFile(filePath), debug)
	case "04":
		Solve04A(util.ReadFileDoubleNewlineSplit(filePath), debug)
		Solve04B(util.ReadFileDoubleNewlineSplit(filePath), debug)
	case "05":
		Solve05A(util.ReadFile(filePath))
		Solve05B(util.ReadFile(filePath))
	case "06":
		Solve06A(util.StringSliceToIntSlice(util.ReadFileCommaSplit(filePath)))
		Solve06B(util.StringSliceToIntSlice(util.ReadFileCommaSplit(filePath)))
	case "07":
		Solve07A(util.StringSliceToIntSlice(util.ReadFileCommaSplit(filePath)))
		Solve07B(util.StringSliceToIntSlice(util.ReadFileCommaSplit(filePath)))
	}
}

func GetDayString() string {
	day := time.Now().Day()
	return fmt.Sprintf("%02d", day)
}
