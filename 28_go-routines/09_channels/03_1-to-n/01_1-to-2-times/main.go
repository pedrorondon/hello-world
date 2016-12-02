package main

import "fmt"

func main() {
	c := make(chan int)     // create a channe that takes an int
	done := make(chan bool) // create another channel that takes a bool and will signify being done
	// launch three Goroutines:
	go func() {
		for i := 0; i < 100000; i++ { // loop 100,000 times
			c <- i // put values onto channel c
		}
		close(c)
	}()
	// thee two functions pull values from channel c
	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true // once you are done pulling values off of channel c, place a bool, true,
		// on channel done
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()
	<-done // waits for channel done to receive a bool so that value can then be removed off of
	// the channel, signalling completion
	<-done
}
