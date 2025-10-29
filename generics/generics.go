package main

import "fmt"

type stack[T any] struct {
	element []T
}

func main() {

	myStack := stack[string]{
		element: []string{"golang"},
	}

	fmt.Println(myStack)
}
