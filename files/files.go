package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("exampole.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()
	fileInfo, err := f.Stat()

	if err != nil {

		panic(err)
	}

	fmt.Println("files", fileInfo.Name())
}
