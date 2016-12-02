package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var x sync.WaitGroup

	go func() { // race condition because this Goroutine is trying to access and add to a
		// shared variable x while the Goroutine on line 22 is trying to do the same
		x.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		x.Done()
	}()
	go func() {
		x.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		x.Done()
	}()
	go func() {
		x.Wait()
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
}
