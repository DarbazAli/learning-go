package main

import "fmt"

func main() {
	// main function begins program execution
	fmt.Println("Hello world")

	x := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(x))

	a, b := f()
	fmt.Println(a, b)
}

// calculate average value func
func average(xs []float64) float64 {

	total := 0.0

	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}

// function with multiple returns
func f() (int, int) {
	return 5, 10
}
