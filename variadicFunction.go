/*
	Variadic Functions

	@author: Darbaz Ali
	@date: Feb 19, 2022

	@description:
	Variadic Function is a special type of function that allows to pass
	multiple optional argumensts
*/
package main

import "fmt"

func main() {
	x := add(1, 2, 3)
	fmt.Println(x)
}

func add(args ...int) int {
	total := 0

	for _, v := range args {
		total += v
	}

	return total
}
