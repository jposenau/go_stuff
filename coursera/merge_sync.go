package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// create 4 separate sort routine to work concurrently on individual splits.
// Take the final results and do a final sort to merge the results
func sort1(A []int) []int {

	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = sort1(left)
	right = sort1(right)
	return merge(left, right)
}
func sort2(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = sort2(left)
	right = sort2(right)

	return merge(left, right)

}
func sort3(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = sort3(left)
	right = sort3(right)

	return merge(left, right)
}
func sort4(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = sort4(left)
	right = sort4(right)

	return merge(left, right)
}
func fsort(A []int) []int {

	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = fsort(left)
	right = fsort(right)
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
func returnList() []string {
	// input a string of integers as a string and convert to individual elements
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string of integers up to 10 to be sorted ==> ")
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	return words
}
func sortA(wg *sync.WaitGroup, AC []int) []int {

	AC = sort1(AC)
	fmt.Println("First sorted group  ", AC)

	wg.Done()
	return AC
}
func sortB(wg *sync.WaitGroup, AD []int) []int {
	AD = sort1(AD)
	fmt.Println("Second sorted group  ", AD)
	wg.Done()
	return AD
}
func sortC(wg *sync.WaitGroup, BC []int) []int {
	BC = sort1(BC)
	fmt.Println("Third sorted group  ", BC)
	wg.Done()
	return BC
}
func sortD(wg *sync.WaitGroup, BD []int) []int {
	BD = sort1(BD)
	fmt.Println("First sorted group  ", BD)
	wg.Done()
	return BD
}
func main() {
	var wg sync.WaitGroup
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
	// splt the array into 4 pieces
	Orig := intList
	A, B := split(Orig)
	AC, AD := split(A)
	BC, BD := split(B)
	fmt.Println(" Initial unsorted splits \n", AC, AD, BC, BD)
	wg.Add(4)
	go sortA(&wg, AC)
	go sortB(&wg, AD)
	go sortC(&wg, BC)
	go sortD(&wg, BD)
	wg.Wait()
	x := merge(BC, BD)
	y := merge(AC, AD)

	z := merge(x, y)
	//time.Sleep(1000 * time.Millisecond)
	//fmt.Println(" Final sorted splits \n", AC, AD, BC, BD)

	fmt.Println(" Final Result", fsort(z))

}
