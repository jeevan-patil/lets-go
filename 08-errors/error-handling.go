package main

import (
	"errors"
	"fmt"
)

type customError struct {
	arg int
	msg string
}

func (ce *customError) Error() string {
	return fmt.Sprintf("%d - %s", ce.arg, ce.msg)
}

func func1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

func func2(arg int) (int, error) {
	if arg == 42 {
		return -1, &customError{arg, "Can't work with this input"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{4, 42} {
		if r, e := func1(i); e != nil {
			fmt.Println("func1 failed :", e)
		} else {
			fmt.Println("func1 worked: ", r)
		}
	}

	for _, i := range []int{4, 42} {
		if r, e := func2(i); e != nil {
			fmt.Println("func2 failed :", e)
		} else {
			fmt.Println("func2 worked: ", r)
		}
	}
}
