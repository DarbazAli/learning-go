package main

import "fmt"

// function main begins program execution
func main() {
	fmt.Print("Enter a degree in C˙: ")
	var degree float64
	fmt.Scanf("%f", &degree)

	// F = (C * 9 / 5 ) + 32
	var ouput = (degree * 9 / 5) + 32

	// print out ouput to standard output
	fmt.Println(ouput)
}
