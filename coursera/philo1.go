package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(phil_num int, c chan int) {
	for i := 0; i < 25; i++ {
		fmt.Println("Phil", phil_num, "Eating")
		c <- phil_num + 1
		fmt.Println("Finsished Eating", phil_num)
	}

}

func main() {
	var f int
	c := make(chan int)
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}

	}
	for i := 0; i < 5; i++ {
		//go philos[i].eat()
		fmt.Println("Calling Eat ", i+1)

		go philos[i].eat(i+1, c)
		f = <-c
		fmt.Println(f)
	}
	fmt.Println(f)
}
