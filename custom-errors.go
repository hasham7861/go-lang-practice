package main

import (
	"errors"
	"fmt"
)

// custom error struct
type argError struct {
	arg int
	message string
}
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)

	var ae *argError
	// .As method checks if its custom advanced version of error with struct reference
	if errors.As(err, &ae) {
		fmt.Println("argError:", ae.arg, ae.message)
	} else {
		fmt.Println("error:", err)
	}


}