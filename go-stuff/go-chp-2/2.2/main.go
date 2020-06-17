package main

import (
	"fmt"
	"os"
	"strconv"

	"go-stuff/go-chp-2/2.2/unitconv"
)

/*Exercise 2.2 -- Based on the cf program write a unit conversion program that reads in input from the command line
  or standard input if there are no arguements and convert each number into feet to meters, and lbs to kilograms
*/
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Enter some args please...\n")
	}
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
		}
		m := unitconv.FeetToMeters(unitconv.Ft(t))
		k := unitconv.PoundsToKilograms(unitconv.Pounds(t))

		fmt.Printf("Feet: %v -> Meters: %v\n", t, m)
		fmt.Printf("Pounds: %v -> Kilograms: %v\n", t, k)
	}
}
