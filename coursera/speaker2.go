package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}
type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name)
}
func main() {
	var s1 Speaker
	d1 := Dog{"Brian"}
	var l *Dog
	fmt.Printf("%T\n", l)
	s1 = d1
	s1.Speak()
}
