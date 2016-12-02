package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1 // error because here you are putting a value on the channel but there is nothing
	// ready to receive it
	fmt.Println(<-c)
}

// This result is a deadlock
