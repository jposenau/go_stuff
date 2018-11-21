package main

import (
	"fmt"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
func beforeTimeSpan(start, end, check time.Time) bool {
	return check.Before(start)
}

func main() {
	//start, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
	//end, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")
	start, _ := time.Parse("2001-01-01", "2015-1-15")
	end, _ := time.Parse("2001-01-01", "2015-2-15")
	fmt.Println(" start time is", start)

	in, _ := time.Parse(time.RFC822, "01 Jan 15 20:00 UTC")
	//out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")
	out := time.Now()

	if inTimeSpan(start, end, in) {
		fmt.Println(in, "is between", start, "and", end, ".")
	}

	if !inTimeSpan(start, end, out) {
		fmt.Println(out, "is not between", start, "and", end, ".")
	}
	if !beforeTimeSpan(start, end, out) {
		fmt.Println(out, "after", start, ".")
	}
	input := "2017-08-31"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017
}
