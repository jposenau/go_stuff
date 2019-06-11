// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
func (c cow) Inter() string {
	return c.name
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

// Build the requests by reading a string converitng to lower case for consistency and returing the values

func returnList() []string {
	// input a string of integers as a string and convert to individual elements
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type your request > ")
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	if len(words) != 3 {
		fmt.Println("invalid input string -- Leaving the Program!")
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
func getCow(a string) cow {
	var ll cow
	ll.name = a
	ll.food = "grass"
	ll.locomotion = "walk"
	ll.noise = "moo"
	return ll
}
func getBird(a string) bird {
	var ll bird
	ll.name = a
	ll.food = "worms"
	ll.locomotion = "fly"
	ll.noise = "peep"
	return ll
}
func getSnake(a string) snake {
	var ll snake
	ll.name = a
	ll.food = "mice"
	ll.locomotion = "slither"
	ll.noise = "hsss"
	return ll
}
func main() {

	p1 := cow{}
	p2 := bird{}
	p3 := snake{}
	for {
		words := returnList()

		aname := words[1]
		b := words[2]

		switch words[0] {
		case "query":
			{
				switch aname {
				case p1.name:
					{
						interpret(p1, aname, b)
					}
				case p2.name:
					{
						interpret(p2, aname, b)
					}
				case p3.name:
					{
						interpret(p3, aname, b)
					}
				}

			}
		case "newanimal":
			{
				switch b {
				case "cow":
					{

						p1 = getCow(aname)
						fmt.Println("Created!")
					}
				case "bird":
					{
						p2 = getBird(aname)
						fmt.Println("Created!")
					}
				case "snake":
					{
						p3 = getSnake(aname)
						fmt.Println("Created!")
					}
				default:
					{
						fmt.Println("bad input")
						break
					}
				}
			}
		default:
			{
				fmt.Println("missing values")
			}
		}

	}
}
