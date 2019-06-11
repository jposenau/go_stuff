package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func teststr(s string) bool {
	// search for string starting with i containing a and ending with n

	return strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.ContainsAny(s, "a")
}
func notfound() {
	fmt.Printf("Not Found\n")
}

func found() {
	fmt.Printf("Found\n")
}

// Read a string from the keyboard and strip the line feed characters
func getName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: \n ")
	text, _ := reader.ReadString('\n')

	//remove windows based line feed characters

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func main() {
	var newstr string
	// Loop continuously until  just CR pressed to exit
	for {
		// Get a string and convert to lower case for consistency
		inval := getName()
		newstr = strings.ToLower(inval)
		if newstr == "" {
			break
		}

		switch teststr(newstr) {
		case true:
			found()
		default:
			notfound()
		} //end switch
	} // end for
}
