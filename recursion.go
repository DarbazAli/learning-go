/*
	Recursion

	@author: Darbaz Ali
	@date: Feb 19, 2022

	@desc
	Recursion is a function that calls itslef.
*/
package main

import "fmt"

func main() {

	fac := factorial(10)
	fmt.Println(fac)
}

func factorial(x uint) uint {
	fmt.Println(x)
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
