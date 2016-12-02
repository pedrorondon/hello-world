package main

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
	}

}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
	}
}

// no concurrency is this example; instead, the code is running sequentially; func foo runs first
// and prints all the values fom 0 to 44 and then func bar runs and prints its values from 0 to 44;
// in other words, func foo has to run to cpmpletion before func bar runs
