package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getName() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r")
	return text
}
func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
func main() {
	intList := make([]int, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanLines)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	fmt.Printf("%d\n", len(words))

	// Do whatever you want here
	for _, eachline := range words {
		tmp, _ := strconv.Atoi(eachline)
		intList = append(intList, tmp)
	}
	bubblesort(intList)
	for _, val := range intList {
		fmt.Printf("%d ", val)
	}
}
