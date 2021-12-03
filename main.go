package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gazab/advent-of-code-2021/util"
)

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
		//Solve03A(util.ReadFile(filePath), debug)
		Solve03B(util.ReadFile(filePath), debug)
	}
}

func GetDayString() string {
	day := time.Now().Day()
	return fmt.Sprintf("%02d", day)
}
