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
	text = strings.TrimSuffix(text, "\r\n")
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

	for scanner.Scan() {
		words := strings.Fields((scanner.Text()))
		//fmt.Printf("%T %T\n", words[0], words[1])
		lcl.fname = truncLongName(words[0])
		lcl.lname = truncLongName(words[1])
		//fmt.Println(lcl)
		hold = append(hold, lcl)
		//}

		file.Close()
	}
	for _, eachline := range hold {
		fmt.Print(eachline.fname, "  ", eachline.lname, "\n")
	}
}
