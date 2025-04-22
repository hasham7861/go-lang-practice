package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	value T
}

func (l *List[T]) Push(value T) {
	newElement := &element[T]{value: value}
	if l.head == nil {
		l.head = newElement
	} else {
		l.tail.next = newElement
	}
	l.tail = newElement
}

// use iterator to traverse the a data structure (linkedlist) next value
func (l *List[T]) All () iter.Seq[T] {
	// importance is this inner function which takes in a yield function to keep going or not
	return func (yield func (T) bool) {
		for e := l.head; e != nil; e = e.next {
			if !yield(e.value) {
				return
			}
		}
	}
}


// create an iterator on top of sequenece like fibonacci, doesnt even have to be a list or datastructure even
func genFib() iter.Seq[int] {
	return func(yield func (int) bool) {
		a, b := 1, 1

		// This will keep going forever for fibonacci
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
} 

func main () {

	lst := &List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e:= range lst.All() {
		fmt.Println(e)
	}

	// you can create a slice from any iterator
	all := slices.Collect(lst.All())
	fmt.Println(all)

	for n := range genFib() {
		if n > 100 {
			break
		}
		fmt.Println(n)
	}
}