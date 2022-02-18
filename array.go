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

	// traversing arrays
	// Go has only on looping construct which is for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	// array traversal with range
	for _, value := range x {
		fmt.Println(value)
	}

	getAverage()
}

func getAverage() {
	grades := [5]float64{88, 89, 99, 77, 66}

	var total float64 = 0
	for _, value := range grades {
		total += value
	}

	fmt.Println(total / float64(len(grades)))
}
