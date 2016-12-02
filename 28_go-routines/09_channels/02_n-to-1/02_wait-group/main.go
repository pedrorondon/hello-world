package main

import (
	"fmt"
	"sync"
)

func main() { // entry point for the program

	c := make(chan int) // create a channel and assign it to the variable c

	var x sync.WaitGroup // create variable x of type sync.WaitGroup; its scope is from this point
	// to the end of the code

	x.Add(2) // add two items to the WaitGroup
	// .Add() was moved outside of the scope of the individual Goroutines

	// the following three Goroutines get launch simultaneously/concurrently:
	go func() { // an anonymous, self-executing function
		for i := 0; i < 10; i++ { // loops 10 times
			c <- i // places a value onto the channel each time it loops
		}
		x.Done() // after completion of the looping, the code will signal to the WaitGroup that
		// it has finished
		// .Done() can remain inside of each respective Goroutine while .Add() was moved to
		// package scope
	}()
	go func() { // same a function above
		for i := 0; i < 10; i++ {
			c <- i
		}
		x.Done()
	}()
	go func() {
		x.Wait() // waiting for the items in the WaitGroup x to complete; waiting for the two
		// Goroutines above to finish
		close(c) // closes the channel
	}()
	for n := range c { // waits to receive values from channel c and even drains any remaining
		// values once the channel has closed; these values will be assigned to variable n
		fmt.Println(n) // print variable n
	}
} // exit function main
