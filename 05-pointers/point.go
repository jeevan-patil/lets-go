package main

import "fmt"

func ZeroValue(val int) {
	val = 0
}

func ZeroPointer(val *int) {
	*val = 0
}

func main() {
	num := 1
	ZeroValue(num)
	fmt.Println("By Value", num)

	ZeroPointer(&num)
	fmt.Println("By reference", num)

	fmt.Println("Num address", &num)
}
