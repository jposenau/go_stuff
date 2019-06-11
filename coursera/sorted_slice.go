package main

import (
	"fmt"
	"sort"
)

func bubbleSort(slice []int) {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}
func main() {

	var input int
	var slice []int
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter an integer %d value : ", i)
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("First Name: %d\n", input)
		slice = append(slice, input)
	}
	fmt.Println(slice)

	bubbleSort(slice)
	fmt.Println("Sorted slice")

	fmt.Println(slice)
}
