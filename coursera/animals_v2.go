package main

import (
	"fmt"
)

type Cow struct{}
type Bird struct{}
type Snake struct{}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (snake Snake) Move() {
	fmt.Println("mice")
}

func (snake Snake) Eat() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var command string
	var animalName, animalType, animalRequest string

	animals := make(map[string]Animal)

	fmt.Println("Enter Q for quit.")

	for {
		fmt.Println(">")

		fmt.Scanf("%s", &command)
		if command == "Q" {
			return
		}

		if command == "newanimal" {
			fmt.Scanf("%s %s", &animalName, &animalType)

			switch animalType {
			case "cow":
				animals[animalName] = Cow{}
				fmt.Println("Created it!")
			case "bird":
				animals[animalName] = Bird{}
				fmt.Println("Created it!")
			case "snake":
				animals[animalName] = Snake{}
				fmt.Println("Created it!")
			default:
				fmt.Println("Unknown animal")
			}
		} else if command == "query" {
			fmt.Scanf("%s %s", &animalName, &animalRequest)

			animal := animals[animalName]

			if animal != nil {
				switch animalRequest {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				}
			} else {
				fmt.Println("Animal not found!")
			}
		}
	}
}
