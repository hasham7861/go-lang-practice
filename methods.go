package main

import "fmt"

type rect struct {
	width, height int
}

// the (r *react) is adding a method onto rect struct, in away creating methods on a class
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height:5}	
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
	
	// create new variable pointing to same memory address of r
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())


}