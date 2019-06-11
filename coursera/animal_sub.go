package main

import (
"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}
	
func main() {
	
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	
	
	for {
		
	    fmt.Print("> ")
	    var ani, info string
	    fmt.Scan(&ani)
	    fmt.Scan(&info)
	    switch {
			case ani == "cow":
			    cow.Know(info)
			case ani == "bird":
			    bird.Know(info)
			case ani == "snake":
			    snake.Know(info) 
			}
		}
}

func (a Animal) Eat() string {
    return a.food
}

func (a Animal) Move() string {
    return a.locomotion
}
	
func (a Animal) Speak() string {
    return a.noise
}

func (a Animal) Know(info string){
	switch {
		case info == "eat":
		    fmt.Println(a.Eat())
		case info == "move":
		    fmt.Println(a.Move())
		case info == "speak":
		    fmt.Println(a.Speak())
		}
}
