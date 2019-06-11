package main

/* concur calls two functions hello and world twice. The program takes an initial global value and sums it 4 times then decrements 4 times
The first time without concurrency, the second with
The same example without the  concurrent execution prints out
hello  1 2 3 4  3 2 1 0
 when they are called with concurrency there is no way to to determine
the correct order and it often prints differently so you see different results on the call
There is no way of predicting the order of function call or code execution within the function
*/
import (
	"fmt"
	"time"
)

var cnt int

func hello1() {
	for i := 1; i < 5; i++ {
		cnt = cnt + 1
		fmt.Print(" ", cnt)
		time.Sleep(300 * time.Millisecond)
	}
}
func world1() {
	for i := 1; i < 5; i++ {
		cnt = cnt - 1
		fmt.Print(" ", cnt)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println()
}

func hello2() {
	for i := 1; i < 5; i++ {
		cnt = cnt + 1
		fmt.Print(" ", cnt)
		time.Sleep(300 * time.Millisecond)
	}
}
func world2() {
	for i := 1; i < 5; i++ {
		cnt = cnt - 1
		fmt.Print(" ", cnt)
		time.Sleep(300 * time.Millisecond)
	}
}
func main() {
	cnt = 0
	hello1()
	world1()
	//fmt.Println("\n Beginning Concurrent calls\n")
	cnt = 0
	go hello2()
	go world2()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\n main terminated")
}
