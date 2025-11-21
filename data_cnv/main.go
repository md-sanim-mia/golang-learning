package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42

	fmt.Println("Number is:", num)

	fmt.Printf("type of num is %T\n", num)

	var data float64 = float64(num)

	data = data + 1.34

	fmt.Println("data is :", data)

	fmt.Printf("data type is %T\n", data)

	num = 40

	str := strconv.Itoa(num)

	fmt.Println("str data:", str)

	fmt.Printf("str type is %T\n", str)

	string_to_number := "54325"

	string_to_numbers, _ := strconv.Atoi(string_to_number)

	fmt.Println("string to numer ", string_to_numbers)

	fmt.Printf("string to number type %T\n", string_to_numbers)

}
