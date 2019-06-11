package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type philosopher struct {
	thinkTime time.Duration
	eatTime   time.Duration
	free      []chan bool
	done      []chan bool
	n         int
}

func newPhil(n int, eatTime time.Duration, thinkTime time.Duration) *philosopher {
	// add in some eating and talking dealys to make the simulation more realistic
	diner := new(philosopher)
	diner.n = n
	// first create n chopstick channels and philo channels as communication between chopsticks and philosophers
	diner.free = make([]chan bool, n)
	diner.done = make([]chan bool, n)

	diner.thinkTime = thinkTime
	diner.eatTime = eatTime

	for i := 0; i < n; i++ {
		diner.free[i] = make(chan bool)
		diner.done[i] = make(chan bool)
	}
	return diner
}

func (diner *philosopher) start() {
	var wg sync.WaitGroup
	// create max number of sessions to control exit 3/phiolospher
	wg.Add(15)
	// run the host sessions
	for i := 0; i < diner.n; i++ {
		go diner.chopsticks(i)
		go diner.host(i, &wg)
	}
	wg.Wait()
	fmt.Println(" \n\nCompleted all Eaters")
	os.Exit(3)
}

func (diner *philosopher) host(p int, wg *sync.WaitGroup) {
	var lmt int
	const MaxIter = 3
	//lmt controls the number of times a phiosopher can eat currently 3
	for {
		// think time is used to delay to let independent processes run
		//fmt.Printf("%d,think\n", p)
		time.Sleep(diner.thinkTime)

		left := false
		right := false
		// check if we can get the left chopstick
		select {
		case left = <-diner.free[p]:
			break
		default:
			left = false
			break
		}

		if !left {
			continue
		}

		//try and get right one
		select {
		case right = <-diner.free[((p + 1) % diner.n)]:
			break
		default:
			right = false
			break
		}

		if !right {
			// then free left as well
			diner.done[p] <- true
		}

		if left && right {
			if lmt < MaxIter {
				lmt++
				// eat and then release allow a little time to eat
				fmt.Printf("starting to Eat %d\n", p+1)
				time.Sleep(diner.eatTime)
				fmt.Println(" finished eating ", p+1)
				// indicate done
				diner.done[p] <- true
				diner.done[((p + 1) % diner.n)] <- true

				wg.Done()
			} else {
				diner.done[p] <- true
				diner.done[((p + 1) % diner.n)] <- true
			}
		}

	}

}

func (diner *philosopher) chopsticks(f int) {

	for {

		// indicate that the chopstick is free
		diner.free[f] <- true

		//fmt.Printf("Waiting for chopstick %d to be done\n",f)
		//wait for it to be used and then released
		select {
		case <-diner.done[f]:
			break
		}
	}

}

func main() {

	diner := newPhil(5, 1*time.Second, 1*time.Second)
	diner.start()

}
