package main

import "fmt"

func main() { // the program enters on main
	// we create two channels and have both working:
	c := make(chan int)
	done := make(chan bool) // creates a new, second channel and calls it done; it takes
	// a bool

	// we launch three Goroutines concurrently:
	go func() {
		for i := 0; i < 10; i++ { // loops 10 times
			c <- i // puts a value on channel c
		}
		done <- true // after the looping has completed, we are placing a bool onto channel
		// done instead of signalling a WaitGroup
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done // instead of using .Wait as seen in the WaitGroup, we to this code to wait
		// for a value to come through and pull it off of channel done; we are waiting to
		// receive values  from the channel
		<-done // <-channel is the syntax for removing/pulling off values from a channel;
		// _ := <-done is not valid; the values from channel done can just be pulled off
		// and thrown away and do NOT need to be assigned to the underscore_ character
		close(c) // after the two bools have been pulled off of channel done, the channel c
		// will close
	}()

	for n := range c { // waits to receive values from unbuffered channel c like a runner waiting
		// for a baton in a relay race; this code blocks because it just waits to receive until
		// someone sends on the channel; range c needs to be paired with close(c) from line 33
		// to give it a stopping point and tell it to stop looking/waiting for values from the
		// channel
		fmt.Println(n) // prints the values from channel c
	}
} // program exits

// no WaitGroup, mutex, or atomicity used; were are being purists and only using channels;
// since we will not be using a WaitGroup, we do not need package sync
