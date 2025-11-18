package main

import "fmt"

func divider(a, b int64) (int64, error) {

	if b == 0 {
		return 0, fmt.Errorf("dennominator must not be not zero")
	}

	return a / b, nil
}

func main() {
	result, err := divider(200, 0)

	if err != nil {

		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
