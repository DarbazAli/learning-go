package main

import "fmt"

// function main begins program execution
func main() {
	fmt.Print("Eneter a number in feets: ")
	var number float32
	fmt.Scanf("%f", &number)

	var output float32 = number * 0.3048

	fmt.Println(output, "meters")
}
