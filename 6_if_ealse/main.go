package main

import "fmt"

func main() {

	age := 90
	if age >= 18 && age < 60 {

		fmt.Println("your the adulth person ")
	} else if age >= 60 {

		fmt.Println("your ar the old person")
	} else {
		fmt.Println("your not adulth person")
	}
}
