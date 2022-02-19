/*
	Closures

	@author: Darbaz Ali
	@date: Feb 19, 2022

	@desc
	A closue is a function that returns another function
	The inner function has access to the local variables in scope of
	outer function

*/
package main

import "fmt"

func main() {

	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
