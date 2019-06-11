package main

import (
	"fmt"
	"time"
)

func main() {
	Orig := []int{4, 7, 6, 9, 3, 2, 1, 6, 8, 9, 5, 43, 7, -1, 5, 3, 4}
	A, B := split(Orig)
	AC, AD := split(A)
	BC, BD := split(B)
	fmt.Println(" Initial splits \n", AC, AD, BC, BD)

	//A := []int{3, 5, 1, 6, 1, 7, 2, 4, 5}
	//B := []int{7, 6, 8, 3, 2, 1, 5}
	//C := []int{3, 5, 1, 6, 8, 4, 5, 7, 1}
	//D := []int{13, 1, 4, 6, 3, 8, 4}
	go fmt.Println(" First split \n", sort1(AC))
	go fmt.Println(" Second split \n", sort2(AD))
	go fmt.Println(" Third split \n", sort3(BC))
	go fmt.Println(" Fourth split \n", sort4(BD))
	x := merge(BC, BD)
	y := merge(AC, AD)

	z := merge(x, y)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(" Final sorted splits \n", AC, AD, BC, BD)

	fmt.Println(" Final Result", sort1(z))

}

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
