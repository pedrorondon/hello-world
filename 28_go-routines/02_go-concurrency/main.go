package main

import "fmt"

func main() {
	go foo() // adding the word go creates concurrency
	go bar()
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

// nothing prints because three threads are running simultaneously and func main ends before
// func foo and func bar can print their values
