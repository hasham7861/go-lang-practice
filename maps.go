package main

import (
	"fmt"
	"maps"
)

func main() {
	// intializing of slices or map you gotta use make
	// to set a map key to int mapping you use this syntax
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	delete(m, "k2")
	fmt.Println("map:", m)

	// removes all the entries
	clear(m)
	fmt.Println("map:", m)

	// second value returned from entry is that if value is present or not
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	// map also equals of equality check on other map
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
