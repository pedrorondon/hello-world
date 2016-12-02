package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	fmt.Println(<-c)
}

// error: only prints 0 instead of 0-9
