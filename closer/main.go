package main

import "fmt"

func counter() func() int {

	count := 0

	return func() int {

		count += 1

		return count

	}

}

func main() {

	incement := counter()

	fmt.Println(incement())

	fmt.Println(incement())
}
