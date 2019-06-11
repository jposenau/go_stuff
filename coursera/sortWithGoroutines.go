package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortSubarray(wg *sync.WaitGroup, s []int, res *[]int) {
	fmt.Println("Next subarray will be sorted:", s)
	sort.Ints(s)
	*res = append(*res, s...)
	wg.Done()
}

func main() {
	fmt.Print("Enter a series of integers separated by a space (at least 4 integers): ")

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	userInputSlice := strings.Split(userInput, " ")

	userInputSliceLen := len(userInputSlice)

	// we should have at least 4 goroutines
	if userInputSliceLen < 4 {
		fmt.Println("The series should be at least 4 integers. Try again.")
		os.Exit(1)
	}

	// convert slice of strings to slice of integers
	var userInputSliceInt = []int{}
	for _, i := range userInputSlice {
		j, _ := strconv.Atoi(i)
		userInputSliceInt = append(userInputSliceInt, j)
	}

	partsCount := 4
	partLenApprox := int(math.Round(float64(userInputSliceLen) / float64(partsCount)))

	// partition the array into 4 parts
	var partsSlice [][]int
	low := 0
	high := partLenApprox
	for i := 1; i <= partsCount; i++ {
		if i != partsCount {
			partsSlice = append(partsSlice, userInputSliceInt[low:high])
			low = high
			high = high + partLenApprox
		} else {
			partsSlice = append(partsSlice, userInputSliceInt[low:])
		}
	}

	// main logic
	var wg sync.WaitGroup
	var resultSlice = []int{}
	for _, j := range partsSlice {
		wg.Add(1)
		go sortSubarray(&wg, j, &resultSlice)
	}
	wg.Wait()
	sort.Ints(resultSlice)
	fmt.Println("Entire sorted list:", resultSlice)
}
