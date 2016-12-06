package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen() // func gen returns a receive only channel, <-chan int and assigns it to
	// variable in

	// Fan Out because multiple functions are reading from the same channel until that channel
	// is closed
	f1 := factorial(in) // the receive only channel, variable n, is passed onto three workers,
	// f1-f3
	f2 := factorial(in)
	f3 := factorial(in)

	// Fan In because multiple channels are wrtiting onto a single channel; we merge the channels
	// from f1-f3 onto a single channel
	for n := range merge(f1, f2, f3) { // will pull values off from the channels; blocks
		// func main from exiting too early
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ { // 100 loops
			for j := 3; j < 13; j++ { // each of the 100 loops will have 10 values; total of
				// 1,000 values
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out2 := make(chan int)
	go func() {
		for n := range in {
			out2 <- fact(n)
		}
		close(out2)
	}()
	return out2
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(in2 ...<-chan int) <-chan int { // func merge takes a variatic parameter, slice of
	// <- chan int
	var x sync.WaitGroup // create a WaitGroup
	out3 := make(chan int)

	/* Alternate method:

	output := func(c <-chan int) {
		for n := range c {
			out3 <- n // merging all values onto channel out3
		}
		x.Done()
	}

	x.Add(len(in2))
	for _, v := range in2 {
		go output(v) launch a Goroutine a pass the current chan int, v
	}
	*/

	x.Add(len(in2)) // add to the WaitGroup the length of your slice from the variatic
	// parameter
	for _, v := range in2 { // range over all of the parameters that came into func merge;
		// note the syntax to range over a slice when you only want the value,
		// for _, v := range in2 {
		go func(c <-chan int) {
			for n := range c {
				out3 <- n
			}
			x.Done()
		}(v)
	}

	go func() { // this Goroutine will close out once all the output Goroutines are done;
		// this pertains to the alternate method that is commented out
		x.Wait()
		close(out3)
	}()
	return out3
}

// the commented code replaces line 80-91;  skip down to 93 to complete program
