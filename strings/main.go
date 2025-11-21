package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "apple,orange,banana"
	pars := strings.Split(data, ",")
	fmt.Println(pars)
	fmt.Println("data is :", data)
	fmt.Printf("data type is %T\n", data)

	str := "one,tow ,five, six, siven, siven, siven, siven"
	count := strings.Count(str, "siven")

	fmt.Println(count)

}
