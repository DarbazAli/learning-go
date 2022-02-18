package main

import "fmt"

func main() {
	// main function begins program execution
	fmt.Println("Hello world")

	x := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(x))
}

// calculate average value func
func average(xs []float64) float64 {

	total := 0.0

	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}
