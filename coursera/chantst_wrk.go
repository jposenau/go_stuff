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
	d := new(philosopher)
	d.n = n
	// first create n chopstick channels and philo channels as communication between chopsticks and philosophers
	d.free = make([]chan bool, n)
	d.done = make([]chan bool, n)

	d.thinkTime = thinkTime
	d.eatTime = eatTime

	for i := 0; i < n; i++ {
		d.free[i] = make(chan bool)
		d.done[i] = make(chan bool)
	}
	return d
}

func (d *philosopher) start() {
	var wg sync.WaitGroup
	// create max number of sessions to control exit 3/phiolospher
	wg.Add(15)
	// run the host sessions
	for i := 0; i < d.n; i++ {
		go d.chopsticks(i)
		go d.host(i, &wg)
	}
	wg.Wait()
	fmt.Println(" \n\nCompleted all Eaters")
	os.Exit(3)
}

func (d *philosopher) host(p int, wg *sync.WaitGroup) {
	var lmt int
	const MaxIter = 3
	//lmt controls the number of times a phiosopher can eat currently 3
	for {
		// think time is used to delay to let independent processes run
		//fmt.Printf("%d,think\n", p)
		time.Sleep(d.thinkTime)

		left := false
		right := false
		// check if we can get the left chopstick
		select {
		case left = <-d.free[p]:
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
		case right = <-d.free[((p + 1) % d.n)]:
			break
		default:
			right = false
			break
		}

		if !right {
			// then free left as well
			d.done[p] <- true
		}

		if left && right {
			if lmt < MaxIter {
				lmt++
				// eat and then release allow a little time to eat
				fmt.Printf("starting to Eat %d\n", p+1)
				time.Sleep(d.eatTime)
				fmt.Println(" finished eating ", p+1)
				// indicate done
				d.done[p] <- true
				d.done[((p + 1) % d.n)] <- true

				wg.Done()
			} else {
				d.done[p] <- true
				d.done[((p + 1) % d.n)] <- true
			}
		}

	}

}

func (d *philosopher) chopsticks(f int) {

	for {

		// indicate that the chopstick is free
		d.free[f] <- true

		//fmt.Printf("Waiting for chopstick %d to be done\n",f)
		//wait for it to be used and then released
		select {
		case <-d.done[f]:
			break
		}
	}

}

func main() {

	d := newPhil(5, 1*time.Second, 1*time.Second)
	d.start()

}
