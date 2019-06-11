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
	//defer table[p.left].Unlock()
	//defer table[p.right].Unlock()

	fmt.Println(" Starting to Eat. ", p.name)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Finished Eating. ", p.name)
	table[p.left].Unlock()
	table[p.right].Unlock()
}

func doit(p []*Philosopher) {
	var wg sync.WaitGroup
	table := make([]sync.Mutex, 5)

	for i := 0; i < 3; i++ {
		for _, philosopher := range p {
			fmt.Println(" philosopher ", philosopher.name)
			wg.Add(1)

			go philosopher.Eat(table)

			wg.Done()
		}
	}

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
	go doit(philosophers)
	time.Sleep(500 * time.Millisecond)
}
