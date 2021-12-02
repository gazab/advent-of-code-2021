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
	dayPtr := flag.String("day", currentDay, "01-24")
	flag.Parse()

	filePath := "./data/" + *dayPtr + "/example.txt"
	debug := *debugPtr

	switch *dayPtr {
	case "01":
		Solve01A(util.StringArrayToIntArray(util.ReadFile((filePath))))
		Solve01B(util.StringArrayToIntArray(util.ReadFile((filePath))))
	case "02":
		Solve02A(util.ReadFile((filePath)), debug)
		Solve02B(util.ReadFile((filePath)), debug)
	}

}

func GetDayString() string {
	day := time.Now().Day()
	return fmt.Sprintf("%02d", day)
}
