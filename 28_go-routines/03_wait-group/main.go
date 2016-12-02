package main

import (
	"fmt"
	"sync"
)

var x sync.WaitGroup // assign type sync.WaitGroup to variable x

func main() {
	x.Add(2) // this line adds two items to the WaitGroup
	go foo()
	go bar()
	x.Wait() // will wait until WaitGroup goes to 0
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
	}
	x.Done() // when func is complete, it will say it is done and take one away from the
	// wait group and go from 2 to 1
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
	}
	x.Done() // when func is complete, it will say it is done and take one away from the
	// wait group and go from 1 to 0
}
