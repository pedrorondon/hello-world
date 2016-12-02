package main

import "fmt"

func main() {
	c1 := incrementor("Foo:") // calls function incrementor and assignes the returned chan int
	// to variable c1
	c2 := incrementor("Bar:")
	c3 := puller(c1) // calls function puller and passes the chan int, c1, and calls the newly
	// returned chan int c3
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4) // pulls values from channels c3 and c4 and adds
	// them together;
	// serves to block func main from exiting since it is waiting for values that will only
	// be produced once the other code has executed
}

func incrementor(s string) chan int { // function that takes a string and return a chan int
	out := make(chan int) // create a chan int and call it out
	go func() {           // launch a Goroutine
		for i := 0; i < 20; i++ { // loop 20 times
			out <- i          // place values 0-19 onto channel out
			fmt.Println(s, i) // print string Foo: and the iteration it is one from 0-19
		}
		close(out) // close the channel
	}()
	return out // return the channel
}

func puller(x chan int) chan int { // passes chan int c1 as an argument can calls it variable x;
	// also returns a chan int
	out2 := make(chan int) // create a new channel and call it out2
	go func() {            // launch Goroutine
		var sum int
		for n := range x { // range over channel c1 known as variable x
			sum += n // accumulator variable; adds up values from channel c1 known as variable x
		}
		out2 <- sum // places sum total on channel out2
		close(out2) // closes channel out2
	}()
	return out2 // returns channel out2
}

// no deadlock because when I am sending through func incrementor, I am simultaneously/concurrently
// receiving with func puller
