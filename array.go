package main

import "fmt"

func main() {
	var arr [5]int
	arr[4] = 100

	fmt.Println(arr)

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))

	// using range
	var total2 float64 = 0
	for _, value := range x {
		total2 += value
	}

	fmt.Println(total2 / float64(len(x)))
}
