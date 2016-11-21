package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int // variable xs is a slice
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n) // append can be used to add an entry to an empty slice
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}
