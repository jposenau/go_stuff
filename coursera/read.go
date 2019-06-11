package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func getName() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r")
	return text
}
func truncLongName(buffer string) string {

	if len(buffer) > 20 {
		return string(buffer[0:20])
	} else {
		return buffer
	}
}
func main() {

	var lcl name
	// create a holder for tructs with initial capacity of 10
	hold := make([]name, 0, 10)

	fmt.Println("Enter file name to read names  ")
	inval := getName()
	file, err := os.Open(inval)
	if err != nil {
		fmt.Println("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//var txtlines []string
	// break the textline into individual words and ensure that they are no longer than 20 characters per assignment
	for scanner.Scan() {
		words := strings.Fields((scanner.Text()))
		lcl.fname = truncLongName(words[0])
		lcl.lname = truncLongName(words[1])
		hold = append(hold, lcl)

		file.Close()
	}
	//print out the structs
	for _, eachline := range hold {
		fmt.Print(eachline.fname, "  ", eachline.lname, "\n")
	}
}
