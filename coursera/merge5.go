package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	slice := generateSlice(50)
	fmt.Println("\n --- unsorted --- \n\n", slice)
	x := mergesortv1(slice)
	fmt.Println("\n--- sorted ---\n\n", x, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

// Runs MergeSort algorithm on a slice single
func mergesortv1(s []int) {
	len := len(s)

	if len > 1 {
		middle := len / 2

		var wg sync.WaitGroup
		wg.Add(2)

		// First half
		go func() {
			defer wg.Done()
			mergesortv1(s[:middle])
		}()

		// Second half
		go func() {
			defer wg.Done()
			mergesortv1(s[middle:])
		}()

		// Wait that the two goroutines are completed
		wg.Wait()
		Merge(s, middle)
	}
}

// Merges left and right slice into newly created slice
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	count := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[count] = left[i]
			count, i = count+1, i+1
		} else {
			slice[count] = right[j]
			count, j = count+1, j+1
		}
	}
	for i < len(left) {
		slice[count] = left[i]
		count, i = count+1, i+1
	}
	for j < len(right) {
		slice[count] = right[j]
		count, j = count+1, j+1
	}

	return slice
}
