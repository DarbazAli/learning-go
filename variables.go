package main

import "fmt"

// package-level variable declaration
var username string = "John Doe"

func main() {
	dogsName := "Max"
	dogsAge := 1
	dogsHight := 0.33

	// constant identifires
	const greet string = "Hello John"
	//greet = "Hello Jane"	// cannot assign to greet (declared const)

	fmt.Printf("%v, %T \n", dogsName, dogsName)
	fmt.Printf("%v, %T \n", dogsAge, dogsAge)
	fmt.Printf("%v, %T \n", dogsHight, dogsHight)
}

func getUsername() {
	// this func has access to the variable "username"
	fmt.Println(username)
}
