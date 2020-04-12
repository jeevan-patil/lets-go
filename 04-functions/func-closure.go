package main

import "fmt"

// this function returns a function
func intSeq() func() int {
	i := 0

	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	increment := intSeq()
	fmt.Println(increment())
}
