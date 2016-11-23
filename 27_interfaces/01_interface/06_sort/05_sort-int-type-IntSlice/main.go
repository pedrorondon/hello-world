package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.IntSlice(n)) // convert slice of int, n, to type IntSlice which gives
	// an interface; sort.Sort then takes that interface and sorts it
	fmt.Println(n)
}
