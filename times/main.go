package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("current time :", currentTime)

	fmt.Printf("current time type %T\n", currentTime)

	formation := currentTime.Format("02-01-2006")

	fmt.Println(formation)
}
