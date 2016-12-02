package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var x sync.WaitGroup // assign type sync.WaitGroup to variable x

func init() { // allows for setup and initialization before running/executing other code;
	// func init runs once and runs first
	runtime.GOMAXPROCS(runtime.NumCPU()) // tells the program to use all of the computer's
	// cores; this line of code is no longer necessary; current versions of Go use more than
	// one core by defaukt
}

func main() {
	x.Add(2) // this line adds two items to the WaitGroup
	go foo()
	go bar()
	x.Wait() // will wait until WaitGroup goes to 0
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	x.Done() // when func is complete, it will say it is done and take one away from the
	// wait group and go from 2 to 1
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	x.Done() // when func is complete, it will say it is done and take one away from the
	// wait group and go from 1 to 0
}
