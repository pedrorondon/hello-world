package main

import "fmt"

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println("Total:", n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; n > 0; i-- { // initializer, condition, post operation
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

// launching concurrently running Goroutines will utilize all of the CPU's cores
// and get calculations processed faster
