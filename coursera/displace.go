package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// read input as text and convert to a real number
func getName() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	f, err := strconv.ParseFloat(words[0], 64)
	if err != nil {
		fmt.Println(f, "  ", err)
	}
	return f
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + (vo * t) + so

	}
	return fn
}

func main() {
	fmt.Print("Enter a value for acceleration ==> ")
	a := getName()
	fmt.Print("Enter a value for initial velocity ==> ")
	vo := getName()
	fmt.Print("Enter a value for initial displacement ==> ")
	so := getName()
	fmt.Print("Enter a value for time ==> ")
	t := getName()
	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(" The resulting displacement travelled after ", t, " seconds is ", fn(t))

}
