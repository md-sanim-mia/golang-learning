package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing taks id", id)
}

func main() {

	for i := 0; i <= 10; i++ {
		go task(i)
	}

	time.Sleep(time.Second * 2)
}
