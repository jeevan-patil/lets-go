package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["One"] = 1
	m["Two"] = 2
	m["Three"] = 3
	m["Four"] = 4

	fmt.Println("Map values:", m)
	fmt.Println("One is:", m["One"])

	delete(m, "Four")
	fmt.Println("After removing Four:", m)

	_, present := m["One"]
	fmt.Println("One present in map?", present)
}
