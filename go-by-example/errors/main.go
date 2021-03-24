package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if result, error := f1(i); error != nil {
			fmt.Println("f1 failed: ", error)
		} else {
			fmt.Println("f1 worked: ", result)
		}
	}

	fmt.Println("----------")

	for _, i := range []int{7, 42} {
		if result, error := f2(i); error != nil {
			fmt.Println("f2 failed: ", error)
		} else {
			fmt.Println("f2 worked: ", result)
		}
	}

	fmt.Println("----------")

	_, error := f2(42)

	if ae, ok := error.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	// More about error handling:
	// https://blog.golang.org/error-handling-and-go
}
