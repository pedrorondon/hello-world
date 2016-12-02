package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var x sync.WaitGroup
var count int // global counter

func main() {
	x.Add(2)
	go incrementor("Foo: ") // variable incrementor is called twice
	go incrementor("Bar: ")
	x.Wait()
	fmt.Println("Final Counter:", count)
}
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := count // global counter value is stored
		// in x
		x++                                                        // increments the counter by 1
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond) // sleep for a while
		count = x                                                  // then take that incremented
		// value and write it back to the global variable
		fmt.Println(s, i, "Counter:", count)
	}
	x.Done()
}

// go run -race main.go checks for race conditions

// a race condition is when you have things running concurrently and different
// processes are trying to access the same variable simultaneously; overwriting
// can occur and you can get mistakes in your code
