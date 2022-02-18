package main

import "fmt"

func main() {

	x := make([]int, 5, 10)
	x = append(x, 1, 2, 3, 4, 5)
	fmt.Println(x)
}
