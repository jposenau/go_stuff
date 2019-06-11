package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getName() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	return text
}

func main() {

	var person = make(map[string]string)

	fmt.Print("Enter Name: \n ")
	name := getName()
	fmt.Print("Enter Address \n ")
	address := getName()
	person["name"] = name
	person["address"] = address
	// encode using json Marshal
	// order encoded alphabetically
	barr, err := json.Marshal(person)
	if err != nil {
		fmt.Println("json error")
	} else {

		fmt.Println("Resulting JSON string for address and name \n ", string(barr))
	}

}
