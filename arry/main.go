package main

import "fmt"

func main() {

	var name [5]string

	name[0] = "Sanim"
	name[1] = "Rohi"

	fmt.Println(name)
	var number = [5]int{3, 4, 5, 67, 7}

	number[2] = 300

	fmt.Println(len(number))

}
