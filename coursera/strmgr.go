package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func teststr(s string) bool {

	return strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.ContainsAny(s, "a")
}
func notfound() {
	fmt.Printf("Not Found\n")
}

func found() {
	fmt.Printf("Found\n")
}

func getName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: \n ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	return text
}

func main() {
	var newstr string

	inval := getName()
	newstr = strings.ToLower(inval)

	switch teststr(newstr) {
	case true:
		found()
	default:
		notfound()
	}
}
