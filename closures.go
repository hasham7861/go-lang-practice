// practicing closures, to avoid global scope polution but still reusing state for memoization and what not
package main

import "fmt"

// intSeq returns another annoymus function
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// only keeps state of a reference of a function
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
