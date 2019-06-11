package main

import (
	"fmt"
)

func cntrl(v1 int, v2 int, c chan int) {
	c <- v1 + v2
}

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go cntrl(1, 2, c)
		go cntrl(4, 5, c)
		a := <-c
		b := <-c
		fmt.Println(" cntrl1 ", a)
		fmt.Println(" cntrl2 ", b)
	}
	fmt.Println("\n main terminated")
}
