package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {

		fmt.Println("number is :", i)
	}

	numbers := []int{10, 30, 50, 60, 70, 80, 90}

	for _, sum := range numbers {

		fmt.Print("numer list :", sum)
	}

}
