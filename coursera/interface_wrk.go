// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Here's a basic interface for geometric shapes.
type Animal interface {
	Eat()
	Move()
	Speak()
}
type cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type combined struct {
	c1 cow
	s1 snake
	b1 bird
}

// Functions to implment the interface

func (c cow) Eat() {
	fmt.Println(c.food)
}

func (c cow) Speak() {
	fmt.Println(c.noise)
}

func (c cow) Move() {
	fmt.Println(c.locomotion)
}
func (c snake) Eat() {
	fmt.Println(c.food)
}

func (c snake) Speak() {
	fmt.Println(c.noise)
}

func (c snake) Move() {
	fmt.Println(c.locomotion)
}
func (c bird) Eat() {
	fmt.Println(c.food)
}

func (c bird) Speak() {
	fmt.Println(c.noise)
}

func (c bird) Move() {
	fmt.Println(c.locomotion)
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `getInfo` function taking advantage of this
// to work on any `geometry`.
func getInfo(g Animal) {

	g.Eat()
	g.Speak()
	g.Move()
}

// Build the 2 requests by reading a string converitng to lower case for consistency and returing the values

func returnList() []string {
	// input a string of integers as a string and convert to individual elements
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type your request > ")
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	if len(words) != 3 {
		fmt.Println("invalid input string ")
		os.Exit(3)
	}
	for i := 0; i < len(words); i++ {
		tmp := strings.ToLower(words[i])
		words[i] = tmp
	}

	return words
}

// take the animal requested animal activity and call the appropriate function
func interpret(p Animal, n string, val string) {

	switch val {
	case "eat":
		p.Eat()
	case "move":
		p.Move()
	case "speak":
		p.Speak()
	default:
		fmt.Println("Unknown activity choice")
	}
}

func main() {
	cw := cow{name: "", food: "grass", locomotion: "walk", noise: "moo"}
	bw := bird{food: "worms", locomotion: "fly", noise: "peep"}
	sw := snake{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		words := returnList()
		fmt.Println(words[0])
		a := words[1]
		b := words[2]
		//fmt.Println(bw, sw)
		switch words[0] {
		case "query":
			{
				if a == cw.name {
					interpret(cw, a, b)
				} else if a == bw.name {
					interpret(bw, a, b)
				} else if a == sw.name {
					interpret(sw, a, b)
				}
			}
		case "newanimal":
			{
				switch b {
				case "cow":
					{
						cw.name = a
						fmt.Println(cw)
					}
				case "bird":
					{
						bw.name = a
						fmt.Println(bw)
					}
				case "snake":
					{
						sw.name = a
						fmt.Println(sw)
					}
				default:
					{
						fmt.Println("bad input")
						break
					}
					fmt.Println("Created It!!")
				}
			}
		default:
			{
				fmt.Println("missing values")
			}
		}

	}
}
