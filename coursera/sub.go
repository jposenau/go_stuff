package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func convertStringToFloat(text string) (float64, error) {
	return strconv.ParseFloat(text, 64)
}

func Contains(a []int, x int) bool {
	fmt.Print("Len : " + strconv.Itoa(len(a)))
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {

	intSlice := make([]int, 0)

	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Please insert Integer")
	fmt.Println("----------------------")

	for reader.Scan() {
		fmt.Print("-> ")
		strInput := reader.Text()
		intInput, _ := strconv.Atoi(strInput)

		if Contains(intSlice, intInput) {
			fmt.Printf(" and Failed : Cannot add new int : %s", strconv.Itoa(intInput))
			fmt.Print("\n")
		} else {

			intSlice = append(intSlice, intInput)
			sort.Ints(intSlice)
			fmt.Print(" and Contents : ")
			fmt.Print(intSlice)
			fmt.Print("\n")

		}

	}

}
