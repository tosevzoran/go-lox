package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tosevzoran/go-lox/golox"
)

func main() {
	args := os.Args[1:]

	var ioReader *bufio.Reader

	if len(args) == 0 {
		ioReader = bufio.NewReader(os.Stdin)
	}

	if len(args) == 1 {
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Printf("error opening file %s, %v\n", args[0], err)
			os.Exit(1)
		}
		defer file.Close()

		ioReader = bufio.NewReader(file)
	}

	if len(args) > 1 {
		fmt.Printf("unsupported number of arguments, should be 0 or 1\n")
		os.Exit(1)
	}

	lox := golox.New(ioReader)

	err := lox.Run()

	if err != nil {
		fmt.Printf("unexpected error %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
