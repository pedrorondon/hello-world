package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // when we're done putting values on the channel, close that channel;
		// when a channel is closed, you can no longer put values onto the channel;
		// note, you can still receive values from the channel;
	}()
	for n := range c { // will loop over the channel until it is closed and empty
		fmt.Println(n)
	}
}

// two different Goroutines passing a value over the channel
