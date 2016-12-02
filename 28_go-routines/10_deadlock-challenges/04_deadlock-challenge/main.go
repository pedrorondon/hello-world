package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for { // creates an eternal loop
		fmt.Println(<-c) // after printing all of the values from channel c, it continues
		// to loop and waits for more values even when no more are available
	}
}

// will print all values 0-9 but creates a deadlock because it is still waiting to receive
// values even when the channel c has drained all of its values
