package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS int
}

func (p Philo) eat(phil_num int, c chan int) {
	for i := 0; i < 5; i++ {
		p.leftCS.Lock()

		p.rightCS.Lock()
		fmt.Println("Phil", phil_num, "Eating")
		//time.Sleep(time.Microsecond * 300)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Println("Finished Eating Phil ", phil_num)
	}

func host(d chan int) {
	//var f int
	//var mu sync.Mutex
	c := make(chan int)
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	//for i := 0; i < 5; i++ {
	for {
		//fmt.Println("Calling Eat ", i+1)
		go philos[0].eat(1, c)
		go philos[1].eat(2, c)
		go philos[2].eat(3, c)
		go philos[3].eat(4, c)
		go philos[4].eat(5, c)

	}

}
func main() {
	d := make(chan int)
	go host(d)
	f := <-d
	fmt.Println("Finished", f)

}
