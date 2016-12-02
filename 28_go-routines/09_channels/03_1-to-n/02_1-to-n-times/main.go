package main

import "fmt"

func main() {
	n := 10                 // how many functions do you want pulling from the same channel? n number or 10
	c := make(chan int)     // create a channel that takes ints
	done := make(chan bool) // create a channel that takes your semaphore/flag which signals
	// completion

	go func() { // launch a Goroutine
		for i := 0; i < 100000; i++ { // loops 100,000 times
			c <- i // places values 0 to 99,999 on channel c
		}
		close(c) // when finished placeing values, close the channel
	}()
	for i := 0; i < n; i++ { // loop n number of times or 10 here
		go func() { // launched n number of times
			for n := range c { // listening to take a value off of channel c
				fmt.Println(n)
			}
			done <- true
		}()
	}
	// this code holds main from exiting until all of the code aboce has finished processing
	for i := 0; i < n; i++ { // loops n number or 10 times
		<-done // pulls the bool off of channel done
	}
}

// one Goroutine putting values onto the channel c and n number or 10 Goroutines pulling
// values off of the channel
