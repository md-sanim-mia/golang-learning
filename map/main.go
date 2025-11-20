package main

import "fmt"

func main() {
	studentGrate := make(map[string]int)

	studentGrate["prince"] = 90

	studentGrate["sanim"] = 95
	studentGrate["rohi"] = 100

	// fmt.Println("all studnent numer :", studentGrate["sanim"])

	for index, value := range studentGrate {
		fmt.Println(index, value)
	}
}
