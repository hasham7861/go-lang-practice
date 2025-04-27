package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i:=0; i<3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// executes function in same thread by default
	f("direct")

	// simply adding go keyword before function call will create lighweight thread to execute function
	go f("goroutine")

	// you can add go keyword before annoymous closure function as well
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// a naive way to wait for async thread operation to finish is to wait some seconds or mins using sleep 
	time.Sleep(time.Second) // should use wait groups to wait for routine to finish (will learn in near future)
	fmt.Println("done")
}