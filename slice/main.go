package main

import "fmt"

func main() {

	// number := []int{10, 30, 70, 40, 43, 43, 43, 4, 5, 5}

	// result := append(number, 9000, 458593, 4435, 5454, 636246)

	// fmt.Println(result)

	numbers := make([]int, 3, 5)

	fmt.Println("slice:", numbers)

	numbers = append(numbers, 4)

	numbers = append(numbers, 8)

	numbers = append(numbers, 90)

	numbers = append(numbers, 70)

	numbers = append(numbers, 76)
	numbers = append(numbers, 90)
	numbers = append(numbers, 90)

	stock := make([]int, 0)

	fmt.Println("lenght:", len(stock))
	fmt.Println("capacity:", cap(stock))
}
