package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Receiver choices for eat, move, and speak
// Take an instance of Animal and retrun requested activity
func (p Animal) Eat() {
	fmt.Println(p.food)
}
func (p Animal) Move() {
	fmt.Println(p.locomotion)
}
func (p Animal) Speak() {
	fmt.Println(p.noise)
}

// Build the 2 requests by reading a string converitng to lower case fro consistency and returing the values
func returnList() []string {
	// input a string of integers as a string and convert to individual elements
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type your request > ")
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	if len(words) != 2 {
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
func interpret(p Animal, ani, val string) {
	fmt.Print("\n The ", ani, " ", val, "s -> ")
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
	// control the requests in a continuous loop
	// a blank line will exit with a bad input message
	//Build three instances of Animals

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		words := returnList()

		switch words[0] {
		case "cow":
			interpret(cow, words[0], words[1])
		case "bird":
			interpret(bird, words[0], words[1])
		case "snake":
			interpret(snake, words[0], words[1])
		default:
			fmt.Println("Requested animal doesn't exist! Try Again ")
		}
		// Limit the number of integers to 10

	}
}
