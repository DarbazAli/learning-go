package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)

	fmt.Println(slice1, slice2)

	// append
	sl1 := []int{1, 2, 3}
	sl2 := append(sl1, 4, 5)

	fmt.Println(sl1, sl2)

	// copy
	slc1 := []int{1, 2, 3}
	slc2 := make([]int, 3)

	copy(slc2, slc1)

	fmt.Println(slc1, slc2)
}
