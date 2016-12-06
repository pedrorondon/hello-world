package main

import (
	"fmt"
	"strconv"
)

func main() { // entry point for executable
	c := incrementor(2) // pass integer 2 as an argument to func incrementor
	// assign the values on the returned chan string out to variable c;

	var count int      // the variable count is integer 0 by default
	for n := range c { // loop over variable c and assign values to variable n; pull values
		// off of the channel; serves to block func main from exiting as it awaits the values
		// from variable c, previously known as channel out that was returned by func incrementor
		count++ // for every value pulled from variable c, add 1 to the variable count,
		// which by default began at integer 0
		fmt.Println(n) // print the chan string values from variable n
	}
	fmt.Println("Final Count:", count)
}

func incrementor(n int) chan string {
	out := make(chan string) // create a channel and name it out
	done := make(chan bool)  // create a second channel and name it done; it will act as a
	// semaphore, letting us know when everthing is finished

	for i := 0; i < n; i++ { // loop two times, 0 and 1
		go func(i int) { // launch a Goroutine
			for j := 0; j < 20; j++ { // loop 20 times
				out <- fmt.Sprint("Foo: "+strconv.Itoa(i)+" Bar: ", j) // here, we are putting
				// a string onto the channel out; we can convert an int to ASCII to concatenate it
				// with the other strings or simply write the variable as we did with j
			}
			done <- true // once 20 loops are complete, place bool true on channel done
		}(i) // using closure
	}

	go func() {
		for i := 0; i < n; i++ { // loop n or two times, 0 and 1
			<-done // pull off bool true from channel done during each loop; serves to block
			// until the code above has executed completely and placed bool true on channel done
			// twice
		}
		close(out) // closes channel out
	}()
	return out
}
