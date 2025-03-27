package main

import "fmt"

// returning multiple values you need to put in brackets
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
