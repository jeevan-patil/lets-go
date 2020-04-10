package main

import "fmt"

func main() {
	// variables without initial values will have zero value.
	// zero value for a string is empty string, and 0 for an int
	var fruit string

	fmt.Println(fruit)

	fruit = "apple"
	fmt.Println(fruit)

	// this is shorthand for declaring and initializing a variable
	months := 12
	fmt.Println(months)
}
