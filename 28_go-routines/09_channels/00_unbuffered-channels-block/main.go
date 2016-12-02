package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // as seen with maps and slices, we use make() to create a channel;
	// here we have a channel over which we will be able to communicate an int;
	// this is an unbuffered channel

	go func() { // first Goroutine
		for i := 0; i < 10; i++ { // we are going from 0 to 9 and taking that number, i,
			// and putting it in the channel
			c <- i // once the number is put into the channel with this notation c <-,
			// the code stops and waits for something to remove that number;
			// like passing a baton during a relay race
		}
	}()

	go func() { // second Goroutine
		for {
			fmt.Println(<-c) // here you are taking the previously placed number off of
			// the channel; you are receiving the value from the channel and printing it;
			// from the chanel, give me whatever value is there
		}
	}()
	time.Sleep(time.Second) // the use of time here is similar to a WaitGroup
}
