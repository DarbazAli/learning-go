package main

import "fmt"

// package-level variable declaration
var username string = "John Doe"

// multiple variables
const (
	firstname string = "John"
	lastname  string = "Doe"
	age       int    = 28
)

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

	// accessing multiple variables form
	fmt.Printf("%v, %T \n", firstname, firstname)
	fmt.Printf("%v, %T \n", lastname, lastname)
	fmt.Printf("%v, %T \n", age, age)
}

func getUsername() {
	// this func has access to the variable "username"
	fmt.Println(username)
}
