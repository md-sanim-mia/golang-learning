package main

import "fmt"

func main() {

	number := []int{10, 30, 70, 40, 43, 43, 43, 4, 5, 5}

	result := append(number, 9000, 458593, 4435, 5454, 636246)

	fmt.Println(result)
}
