package main

import "fmt"

func main() {
	/*

			A race condition is when two or more dependent blocks of code run in a non-deterministic fashion
			producing a result that is impossible to predict.  This can happen if dependent threads are not
		    waiting properly.

		    For example, if two functions are using the same data, but are not being executed in a specific order
	*/

	var i int
	go func() {
		i = 5
	}()

	go func() {
		i = 6
	}()
	fmt.Println(i)
}
