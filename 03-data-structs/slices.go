package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "Hello"
	s[1] = "world"
	s[2] = "!!"
	fmt.Println(s)

	s = append(s, "let's")
	s = append(s, "learn", "go")
	fmt.Println(s)

	fmt.Println("from position 2 to 4", (s[2:4]))
	fmt.Print("from position 3 to end", (s[3:]))
}
