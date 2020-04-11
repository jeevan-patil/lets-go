package main

import "fmt"

func main() {
	nums := []int{11, 12, 13}

	for in, val := range nums {
		fmt.Println(in, val)
	}

	fruits := map[string]string{"a": "Apple", "b": "Banana"}

	for key, val := range fruits {
		fmt.Println(key, "for", val)
	}

	for i, c := range "jeevan" {
		fmt.Println(i, string(c))
	}
}
