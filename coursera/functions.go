/*

Let us assume the following formula for displacement s as a
function of time t, acceleration a, initial velocity vo, and
initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for
acceleration, initial velocity, and initial displacement.

 Then the program should prompt the user to enter a value for time
 and the program should compute the displacement after the
 entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a,
initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement
 as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
  The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64
  argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for
acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a
function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fAccel float64
var fVel float64
var fDisp float64
var fTime float64
var results []string
var scanner *bufio.Scanner

func GenDisplaceFn(a, vot, so float64) func(float64) float64 {

	//t2 := 10.0
	//	s =½ a t2 + vot + so
	//s := ((a * t2) * .5) + vot + so
	fn := func(t2 float64) float64 {
		return (((a * t2) * .5) + vot + so)
	}
	return (fn)
}

func main() {
	fmt.Printf("Inital Acceleration: ")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		results = strings.Fields(scanner.Text())
		fAccel, _ = strconv.ParseFloat(results[0], 32)

	}
	fmt.Printf("Inital Displacement: ")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		results = strings.Fields(scanner.Text())
		fDisp, _ = strconv.ParseFloat(results[0], 32)

	}

	fmt.Printf("Inital Velocity: ")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		results = strings.Fields(scanner.Text())
		fVel, _ = strconv.ParseFloat(results[0], 32)

	}
	fmt.Printf("Time: ")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		results = strings.Fields(scanner.Text())
		fTime, _ = strconv.ParseFloat(results[0], 32)

	}

	fn := GenDisplaceFn(fAccel, fVel, fDisp)

	fmt.Printf("Displacement after %f seconds: %f\n", fTime, fn(fTime))

}
