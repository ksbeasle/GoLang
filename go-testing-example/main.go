package main

import "fmt"

func AddTwoNums(x int, y int) int {
	res := x + y
	return res
}
func main() {
	a := 1
	b := 10

	fmt.Println(AddTwoNums(a, b))
}
