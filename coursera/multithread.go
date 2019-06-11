package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MultiMergeSort(data []float64, res chan []float64) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []float64)
	rightChan := make(chan []float64)
	middle := len(data) / 2

	go MultiMergeSort(data[:middle], leftChan)
	go MultiMergeSort(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- Merge(ldata, rdata)
	return
}

func RunMultiMergeSort(data []float64) (multiResult []float64) {
	res := make(chan []float64)
	go MultiMergeSort(data, res)
	multiResult = <-res
	return
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

func main() {

	slice := generateSlice(50)
	fmt.Println("\n --- unsorted --- \n\n", slice)
	fmt.Println("\n--- sorted ---\n\n", MergeSort(slice), "\n")
}
