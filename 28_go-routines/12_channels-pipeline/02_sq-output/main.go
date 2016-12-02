package main

import "fmt"

func main() {
	for n := range sq(gen(3, 4)) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int { // ...int is a variatic parameter that takes an unlimited amount
	// of ints and assigns them to variable nums; then returns a chan of ints
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out2 := make(chan int)
	go func() {
		for n := range in {
			out2 <- n * n
		}
		close(out2)
	}()
	return out2
}
