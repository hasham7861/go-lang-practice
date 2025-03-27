package main

import "fmt"

func main() {
	// you can use helper method range to iterate over multiple built in types
	//
	nums := []int{2, 3, 4}
	sum := 0

	// iterate over an array
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// iterate over a map
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// iterate over string
	for i, c := range "go" {
		fmt.Println("index:", i, "character:", c)
	}
}
