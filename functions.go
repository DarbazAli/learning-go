package main

import "fmt"

func main() {
	//
	result := div(5, 2)
	fmt.Println(result)
}

func div(numerator, demoniator int) int {
	if demoniator == 0 {
		return 0
	}

	return numerator / demoniator
}
