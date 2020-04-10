package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for cnt := 0; cnt < 5; cnt++ {
		fmt.Println(cnt)
	}

	for {
		fmt.Println("infinite unless and until break or return")
		break
	}

	for j := 1; j < 5; j++ {
		if j%2 == 0 {
			continue
		}

		fmt.Println(j)
	}
}
