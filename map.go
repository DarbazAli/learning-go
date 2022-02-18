package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["One"] = 1
	x["Tow"] = 2
	x["Three"] = 3

	fmt.Println(x, len(x))
	fmt.Println(x["Three"])

	delete(x, "One")

	fmt.Println(x)
}
