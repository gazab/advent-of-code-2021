package main

func main() {

	//filePath := "./01/example.txt"
	filePath := "./01/input.txt"

	Solve01A(StringArrayToIntArray(ReadFile((filePath))))
	Solve01B(StringArrayToIntArray(ReadFile((filePath))))
}
