package main

import (
	"fmt"
	"os"
)

func main() {
	//Exercise 1.2 -- Modify the echo program to print the index and value of os.Args
	for index, value := range os.Args[1:] {
		fmt.Printf("index: %d\tvalue: %s\n", index, value)
	}
}
