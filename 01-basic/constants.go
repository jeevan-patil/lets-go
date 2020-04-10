package main

import (
	"fmt"
	"math"
)

const language string = "go"

func main() {
	fmt.Println(language)

	// constant without data type
	const num = 100
	fmt.Println(num)

	// num will be inferred as float64 based on the context
	fmt.Println(math.Sin(num))
}
