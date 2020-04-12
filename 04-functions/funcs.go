package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add3Nums(a, b, c int) int {
	return a + b + c
}

func sub(a int, b int) int {
	return a - b
}

func addAll(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}

func main() {
	fmt.Println("1 + 2 =", (add(1, 2)))
	fmt.Println("1 + 2 + 3 =", (add3Nums(1, 2, 3)))
	fmt.Println("2 - 3 =", (sub(2, 3)))
	fmt.Println("2 + 3 + 1 + 4 - 5 =", addAll(2, 3, 1, 4, -5))
}
