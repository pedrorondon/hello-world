package main

import "fmt"

func main() { // enter main
	c := incrementor() // call incrementor function which will run and return a channel out
	// and then assign it to variable c
	cSum := puller(c) // passing the channel c, a chan int, into the puller function which takes
	// a chan int
	for n := range cSum { // ranges over variable cSum (previously channel out2)
		// also, blocks the main function from exiting until it receives values which will only
		// occur after the previous Goroutines have completed
		fmt.Println(n) // prints value
	}
}

func incrementor() chan int { // func incrementor returns a chan that takes an int
	out := make(chan int) // creates a channel and assigns it to variable out
	go func() {
		for i := 0; i < 10; i++ { // loops 10 times
			out <- i // places 10 values on channel out
		}
		close(out) // after all of the values have been placed on the channel, it is closed
	}()
	return out // return that channel full of its 10 values
}

func puller(c chan int) chan int { // takes/receives an argument that is defined with parameter
	// of being a channel that takes an int, chan int; then chan int is assigned to variable c;
	// func puller takes the chan int created by func incrementor
	out2 := make(chan int) // creates another channel and calls it out2
	go func() {
		var sum int
		for n := range c {
			sum += n // used to accumulate values; adds up all the values stored in channel c
		}
		out2 <- sum // the sum total is placed on channel out2
		close(out2) // closed the channel
	}()
	return out2 // return that channel which will be assigned to variable cSum on line 8 above
}

// a great tip when reading and understanding code that uses concurrency to to follow the flow of
// execution

// a deadlock occurs when there is a mismatch between the passing and receiving pair; one
// not there to pass or the other is not there to receive
