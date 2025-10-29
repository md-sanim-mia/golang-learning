package main

import "fmt"

// func add(a int, b int) int {

// 	return a + b
// }

func sum(nums ...int) int {

	total := 0

	for _, num := range nums {

		total = total + num

	}
	return total
}

func main() {

	result := sum(30, 40, 60, 80, 90, 20)

	fmt.Println(result)
}
