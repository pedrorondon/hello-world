package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // must close the channel after looping to stop the ranging below from looking for
		// more values
	}()

	for n := range c { // must be paired with close(c) to stop it from ranging further
		fmt.Println(n)
	}
}
