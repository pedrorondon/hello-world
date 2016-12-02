package main

import "fmt"

func main() { // entry point

	n := 10 // you want 10 things running
	// create two channels:
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ { // looping n times which here is 10, from 0 to 9
		go func() { // this Goroutine is launched each time we loop
			for i := 0; i < 10; i++ { // loops 10 times
				c <- i // puts a value in channel c with each loop
			}
			done <- true // puts bool into channel done after looping of Goroutine has completed
		}()
	}
	/* instead of:
	   go func() {
	   	<-done
	   	<-done
	   	close(c)

	   	we use:
	*/
	go func() {
		for i := 0; i < n; i++ { // loops n times
			<-done // removes the bool from channel done for howver many, n, you have running
		}
		close(c) // closes however many, n, you have running
	}()
	for n := range c { // as c <- i puts values on the channel, this takes values off
		fmt.Println(n)
	}
}
