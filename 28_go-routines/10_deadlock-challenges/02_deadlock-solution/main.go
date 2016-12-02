package main

import "fmt"

func main() {
	c := make(chan int)
	go func() { // put code into a Goroutine that has it running off on its and
		// moves it out of the way
		c <- 1
	}()
	fmt.Println(<-c) // since Goroutine is running independently, <-c is free and ready
	// to receive values from that channel
}
