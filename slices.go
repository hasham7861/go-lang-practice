package main

import (
	"fmt"
	"slices"
	// "slices"
)

func main() {

	var s []string

	fmt.Println("unint:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// for slices you can create a new instance to it
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// you can create new slice same slice as another and copy items to it
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slices also support "slicing"
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	// You can check equality of arrays like so
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can have inner slices and sizes of inner slices can vary. Wow
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	x := make([]string, 10)
	x = append(x, "2323232")
	fmt.Println(x)

	var a1 [5]int

}
