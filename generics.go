package main

import (
	"fmt"
)

// S is a generic for slice type, and E is a generic type that is used to compare on top of slice ~[]



func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}


type List[T any] struct {
	head, tail *element[T]
}

func (l *List[T]) Print() {
	for e := l.head; e != nil; e = e.next {
		fmt.Print(e.value)
		fmt.Print("->")
	}
	fmt.Println("null")
}

func (l *List[T]) Append(v T) {
	if l.head == nil {
		l.head = &element[T]{value: v}
		l.tail = l.head
	} else if l.tail != nil {
		l.tail.next = &element[T]{value: v}
		l.tail = l.tail.next
	} else {
		l.tail = &element[T]{value: v}
		l.head = l.tail
	}
}

type element[T any] struct {
	next *element[T]
	value T
}


func main () {
	list := List[int] {}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	
	list.Print()
}

