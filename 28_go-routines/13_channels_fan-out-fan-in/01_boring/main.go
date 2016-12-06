package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(awesome("Dame"), awesome("Brenn"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both awesome!!")
}

func awesome(msg string) <-chan string { // since two strings are passed as arguments to
	// func awesome(), this function returns receive only channels, <-chan string
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s %d", msg, i) // Sprint means string print
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return out
}

func fanIn(input1, input2 <-chan string) <-chan string { // this functions takes the two
	// <-chan string returned by func awesome() and names them input1 and input2
	out2 := make(chan string)
	go func() {
		for {
			out2 <- <-input1 // one of two channels writing to channel out2
		}
	}()
	go func() {
		for {
			out2 <- <-input2 // second of two channels writing to same channel out2,
			// casuing fan in to occur
		}
	}()
	return out2
}

// Fan In: when multiple channels all write to the same individual channel
// Fan Out: when multiple funcs read from a single channel
