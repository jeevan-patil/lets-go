package main

import "fmt"

type rectangle struct {
	width, height int
}

func (rect *rectangle) area() int {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() int {
	return 2*rect.width + 2*rect.height
}

func main() {
	rect := rectangle{width: 10, height: 4}
	fmt.Println("area", rect.area())
	fmt.Println("perimeter", rect.perimeter())
}
