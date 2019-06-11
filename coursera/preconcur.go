package main

import (
	"fmt"
	"time"
)

var cnt int

func hello() {
	for i := 1; i < 5; i++ {
		cnt = cnt + 1
		fmt.Print(" ", cnt)
	}
}
func world() {
	for i := 1; i < 5; i++ {
		cnt = cnt - 1
		fmt.Print(" ", cnt)
	}
}
func main() {
	cnt = 0
	hello()
	world()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\n main terminated")
}
