package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 int = 2
	var x2 int = 5
	var output int64 = int64(math.Max(float64(x1), float64(x2)))
	fmt.Println(output)

	var x3 uint8 = 127
	fmt.Println("printing", x3)

	fmt.Println("Test lateny.......")
}
