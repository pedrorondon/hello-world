package main

import "fmt"

func mx(nums ...int) int { // variadic parameter ...int makes it a variadic function; allows function
	// to receive a bunch of integers
	// nums becomes a slice[]int when the list of numbers are passed through
	var lrg int
	for _, v := range nums {
		if v > lrg {
			lrg = v
		}
	}
	return lrg
}

func main() {
	grt := mx(4, 7, 9, 123, 543, 23, 435, 53, 125) // this line is a func expression
	fmt.Println(grt)
}
