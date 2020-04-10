package main

import "fmt"

// there is no ternary if in go
func main() {
	if 8%2 == 0 {
		fmt.Println("Even number")
	}

	// declaration and usage. num is available in all branches
	if num := 5; num > 8 {
		fmt.Println(num, "is greater than 8")
	} else if num < 3 {
		fmt.Println(num, "greater than 3")
	} else {
		fmt.Println("somewhere in-between")
	}
}
