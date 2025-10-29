package main

import (
	"fmt"
	"time"
)

type order struct {
	id string

	amount int

	status string

	createAtd time.Time
}

func main() {
	order := order{
		id:     "1",
		amount: 543,
		status: "pending",
	}

	order.createAtd = time.Now()

	fmt.Println(order)
}
