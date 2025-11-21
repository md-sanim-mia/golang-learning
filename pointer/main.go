package main

import "fmt"

func modifyByreferance(a *int) {

	*a = *a + 200
}

func main() {
	var nums int

	var prt *int

	prt = &nums

	modifyByreferance(&nums)
	fmt.Println("Nums has value", nums)

	fmt.Println(*prt)

}
