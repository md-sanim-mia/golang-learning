package main

import (
	"bufio"
	"os"
)

func main() {

	// f, err := os.Create("exampole2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()
	// fileInfo, err := f.Stat()

	// if err != nil {

	// 	panic(err)
	// }

	var scourceFiles *os.File

	scourceFiles, err := os.Open("exampole.txt")

	if err != nil {
		panic(err)
	}

	defer scourceFiles.Close()

	destFile, err := os.Create("exmpole.txt")

	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reander := bufio.NewReader(scourceFiles)

	writer := bufio.NewWriter(destFile)

	for {
		b, err := reander.ReadByte()

		if err != nil {

			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		writer.WriteByte(b)

	}

}
