package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	p := Person{"John Doe", 30}
	fmt.Println(p.name)
	p2 := Person{
		name: "John Denzel",
		age: 30,
	}
	fmt.Println(p2.name)

	// can do one type struct typing and then assign value to it
	p3 := struct {
		name string
		age  int
	} {
		name: "John Doe",
		age: 30,
	}
}
