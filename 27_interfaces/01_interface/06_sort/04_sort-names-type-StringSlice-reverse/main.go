package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(x)
	sort.Sort(sort.Reverse(sort.StringSlice(x))) // slice of strings has been converted to
	// type  StringSlice
	fmt.Println(x)
}
