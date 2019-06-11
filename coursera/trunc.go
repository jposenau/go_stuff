package main

import "fmt"

var fl float32 // get a float
var rslt int32 //return an integer
// Build a quick wat to test floating point trucation in a loop
// no error testing was performed
// initial loop set to 2 for exercise requirements
func main() {
	for i := 0; i < 2; i++ {
		fmt.Printf("Enter a floating point value\n")
		fmt.Scanf("%f \n", &fl)
		rslt = int32(fl)
		fmt.Printf("the truncated values is %d\n", rslt)
	}
}
