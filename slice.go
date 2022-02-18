package main

import "fmt"

func main() {

	x := make([]int, 5, 10)
	x = append(x, 1, 2, 3, 4, 5)
	fmt.Println(x)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
