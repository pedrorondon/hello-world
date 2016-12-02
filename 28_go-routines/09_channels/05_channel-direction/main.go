package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int { // <-operator specifies the channel direction, send or receive;
	// <-chan is notation that indicates that this is a receive only channel; you can only
	// receive values from it since it closes after all the values have been placed in it
	// <-chan int means that some variable can receive ints from the channel
	// chan<- is a send only channel
	// if not direction is given, the channel is bidirectional
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
