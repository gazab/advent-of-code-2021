package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	currentDay := GetDayString()
	debugPtr := flag.Bool("debug", false, "true/false")
	dayPtr := flag.String("day", currentDay, "01-24")
	flag.Parse()

	filePath := "./" + *dayPtr + "/example.txt"
	debug := *debugPtr

	switch *dayPtr {
	case "01":
		Solve01A(StringArrayToIntArray(ReadFile((filePath))))
		Solve01B(StringArrayToIntArray(ReadFile((filePath))))
	case "02":
		Solve02A((ReadFile((filePath))), debug)
		Solve02B((ReadFile((filePath))), debug)
	}

}

func GetDayString() string {
	day := time.Now().Day()
	return fmt.Sprintf("%02d", day)
}
