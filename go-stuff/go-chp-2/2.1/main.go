package main

import (
	"fmt"

	"tempconv/tempconv"
)

//Exercise 2.1 -- Add types/constants and functions to tempconv for processing temperatures in Kelvin
func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Brrrr! %v\n", tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("F to K = %v\n", tempconv.FToK(tempconv.Fahrenheit(32)))
}
