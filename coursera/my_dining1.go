package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name  string
	left  int
	right int
}

func (p *Philosopher) Eat(table []sync.Mutex) {
	table[p.left].Lock()
	table[p.right].Lock()
	defer table[p.left].Unlock()
	defer table[p.right].Unlock()

	fmt.Println(p.name, "is eating.")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(p.name, "is done eating.")
}

func host(p *Philosopher) {
	var wg sync.WaitGroup
	table := make([]sync.Mutex, 5)

	//fmt.Println(" to function = ", *p)
	//fmt.Println("leftCS = ",p.left)
	wg.Add(1)
	go p.Eat(table)
	wg.Done()

}
func main() {
	//var wg sync.WaitGroup
	philosophers := []*Philosopher{
		&Philosopher{"1", 0, 1},
		&Philosopher{"2", 1, 2},
		&Philosopher{"3", 2, 3},
		&Philosopher{"4", 3, 4},
		&Philosopher{"5", 4, 0},
	}
	for i := 0; i < 3; i++ {
		for _, philosopher := range philosophers {
			//fmt.Println(" value = ",*philosopher)

			go host(philosopher)
			time.Sleep(200 * time.Millisecond)
		}
	}

}
