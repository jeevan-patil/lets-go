package main

import "fmt"

func add(inputs ...int) int {
	res := 0

	for _, num := range inputs {
		res = res + num
	}

	return res
}

func main() {
	inputs := [](int){1, 2, 3, 4}
	fmt.Println(add(inputs...))
}
