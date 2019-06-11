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
	time.Sleep(time.Second)
	fmt.Println(p.name, "is done eating.")
}

func main() {
	philosophers := []*Philosopher{
		&Philosopher{"Baruch Spinoza", 0, 1},
		&Philosopher{"Gilles Deleuze", 1, 2},
		&Philosopher{"Karl Marx", 2, 3},
		&Philosopher{"Friedrich Nietzsche", 3, 4},
		&Philosopher{"Michel Foucault", 4, 0},
	}

	table := make([]sync.Mutex, len(philosophers))

	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go func(p *Philosopher) {
			p.Eat(table)
			wg.Done()
		}(philosopher)
	}
	wg.Wait()
}
