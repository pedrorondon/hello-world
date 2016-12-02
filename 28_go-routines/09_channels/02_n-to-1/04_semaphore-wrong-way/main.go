package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	<-done // we are blocked here until a bool is placed on channel done via done <- true
	// so the bool can then be removed with this code but this will never happen because
	// the values in channel c are not being received by range c below;
	// to prevent this deadlock, these must be launched within their own anonymous,
	// self-executing functions and make them go routines by adding the prefix go to func();
	// since the Goroutine is out running on itws own, the range c code in line 35 is free
	// to receive the values placed in channel c from the two Goroutine above
	<-done
	close(c)

	for n := range c { // this will take values off of channel c but we never get here due to
		// the block above
		fmt.Println(n)
	}
}
