package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// if you have multiple consective args as same type you can do this
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1,2)
	fmt.Println("1+2=", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3=", res)
}
