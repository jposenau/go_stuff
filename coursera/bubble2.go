package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func returnList() []string {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string of integers up to 10 to be sorted ==> ")
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	return words
}
func Swap(i int, items []int) {
	items[i+1], items[i] = items[i], items[i+1]
}
func BubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				Swap(i, items)
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

	words := returnList()
	end := len(words)
	// Limit the number of integers to 10
	if end > 10 {
		end = 10
	}

	// Do wthe work here in bubble sort
	// Build the integer array from a text string
	for val := 0; val <= end-1; val++ {
		tmp, _ := strconv.Atoi(words[val])
		intList = append(intList, tmp)
	}
	fmt.Print("Sorted List ==>")
	BubbleSort(intList)
	for _, val := range intList {
		fmt.Printf("%d ", val)
	}
}
