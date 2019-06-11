package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// random merge sort routine
func sort1(A []int) []int {

	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = sort1(left)
	right = sort1(right)
	return merge(left, right)
}

// split array into two
func split(A []int) ([]int, []int) {
	return A[0 : len(A)/2], A[len(A)/2:]
}

// assumes that A and B are sorted
func merge(A, B []int) []int {
	arr := make([]int, len(A)+len(B))

	// index j for A, k for B
	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		// fix for index out of range without using sentinel
		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		}
		// default loop condition
		if A[j] > B[k] {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}

	return arr
}

// Take a strin of integer representations and return count values
func returnList() []string {
	// input a string of integers as a string and convert to individual elements
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string of integers up to 20 to be sorted ==> ")
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	return words
}

// Build the four goroutines to mange the subsections of sorts
func sortA(wg *sync.WaitGroup, AC []int) []int {
	fmt.Println(" Func SortA list ", AC)
	AC = sort1(AC)
	fmt.Println("First sorted group  ", AC)

	wg.Done()
	return AC
}
func sortB(wg *sync.WaitGroup, AD []int) []int {
	fmt.Println(" Func SortB list ", AD)
	AD = sort1(AD)
	fmt.Println("Second sorted group  ", AD)
	wg.Done()
	return AD
}
func sortC(wg *sync.WaitGroup, BC []int) []int {
	fmt.Println(" Func SortC list ", BC)
	BC = sort1(BC)
	fmt.Println("Third sorted group  ", BC)
	wg.Done()
	return BC
}
func sortD(wg *sync.WaitGroup, BD []int) []int {
	fmt.Println(" Func SortD list ", BD)
	BD = sort1(BD)
	fmt.Println("Forth sorted group  ", BD)
	wg.Done()
	return BD
}
func main() {
	var wg sync.WaitGroup
	intList := make([]int, 0, 20)

	words := returnList()
	end := len(words)
	// Limit the number of integers to 20
	if end > 20 {
		end = 20
	}

	// Build the integer array from a text string
	for val := 0; val <= end-1; val++ {
		tmp, _ := strconv.Atoi(words[val])
		intList = append(intList, tmp)
	}
	// splt the array into 4 chuck for individual calls

	Orig := intList
	A, B := split(Orig)
	AC, AD := split(A)
	BC, BD := split(B)
	//fmt.Println(" Initial unsorted splits \n", AC, AD, BC, BD)
	// Set up synchronization fro 4 routines
	//and call the functions with a wati for completion
	wg.Add(4)
	go sortA(&wg, AC)
	go sortB(&wg, AD)
	go sortC(&wg, BC)
	go sortD(&wg, BD)
	wg.Wait()
	x := merge(BC, BD)
	y := merge(AC, AD)

	z := merge(x, y)

	// Merge final results combine and print

	fmt.Println(" Final Merged Results", sort1(z))

}
