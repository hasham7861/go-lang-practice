package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		// error.New creats basic new error value with given error message
		return -1, errors.New("can't work with 42")
	}
	
	return arg + 3, nil
}

// a way to declare error variables
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 0 {
		return ErrOutOfTea
	} else if arg == 4 {
		// You can wrap an error with another error
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {

	for _, i := range []int{7, 42}{
		// common to use inline error checks with if statement
		if r, e := f (i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// you can compare error type
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("We should fix the kettle!")
			} else {
				fmt.Println("unknown error type:", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}

