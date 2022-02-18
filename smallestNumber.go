/*
	Find smallest number in a list

	@author: Darbaz Ali
	@date: Feb 18, 2022

*/
package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 0,
	}

	var smallest = x[0]

	for _, item := range x {
		if item <= smallest {
			smallest = item
		}
	}

	fmt.Println(smallest)
}
