package main

import "fmt"

func main() {
	c := incrementor()
	/*
			   instead of:
			       cSum := puller(c) instead of assigning result puller(c) to variable cSum and then
		                             ranging over variable cSum in the next line of code, we can use
		                             code substitution
			   	for n := range cSum {
			   we can use:
	*/
	for n := range puller(c) { // we use code substitution and the concept of equality; process
		// of shortening is making code better is called refactoring
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out2 := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out2 <- sum
		close(out2)
	}()
	return out2
}
