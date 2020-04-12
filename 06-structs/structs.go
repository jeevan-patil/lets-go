package main

import "fmt"

type person struct {
	name string
	age  int
}

func createNew(name string) *person {
	pers := person{name, 42}
	return &pers
}

func main() {
	fmt.Println(person{"Rahul", 39})
	fmt.Println(person{age: 41, name: "Sachin"})
	fmt.Println(person{name: "Saurav"})
	fmt.Println(&person{name: "Irfan", age: 40})
	fmt.Println(createNew("Zaheer"))

	laxman := person{name: "VVS", age: 38}
	fmt.Println(laxman.name, laxman.age)

	vvs := &laxman
	fmt.Println(vvs.name)

	// structs are mutable
	vvs.name = "VVS Laxman"
	fmt.Println(vvs)
	fmt.Println(laxman)
}
