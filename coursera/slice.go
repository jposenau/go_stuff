package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Read a string from the keyboard and strip the line feed characters
func getName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter integer value X to finish: \n ")
	text, _ := reader.ReadString('\n')

	//remove windows based line feed characters

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func main() {
	var newstr string
	//create a slice of capacity 3 with an inital length of 1
	var numbers = make([]int, 1, 3) // create a slice of size 3
	// Loop continuously until  just CR pressed to exit

	i := 0
	for {
		// Get a string and convert to lower case for consistency
		inval := getName()
		newstr = strings.ToLower(inval)
		//is the value an X leave
		if newstr == "x" {
			break
		}
		val, _ := strconv.Atoi(newstr)

		if i > len(numbers)-1 {
			numbers = append(numbers, val)
		} else {
			numbers[0] = val
		}
		//	fmt.Println(" string length = ", len(numbers))
		sort.Ints(numbers)
		//sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

		for _, value := range numbers {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
		i++

	} // end for
}
