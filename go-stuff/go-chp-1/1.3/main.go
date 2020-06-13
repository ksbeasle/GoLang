package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//Exercise 1.3 -- Measure difference in running time between inefficient vs efficient
	begin := time.Now()
	fmt.Println("----------Inefficient Start----------")
	str, sep := "", " "
	for _, arg := range os.Args[1:] {
		str += sep + arg
	}
	fmt.Println(str)
	fmt.Println(time.Since(begin))
	fmt.Println("----------Inefficient End----------")

	fmt.Println("----------Efficient Start----------")
	begin = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(begin))
	fmt.Println("----------Efficient End----------")
}
