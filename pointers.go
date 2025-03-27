package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// the value of i shouldn't change
	zeroval(i)
	fmt.Println("zeroval:", i)

	// the value of i changes as you are passing reference of i rather than value
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// &i is syntax for passing the reference aka memory address
	fmt.Println("pointer:", &i)
}
