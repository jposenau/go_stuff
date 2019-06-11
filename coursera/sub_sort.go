package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const numOfGoroutines = 4

func main() {
	intsToSort, err := readInts()
	if err != nil {
		fmt.Println(err)
		return
	}

	chunkSize := (len(intsToSort) + numOfGoroutines - 1) / numOfGoroutines
	numOfWorkers := len(intsToSort) / chunkSize

	// buffered channel to not block write in sort till merging
	subarrs := make(chan []int, numOfWorkers)

	var sorted []int
	// worker pool - sort sub array in a different goroutine
	for i := 0; i < len(intsToSort); i += chunkSize {
		end := i + chunkSize
		if end > len(intsToSort) {
			end = len(intsToSort)
		}
		go sort(intsToSort[i:end], subarrs)
	}

	// merge subarrays in main goroutine
	for i := 0; i < numOfWorkers; i++ {
		if len(sorted) == 0 {
			sorted = append(sorted, <-subarrs...)
			continue
		}
		var newArr []int
		subarr := <-subarrs
		i, j := 0, 0
	inner:
		for {
			if i < len(sorted) && j < len(subarr) && sorted[i] < subarr[j] {
				newArr = append(newArr, sorted[i])
				i++
				continue inner
			} else if i < len(sorted) && j >= len(subarr) {
				newArr = append(newArr, sorted[i])
				i++
				continue inner
			} else if j < len(subarr) {
				newArr = append(newArr, subarr[j])
				j++
				continue inner
			}
			break inner
		}

		sorted = newArr
	}

	fmt.Println("sorted array", sorted)
}

func sort(subarr []int, ch chan<- []int) {
	arr := make([]int, len(subarr))
	copy(arr[:], subarr)
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			swap := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = swap
		}
	}
	fmt.Println("subarray: ", arr)
	ch <- arr
}

func readInts() ([]int, error) {
	fmt.Print("Enter Numbers: ")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("cannot read from console: %+v", err)
	}
	fmt.Println(line)
	ints, err := parseInts(line)
	if err != nil {
		return nil, fmt.Errorf("cannot parse integers: %+v", err)
	}
	return ints, nil
}

func parseInts(s string) ([]int, error) {
	var ints []int
	for _, r := range s {
		if r == '\n' {
			break
		}
		fmt.Println(s)
		value, err := strconv.Atoi(string(r))
		fmt.Println(" current value", value, err)
		if err != nil {
			return nil, fmt.Errorf("cannot convert %d to int: %+v", value, err)
		}
		ints = append(ints, value)
		fmt.Printf("%b ", ints)
	}
	return ints, nil
}
