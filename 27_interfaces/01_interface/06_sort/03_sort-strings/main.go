package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(x) // package sort with function Strings takes a slice of strings x
	fmt.Println(x)

}
