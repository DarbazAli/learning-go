package main

import "fmt"

func main() {

	// array creation - first form
	var x [5]int
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	fmt.Println(x)

	// array creating - short syntax
	arr := [5]int{83, 92, 82, 32, 65}

	fmt.Println(arr)
}
