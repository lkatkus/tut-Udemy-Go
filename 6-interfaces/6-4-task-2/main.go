package main

import (
	"fmt"
	"io"
	"os"
)

type contentPrinter struct{}

func main() {
	programArgs := os.Args
	fileName := programArgs[1]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cp := contentPrinter{}

	io.Copy(cp, file)
}

func (contentPrinter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
