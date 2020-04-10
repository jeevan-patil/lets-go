package main

import (
	"fmt"
	"time"
)

func main() {
	num := 2
	switch num {
	case 1:
		fmt.Println(num, " written as one")
	case 2:
		fmt.Println(num, "written as two")
	case 3:
		fmt.Println(num, "written as three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	findType := func(par interface{}) {
		switch par.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("god knows")
		}
	}

	findType(true)
	findType("G1")
	findType(12.22)
}
