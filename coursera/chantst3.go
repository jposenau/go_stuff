package main

import (
	"fmt"
	"sync"
)

var fork [5]int
var waiter sync.Mutex
var exit int

/*
 Method implements the waiter's tasks,
 Either to pick up the fork
 Or to drop them.
*/
func waiter_task(pick bool, pid int) bool {
	var i, j = pid, (pid + 1) % 5
	var has_value int = 1
	if pick {
		has_value = 0
	}
	if fork[i] == has_value && fork[j] == has_value {
		fork[i] = (has_value + 1) % 2
		fork[j] = (has_value + 1) % 2
		return true
	}
	return false
}

/*
 Method implements the actions to be
 Performed by the waiter.
*/
func Action(pick bool, pid int) {
	for true {
		waiter.Lock()
		if waiter_task(pick, pid) {
			waiter.Unlock()
			break
		}
		waiter.Unlock()
	}
}

/*
 Method implements the actions
 Of the philosopher.
*/
func Philosopher(id int) {
	Action(true, id)
	fmt.Println("Philosopher ", id, "pick up the forks, ate and dropped them.")
	Action(false, id)
	exit++
}

// Main
func main() {
	exit = 0
	fmt.Println("Dining Philosophers Problem - Arbitrator Solution")
	// Running five threads and supplying the IDs
	for j := 0; j < 3; j++ {
		for i := 0; i < 5; i++ {
			go Philosopher(i)
		}
	}
	for exit != 5 { /* wait for threads to get over */
	}
}
