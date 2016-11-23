package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(x)
	sort.StringSlice(x).Sort() // convert x which is a slice of string to a StringSlice so you will have the
	// sort method available to you
	// sort.Sort(sort.StringSlice(x)) is an alternate way to sort the information; the first sort is the package
	// name; the second Sort is the function name; in other words, from package sort, give me function Sort;
	// function Sort takes an interface StringSlice(x)
	fmt.Println(x)
}
