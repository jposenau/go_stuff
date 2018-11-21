package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()

	ttime := "11-19-2018"
	fmt.Println(ttime)

	//Format MM-DD-YYYY
	fmt.Println("time.Time format", dt.Format("2006-01-20"))

	fmt.Printf("%v converted value type \n", dt)
	//Format MM-DD-YYYY hh:mm:ss
	fmt.Println("shortened time format", dt.Format("01-02-2006 15:04:05"))

	//With short weekday (Mon)
	fmt.Println(dt.Format("01-02-2006 15:04:05 Mon"))

	//With weekday (Monday)
	fmt.Println(dt.Format("01-02-2006 15:04:05 Monday"))

	//Include micro seconds
	fmt.Println(dt.Format("01-02-2006 15:04:05.000000"))

	//Include nano seconds
	fmt.Println(dt.Format("01-02-2006 15:04:05.000000000"))

	fmt.Println("******************************************************")
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	myear := year
	myval := "2015-11-11"

	fmt.Printf("year = %v\n", year)
	fmt.Printf("my year %v\n", myear)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)
	t, err := time.Parse("2006-01-02", myval)
	fmt.Println(" parsed value is:", t, err)

}
