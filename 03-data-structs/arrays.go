package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println("arr:", a)

	a[0] = 1
	fmt.Println(a[0])

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	fmt.Println("length of b array is:", len(b))
}
