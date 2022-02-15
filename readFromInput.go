package main

import "fmt"

func main() {
	// promt user to enter a number
	fmt.Print("Enter a number: ")
	var input float64

	// read from standard input, and put it inside "input" variable
	fmt.Scanf("%f", &input)

	// declare a new variable and assign to input * 2
	output := input * 2

	// print output to standard output
	fmt.Println(output)
}
