package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// go infers type by default
	var d = true
	fmt.Println(d)

	// default value for int is 0
	var e int
	fmt.Println(e)

	// short hand syntax for declaring and intializing
	f := "apple"
	fmt.Println(f)
}
