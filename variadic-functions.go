package main

import "fmt"

// to deconstruct args with endless number of ints
func sum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		// fmt.Println( num)
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}

	// to deconstract values from array
	sum(nums...)
}
